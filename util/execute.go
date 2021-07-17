package util

import "fmt"

func ExecutePipeline(cfg *Config) error {
	renderedPipeline, err := RenderPipeline(cfg)
	if err != nil {
		return err
	}

	for _, pipeline := range renderedPipeline {
		if pipeline.Kind == "pipeline" {
			fmt.Printf("==> Execute Pipeline: %s \n", pipeline.Metadata.Name)
			fmt.Println(Indent(pipeline.Metadata.Description, "    "))
			continue
		}
		if pipeline.Kind == "step" {
			fmt.Printf("    ==> Execute Step: %s \n", pipeline.Metadata.Name)
			fmt.Println(Indent(pipeline.Metadata.Description, "        "))
			step := pipeline.ToStep()
			for index, job := range step.Spec.Jobs {
				fmt.Printf("        ==> Execute job: %d (Pre-condition)\n", index)
				fmt.Println(Indent(job.PreCondition.Description, "                "))

				fmt.Printf("        ==> Execute job: %d (Action)\n", index)
				fmt.Println(Indent(job.PreCondition.Description, "                "))

				fmt.Printf("        ==> Execute job: %d (Error Handling)\n", index)
				fmt.Println(Indent(job.PreCondition.Description, "                "))
			}
			continue
		}
	}

	return nil
}
