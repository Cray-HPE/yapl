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
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	Blue.Printf("==============> Command output: (%s)\n", prefix)
	err := command.Run()
	if err != nil {
		return err
	}
	Green.Println("Done")
	return nil
}

func executePipeline(pipeline model.GenericYAML) error {
	Blue.Printf("==> Pipeline: %s \n", pipeline.Metadata.Name)
	if debug {
		Yellow.Println(Indent(pipeline.Metadata.Description, "    "))
	}
	return nil
}

func executeStep(pipeline model.GenericYAML) error {
	Blue.Printf("======> Step: %s \n", pipeline.Metadata.Name)
	if debug {
		Yellow.Println(Indent(pipeline.Metadata.Description, "        "))
	}
	step := pipeline.ToStep()
	for index, job := range step.Spec.Jobs {
		Blue.Printf("==========> job: %d (Pre-condition)\n", index)
		if debug {
			Yellow.Println(Indent(job.PreCondition.Description, "                "))
		}
		err := runCommand(job.PreCondition.Command, "Precondition")
		if err != nil {
			Red.Println("ERROR: Pre condition failed, stop pipeline")
			return err
		}

		if debug {
			Yellow.Println(Indent(job.Action.Description, "                "))
		}
		err = runCommand(job.Action.Command, "Action")
		if err != nil {
			Red.Println("ERROR: Action failed!!!")
			fmt.Println(MarkdownToText(job.ErrorHandling.Description))
			return runCommand(job.ErrorHandling.Command, "Error Handling")
		}

	}
	return nil
}
