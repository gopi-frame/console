package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestUintSliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetUintSlice("amounts")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		for i, amount := range amounts {
			buf.WriteString(fmt.Sprintf("amounts[%d] = %d\n", i, amount))
		}

	}
	flag := NewUintSliceFlag("amounts", "a", "amounts", []uint{1, 2, 3})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[0] = 1\namounts[1] = 2\namounts[2] = 3\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock", "--amounts", "4", "--amounts", "5", "--amounts", "6"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[0] = 4\namounts[1] = 5\namounts[2] = 6\n", buf.String())
	})
}
