package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v2"
)

type TmplVars struct {
	Vars map[string]interface{}
}

// TemplateFilter is the type of the yapl Template Filter which include custom variables and functions.
type TemplateFilter func([]byte) ([]byte, error)

// NewTemplateFilter creates a new Template Filter based in the file and inline variables.
func NewTemplateFilter(varsFile string) (func([]byte) ([]byte, error), error) {
	vars, err := loadVars(varsFile)
	if err != nil {
		return nil, fmt.Errorf("failed while loading vars file %q: %v", varsFile, err)
	}

	tVars := &TmplVars{Vars: vars}

	f := func(data []byte) ([]byte, error) {
		t := template.New("test").Funcs(sprig.TxtFuncMap()).Funcs(funcMap)

		tmpl, err := t.Parse(string(data))
		if err != nil {
			return []byte{}, err
		}

		tmpl.Option("missingkey=error")
		var doc bytes.Buffer

		err = tmpl.Execute(&doc, tVars)
		if err != nil {
			return []byte{}, err
		}

		return doc.Bytes(), nil
	}

	return f, nil
}

func mkSlice(args ...interface{}) []interface{} {
	return args
}

func readFile(f string) (string, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err

	}
	return strings.TrimSpace(string(b)), nil
}

func getEnv(key string, def ...string) string {
	val := os.Getenv(key)
	if val == "" && len(def) > 0 {
		return def[0]
	}

	return os.Getenv(key)
}

var funcMap = template.FuncMap{
	"mkSlice":  mkSlice,
	"readFile": readFile,
	"getEnv":   getEnv,
}

func loadVars(varsFile string) (map[string]interface{}, error) {
	vars := make(map[string]interface{})
	if varsFile == "" {
		return vars, nil
	}
	data, err := ioutil.ReadFile(varsFile)
	if err != nil {
		return vars, err
	}
	if err := yaml.Unmarshal(data, &vars); err != nil {
		return vars, err
	}
	return vars, nil
}
