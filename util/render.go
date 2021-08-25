package util

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/Cray-HPE/yapl/model"
	"github.com/pterm/pterm"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

var currentTemplateFilter TemplateFilter
var validate *validator.Validate
var pipelineCounter SafeCounter

func RenderPipeline(cfg *Config) (int, string, error) {
	pipelineCounter = SafeCounter{}

	validate = validator.New()
	currentTemplateFilter, _ = NewTemplateFilter(cfg.Vars)

	tmpYaml, err := ReadYAML(cfg.File)
	if err != nil {
		return 0, "", err
	}

	err = mergeYAMLData(&tmpYaml, 0, filepath.Dir(cfg.File))
	if err != nil {
		return 0, "", err
	}

	return pipelineCounter.Value(), tmpYaml.Metadata.Id, nil
}

func ReadYAML(filePath string) (model.GenericYAML, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return model.GenericYAML{}, fmt.Errorf("file error: %v", err)
	}
	return readYAMLData(file)
}

func unmarshalYAML(data []byte, v interface{}) error {
	err := yaml.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("could not unmarshal %q as YAML data: %s", string(data), err)
	}

	return nil
}

func mergeYAMLData(genericYAML *model.GenericYAML, depth int, path string) error {
	MAX_DEPTH := 50
	if genericYAML.Kind != "step" {

		depth++
		if depth >= MAX_DEPTH {
			return fmt.Errorf("max depth of %d reached, possibly due to dependency loop in yapl pipeline file", MAX_DEPTH)
		}

		pipeline, _ := genericYAML.ToPipeline()

		for _, step := range pipeline.Spec.Steps {
			step := step
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
				err = validateAndFillDefaultValues(&j)
				if err != nil {
					pterm.Error.Printf("ERROR: validation error in: %s\n", match)
					return err
				}

				err = mergeYAMLData(&j, depth, fdir)
				if err != nil {
					return err
				}
				genericYAML.Metadata.ChildrenIds = append(genericYAML.Metadata.ChildrenIds, j.Metadata.Id)
			}
		}

	}
	genericYAML.Metadata.Id = fmt.Sprint(pipelineCounter.Value())
	pipelineCounter.Inc()
	data, _ := yaml.Marshal(genericYAML)
	genericYAML.Metadata.Id = fmt.Sprintf("%x", md5.Sum(data))
	storePipelineToDisk(*genericYAML)
	pipelineCounter.Lock()
	pterm.Debug.Printf("Store: %s\n", genericYAML.Metadata.Name)
	pipelineCounter.Unlock()
	return nil
}

func validateAndFillDefaultValues(genericYAML *model.GenericYAML) error {
	genericYAML.Metadata.Completed = false
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

func storePipelineToDisk(genericYAML model.GenericYAML) {
	if !IsCached(genericYAML.Metadata.Id) {
		pterm.Debug.Printf("caching: %s - %s\n", genericYAML.Metadata.Name, genericYAML.Metadata.Id)
		PushToCache(genericYAML)
	}
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
