package model

type Step struct {
	Kind     string   `yaml:"kind,omitempty"`
	Metadata Metadata `yaml:"metadata,omitempty"`
	Jobs     []Job    `yaml:"jobs,omitempty"`
}

type Job struct {
	PreCondition  Runnable `yaml:"preCondition,omitempty"`
	Action        Runnable `yaml:"action,omitempty"`
	ErrorHandling Runnable `yaml:"errorHandling,omitempty"`
}

type Runnable struct {
	Description string `yaml:"description,omitempty"`
	Command     string `yaml:"command,omitempty"`
}

func NewStep() *Step {
	return &Step{}
}
