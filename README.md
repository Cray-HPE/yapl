# yapl
Yet Another Pipeline Language

## Introduction
----
### What is YAPL?

Goss is a Yaml based pipeline orchestration tool for executing/managing pipeline workflows. It is best suited for a complicated installation process that involves many moving parts

### Why YAPL

- You can embed your user manual into your pipeline definition. When things go sideways, you can show rendered markdown doc that tells your customer what to do in the same terminal console.
- You can visualize your pipeline in a bird's view. It helps you to understand where can you optimize your flow.
- You can visualize progress of pipeline executing in a very complicated process
- You can define a reusable `Step` that can be imported anywhere

## Quick Start Example

Check the [example](example/pipelines/demo.yaml) here.  It is a simple pipeline that has 4 steps and also imports another pipeline.

```bash
go run cmd/yapl.go -f example/pipelines/demo.yaml --vars example/vars.yaml execute
```

## Usage
```
NAME:
   yapl - Yet another pipeline

USAGE:
   yapl [global options] command [command options] [arguments...]

COMMANDS:
   render, r  render yapl after imports
   execute    execute yapl after imports
   doc        generate doc after imports
   dep        generate dependency graph after imports
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value  Pipeline file to read from (default: "./pipeline.yaml") [$PIPELINE_FILE]
   --vars value            json/yaml file containing variables for template [$GOSS_VARS]
   --no-color              
   --help, -h              show help
```
## Road Map
- [x] Pipeline: should be able to import/reuse another pipeline
- [x] Render: Go template support
- [x] Execute: basic workflow
- [x] Dependency graph: Generate dot graph
- [ ] Doc: should be able to generate doc from pipeline/step definitions
