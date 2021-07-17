package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/dzou-hpe/yapl/model"
	"gopkg.in/yaml.v2"
)

func RenderPipeline(cfg *Config) (string, error) {
	j, err := readYAML(cfg.File)
	if err != nil {
		return "", err
	}

	pipeline, err := mergeYAMLData(j, 0, filepath.Dir(cfg.File))
	if err != nil {
		return "", err
	}
	b, err := yaml.Marshal(pipeline)
	if err != nil {
		return "", fmt.Errorf("rendering failed: %v", err)
	}
	return string(b), nil
}

func readYAML(filePath string) (model.Pipeline, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return model.Pipeline{}, fmt.Errorf("file error: %v", err)
	}

	return readYAMLData(file, false)
}

func readYAMLData(data []byte, detectFormat bool) (model.Pipeline, error) {
	// var err error
	// if currentTemplateFilter != nil {
	// 	data, err = currentTemplateFilter(data)
	// 	if err != nil {
	// 		return model.Pipeline{}, err
	// 	}
	// }

	pipeline := model.NewPipeline()
	// Horrible, but will do for now
	if err := unmarshalYAML(data, pipeline); err != nil {
		return *pipeline, err
	}

	return *pipeline, nil
}

func unmarshalYAML(data []byte, v interface{}) error {
	err := yaml.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("could not unmarshal %q as YAML data: %s", string(data), err)
	}

	return nil
}

func mergeYAMLData(pipeline model.Pipeline, depth int, path string) (model.Pipeline, error) {
	b, _ := yaml.Marshal(pipeline)
	fmt.Println("---")
	fmt.Println(string(b))
	if pipeline.Kind == "step" {
		return pipeline, nil
	}

	depth++
	if depth >= 50 {
		return model.Pipeline{}, fmt.Errorf("max depth of 50 reached, possibly due to dependency loop in goss file")
	}
	// Our return pipeline
	ret := *model.NewPipeline()
	//ret = mergePipeline(ret, pipeline)
	ret = pipeline

	// // Sort the gossfiles to ensure consistent ordering
	// var keys []string
	// for k := range pipeline.Steps {
	// 	keys = append(keys, k)
	// }
	// sort.Strings(keys)

	// Merge gossfiles in sorted order
	for _, step := range pipeline.Steps {
		fpath := filepath.Join(path, step)
		matches, err := filepath.Glob(fpath)
		if err != nil {
			return ret, fmt.Errorf("error in expanding glob pattern: %q", err)
		}
		if matches == nil {
			return ret, fmt.Errorf("no matched files were found: %q", fpath)
		}
		for _, match := range matches {
			fdir := filepath.Dir(match)
			j, err := readYAML(match)
			if err != nil {
				return model.Pipeline{}, fmt.Errorf("could not read json data in %s: %s", match, err)
			}
			j, err = mergeYAMLData(j, depth, fdir)
			if err != nil {
				return ret, fmt.Errorf("could not write json data: %s", err)
			}
		}
	}
	return ret, nil
}

// func (c *model.Pipeline) Merge(g2 interface{}) {
// 	return
// }
