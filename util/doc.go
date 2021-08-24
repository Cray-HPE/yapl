package util

import (
	"fmt"
	"os"
	"time"

	"github.com/Cray-HPE/yapl/model"
	"github.com/pterm/pterm"
)

var outputDir string
var p *pterm.ProgressbarPrinter

func DocGenFromPipeline(cfg *Config) error {
	// outputDir = cfg.OutputDir
	// os.MkdirAll(outputDir, os.ModePerm)
	// renderedPipeline, err := RenderPipeline(cfg)
	// if err != nil {
	// 	return err
	// }
	// // Create progressbar as fork from the default progressbar.
	// p, _ = pterm.DefaultProgressbar.WithTotal(len(renderedPipeline)).WithTitle("Generating Documents").Start()

	// for _, pipeline := range renderedPipeline {
	// 	if pipeline.Kind == "pipeline" {
	// 		pterm.Debug.Println(pterm.DefaultSection.WithLevel(1).WithIndentCharacter("==").Sprintf("Pipeline - %s\n", pipeline.Metadata.Name))
	// 		pterm.Debug.Println(MarkdownToText(pipeline.Metadata.Description))
	// 		writeDocToFile(pipeline.Metadata.Name, pipeline.Metadata.Description)
	// 		continue
	// 	}
	// 	if pipeline.Kind == "step" {
	// 		pterm.Debug.Println(pterm.DefaultSection.WithLevel(2).WithIndentCharacter("==").Sprintf("Step - %s\n", pipeline.Metadata.Name))
	// 		docGenFromStep(pipeline)
	// 		continue
	// 	}
	// }
	return nil
}

func docGenFromStep(pipeline model.GenericYAML) {
	step, _ := pipeline.ToStep()
	content := fmt.Sprintf("%s\n", step.Metadata.Description)
	for _, job := range step.Spec.Jobs {
		pterm.Debug.Println(pterm.DefaultSection.WithLevel(3).WithIndentCharacter("==").Sprint("Pre condition\n"))
		pterm.Debug.Println(MarkdownToText(job.PreCondition.Description))
		pterm.Debug.Println(pterm.DefaultSection.WithLevel(3).WithIndentCharacter("==").Sprint("Action\n"))
		pterm.Debug.Println(MarkdownToText(job.Action.Description))
		pterm.Debug.Println(pterm.DefaultSection.WithLevel(3).WithIndentCharacter("==").Sprint("Post Validation\n"))
		pterm.Debug.Println(MarkdownToText(job.PostValidation.Description))
		content += fmt.Sprintf("# Precondtion \n %s \n# Action \n %s \n# Post Validation \n %s \n", job.PreCondition.Description, job.Action.Description, job.PostValidation.Description)
	}
	writeDocToFile(step.Metadata.Name, content)
}

func writeDocToFile(filename string, content string) error {
	p.Title = "Generating: " + filename
	pterm.Success.Println("Generated: " + filename)
	f, err := os.Create(outputDir + "/" + filename + ".md")
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	p.Increment()
	time.Sleep(time.Millisecond * 350)
	return nil
}
