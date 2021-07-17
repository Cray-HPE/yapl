package model

type Pipeline struct {
	Kind     string   `yaml:"kind,omitempty"`
	Metadata Metadata `yaml:"metadata,omitempty"`
	Steps    []string `yaml:"steps,omitempty"`
}
type StepsMap map[string]*string

func NewPipeline() *Pipeline {
	return &Pipeline{}
}
