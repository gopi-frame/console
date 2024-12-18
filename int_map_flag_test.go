package console

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestIntMapFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetIntMap("amounts")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		var keys []string
		for key := range amounts {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			buf.WriteString(fmt.Sprintf("amounts[%s]=%d\n", key, amounts[key]))
		}

	}
	flag := NewIntMapFlag("amounts", "a", "amounts", map[string]int{
		"sku1": 10,
		"sku2": 20,
	})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[sku1]=10\namounts[sku2]=20\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock", "--amounts", "sku1=100", "--amounts", "sku2=200"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[sku1]=100\namounts[sku2]=200\n", buf.String())
	})
}
