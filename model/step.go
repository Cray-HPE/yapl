package model

import "gopkg.in/yaml.v2"

type Step struct {
	Kind     string    `yaml:"kind,omitempty" validate:"required"`
	Metadata *Metadata `yaml:"metadata,omitempty" validate:"required,dive"`
	Spec     *JobsMap  `yaml:"spec,omitempty" validate:"required,dive"`
}

type JobsMap struct {
	Jobs []Job `yaml:"jobs,omitempty" validate:"required,dive"`
}
type Job struct {
	PreCondition   *Runnable `yaml:"preCondition,omitempty" validate:"required,dive"`
	Action         *Runnable `yaml:"action,omitempty" validate:"required,dive"`
	PostValidation *Runnable `yaml:"postValidation,omitempty" validate:"required,dive"`
}

type Runnable struct {
	Description     string `yaml:"description,omitempty" validate:"required"`
	Command         string `yaml:"command,omitempty"`
	Troubleshooting string `yaml:"troubleshooting,omitempty" validate:"required"`
	Output          string
}

func NewStep() *Step {
	return &Step{}
}

func (stepYaml *Step) ToGeneric() GenericYAML {
	genericYAMLBytes, _ := yaml.Marshal(stepYaml)
	genericYaml := NewGenericYAML()
	err := yaml.Unmarshal(genericYAMLBytes, genericYaml)
	if err != nil {
		return *NewGenericYAML()
	}
	return *genericYaml
}
