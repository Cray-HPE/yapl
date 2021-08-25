package model

type Metadata struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Id          string
	ChildrenIds []string
	Completed   bool
}
