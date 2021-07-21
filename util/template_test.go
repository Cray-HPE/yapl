package util

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadVars(t *testing.T) {

	type args struct {
		varsFile string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "file_empty",
			args: args{
				varsFile: "example/pipelines/empty.yaml",
			},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "simple",
			args: args{
				varsFile: "example/pipelines/simple.yaml",
			},
			want: map[string]interface{}{
				"kind": "pipeline",
				"metadata": map[interface{}]interface{}{
					"name":        "Simple Pipeline",
					"description": "# Simple pipeline\nThis is a simple pipeline that does nothing\n\nIt is just an exmaple of what the framework can do\n## Hello\nEget felis eget nunc lobortis mattis aliquam faucibus purus in. Vulputate\nenim nulla aliquet porttitor lacus luctus accumsan tortor posuere.\nUltrices dui sapien eget mi proin sed. Neque aliquam vestibulum morbi\nblandit cursus risus at ultrices. Etiam dignissim diam quis enim lobortis\nscelerisque fermentum dui. Nulla posuere sollicitudin aliquam ultrices\nsagittis. Urna nec tincidunt praesent semper. Enim nunc faucibus a\npellentesque sit.\n\n## Bye\nEget felis eget nunc lobortis mattis aliquam faucibus purus in. Vulputate\nenim nulla aliquet porttitor lacus luctus accumsan tortor posuere.\nUltrices dui sapien eget mi proin sed. Neque aliquam vestibulum morbi\nblandit cursus risus at ultrices. Etiam dignissim diam quis enim lobortis\nscelerisque fermentum dui. Nulla posuere sollicitudin aliquam ultrices\nsagittis. Urna nec tincidunt praesent semper. Enim nunc faucibus a\npellentesque sit.\n\n![logo](https://ca.slack-edge.com/E01LD9FH0JZ-U01T66ZG20J-369a4b616e69-72)",
				},
				"spec": map[interface{}]interface{}{
					"steps": nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, filename, _, _ := runtime.Caller(0)
			t.Logf("Current test filename: %s", filename)
			fpath := filepath.Join(filename, "../..", tt.args.varsFile)
			got, err := loadVars(fpath)

			assert.Equal(t, tt.want, got, "map contents")
			assert.Equal(t, tt.wantErr, err != nil, "has error")
		})
	}
}
