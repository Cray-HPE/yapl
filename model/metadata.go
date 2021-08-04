package model

type Metadata struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Parent      int
	Id          string
	OrderId     int
	Children    []int
	Completed   bool
}
