package util

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Cray-HPE/yapl/model"
	"github.com/stretchr/testify/assert"
)

func Test_runCommand(t *testing.T) {
	os.Setenv("CACHE_DIR", "/tmp/"+fmt.Sprint(time.Now().Unix()))
	tests := []struct {
		name    string
		cmd     string
		want    model.GenericYAML
		wantErr bool
	}{
		{
			name:    "simple pwd",
			cmd:     "pwd",
			wantErr: false,
		},
		{
			name:    "simple error",
			cmd:     "echo 'a' | grep b",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := runCommand(tt.cmd, nil)

			assert.Equal(t, tt.wantErr, err != nil, "has error")
			t.Log(err)
		})
	}
}

func Test_executeStep(t *testing.T) {
	os.Setenv("CACHE_DIR", "/tmp/"+fmt.Sprint(time.Now().Unix()))
	tests := []struct {
		name    string
		step    model.GenericYAML
		want    string
		wantErr bool
	}{
		{
			name: "fail pre condition",
			step: model.GenericYAML{
				Kind:     "step",
				Metadata: &model.Metadata{},
				Spec: &map[string]interface{}{
					"jobs": []map[string]interface{}{
						{
							"preCondition": model.Runnable{
								Description: "this is a description",
								Command:     "exit 1",
							},
							"action":        model.Runnable{},
							"errorHandling": model.Runnable{},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail action",
			step: model.GenericYAML{
				Kind:     "step",
				Metadata: &model.Metadata{},
				Spec: &map[string]interface{}{
					"jobs": []map[string]interface{}{
						{
							"preCondition": model.Runnable{
								Description: "this is a description",
								Command:     "exit 0",
							},
							"action": model.Runnable{
								Description: "this is a description",
								Command:     "exit 1",
							},
							"errorHandling": model.Runnable{},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "happy path",
			step: model.GenericYAML{
				Kind: "step",
				Metadata: &model.Metadata{
					Id: "happy path",
				},
				Spec: &map[string]interface{}{
					"jobs": []map[string]interface{}{
						{
							"preCondition": model.Runnable{
								Description: "this is a description",
								Command:     "exit 0",
							},
							"action": model.Runnable{
								Description: "this is a description",
								Command:     "exit 0",
							},
							"postValidation": model.Runnable{Description: ""},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := executeStep(&tt.step)

			assert.Equal(t, tt.wantErr, err != nil, "has error")
			t.Log(err)

		})
	}
}
