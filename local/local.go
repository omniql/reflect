package local

import (
	"github.com/omniql/reflect/spec"
	"github.com/omniql/reflect"
	"github.com/iancoleman/strcase"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
	"strings"
	"github.com/mitchellh/mapstructure"
)

//LocalBuilder build a reflection object from a local file path
type Builder struct {
	app   *applicationContainer
	store map[string]*oType
	files map[string]interface{}
	externalApps map[string]*externalApplicationContainer
	resources map[string]*resourceContainer

}

func (lb *Builder) parseField(inputField *spec.Field, parent *oType, parentID string, position hybrids.FieldNumber) (fc *fieldContainer, err error) {
	fc = &fieldContainer{}
	fc.name = strcase.ToCamel(inputField.Name)
	fc.position = position
	fc.id = parentID + fmt.Sprintf("Field/%d", position)
	fc.app = lb.app
	fc.parent = parent
	switch reflect.ToLower(inputField.Type) {
	case "vector":

	default:
		fc.value, fc.hybridsType, err = lb.parseFieldType(reflect.ToLower(inputField.Type))
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
	default:
		o, h, err = lb.parseType(str)
	}
	return
}

func (lb *Builder) parseType(str string) (o *oType, h hybrids.Types, err error) {
	var ok bool

	if !strings.Contains(str, ".") {
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
					h = hybrids.Union
				default:
					err = fmt.Errorf("union of type %d not recognized", o.union.itemsKind)
				}
			case reflect.Enumeration:
				h = o.enum.hType
			case reflect.Resource:
				h = hybrids.ResourceID
			case reflect.ExternalResource:
				h = hybrids.Table
			default:
				err = fmt.Errorf("otype of type %s not recognized", o.kind)
			}
			return
		}

		file, ok := lb.files[str]

		if !ok {
			err = fmt.Errorf("type %s not found in application", str)
			return
		}

		meta, ok := file["metadata"].(map[string]string)
		if !ok {
			err = fmt.Errorf("invalid file definition for  %s", str)
			return
		}

		switch  reflect.ToLower(meta["kind"]) {
		case "struct":
			is := &spec.Struct{}
			err = mapstructure.Decode(is, file)
			if err != nil {
				return
			}
			_, o, err = lb.parseStruct(is)
			h = hybrids.Struct
			return
		case "table":
			it := &spec.Table{}
			err = mapstructure.Decode(it, file)
			if err != nil {
				return
			}
			_, o, err = lb.parseTable(it)
			h = hybrids.Table
			return
		case "union":
			ut := &spec.Union{}
			err = mapstructure.Decode(ut, file)
			if err != nil {
				return
			}
			_, o, err = lb.parseUnion(ut)
			h = hybrids.Union
			return
		case "enumeration":
			et := &spec.Enumeration{}
			err = mapstructure.Decode(et, file)
			if err != nil {
				return
			}
			_, o, err = lb.parseEnumeration(et)
			if err != nil {
				h = o.enum.hType
			}
			return
		default:
			err = fmt.Errorf("definition of type %s, not supported", reflect.ToLower(meta["itemsKind"]))
			return
		}

	} else {
		//is external resource

	}
	return
}

func (lb *Builder) parseStruct(inputStruct *spec.Struct) (s *structContainer, o *oType, err error) {
	var fc *fieldContainer
	s = &structContainer{}
	s.name = strcase.ToCamel(inputStruct.Metadata.Name)
	s.id = "Struct/" + s.name
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
		s.fieldIndex[pos] = fc
		s.fieldMap[reflect.ToLower(fc.name)] = fc
	}
	return
}

func (lb *Builder) parseTable(inputTable *spec.Table) (t *tableContainer, o *oType, err error) {
	var fc *fieldContainer
	t = &tableContainer{}
	t.name = strcase.ToCamel(inputTable.Metadata.Name)
	t.id = "Table/" + t.name
	t.app = lb.app
	o = &oType{}
	o.kind = reflect.Table
	o.table = t
	t.fieldIndex = make([]*fieldContainer, 0, len(inputTable.Fields))
	t.fieldMap = make(map[string]*fieldContainer)

	lb.store[reflect.ToLower(inputTable.Metadata.Name)] = o

	for pos, field := range inputTable.Fields {
		fc, err = lb.parseField(field, o, t.id, hybrids.FieldNumber(pos))
		if err != nil {
			return
		}
		t.fieldIndex[pos] = fc
		t.fieldMap[reflect.ToLower(fc.name)] = fc
	}
	return
}

func (lb *Builder) parseEnumeration(inputEnum *spec.Enumeration) (e *enumerationContainer, o *oType, err error) {
	e = &enumerationContainer{}
	e.app = lb.app
	e.name = strcase.ToCamel(inputEnum.Metadata.Name)
	e.id = "Enumeration/" + e.name
	e.stringMap = map[string]uint16{}
	e.stringIndex = make([]string, 0, len(inputEnum.Items))
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
	case "uin16":
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
		e.stringIndex = append(e.stringIndex, strcase.ToCamel(item.Name))
		e.stringMap[reflect.ToLower(item.Name)] = uint16(position)
	}
	return
}

func (lb *Builder) parseUnion(inputUnion *spec.Union) (u *unionContainer, o *oType, err error) {
	u = &unionContainer{}
	u.app = lb.app
	u.name = strcase.ToCamel(inputUnion.Metadata.Name)
	u.id = "Union/" + u.name
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
		f.value, f.hybridsType, err = lb.parseType(item.Type)
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
			f.name = strcase.ToCamel(item.FieldName)
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
