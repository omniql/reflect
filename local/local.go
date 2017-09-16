package local

import (
	"github.com/omniql/reflect/spec"
	"github.com/omniql/reflect"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
	"strings"
)

//LocalBuilder build a reflection object from a local file path
type Builder struct {
	app          *applicationContainer
	store        map[string]*oType
	files        map[string]interface{}
	files_types  map[string]string
	externalApps map[string]*externalApplicationContainer
	resources    map[string]*resourceContainer
}

func (lb *Builder) parseField(inputField *spec.Field, parent *oType, parentID string, position hybrids.FieldNumber) (fc *fieldContainer, err error) {
	fc = &fieldContainer{}
	fc.name = inputField.Name
	fc.position = position
	fc.id = parentID + fmt.Sprintf("/Field/%d", position)
	fc.app = lb.app
	fc.parent = parent

	switch reflect.ToLower(inputField.Type) {
	case "vector":
		fc.items = &itemsContainer{}
		fc.items.value, fc.items.hType, err = lb.parseFieldType(strings.ToLower(inputField.Items))
		if err != nil {
			return
		}
		switch fc.items.hType {
		case hybrids.Boolean:
			fc.hybridsType = hybrids.VectorBoolean
		case hybrids.Int8:
			fc.hybridsType = hybrids.VectorInt8
		case hybrids.Uint8:
			fc.hybridsType = hybrids.VectorUint8
		case hybrids.Uint16:
			fc.hybridsType = hybrids.VectorUint16
		case hybrids.Int32:
			fc.hybridsType = hybrids.VectorInt32
		case hybrids.Uint32:
			fc.hybridsType = hybrids.VectorUint32
		case hybrids.Int64:
			fc.hybridsType = hybrids.VectorInt64
		case hybrids.Float32:
			fc.hybridsType = hybrids.VectorFloat32
		case hybrids.Float64:
			fc.hybridsType = hybrids.VectorFloat64
		case hybrids.Struct:
			fc.hybridsType = hybrids.VectorStruct
		case hybrids.String:
			fc.hybridsType = hybrids.VectorString
		case hybrids.Byte:
			fc.hybridsType = hybrids.VectorByte
		case hybrids.ResourceID:
			fc.hybridsType = hybrids.VectorResourceID
		case hybrids.Table:
			fc.hybridsType = hybrids.VectorTable
		case hybrids.Union:
			fc.hybridsType = hybrids.VectorUnion
		}

	default:
		fc.value, fc.hybridsType, err = lb.parseFieldType(strings.ToLower(inputField.Type))
		if err != nil {
			return
		}
	}

	return
}

func (lb *Builder) parseFieldType(str string) (o *oType, h hybrids.Types, err error) {
	switch str {
	case "boolean":
		h = hybrids.Boolean
	case "int8":
		h = hybrids.Int8
	case "uint8":
		h = hybrids.Uint8
	case "int16":
		h = hybrids.Int16
	case "uint16":
		h = hybrids.Uint16
	case "int32":
		h = hybrids.Int32
	case "uint32":
		h = hybrids.Uint32
	case "int64":
		h = hybrids.Int64
	case "uint64":
		h = hybrids.Uint64
	case "float32":
		h = hybrids.Float32
	case "float64":
		h = hybrids.Float64
	case "string":
		h = hybrids.String
	case "byte":
		h = hybrids.Byte
	default:
		o, h, err = lb.parseType(str)
	}
	return
}

func (lb *Builder) parseType(str string) (o *oType, h hybrids.Types, err error) {
	var ok bool
	o, ok = lb.store[str]

	if ok {
		switch o.Kind() {
		case reflect.Table:
			h = hybrids.Table
		case reflect.Struct:
			h = hybrids.Struct
		case reflect.Union:
			switch o.union.itemsKind {
			case reflect.UnionOfTables:
				h = hybrids.Union
			case reflect.UnionOfResources:
				h = hybrids.ResourceID
			case reflect.UnionOfExternalResources:
				h = hybrids.Byte
			default:
				err = fmt.Errorf("union of type %d not recognized", o.union.itemsKind)
			}
		case reflect.Enumeration:
			h = o.enum.hType
		case reflect.Resource:
			h = hybrids.ResourceID
		case reflect.ExternalResource:
			h = hybrids.Byte
		default:
			err = fmt.Errorf("otype of type %s not recognized", o.kind)
		}
		return
	}

	if !strings.Contains(str, ".") {

		bType, ok := lb.files_types[str]
		if !ok {
			err = fmt.Errorf("type definition for %s, not found", str)
			return
		}
		itf := lb.files[str]

		switch  bType {
		case "struct":
			is := itf.(*spec.Struct)
			_, o, err = lb.parseStruct(is)
			h = hybrids.Struct
			return
		case "table":
			it := itf.(*spec.Table)
			_, o, err = lb.parseTable(it, false)
			h = hybrids.Table
			return
		case "union":
			ut := itf.(*spec.Union)
			_, o, err = lb.parseUnion(ut)
			h = hybrids.Union
			return
		case "enumeration":
			et := itf.(*spec.Enumeration)
			_, o, err = lb.parseEnumeration(et)
			if err != nil {
				return
			}
			h = o.enum.hType
			return
		default:
			err = fmt.Errorf("definition of type %s, not supported", bType)
			return
		}

	} else {
		h = hybrids.Byte
		//is external resource
		result := strings.Split(str, ".")
		ea, ok := lb.app.externalAppMap[reflect.ToLower(result[0])]
		if !ok {
			err = fmt.Errorf("application with alias %s not found", result[0])
			return
		}
		er := &externalResourceContainer{}
		er.name = result[1]
		er.app = ea
		er.id = ea.path + "/Resource/" + er.name
		ea.usedResources = append(ea.usedResources, er)
		oer := &oType{}
		oer.kind = reflect.ExternalResource
		oer.er = er
		lb.store[str] = oer

	}
	return
}

func (lb *Builder) parseStruct(inputStruct *spec.Struct) (s *structContainer, o *oType, err error) {
	var fc *fieldContainer
	s = &structContainer{}
	s.name = inputStruct.Metadata.Name
	s.id = lb.app.path + "/Struct/" + s.name
	s.app = lb.app
	o = &oType{}
	o.kind = reflect.Struct
	o.str = s
	s.fieldIndex = make([]*fieldContainer, 0, len(inputStruct.Fields))
	s.fieldMap = make(map[string]*fieldContainer)

	lb.store[reflect.ToLower(inputStruct.Metadata.Name)] = o

	for pos, field := range inputStruct.Fields {
		fc, err = lb.parseField(field, o, s.id, hybrids.FieldNumber(pos))
		if err != nil {
			return
		}
		if !fc.hybridsType.IsScalar() {
			err = fmt.Errorf("field with id %s and name %s is not a scalar", fc.id, fc.name)
			return
		}
		s.fieldIndex = append(s.fieldIndex, fc)
		s.fieldMap[reflect.ToLower(fc.name)] = fc
	}
	return
}

func (lb *Builder) parseTable(inputTable *spec.Table, asResource bool) (t *tableContainer, o *oType, err error) {
	var fc *fieldContainer
	t = &tableContainer{}
	t.name = inputTable.Metadata.Name
	t.id = lb.app.path + "/Table/" + t.name
	t.app = lb.app
	o = &oType{}
	o.kind = reflect.Table
	o.table = t
	t.fieldIndex = make([]*fieldContainer, 0, len(inputTable.Fields))
	t.fieldMap = make(map[string]*fieldContainer)

	if !asResource {
		lb.store[reflect.ToLower(inputTable.Metadata.Name)] = o
	}

	for pos, field := range inputTable.Fields {
		fc, err = lb.parseField(field, o, t.id, hybrids.FieldNumber(pos))
		if err != nil {
			return
		}
		t.fieldIndex = append(t.fieldIndex, fc)
		t.fieldMap[reflect.ToLower(fc.name)] = fc
	}
	return
}

func (lb *Builder) parseEnumeration(inputEnum *spec.Enumeration) (e *enumerationContainer, o *oType, err error) {
	e = &enumerationContainer{}
	e.app = lb.app
	e.name = inputEnum.Metadata.Name
	e.id = lb.app.path + "/Enumeration/" + e.name
	e.stringMap = map[string]uint16{}
	e.stringIndex = make([]string, 0, len(inputEnum.Items)+1)
	e.stringIndex = append(e.stringIndex, "None")
	e.stringMap["none"] = 0

	o = &oType{}
	o.kind = reflect.Enumeration
	o.enum = e
	lb.store[reflect.ToLower(inputEnum.Metadata.Name)] = o

	switch reflect.ToLower(inputEnum.Kind) {
	case "uint8":
		e.hType = hybrids.Uint8
		if len(inputEnum.Items) > hybrids.MaxUint8 {
			err = fmt.Errorf("enumeration of type %s only supports %d items", inputEnum.Kind, hybrids.MaxUint8)
			return
		}
	case "uint16":
		e.hType = hybrids.Uint16
		if len(inputEnum.Items) > hybrids.MaxUint16 {
			err = fmt.Errorf("enumeration of type %s only supports %d items", inputEnum.Kind, hybrids.MaxUint16)
			return
		}
	default:
		err = fmt.Errorf("enumeration of type %s not supported", inputEnum.Kind)
		return
	}

	for position, item := range inputEnum.Items {
		e.stringIndex = append(e.stringIndex, item.Name)
		e.stringMap[reflect.ToLower(item.Name)] = uint16(position + 1)
	}
	return
}

func (lb *Builder) parseUnion(inputUnion *spec.Union) (u *unionContainer, o *oType, err error) {
	u = &unionContainer{}
	u.app = lb.app
	u.name = inputUnion.Metadata.Name
	u.id = lb.app.path + "/Union/" + u.name
	o = &oType{}
	o.kind = reflect.Union
	o.union = u

	lb.store[reflect.ToLower(inputUnion.Metadata.Name)] = o

	switch reflect.ToLower(inputUnion.ItemsKind) {
	case "table":
		u.itemsKind = reflect.UnionOfTables
	case "resource":
		u.itemsKind = reflect.UnionOfResources
	case "externalresource":
		u.itemsKind = reflect.UnionOfExternalResources
	default:
		err = fmt.Errorf("union items of type %s don't found ", inputUnion.ItemsKind)
		return
	}
	u.fieldIndex = make([]*fieldContainer, 0, len(inputUnion.Items))
	u.fieldMap = map[string]*fieldContainer{}

	for position, item := range inputUnion.Items {
		//create field
		f := &fieldContainer{}
		f.parent = o
		f.id = u.id + fmt.Sprintf("/Field/%d", position)
		f.position = hybrids.FieldNumber(position)
		f.app = lb.app
		f.value, f.hybridsType, err = lb.parseType(reflect.ToLower(item.Type))
		if err != nil {
			return
		}
		switch u.itemsKind {
		case reflect.UnionOfTables:
			if f.value.kind != reflect.Table {
				err = fmt.Errorf("union of tables only accepts tables as items, instead I found %s", f.value.kind.String())
				return
			}
		case reflect.UnionOfResources:
			if f.value.kind != reflect.Resource {
				err = fmt.Errorf("union of resources only accepts resources as items, instead I found %s", f.value.kind.String())
				return
			}
		case reflect.UnionOfExternalResources:
			if f.value.kind != reflect.ExternalResource {
				err = fmt.Errorf("union of external resources only accepts external resources as items, instead I found %s", f.value.kind.String())
				return
			}
		}

		if item.FieldName != "" {
			f.name = item.FieldName
		} else {
			switch u.itemsKind {
			case reflect.UnionOfTables:
				f.name = f.value.table.name
			case reflect.UnionOfResources:
				f.name = f.value.res.name
			case reflect.UnionOfExternalResources:
				f.name = f.value.er.name
			}
		}

		u.fieldIndex = append(u.fieldIndex, f)
		u.fieldMap[reflect.ToLower(f.name)] = f
	}
	return
}
