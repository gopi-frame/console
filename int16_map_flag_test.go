package console

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestInt16MapFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetInt16Map("amounts")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		var keys []string
		for k := range amounts {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			buf.WriteString(fmt.Sprintf("%s: %d\n", k, amounts[k]))
		}

	}
	flag := NewInt16MapFlag("amounts", "a", "amounts", map[string]int16{
		"sku1": 100,
		"sku2": 200,
	})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "sku1: 100\nsku2: 200\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "sku1=1000"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "sku1: 1000\n", buf.String())
	})
}
