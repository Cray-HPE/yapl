package util

import (
	"os"

	"github.com/Cray-HPE/yapl/model"
	"gopkg.in/yaml.v2"
)

var CACHE_DIR = "/etc/cray/yapl/.cache"

func PushToCache(genericYAML model.GenericYAML) error {
	os.MkdirAll(CACHE_DIR, os.ModePerm)
	f, err := os.Create(CACHE_DIR + "/" + string(genericYAML.Metadata.Id))
	if err != nil {
		return err
	}

	defer f.Close()
	out, _ := yaml.Marshal(genericYAML)

	_, err = f.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func PopFromCache(id string) (model.GenericYAML, error) {
	ret, err := ReadYAML(CACHE_DIR + "/" + id)
	return ret, err
}

func HasRunAlready(id string) bool {
	genericYAML, _ := PopFromCache(id)
	return genericYAML.Metadata.Completed
}

func ClearCache() error {
	return os.RemoveAll(CACHE_DIR)
}

func IsCached(id string) bool {
	if _, err := os.Stat(CACHE_DIR + "/" + id); os.IsNotExist(err) {
		return false
	}
	return true
}
