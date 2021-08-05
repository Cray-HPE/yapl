package main

import (
	"log"
	"os"

	"github.com/Cray-HPE/yapl/util"
	"github.com/fatih/color"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

var version string

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = version
	app.Name = "yapl"
	app.Usage = "Yet another pipeline"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "file,f",
			Value:  "./pipeline.yaml",
			Usage:  "Pipeline file to read from",
			EnvVar: "PIPELINE_FILE",
		},
		cli.StringFlag{
			Name:   "vars",
			Usage:  "json/yaml file containing variables for template",
			EnvVar: "GOSS_VARS",
		},
		cli.BoolFlag{
			Name: "no-color",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "validate",
			Aliases: []string{"r"},
			Usage:   "validate yapl pipeline definition after imports",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "debug, d",
					Usage: "Print debugging info when validating",
				},
			},
			Action: func(c *cli.Context) error {
				_, err := util.RenderPipeline(newRuntimeConfigFromCLI(c))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "execute",
			Usage: "execute yapl after imports",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "debug, d",
					Usage: "Print debugging info when executing",
				},
				cli.BoolFlag{
					Name:   "no-cache",
					Usage:  "ignore cached execution status",
					EnvVar: "NO_CACHE",
				},
			},
			Action: func(c *cli.Context) error {
				err := util.ExecutePipeline(newRuntimeConfigFromCLI(c))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "doc",
			Usage: "generate doc after imports",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "debug, d",
					Usage: "Print debugging info when generating doc",
				},
				cli.StringFlag{
					Name:  "output-dir",
					Usage: "Output directory of docs",
					Value: "dist",
				},
			},
			Action: func(c *cli.Context) error {
				err := util.DocGenFromPipeline(newRuntimeConfigFromCLI(c))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "server",
			Aliases: []string{"r"},
			Usage:   "web representation of yapl pipeline after imports",
			Action: func(c *cli.Context) error {
				err := util.Serve(newRuntimeConfigFromCLI(c))
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// converts a cli context into a goss Config
func newRuntimeConfigFromCLI(c *cli.Context) *util.Config {
	cfg := &util.Config{
		File:      c.GlobalString("file"),
		Debug:     c.Bool("debug"),
		NoColor:   c.GlobalBool("no-color"),
		Vars:      c.GlobalString("vars"),
		OutputDir: c.String("output-dir"),
		NoCache:   c.Bool("no-cache"),
	}

	if cfg.NoColor {
		color.NoColor = true
		pterm.DisableColor()
	}
	if cfg.Debug {
		pterm.EnableDebugMessages()
	} else {
		pterm.DisableDebugMessages()
	}
	if cfg.NoCache {
		util.ClearCache()
	}

	return cfg
}
