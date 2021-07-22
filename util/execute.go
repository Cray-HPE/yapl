package util

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dzou-hpe/yapl/model"
	"github.com/pterm/pterm"
)

var debug bool

func ExecutePipeline(cfg *Config) error {
	debug = cfg.Debug
	renderedPipeline, err := RenderPipeline(cfg)
	if err != nil {
		return err
	}

	for _, pipeline := range renderedPipeline {
		if pipeline.Kind == "pipeline" {
			executePipeline(pipeline)
			continue
		}
		if pipeline.Kind == "step" {
			err := executeStep(pipeline)
			if err != nil {
				return err
			}
			continue
		}
	}

	return nil
}

func runCommand(cmd string) error {
	command := exec.Command("sh", "-c", cmd)
	if debug {
		command = exec.Command("sh", "-cx", cmd)
	}
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()
}

func executePipeline(pipeline model.GenericYAML) error {
	pterm.FgBlue.Printf("==> Pipeline: %s \n", pipeline.Metadata.Name)
	if debug {
		fmt.Println(MarkdownToText(pipeline.Metadata.Description))
	}
	return nil
}

func executeStep(pipeline model.GenericYAML) error {
	// Print a section with level two.
	pterm.FgBlue.Printf("======> Step: %s \n", pipeline.Metadata.Name)
	if debug {
		fmt.Println(MarkdownToText(pipeline.Metadata.Description))
	}
	step := pipeline.ToStep()
	for _, job := range step.Spec.Jobs {
		pterm.FgBlue.Println("==========> job:")
		pterm.FgYellow.Println("==============> Precondition")
		if debug {
			fmt.Println(MarkdownToText(job.PreCondition.Description))
		}
		err := runCommand(job.PreCondition.Command)
		if err != nil {
			pterm.FgRed.Println("ERROR: Pre condition failed, stop pipeline")
			return err
		}
		pterm.FgGreen.Println("==============> Precondition: Done")

		pterm.FgYellow.Println("==============> Action")
		if debug {
			fmt.Println(MarkdownToText(job.Action.Description))
		}
		err = runCommand(job.Action.Command)
		if err != nil {
			pterm.FgBlue.Println("==============> Error Handling")
			runCommand(job.ErrorHandling.Command)
			pterm.FgGreen.Println("==============> Error Handling: Done")
			pterm.FgRed.Println("==============> Action: ERROR, Action failed!!! Error handling has been executed")
			fmt.Println()
			pterm.FgRed.Println("Check the doc below for troubleshooting:")
			fmt.Println(MarkdownToText(job.ErrorHandling.Description))
			return err
		}
		pterm.FgGreen.Println("==============> Action: Done")

	}
	return nil
}
