package model

type Metadata struct {
	Name        string   `yaml:"name,omitempty" json:"name"`
	Description string   `yaml:"description,omitempty"  json:"description"`
	RenderedId  int      `json:"renderedId"`
	ParentId    string   `json:"parentId"`
	Id          string   `json:"id"`
	ChildrenIds []string `json:"childIds"`
	Status      string   `json:"status"`
}
