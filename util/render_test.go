package util

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Cray-HPE/yapl/model"
	"github.com/stretchr/testify/assert"
)

func Test_RenderPipeline(t *testing.T) {

	type args struct {
		file     string
		varsFile string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.GenericYAML
		wantErr bool
	}{
		// {
		// 	name: "empty",
		// 	args: args{
		// 		file: "example/pipelines/empty.yaml",
		// 	},
		// 	want: []model.GenericYAML{
		// 		{
		// 			Kind: "",
		// 			Metadata: model.Metadata{
		// 				Name:        "",
		// 				Description: "",
		// 				Parent:      "",
		// 				Id:          "8a80554c91d9fca8acb82f023de02f11",
		// 				OrderId:     0,
		// 			},
		// 			Spec: map[string]interface{}{},
		// 		},
		// 	},
		// 	wantErr: false,
		// },
		{
			name: "simple",
			args: args{
				file: "example/pipelines/simple.yaml",
			},
			want: []model.GenericYAML{
				{
					Kind: "pipeline",
					Metadata: model.Metadata{
						Name:        "Simple Pipeline",
						Description: "# Simple pipeline\nThis is a simple pipeline that does nothing\n\nIt is just an exmaple of what the framework can do\n## Hello\nEget felis eget nunc lobortis mattis aliquam faucibus purus in. Vulputate\nenim nulla aliquet porttitor lacus luctus accumsan tortor posuere.\nUltrices dui sapien eget mi proin sed. Neque aliquam vestibulum morbi\nblandit cursus risus at ultrices. Etiam dignissim diam quis enim lobortis\nscelerisque fermentum dui. Nulla posuere sollicitudin aliquam ultrices\nsagittis. Urna nec tincidunt praesent semper. Enim nunc faucibus a\npellentesque sit.\n\n## Bye\nEget felis eget nunc lobortis mattis aliquam faucibus purus in. Vulputate\nenim nulla aliquet porttitor lacus luctus accumsan tortor posuere.\nUltrices dui sapien eget mi proin sed. Neque aliquam vestibulum morbi\nblandit cursus risus at ultrices. Etiam dignissim diam quis enim lobortis\nscelerisque fermentum dui. Nulla posuere sollicitudin aliquam ultrices\nsagittis. Urna nec tincidunt praesent semper. Enim nunc faucibus a\npellentesque sit.\n\n![logo](https://ca.slack-edge.com/E01LD9FH0JZ-U01T66ZG20J-369a4b616e69-72)",
						Id:          "f4fa570d4735f2ee74081b1add49d131",
						OrderId:     0,
					},
					Spec: map[string]interface{}{
						"steps": nil,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "demo",
			args: args{
				file:     "example/pipelines/demo.yaml",
				varsFile: "example/vars.yaml",
			},
			want: []model.GenericYAML{
				{
					Kind: "pipeline",
					Metadata: model.Metadata{
						Name:        "Demo Pipeline",
						Description: "# Demo pipeline\nThis is a demo pipeline that does some simple file operations\n\nIt is just an exmaple of what the framework can do\n## create file\nEget felis eget nunc lobortis mattis aliquam faucibus purus in. Vulputate\nenim nulla aliquet porttitor lacus luctus accumsan tortor posuere.\nUltrices dui sapien eget mi proin sed. Neque aliquam vestibulum morbi\nblandit cursus risus at ultrices. Etiam dignissim diam quis enim lobortis\nscelerisque fermentum dui. Nulla posuere sollicitudin aliquam ultrices\nsagittis. Urna nec tincidunt praesent semper. Enim nunc faucibus a\npellentesque sit.\n\n## append text to the file\nEget felis eget nunc lobortis mattis aliquam faucibus purus in. Vulputate\nenim nulla aliquet porttitor lacus luctus accumsan tortor posuere.\nUltrices dui sapien eget mi proin sed. Neque aliquam vestibulum morbi\nblandit cursus risus at ultrices. Etiam dignissim diam quis enim lobortis\nscelerisque fermentum dui. Nulla posuere sollicitudin aliquam ultrices\nsagittis. Urna nec tincidunt praesent semper. Enim nunc faucibus a\npellentesque sit.\n\n## delete file\nEget felis eget nunc lobortis mattis aliquam faucibus purus in. Vulputate\nenim nulla aliquet porttitor lacus luctus accumsan tortor posuere.\nUltrices dui sapien eget mi proin sed. Neque aliquam vestibulum morbi\nblandit cursus risus at ultrices. Etiam dignissim diam quis enim lobortis\nscelerisque fermentum dui. Nulla posuere sollicitudin aliquam ultrices\nsagittis. Urna nec tincidunt praesent semper. Enim nunc faucibus a\npellentesque sit.\n\n## delete file again\nthis will fail because file doesn't exist\n\n![logo](https://ca.slack-edge.com/E01LD9FH0JZ-U01T66ZG20J-369a4b616e69-72)",
						Id:          "ac0503d5a750aa6e5ee976618f507258",
						OrderId:     0,
					},
					Spec: map[string]interface{}{
						"steps": []interface{}{
							"../steps/create-file.yaml",
							"../steps/append-text.yaml",
							"../steps/delete-file.yaml",
							"../steps/delete-file.yaml",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "wrong file",
			args: args{
				file: "/tmp/empty.yaml",
			},
			want:    []model.GenericYAML{{}},
			wantErr: true,
		},
		{
			name: "cyclic pipeline",
			args: args{
				file: "example/pipelines/cyclic.yaml",
			},
			want:    []model.GenericYAML{{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, filename, _, _ := runtime.Caller(0)
			fPath := filepath.Join(filename, "../..", tt.args.file)
			varPath := ""
			if tt.args.varsFile != "" {
				varPath = filepath.Join(filename, "../..", tt.args.varsFile)
			}
			cfg := &Config{
				File:  fPath,
				Vars:  varPath,
				Debug: true,
			}
			got, err := RenderPipeline(cfg)

			assert.Equal(t, tt.want[0], got[0], "map contents")
			assert.Equal(t, tt.wantErr, err != nil, "has error")
			t.Log(err)
		})
	}
}

func Test_readYAMLData(t *testing.T) {

	tests := []struct {
		name    string
		data    []byte
		want    model.GenericYAML
		wantErr bool
	}{
		{
			name:    "empty",
			data:    []byte("asdf"),
			want:    model.GenericYAML{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readYAMLData(tt.data)

			assert.Equal(t, tt.want, got, "map contents")
			assert.Equal(t, tt.wantErr, err != nil, "has error")
			t.Log(err)
		})
	}
}
