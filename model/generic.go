package model

import "gopkg.in/yaml.v2"

type GenericYAML struct {
	Kind     string                 `yaml:"kind,omitempty"`
	Metadata Metadata               `yaml:"metadata,omitempty"`
	Spec     map[string]interface{} `yaml:"spec,omitempty"`
}

func NewGenericYAML() *GenericYAML {
	return &GenericYAML{}
}

func (genericYAML *GenericYAML) ToPipeline() Pipeline {
	genericYAMLBytes, _ := yaml.Marshal(genericYAML)
	pipeline := NewPipeline()
	yaml.Unmarshal(genericYAMLBytes, pipeline)
	return *pipeline
}

func (genericYAML *GenericYAML) ToStep() Step {
	genericYAMLBytes, _ := yaml.Marshal(genericYAML)
	step := NewStep()
	yaml.Unmarshal(genericYAMLBytes, step)
	return *step
}
