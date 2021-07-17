package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dzou-hpe/yapl/util"
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
	}
	app.Commands = []cli.Command{
		{
			Name:    "render",
			Aliases: []string{"r"},
			Usage:   "render yapl after imports",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "debug, d",
					Usage: fmt.Sprintf("Print debugging info when rendering"),
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
					Usage: fmt.Sprintf("Print debugging info when rendering"),
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
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// converts a cli context into a goss Config
func newRuntimeConfigFromCLI(c *cli.Context) *util.Config {
	cfg := &util.Config{
		File:  c.GlobalString("file"),
		Debug: c.Bool("debug"),
	}

	return cfg
}
