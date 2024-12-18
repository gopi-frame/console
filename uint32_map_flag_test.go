package console

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestUint32MapFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetUint32Map("amounts")
		if err != nil {
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
	flag := NewUint32MapFlag("amounts", "a", "amounts", map[string]uint32{
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
		assert.Equal(t, "amounts[sku1] = 100\namounts[sku2] = 200\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "sku3=300", "--amounts", "sku4=400"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[sku3] = 300\namounts[sku4] = 400\n", buf.String())
	})
}
