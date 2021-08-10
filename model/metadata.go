package model

type Metadata struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	RenderedId  int
	ParentId    string
	Id          string
	ChildrenIds []string
	Status      string
}
