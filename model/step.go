package model

type Step struct {
	Kind     string   `yaml:"kind,omitempty"`
	Metadata Metadata `yaml:"metadata,omitempty"`
	Spec     JobsMap  `yaml:"spec,omitempty"`
}

type JobsMap struct {
	Jobs []Job `yaml:"jobs,omitempty"`
}
type Job struct {
	Name           string   `yaml:"name"`
	TargetHost     string   `yaml:"targetHost,omitempty"`
	PreCondition   Runnable `yaml:"preCondition,omitempty"`
	Action         Runnable `yaml:"action,omitempty"`
	PostValidation Runnable `yaml:"postValidation,omitempty"`
}

type Runnable struct {
	Description     string `yaml:"description,omitempty"`
	Command         string `yaml:"command,omitempty"`
	Troubleshooting string `yaml:"troubleshooting,omitempty"`
}

func NewStep() *Step {
	return &Step{}
}
