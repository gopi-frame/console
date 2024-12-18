package console

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestFloat64MapFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetFloat64Map("amounts")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		var keys []string
		for k := range amounts {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			buf.WriteString(fmt.Sprintf("%s: %.2f\n", k, amounts[k]))
		}

	}
	flag := NewFloat64MapFlag("amounts", "a", "amounts", map[string]float64{
		"sku1": 10.0,
		"sku2": 20.0,
	})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock"); !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "sku1: 10.00\nsku2: 20.00\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "-a", "sku2=10.0,sku3=30.0"); !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "sku2: 10.00\nsku3: 30.00\n", buf.String())
	})
}
