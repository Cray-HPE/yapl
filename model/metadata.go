package model

type Metadata struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	ParentId    string
	Id          string
	ChildrenIds []string
	Completed   bool
}
