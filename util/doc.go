package util

import (
	"fmt"

	"github.com/dzou-hpe/yapl/model"
	"github.com/pterm/pterm"
)

func DocGenFromPipeline(cfg *Config) error {
	renderedPipeline, err := RenderPipeline(cfg)
	if err != nil {
		return err
	}

	for _, pipeline := range renderedPipeline {
		if pipeline.Kind == "pipeline" {
			pterm.DefaultSection.WithLevel(1).WithIndentCharacter("==").Printf("Pipeline - %s\n", pipeline.Metadata.Name)
			fmt.Println(MarkdownToText(pipeline.Metadata.Description))
			continue
		}
		if pipeline.Kind == "step" {
			pterm.DefaultSection.WithLevel(2).WithIndentCharacter("==").Printf("Step - %s\n", pipeline.Metadata.Name)
			docGenFromStep(pipeline)
			continue
		}
	}
	return nil
}

func docGenFromStep(pipeline model.GenericYAML) {
	step, _ := pipeline.ToStep()
	for _, job := range step.Spec.Jobs {
		pterm.DefaultSection.WithLevel(3).WithIndentCharacter("==").Println("Pre condition")
		fmt.Println(MarkdownToText(job.PreCondition.Description))
		pterm.DefaultSection.WithLevel(3).WithIndentCharacter("==").Println("Action")
		fmt.Println(MarkdownToText(job.Action.Description))
		pterm.DefaultSection.WithLevel(3).WithIndentCharacter("==").Println("Post Validation")
		fmt.Println(MarkdownToText(job.PostValidation.Description))
	}
}
