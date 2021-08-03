package util

import (
	"github.com/Cray-HPE/yapl/model"
	"github.com/pterm/pterm"
)

func DocGenFromPipeline(cfg *Config) error {
	renderedPipeline, err := RenderPipeline(cfg)
	if err != nil {
		return err
	}

	for _, pipeline := range renderedPipeline {
		if pipeline.Kind == "pipeline" {
			pterm.Debug.Println(pterm.DefaultSection.WithLevel(1).WithIndentCharacter("==").Sprintf("Pipeline - %s\n", pipeline.Metadata.Name))
			pterm.Debug.Println(MarkdownToText(pipeline.Metadata.Description))
			continue
		}
		if pipeline.Kind == "step" {
			pterm.Debug.Println(pterm.DefaultSection.WithLevel(2).WithIndentCharacter("==").Sprintf("Step - %s\n", pipeline.Metadata.Name))
			docGenFromStep(pipeline)
			continue
		}
	}
	return nil
}

func docGenFromStep(pipeline model.GenericYAML) {
	step, _ := pipeline.ToStep()
	for _, job := range step.Spec.Jobs {
		pterm.Debug.Println(pterm.DefaultSection.WithLevel(3).WithIndentCharacter("==").Sprint("Pre condition\n"))
		pterm.Debug.Println(MarkdownToText(job.PreCondition.Description))
		pterm.Debug.Println(pterm.DefaultSection.WithLevel(3).WithIndentCharacter("==").Sprint("Action\n"))
		pterm.Debug.Println(MarkdownToText(job.Action.Description))
		pterm.Debug.Println(pterm.DefaultSection.WithLevel(3).WithIndentCharacter("==").Sprint("Post Validation\n"))
		pterm.Debug.Println(MarkdownToText(job.PostValidation.Description))
	}
}
