package util

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dzou-hpe/yapl/model"
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
	Blue.Printf("==> Pipeline: %s \n", pipeline.Metadata.Name)
	if debug {
		fmt.Println(MarkdownToText(pipeline.Metadata.Description))
	}
	return nil
}

func executeStep(pipeline model.GenericYAML) error {
	Blue.Printf("======> Step: %s \n", pipeline.Metadata.Name)
	if debug {
		fmt.Println(MarkdownToText(pipeline.Metadata.Description))
	}
	step := pipeline.ToStep()
	for _, job := range step.Spec.Jobs {
		Blue.Println("==========> job:")
		Yellow.Println("==============> Precondition")
		if debug {
			fmt.Println(MarkdownToText(job.PreCondition.Description))
		}
		err := runCommand(job.PreCondition.Command)
		if err != nil {
			Red.Println("ERROR: Pre condition failed, stop pipeline")
			return err
		}
		Green.Println("==============> Precondition: Done")

		Yellow.Println("==============> Action")
		if debug {
			fmt.Println(MarkdownToText(job.Action.Description))
		}
		err = runCommand(job.Action.Command)
		if err != nil {
			Blue.Println("==============> Error Handling")
			runCommand(job.ErrorHandling.Command)
			Green.Println("==============> Error Handling: Done")
			Red.Println("==============> Action: ERROR, Action failed!!! Error handling has been executed")
			fmt.Println()
			Red.Println("Check the doc below for troubleshooting:")
			fmt.Println(MarkdownToText(job.ErrorHandling.Description))
			return err
		}
		Green.Println("==============> Action: Done")

	}
	return nil
}
