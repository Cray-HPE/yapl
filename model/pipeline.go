package model

type Pipeline struct {
	Kind     string   `yaml:"kind,omitempty"`
	Metadata Metadata `yaml:"metadata,omitempty"`
	Spec     StepsMap `yaml:"spec,omitempty"`
}
type StepsMap struct {
	Steps []string `yaml:"steps,omitempty"`
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}
