package util

import (
	"fmt"

	"github.com/dzou-hpe/yapl/model"
)

func DocGenFromPipeline(cfg *Config) error {
	renderedPipeline, err := RenderPipeline(cfg)
	if err != nil {
		return err
	}

	for _, pipeline := range renderedPipeline {
		if pipeline.Kind == "pipeline" {
			Blue.Printf("Pipeline - %s\n", pipeline.Metadata.Name)
			fmt.Println(MarkdownToText(pipeline.Metadata.Description))
			continue
		}
		if pipeline.Kind == "step" {
			Blue.Printf("  Step - %s\n", pipeline.Metadata.Name)
			err := docGenFromStep(pipeline)
			if err != nil {
				return err
			}
			continue
		}
	}
	return nil
}

func docGenFromStep(pipeline model.GenericYAML) error {
	step := pipeline.ToStep()
	for _, job := range step.Spec.Jobs {
		Yellow.Println("    Pre condition")
		fmt.Println(MarkdownToText(job.PreCondition.Description))
		Yellow.Println("    Action")
		fmt.Println(MarkdownToText(job.Action.Description))
		Yellow.Println("    Error Handling")
		fmt.Println(MarkdownToText(job.ErrorHandling.Description))
	}
	return nil
}
