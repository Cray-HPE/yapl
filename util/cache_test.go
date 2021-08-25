package util

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Cray-HPE/yapl/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CachePushAndPop(t *testing.T) {
	os.Setenv("CACHE_DIR", "/tmp/"+fmt.Sprint(time.Now().Unix()))
	tests := []struct {
		name    string
		step    model.GenericYAML
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			step: model.GenericYAML{
				Kind: "step",
				Metadata: model.Metadata{
					Id:          uuid.NewString(),
					ChildrenIds: []string{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClearCache()
			err := PushToCache(tt.step)
			assert.Equal(t, tt.wantErr, err != nil, "has error")
			cached, _ := PopFromCache(tt.step.Metadata.Id)
			assert.Equal(t, tt.step, cached)
			t.Log(err)

		})
	}
}

func Test_CacheHasRunAlready(t *testing.T) {
	os.Setenv("CACHE_DIR", "/tmp/"+fmt.Sprint(time.Now().Unix()))
	tests := []struct {
		name    string
		step    model.GenericYAML
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			step: model.GenericYAML{
				Kind: "step",
				Metadata: model.Metadata{
					Id:          uuid.NewString(),
					ChildrenIds: []string{},
					Completed:   false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClearCache()
			err := PushToCache(tt.step)
			assert.Equal(t, tt.wantErr, err != nil, "has error")
			assert.Equal(t, false, HasRunAlready(tt.step.Metadata.Id))

			tt.step.Metadata.Completed = true
			PushToCache(tt.step)
			assert.Equal(t, true, HasRunAlready(tt.step.Metadata.Id))

			t.Log(err)
		})
	}
}
