package util

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Cray-HPE/yapl/model"
	"github.com/pterm/pterm"
)

var outputDir string
var p *pterm.ProgressbarPrinter

func DocGenFromPipeline(cfg *Config) error {
	outputDir = cfg.OutputDir
	os.MkdirAll(outputDir, os.ModePerm)
	numOfPipelines, rootId, err := RenderPipeline(cfg)
	if err != nil {
		return err
	}
	// Create progressbar as fork from the default progressbar.
	p, _ = pterm.DefaultProgressbar.WithTotal(numOfPipelines).WithTitle("Generating Documents").Start()
	docGenFromPipeline(rootId)
	return nil
}

func docGenFromPipeline(id string) {
	pipeline, _ := PopFromCache(id)
	if pipeline.Kind == "pipeline" {
		writeDocToFile(pipeline.Metadata.Name, pipeline.Metadata.Description)
		var wg sync.WaitGroup
		for _, childId := range pipeline.Metadata.ChildrenIds {
			wg.Add(1)
			go func(id string) {
				docGenFromPipeline(id)
				wg.Done()
			}(childId)
		}
		wg.Wait()
	}
	if pipeline.Kind == "step" {
		docGenFromStep(pipeline)
	}
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
