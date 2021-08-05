package util

import (
	"testing"

	"github.com/Cray-HPE/yapl/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CachePushAndPop(t *testing.T) {
	CACHE_DIR = ".cache"
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

			err := pushToCache(tt.step)
			assert.Equal(t, tt.wantErr, err != nil, "has error")
			cached, _ := popFromCache(tt.step.Metadata.Id)
			assert.Equal(t, tt.step, cached)
			t.Log(err)

		})
	}
}

func Test_CacheHasRunAlready(t *testing.T) {
	CACHE_DIR = ".cache"
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
			err := pushToCache(tt.step)
			assert.Equal(t, tt.wantErr, err != nil, "has error")
			assert.Equal(t, false, hasRunAlready(tt.step.Metadata.Id))

			tt.step.Metadata.Completed = true
			pushToCache(tt.step)
			assert.Equal(t, true, hasRunAlready(tt.step.Metadata.Id))

			t.Log(err)
		})
	}
}
