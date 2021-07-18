package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/dzou-hpe/yapl/model"
	"gopkg.in/yaml.v2"
)

var renderedPipeline []model.GenericYAML

func RenderPipeline(cfg *Config) ([]model.GenericYAML, error) {
	renderedPipeline = []model.GenericYAML{}

	tmpYaml, err := readYAML(cfg.File)
	if err != nil {
		return []model.GenericYAML{}, err
	}

	err = mergeYAMLData(tmpYaml, 0, filepath.Dir(cfg.File))
	if err != nil {
		return []model.GenericYAML{}, err
	}

	if cfg.Debug {
		for _, rendered := range renderedPipeline {
			Blue.Println("---")
			renderedData, _ := yaml.Marshal(rendered)
			Blue.Printf("%v", string(renderedData))
		}
	}
	return renderedPipeline, nil
}

func readYAML(filePath string) (model.GenericYAML, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return model.GenericYAML{}, fmt.Errorf("file error: %v", err)
	}

	return readYAMLData(file, false)
}

func readYAMLData(data []byte, detectFormat bool) (model.GenericYAML, error) {
	// var err error
	// if currentTemplateFilter != nil {
	// 	data, err = currentTemplateFilter(data)
	// 	if err != nil {
	// 		return model.Pipeline{}, err
	// 	}
	// }

	genericYAML := model.NewGenericYAML()
	// Horrible, but will do for now
	if err := unmarshalYAML(data, genericYAML); err != nil {
		return *genericYAML, err
	}

	return *genericYAML, nil
}

func unmarshalYAML(data []byte, v interface{}) error {
	err := yaml.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("could not unmarshal %q as YAML data: %s", string(data), err)
	}

	return nil
}

func mergeYAMLData(genericYAML model.GenericYAML, depth int, path string) error {
	renderedPipeline = append(renderedPipeline, genericYAML)

	if genericYAML.Kind == "step" {
		return nil
	}

	depth++
	if depth >= 50 {
		return fmt.Errorf("max depth of 50 reached, possibly due to dependency loop in goss file")
	}

	pipeline := genericYAML.ToPipeline()

	// Merge gossfiles in sorted order
	for _, step := range pipeline.Spec.Steps {
		fpath := filepath.Join(path, step)
		matches, err := filepath.Glob(fpath)
		if err != nil {
			return fmt.Errorf("error in expanding glob pattern: %q", err)
		}
		if matches == nil {
			return fmt.Errorf("no matched files were found: %q", fpath)
		}
		for _, match := range matches {
			fdir := filepath.Dir(match)
			j, err := readYAML(match)
			if err != nil {
				return fmt.Errorf("could not read json data in %s: %s", match, err)
			}
			j.Metadata.Parent = genericYAML.Metadata.Name
			err = mergeYAMLData(j, depth, fdir)
			if err != nil {
				return fmt.Errorf("could not write json data: %s", err)
			}
		}
	}
	return nil
}
