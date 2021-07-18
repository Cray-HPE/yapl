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

func runCommand(cmd string, prefix string) error {
	command := exec.Command("sh", "-c", cmd)
	if debug {
		command = exec.Command("sh", "-cx", cmd)
	}
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	Blue.Printf("==============> Command output: (%s) .....\n", prefix)
	err := command.Run()
	Blue.Printf("==============> Command output: (%s) - ", prefix)
	if err != nil {
		Red.Println("Error")
		return err
	}
	Green.Println("Done")
	return nil
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
		if debug {
			Blue.Println("==============> Command (Precondition):")
			fmt.Println(MarkdownToText(job.PreCondition.Description))
		}
		err := runCommand(job.PreCondition.Command, "Precondition")
		if err != nil {
			Red.Println("ERROR: Pre condition failed, stop pipeline")
			return err
		}

		if debug {
			Blue.Println("==============> Command (Action):")
			fmt.Println(MarkdownToText(job.Action.Description))
		}
		err = runCommand(job.Action.Command, "Action")
		if err != nil {
			err := runCommand(job.ErrorHandling.Command, "Error Handling")
			Red.Println("ERROR: Action failed!!!")
			fmt.Println(MarkdownToText(job.ErrorHandling.Description))
			return err
		}

	}
	return nil
}
