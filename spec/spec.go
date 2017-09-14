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
	Metadata *Metadata          `json:"meta"`
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

//Table definition
type Resource struct {
	Metadata *Metadata `json:"metadata"`
	Fields   []*Field  `json:"fields"`
}

//Table definition
type Union struct {
	Metadata *Metadata `json:"metadata"`
	Fields   []*Field  `json:"fields"`
	Kind     string    `json:"fields"`
	Items    []string  `json:"items"`
}

//Application ...
type Application struct {
	Metadata       *Metadata              `json:"metadata"`
	ResourceIdType string                 `json:"resource_id_type"`
	Imports        []ApplicationReference `json:"imports"`
}

//ApplicationReference
type ApplicationReference struct {
	Alias   string `json:"alias"`
	Path    string `json:"path"`
	Version string `json:"version"`
}
