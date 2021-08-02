package util

import (
	"bytes"
	"fmt"
	"log"

	"github.com/Cray-HPE/yapl/model"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

var g graphviz.Graphviz
var graph *cgraph.Graph
var nodes map[string]cgraph.Node

func DepGenFromPipeline(cfg *Config) error {
	g := graphviz.New()
	graph, _ = graphviz.New().Graph()
	nodes = map[string]cgraph.Node{}

	renderedPipeline, err := RenderPipeline(cfg)
	if err != nil {
		return err
	}

	for _, pipeline := range renderedPipeline {
		if pipeline.Kind == "pipeline" {
			buildGraph(pipeline)
			continue
		}
		if pipeline.Kind == "step" {
			buildGraph(pipeline)
			continue
		}
	}
	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Println(buf.String())
	return nil
}

func buildGraph(genericPipeline model.GenericYAML) {
	node, _ := graph.CreateNode(genericPipeline.Metadata.Id)
	if genericPipeline.Kind == "pipeline" {
		node.SetShape("circle")
	} else {
		node.SetShape("box")
	}
	node.SetLabel(fmt.Sprintf("#%d - %s", genericPipeline.Metadata.OrderId, genericPipeline.Metadata.Name))
	node.SetTooltip(genericPipeline.Metadata.Description)

	nodes[genericPipeline.Metadata.Id] = *node
	if genericPipeline.Metadata.Parent != "" {
		parentNode := nodes[genericPipeline.Metadata.Parent]
		edge, _ := graph.CreateEdge("lol", &parentNode, node)
		edge.SetDir("forward")
	}

	if genericPipeline.Kind == "step" {
		step, _ := genericPipeline.ToStep()
		for _, job := range step.Spec.Jobs {
			renderStepGraph(step, job, "Pre Condition")
			renderStepGraph(step, job, "Action")
			renderStepGraph(step, job, "Post Validation")
		}
	}
}
func renderStepGraph(step model.Step, job model.Job, text string) {
	node, _ := graph.CreateNode(step.Metadata.Id + " " + text)
	node.SetShape("cylinder")
	node.SetLabel(text)
	node.SetTooltip(job.PreCondition.Command)
	parentNode := nodes[step.Metadata.Id]
	edge, _ := graph.CreateEdge("lol", &parentNode, node)
	edge.SetDir("forward")
}
