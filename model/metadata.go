package model

type Metadata struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Parent      string `yaml:"parent,omitempty"`
	Id          string
}
