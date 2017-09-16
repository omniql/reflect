package spec

//Metadata definition
type Metadata struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

//Field definition
type Field struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Items string `json:"items"`
}

//Struct definition
type Struct struct {
	Metadata *Metadata `json:"metadata"`
	Fields   []*Field  `json:"fields"`
}

//Enumeration ...
type Enumeration struct {
	Metadata *Metadata          `json:"metadata"`
	Kind     string             `json:"kind"`
	Items    []*EnumerationItem `json:"items"`
}

//EnumerationItem ...
type EnumerationItem struct {
	Name string `json:"name"`
}

//Table definition
type Table struct {
	Metadata *Metadata `json:"metadata"`
	Fields   []*Field  `json:"fields"`
}

//Union definition
type Union struct {
	Metadata  *Metadata    `json:"metadata"`
	ItemsKind string       `json:"items_kind"`
	Items     []UnionItems `json:"items"`
}

//UnionItems definition
type UnionItems struct {
	FieldName string `json:"field_name"`
	Type      string `json:"type"`
}

//Application ...
type Application struct {
	Metadata       *Metadata              `json:"metadata"`
	ResourceIdType string                 `json:"resource_id_type"`
	Imports        []ApplicationReference `json:"imports"`
	Resources      []ResourcesDefinition  `json:"resources"`
	Version        string                 `json:"version"`
	Path           string                 `json:"path"`
}

//ResourcesDefinition ...
type ResourcesDefinition struct {
	Table string `json:"table"`
}

//ApplicationReference
type ApplicationReference struct {
	Alias   string `json:"alias"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

//YamlFile ...
type YamlFile struct {
	Api  string      `json:"api"`
	RID  string      `json:"rid"`
	Spec interface{} `json:"spec"`
}
