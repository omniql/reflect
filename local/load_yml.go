package local

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"github.com/mitchellh/mapstructure"
	"github.com/omniql/reflect/spec"
	"github.com/ghodss/yaml"
	"strings"
	"github.com/omniql/reflect"
	"fmt"
	"github.com/nebtex/hybrids/golang/hybrids"
)

func LoadApplication(path string) (app reflect.ApplicationContainer, oErr error) {
	var table *tableContainer
	builder := &Builder{}
	rawApp := &spec.Application{}
	appFound := false
	builder.files = map[string]interface{}{}
	builder.files_types = map[string]string{}

	oErr = filepath.Walk(path, func(subPath string, info os.FileInfo, walkErr error) (wrErr error) {
		var types []string
		var data []byte
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}
		data, wrErr = ioutil.ReadFile(subPath)
		if wrErr != nil {
			return
		}
		yf := &spec.YamlFile{}
		yaml.Unmarshal(data, yf)
		types = strings.Split(yf.RID, "/")

		switch reflect.ToLower(types[len(types)-2]) {
		case "table":
			table := &spec.Table{}
			wrErr = mapstructure.Decode(yf.Spec, table)
			if wrErr != nil {
				return
			}
			builder.files[reflect.ToLower(table.Metadata.Name)] = table
			builder.files_types[reflect.ToLower(table.Metadata.Name)] = "table"
		case "struct":
			str := &spec.Struct{}
			wrErr = mapstructure.Decode(yf.Spec, str)
			if wrErr != nil {
				return
			}
			builder.files[reflect.ToLower(str.Metadata.Name)] = str
			builder.files_types[reflect.ToLower(str.Metadata.Name)] = "struct"

		case "enumeration":
			enum := &spec.Enumeration{}
			wrErr = mapstructure.Decode(yf.Spec, enum)
			if wrErr != nil {
				return
			}
			builder.files[reflect.ToLower(enum.Metadata.Name)] = enum
			builder.files_types[reflect.ToLower(enum.Metadata.Name)] = "enumeration"

		case "union":
			union := &spec.Union{}
			wrErr = mapstructure.Decode(yf.Spec, union)
			if wrErr != nil {
				return
			}
			builder.files[reflect.ToLower(union.Metadata.Name)] = union
			builder.files_types[reflect.ToLower(union.Metadata.Name)] = "union"

		case "application":
			if appFound {
				wrErr = fmt.Errorf("only an application definition file per application is allowed")
				return
			}
			wrErr = mapstructure.Decode(yf.Spec, rawApp)
			if wrErr != nil {
				return
			}
			appFound = true

		}
		return nil

	})
	if oErr != nil {
		return
	}

	appC := &applicationContainer{}
	appC.version = rawApp.Version
	appC.path = rawApp.Path

	switch reflect.ToLower(rawApp.ResourceIdType) {
	case "string":
		appC.rType = hybrids.ResourceIDTypeString
	case "int16":
		appC.rType = hybrids.ResourceIDTypeInt16
	case "uint16":
		appC.rType = hybrids.ResourceIDTypeUint16
	case "int32":
		appC.rType = hybrids.ResourceIDTypeInt32
	case "uint32":
		appC.rType = hybrids.ResourceIDTypeUint32
	case "int64":
		appC.rType = hybrids.ResourceIDTypeInt64
	case "uint64":
		appC.rType = hybrids.ResourceIDTypeUint64
	default:
		oErr = fmt.Errorf("resource id of type %s don't supported", rawApp.ResourceIdType)
		return
	}

	appC.resourceIndex = make([]*resourceContainer, 0, len(rawApp.Resources))
	appC.resourceMap = map[string]*resourceContainer{}

	appC.externalAppIndex = make([]*externalApplicationContainer, 0, len(rawApp.Imports))
	appC.externalAppMap = map[string]*externalApplicationContainer{}

	builder.app = appC
	builder.store = map[string]*oType{}
	builder.externalApps = map[string]*externalApplicationContainer{}
	//load external app
	for _, imp := range rawApp.Imports {
		ea := &externalApplicationContainer{}
		ea.alias = imp.Alias
		ea.version = imp.Version
		ea.path = imp.Path
		builder.externalApps[reflect.ToLower(ea.alias)] = ea
		appC.externalAppIndex = append(appC.externalAppIndex, ea)
		appC.externalAppMap[reflect.ToLower(ea.alias)] = ea

	}

	//load resources [initial]
	for pos, resource := range rawApp.Resources {
		res := &resourceContainer{}
		res.name = resource.Table
		res.id = appC.path + "/" + fmt.Sprintf("Resource/%s", res.name)
		res.app = appC
		res.position = uint16(pos)
		or := &oType{}
		or.res = res
		or.kind = reflect.Resource
		builder.store[reflect.ToLower(resource.Table)] = or
		appC.resourceIndex = append(appC.resourceIndex, res)
		appC.resourceMap[reflect.ToLower(res.name)] = res

	}

	//Add to each table to their resource
	for _, res := range appC.resourceIndex {
		inf, ok := builder.files[reflect.ToLower(res.name)]
		if !ok {
			oErr = fmt.Errorf("table definition not found for resource %s", res.name)
			return
		}
		inputTable, ok := inf.(*spec.Table)
		if !ok {
			oErr = fmt.Errorf("resources only can be created from tables, Resource: %s", res.name)
			return
		}
		table, _, oErr = builder.parseTable(inputTable, true)
		if oErr != nil {
			return
		}
		res.table = table
	}

	app = appC
	return

}
