package model

type Pipeline struct {
	Kind     string   `yaml:"kind,omitempty" validate:"required" json:"kind"`
	Metadata Metadata `yaml:"metadata,omitempty" validate:"required"  json:"metadata"`
	Spec     StepsMap `yaml:"spec,omitempty" validate:"required"  json:"spec"`
}
type StepsMap struct {
	Steps []string `yaml:"steps,omitempty" validate:"required"`
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}
