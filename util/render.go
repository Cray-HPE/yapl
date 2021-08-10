package util

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/Cray-HPE/yapl/model"
	"github.com/pterm/pterm"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

var renderedPipeline []model.GenericYAML
var currentTemplateFilter TemplateFilter
var validate *validator.Validate

func RenderPipeline(cfg *Config) ([]model.GenericYAML, error) {
	renderedPipeline = []model.GenericYAML{}

	validate = validator.New()
	currentTemplateFilter, _ = NewTemplateFilter(cfg.Vars)

	tmpYaml, err := ReadYAML(cfg.File)
	if err != nil {
		return []model.GenericYAML{{}}, err
	}

	err = mergeYAMLData(tmpYaml, 0, filepath.Dir(cfg.File))
	if err != nil {
		return []model.GenericYAML{{}}, err
	}

	if cfg.Debug {
		for _, rendered := range renderedPipeline {
			fmt.Println("---")
			renderedData, _ := yaml.Marshal(rendered)
			fmt.Printf("%v", string(renderedData))
		}
	}

	var wg sync.WaitGroup
	for _, pipeline := range renderedPipeline {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to save cache
		go func(pipeline model.GenericYAML) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// write to cache
			if !isCached(pipeline.Metadata.Id) {
				pterm.Debug.Printf("caching: %s - %s\n", pipeline.Metadata.Name, pipeline.Metadata.Id)
				pushToCache(pipeline)
			}
		}(pipeline)
	}
	// Wait for all caching to complete.
	wg.Wait()

	return renderedPipeline, nil
}

func ReadYAML(filePath string) (model.GenericYAML, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return model.GenericYAML{}, fmt.Errorf("file error: %v", err)
	}
	return readYAMLData(file)
}

func readYAMLData(data []byte) (model.GenericYAML, error) {
	var err error
	if currentTemplateFilter != nil {
		data, err = currentTemplateFilter(data)
		if err != nil {
			return model.GenericYAML{}, err
		}
	}

	genericYAML := model.NewGenericYAML()
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
	genericYAML.Metadata.RenderedId = len(renderedPipeline)
	data, _ := yaml.Marshal(genericYAML)
	genericYAML.Metadata.Id = fmt.Sprintf("%x", md5.Sum(data))
	renderedPipeline = append(renderedPipeline, genericYAML)

	if genericYAML.Kind == "step" {
		return nil
	}

	depth++
	if depth >= 50 {
		return fmt.Errorf("max depth of 50 reached, possibly due to dependency loop in goss file")
	}

	pipeline, _ := genericYAML.ToPipeline()

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
			j, err := ReadYAML(match)
			if err != nil {
				return fmt.Errorf("could not read json data in %s: %s", match, err)
			}
			j.Metadata.ParentId = genericYAML.Metadata.Id
			genericYAML.Metadata.ChildrenIds = append(genericYAML.Metadata.ChildrenIds, j.Metadata.Id)
			err = validateAndFillDefaultValues(&j)
			if err != nil {
				pterm.Error.Printf("ERROR: validation error in: %s\n", match)
				return err
			}

			err = mergeYAMLData(j, depth, fdir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateAndFillDefaultValues(genericYAML *model.GenericYAML) error {
	genericYAML.Metadata.Status = "Not Started"
	switch genericYAML.Kind {
	case "pipeline":
		_, err := genericYAML.ToPipeline()
		return err
	case "step":
		stepYaml, err := genericYAML.ToStep()
		if err != nil {
			return err
		}

		err = validate.Struct(stepYaml)
		return err
	default:
		return fmt.Errorf("ERROR: Kind Must Be pipeline or step, get: %s", genericYAML.Kind)
	}
}
