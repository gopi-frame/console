package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestFloat64SliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetFloat64Slice("amounts")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		for i, a := range amounts {
			buf.WriteString(fmt.Sprintf("amounts[%d] = %.2f\n", i, a))
		}

	}
	flag := NewFloat64SliceFlag("amounts", "a", "amounts", []float64{10.0, 20.0})
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
		assert.Equal(t, "amounts[0] = 10.00\namounts[1] = 20.00\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "11.0", "--amounts", "21.0"); !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[0] = 11.00\namounts[1] = 21.00\n", buf.String())
	})
}
