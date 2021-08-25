package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/Cray-HPE/yapl/model"
	"github.com/stretchr/testify/assert"
)

func Test_RenderPipeline(t *testing.T) {
	os.Setenv("CACHE_DIR", "/tmp/"+fmt.Sprint(time.Now().Unix()))
	type args struct {
		file     string
		varsFile string
	}
	type wanted struct {
		file string
		err  bool
	}
	tests := []struct {
		name string
		args args
		want wanted
	}{
		{
			name: "simple",
			args: args{
				file: "example/pipelines/simple.yaml",
			},
			want: wanted{
				file: "707ac5487d053ed1e5e80fc832f2f92c",
				err:  false,
			},
		},
		{
			name: "demo",
			args: args{
				file:     "example/pipelines/demo.yaml",
				varsFile: "example/vars.yaml",
			},
			want: wanted{
				file: "a857af958e79d27433a4f88d70260430",
				err:  false,
			},
		},
		{
			name: "wrong file",
			args: args{
				file: "/tmp/empty.yaml",
			},
			want: wanted{
				file: "",
				err:  true,
			},
		},
		{
			name: "cyclic pipeline",
			args: args{
				file: "example/pipelines/cyclic.yaml",
			},
			want: wanted{
				file: "",
				err:  true,
			},
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
				File:    fPath,
				Vars:    varPath,
				NoColor: true,
			}
			_, _, err := RenderPipeline(cfg)

			assert.Equal(t, tt.want.err, err != nil, "has error")

			if !tt.want.err {
				content, err := ioutil.ReadFile(os.Getenv("CACHE_DIR") + "/" + tt.want.file)
				if err != nil {
					t.Fatalf("Error loading golden file: %s", err)
				}
				got := string(content)

				content, err = ioutil.ReadFile("testdata/" + tt.want.file + ".golden")
				if err != nil {
					t.Fatalf("Error loading golden file: %s", err)
				}
				want := string(content)

				if got != want {
					t.Errorf("Want:\n%s\nGot:\n%s", want, got)
				}
			}
		})
	}
}

func Test_readYAMLData(t *testing.T) {
	os.Setenv("CACHE_DIR", "/tmp/"+fmt.Sprint(time.Now().Unix()))
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
