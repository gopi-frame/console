package console

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestFloat32MapFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetFloat32Map("amounts")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		var keys []string
		for key := range amounts {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			buf.WriteString(fmt.Sprintf("%s=%v\n", key, amounts[key]))
		}

	}
	flag := NewFloat32MapFlag("amounts", "a", "amounts", map[string]float32{
		"sku1": 100.0,
		"sku2": 200.01,
	})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	kernel.SetArgs([]string{"mock"})
	t.Run("without-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		err := kernel.Execute()
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "sku1=100\nsku2=200.01\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "--amounts", "sku1=100.01,sku2=200.02"})
		err := kernel.Execute()
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "sku1=100.01\nsku2=200.02\n", buf.String())
	})
}
