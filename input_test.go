package console

import (
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func getFlagMockHandler[T any](t *testing.T, excepted T) func(input console.Input) {
	t.Helper()
	return func(input console.Input) {
		type Data struct {
			Value T `flag:"value"`
		}
		var data Data
		if err := input.Unmarshal(&data); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, excepted, data.Value)
	}
}

func getArgMockHandler[T any](t *testing.T, excepted T) func(input console.Input) {
	t.Helper()
	return func(input console.Input) {
		type Data struct {
			Value T `arg:"0"`
		}
		var data Data
		if err := input.Unmarshal(&data); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, excepted, data.Value)

	}
}
