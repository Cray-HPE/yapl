package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
	step, _ := pipeline.ToStep()
	for _, job := range step.Spec.Jobs {
		fmt.Println()
		pterm.Info.Printf("Step: %s\n", pipeline.Metadata.Name)

		var stdoutBuf, stderrBuf bytes.Buffer
		stdout := bufio.NewWriter(&stdoutBuf)
		stderr := bufio.NewWriter(&stderrBuf)

		preConditionSpinner, _ := pterm.DefaultSpinner.Start("Checking Precondition ...")
		if debug {
			fmt.Println(MarkdownToText(job.PreCondition.Description))
		}
		err := runCommand(job.PreCondition.Command, stdout, stderr)
		if err != nil {
			preConditionSpinner.Fail()
			fmt.Println(stderrBuf.String())
			pterm.FgRed.Println("ERROR: Pre condition failed, stop pipeline")
			fmt.Println(MarkdownToText(job.PreCondition.Troubleshooting))
			return err
		}
		preConditionSpinner.Success()

		actionSpinner, _ := pterm.DefaultSpinner.Start("Executing Action ...")
		if debug {
			fmt.Println(MarkdownToText(job.Action.Description))
		}
		err = runCommand(job.Action.Command, stdout, stderr)
		if err != nil {
			actionSpinner.Fail()
			fmt.Println(stderrBuf.String())
			pterm.FgRed.Println("Check the doc below for troubleshooting:")
			fmt.Println(MarkdownToText(job.Action.Troubleshooting))
			return err
		}
		actionSpinner.Success()

		postValidationSpinner, _ := pterm.DefaultSpinner.Start("Post action validation ...")
		if debug {
			fmt.Println(MarkdownToText(job.PostValidation.Description))
		}
		err = runCommand(job.PostValidation.Command, stdout, stderr)
		if err != nil {
			postValidationSpinner.Fail()
			fmt.Println(stderrBuf.String())
			pterm.FgRed.Println("Check the doc below for troubleshooting:")
			fmt.Println(MarkdownToText(job.PostValidation.Troubleshooting))
			return err
		}
		postValidationSpinner.Success()

	}
	return nil
}
