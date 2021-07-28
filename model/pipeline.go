package model

type Pipeline struct {
	Kind     string   `yaml:"kind,omitempty" validate:"required"`
	Metadata Metadata `yaml:"metadata,omitempty" validate:"required"`
	Spec     StepsMap `yaml:"spec,omitempty" validate:"required"`
}
type StepsMap struct {
	Steps []string `yaml:"steps,omitempty" validate:"required"`
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}
