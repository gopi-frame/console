package console

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestUint8MapFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetUint8Map("amounts")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		var keys []string
		for k := range amounts {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			buf.WriteString(fmt.Sprintf("amounts[%s] = %d\n", k, amounts[k]))
		}

	}
	flag := NewUint8MapFlag("amounts", "a", "amounts", map[string]uint8{
		"sku1": 10,
		"sku2": 20,
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
		assert.Equal(t, "amounts[sku1] = 10\namounts[sku2] = 20\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "sku1=100,sku2=21"); !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[sku1] = 100\namounts[sku2] = 21\n", buf.String())
	})
}
