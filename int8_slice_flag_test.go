package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestInt8SliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetInt8Slice("amounts")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		for i, amount := range amounts {
			buf.WriteString(fmt.Sprintf("amounts[%d] = %d\n", i, amount))
		}

	}
	flag := NewInt8SliceFlag("amounts", "a", "amounts", []int8{10, 20})
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
		assert.Equal(t, "amounts[0] = 10\namounts[1] = 20\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "100", "--amounts", "21"); !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[0] = 100\namounts[1] = 21\n", buf.String())
	})
}
