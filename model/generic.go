package model

import "gopkg.in/yaml.v2"

type GenericYAML struct {
	Kind     string                  `yaml:"kind,omitempty"`
	Metadata *Metadata               `yaml:"metadata,omitempty"`
	Spec     *map[string]interface{} `yaml:"spec,omitempty"`
}

func NewGenericYAML() *GenericYAML {
	return &GenericYAML{}
}

func (genericYAML *GenericYAML) ToPipeline() (Pipeline, error) {
	genericYAMLBytes, _ := yaml.Marshal(genericYAML)
	pipeline := NewPipeline()
	err := yaml.Unmarshal(genericYAMLBytes, pipeline)
	return *pipeline, err
}

func (genericYAML *GenericYAML) ToStep() (Step, error) {
	genericYAMLBytes, _ := yaml.Marshal(genericYAML)
	step := NewStep()
	err := yaml.Unmarshal(genericYAMLBytes, step)
	return *step, err
}
