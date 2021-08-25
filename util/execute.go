package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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

	executePipeline(rootId)

	return nil
}

func executePipeline(pipelineId string) error {
	pipeline, err := PopFromCache(pipelineId)
	if err != nil {
		return err
	}

	if pipeline.Kind == "step" {
		err := executeStep(pipeline)
		if err != nil {
			pterm.Info.Printf("Failed Pipeline/Step id: %s\n", pipeline.Metadata.Id)
		}
	} else {
		pterm.DefaultHeader.Printf("Pipeline: %s \n", pipeline.Metadata.Name)
		pterm.Debug.Println(MarkdownToText(pipeline.Metadata.Description))
		for _, chilePipelineId := range pipeline.Metadata.ChildrenIds {
			executePipeline(chilePipelineId)
		}
		pipeline.Metadata.Completed = true
		PushToCache(pipeline)
	}
	return nil
}

func executeStep(pipeline model.GenericYAML) error {
	step, _ := pipeline.ToStep()
	for _, job := range step.Spec.Jobs {
		fmt.Println()
		pterm.Debug.Println(MarkdownToText(pipeline.Metadata.Description))

		err := execute(job.PreCondition, "Step: "+pipeline.Metadata.Name+" --- Checking Precondition")
		if err != nil {
			return err
		}

		err = execute(job.Action, "Step: "+pipeline.Metadata.Name+" --- Executing Action")
		if err != nil {
			return err
		}

		err = execute(job.PostValidation, "Step: "+pipeline.Metadata.Name+" --- Post action validation")
		if err != nil {
			return err
		}
	}
	pipeline.Metadata.Completed = true
	PushToCache(pipeline)
	return nil
}

func execute(runnable model.Runnable, name string) error {

	var stdoutBuf, stderrBuf bytes.Buffer
	stdout := bufio.NewWriter(&stdoutBuf)
	stderr := bufio.NewWriter(&stderrBuf)
	spinner, _ := pterm.DefaultSpinner.WithShowTimer(true).Start(name)

	pterm.Debug.Println(MarkdownToText(runnable.Description))

	err := runCommand(runnable.Command, stdout, stderr)
	if err != nil {
		spinner.Fail()
		fmt.Println(stderrBuf.String())
		pterm.FgRed.Println("Check the doc below for troubleshooting:")
		fmt.Println(MarkdownToText(runnable.Troubleshooting))
		return err
	}
	spinner.Success()

	return nil
}

func runCommand(cmd string, stdout io.Writer, stderr io.Writer) error {
	command := exec.Command("sh", "-c", cmd)
	if debug {
		command = exec.Command("sh", "-cx", cmd)
	}
	command.Stdin = nil
	command.Stdout = stdout
	command.Stderr = stderr

	return command.Run()
}
