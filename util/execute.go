package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/Cray-HPE/yapl/model"
	"github.com/pterm/pterm"
)

var debug bool

func ExecutePipeline(cfg *Config) error {
	debug = cfg.Debug
	_, rootId, err := RenderPipeline(cfg)
	if err != nil {
		return err
	}

	return executePipeline(rootId)

}

func executePipeline(pipelineId string) error {
	pipeline, err := PopFromCache(pipelineId)
	if err != nil {
		return err
	}

	if pipeline.Kind == "step" {
		err := executeStep(&pipeline)
		if err != nil {
			pterm.Info.Printf("Failed Pipeline/Step id: %s\n", pipeline.Metadata.Id)
		}
	} else {
		pterm.DefaultHeader.Printf("Pipeline: %s \n", pipeline.Metadata.Name)
		pterm.Debug.Println(MarkdownToText(pipeline.Metadata.Description))
		for _, chilePipelineId := range pipeline.Metadata.ChildrenIds {
			if err := executePipeline(chilePipelineId); err != nil {
				return err
			}
		}
		pipeline.Metadata.Completed = true
		if err := PushToCache(pipeline); err != nil {
			return err
		}
	}
	return nil
}

func executeStep(pipeline *model.GenericYAML) error {
	step, _ := pipeline.ToStep()
	for _, job := range step.Spec.Jobs {
		fmt.Println()
		pterm.Debug.Println(MarkdownToText(pipeline.Metadata.Description))

		err := execute(job.PreCondition, "Step: "+pipeline.Metadata.Name+" --- Checking Precondition")
		PushToCache(step.ToGeneric())
		if err != nil {
			return err
		}

		err = execute(job.Action, "Step: "+pipeline.Metadata.Name+" --- Executing Action")
		PushToCache(step.ToGeneric())
		if err != nil {
			return err
		}

		err = execute(job.PostValidation, "Step: "+pipeline.Metadata.Name+" --- Post action validation")
		PushToCache(step.ToGeneric())
		if err != nil {
			return err
		}
	}
	pipeline.Metadata.Completed = true
	PushToCache(step.ToGeneric())
	return nil
}

func execute(runnable *model.Runnable, name string) error {

	var outputBuf bytes.Buffer
	output := bufio.NewWriter(&outputBuf)
	spinner, _ := pterm.DefaultSpinner.WithShowTimer(true).Start(name)

	pterm.Debug.Println(MarkdownToText(runnable.Description))

	err := runCommand(runnable.Command, output)
	runnable.Output = outputBuf.String()
	if err != nil {
		spinner.Fail()
		fmt.Println(runnable.Output)
		pterm.FgRed.Println("Check the doc below for troubleshooting:")
		fmt.Println(MarkdownToText(runnable.Troubleshooting))
		return err
	}
	spinner.Success()
	if os.Getenv("ConsoleOutput") == "true" {
		fmt.Println(runnable.Output)
	}
	return nil
}

func runCommand(cmd string, output io.Writer) error {
	command := exec.Command("sh", "-c", cmd)
	if debug {
		command = exec.Command("sh", "-cx", cmd)
	}
	command.Stdin = nil
	command.Stdout = output
	command.Stderr = output

	return command.Run()
}
