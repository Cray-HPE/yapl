package util

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_docGenFromPipeline(t *testing.T) {

	tests := []struct {
		name    string
		cfg     *Config
		want    []string
		wantErr bool
	}{
		{
			name: "demo pipeline",
			cfg: &Config{
				File:      "../example/pipelines/demo.yaml",
				Vars:      "../example/vars.yaml",
				OutputDir: "./dist",
			},
			want:    []string{"dist/Append file.md", "dist/Create file.md", "dist/Delete file.md", "dist/Demo Pipeline.md"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DocGenFromPipeline(tt.cfg)
			for _, file := range tt.want {
				if _, err := os.Stat(file); os.IsNotExist(err) {
					assert.Error(t, err, fmt.Sprintf("file: %s isn't created", file))
				}
			}
			assert.Equal(t, tt.wantErr, err != nil, "has error")
			t.Log(err)
		})
	}
}
