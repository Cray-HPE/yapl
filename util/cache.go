package util

import (
	"os"

	"github.com/Cray-HPE/yapl/model"
	"gopkg.in/yaml.v2"
)

func getCacheDir() string {
	res := "/etc/cray/yapl/.cache"
	if mp := os.Getenv("CACHE_DIR"); mp != "" {
		res = mp
	}
	return res
}

func PushToCache(genericYAML model.GenericYAML) error {
	if err := os.MkdirAll(getCacheDir(), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(getCacheDir() + "/" + string(genericYAML.Metadata.Id))
	if err != nil {
		return err
	}

	defer f.Close()
	out, _ := yaml.Marshal(genericYAML)

	if _, err = f.Write(out); err != nil {
		return err
	}
	return nil
}

func PopFromCache(id string) (model.GenericYAML, error) {
	ret, err := ReadYAML(getCacheDir() + "/" + id)
	return ret, err
}

func HasRunAlready(id string) bool {
	genericYAML, _ := PopFromCache(id)
	return genericYAML.Metadata.Completed
}

func ClearCache() error {
	return os.RemoveAll(getCacheDir())
}

func IsCached(id string) bool {
	if _, err := os.Stat(getCacheDir() + "/" + id); os.IsNotExist(err) {
		return false
	}
	return true
}
