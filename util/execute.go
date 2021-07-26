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
	command.Stdout = nil //os.Stdout
	command.Stderr = os.Stderr

	return command.Run()
}

func executePipeline(pipeline model.GenericYAML) error {

	pterm.DefaultHeader.Printf("Pipeline: %s \n", pipeline.Metadata.Name)
	if debug {
		fmt.Println(MarkdownToText(pipeline.Metadata.Description))
	}
	return nil
}

func executeStep(pipeline model.GenericYAML) error {
	if debug {
		pterm.Info.Printf("Step: %s \n", pipeline.Metadata.Name)
		fmt.Println(MarkdownToText(pipeline.Metadata.Description))
	}
	step := pipeline.ToStep()
	for _, job := range step.Spec.Jobs {
		pterm.Info.Printf("Step: %s \n  Job: %s\n  Target host:%s\n", pipeline.Metadata.Name, job.Name, job.TargetHost)

		preConditionSpinner, _ := pterm.DefaultSpinner.Start("Checking Precondition ...")
		if debug {
			fmt.Println(MarkdownToText(job.PreCondition.Description))
		}
		err := runCommand(job.PreCondition.Command)
		if err != nil {
			preConditionSpinner.Fail()
			pterm.FgRed.Println("ERROR: Pre condition failed, stop pipeline")
			fmt.Println(MarkdownToText(job.PreCondition.Troubleshooting))
			return err
		}
		preConditionSpinner.Success()

		actionSpinner, _ := pterm.DefaultSpinner.Start("Executing Action ...")
		if debug {
			fmt.Println(MarkdownToText(job.Action.Description))
		}
		err = runCommand(job.Action.Command)
		if err != nil {
			actionSpinner.Fail()
			pterm.FgRed.Println("Check the doc below for troubleshooting:")
			fmt.Println(MarkdownToText(job.Action.Troubleshooting))
			return err
		}
		actionSpinner.Success()

		postValidationSpinner, _ := pterm.DefaultSpinner.Start("Post action validation ...")
		if debug {
			fmt.Println(MarkdownToText(job.PostValidation.Description))
		}
		err = runCommand(job.PostValidation.Command)
		if err != nil {
			postValidationSpinner.Fail()
			pterm.FgRed.Println("Check the doc below for troubleshooting:")
			fmt.Println(MarkdownToText(job.PostValidation.Troubleshooting))
			return err
		}
		postValidationSpinner.Success()

	}
	return nil
}
