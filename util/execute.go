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
	step, _ := pipeline.ToStep()
	for _, job := range step.Spec.Jobs {
		fmt.Println()
		pterm.Info.Printf("Step: %s\n", pipeline.Metadata.Name)
		if debug {
			fmt.Println(MarkdownToText(pipeline.Metadata.Description))
		}
		err := execute(job.PreCondition, "Checking Precondition")
		if err != nil {
			return err
		}

		err = execute(job.Action, "Executing Action")
		if err != nil {
			return err
		}

		err = execute(job.PostValidation, "Post action validation")
		if err != nil {
			return err
		}
	}
	return nil
}

func execute(runnable model.Runnable, name string) error {

	var stdoutBuf, stderrBuf bytes.Buffer
	stdout := bufio.NewWriter(&stdoutBuf)
	stderr := bufio.NewWriter(&stderrBuf)
	spinner, _ := pterm.DefaultSpinner.Start(name + " ...")

	err := runCommand(runnable.Command, stdout, stderr)
	if err != nil {
		spinner.Fail()
		fmt.Println(stderrBuf.String())
		pterm.FgRed.Println("Check the doc below for troubleshooting:")
		fmt.Println(MarkdownToText(runnable.Troubleshooting))
		return err
	}
	spinner.Success()
	if debug {
		fmt.Println(MarkdownToText(runnable.Description))
	}
	return nil
}