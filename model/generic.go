package model

type GenericYAML map[string]interface{}

func NewGenericYAML() *GenericYAML {
	return &GenericYAML{}
}
