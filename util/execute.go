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
var progressBar *pterm.ProgressbarPrinter

func ExecutePipeline(cfg *Config) error {
	debug = cfg.Debug
	counter, rootId, err := RenderPipeline(cfg)
	if err != nil {
		return err
	}
	progressBar, _ = pterm.DefaultProgressbar.WithTotal(counter).Start()

	err = executePipeline(rootId)
	//progressBar.Stop()
	return err

}

func executePipeline(pipelineId string) error {
	pipeline, err := PopFromCache(pipelineId)
	if err != nil {
		return err
	}

	if pipeline.Kind == "step" {
		progressBar.UpdateTitle("Running: " + pipeline.Metadata.Name)
		err := executeStep(&pipeline)
		progressBar.UpdateTitle("Done: " + pipeline.Metadata.Name)
		if err != nil {
			pterm.Info.Printf("Failed Pipeline/Step id: %s\n", pipeline.Metadata.Id)
			return err
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
	if progressBar != nil {
		progressBar.Increment()
	}
	step, _ := pipeline.ToStep()

	for _, job := range step.Spec.Jobs {
		fmt.Println()
		pterm.Debug.Println(MarkdownToText(pipeline.Metadata.Description))

		err := execute(job.PreCondition, "Step: "+pipeline.Metadata.Name+" --- Checking Precondition")
		PushToCache(step.ToGeneric()) //nolint
		if err != nil {
			return err
		}

		err = execute(job.Action, "Step: "+pipeline.Metadata.Name+" --- Executing Action")
		PushToCache(step.ToGeneric()) //nolint
		if err != nil {
			return err
		}

		err = execute(job.PostValidation, "Step: "+pipeline.Metadata.Name+" --- Post action validation")
		PushToCache(step.ToGeneric()) //nolint
		if err != nil {
			return err
		}
	}
	pipeline.Metadata.Completed = true
	PushToCache(step.ToGeneric()) //nolint
	return nil
}

func execute(runnable *model.Runnable, name string) error {

	var outputBuf bytes.Buffer
	var output io.Writer
	if os.Getenv("ConsoleOutput") == "true" {
		output = io.MultiWriter(os.Stdout, &outputBuf)
	} else {
		output = bufio.NewWriter(&outputBuf)
	}
	pterm.Debug.Println(MarkdownToText(runnable.Description))

	err := runCommand(runnable.Command, output)
	runnable.Output = outputBuf.String()
	writeConsoleOutputTofile(outputBuf, name)
	if err != nil {
		pterm.Error.Println(name)
		if os.Getenv("ConsoleOutput") != "true" {
			fmt.Println(runnable.Output)
		}
		pterm.FgRed.Println("Check the doc below for troubleshooting:")
		fmt.Println(MarkdownToText(runnable.Troubleshooting))
		return err
	}
	pterm.Success.Println(name)
	return nil
}

func writeConsoleOutputTofile(outputBuf bytes.Buffer, name string) error {
	// write console log to file
	f, err := os.OpenFile("yapl.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString("==> " + name + "\n" + outputBuf.String() + "\n\n"); err != nil {
		return err
	}
	return nil
}

func runCommand(cmd string, output io.Writer) error {
	command := exec.Command("sh", "-ceu", cmd)
	if debug {
		command = exec.Command("sh", "-cxeu", cmd)
	}
	command.Stdin = os.Stdin
	command.Stdout = output
	command.Stderr = output

	return command.Run()
}
