package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestInt16SliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetInt16Slice("amounts")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		for i, v := range amounts {
			buf.WriteString(fmt.Sprintf("amounts[%d] = %d\n", i, v))
		}

	}
	flag := NewInt16SliceFlag("amounts", "a", "amounts", []int16{10, 20})
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
		assert.Equal(t, "amounts[0] = 10\namounts[1] = 20\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "100", "--amounts", "200"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[0] = 100\namounts[1] = 200\n", buf.String())
	})
}
