package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestFloat32SliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetFloat32Slice("amounts")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		for i, amount := range amounts {
			buf.WriteString(fmt.Sprintf("amounts[%d]=%.2f\n", i, amount))
		}

	}
	flag := NewFloat32SliceFlag("amounts", "a", "amounts", []float32{1.0, 2.0, 3.0})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock"})
		if err := kernel.Execute(); !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[0]=1.00\namounts[1]=2.00\namounts[2]=3.00\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "--amounts", "4.0", "--amounts", "5.0", "--amounts", "6.0"})
		if err := kernel.Execute(); !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[0]=4.00\namounts[1]=5.00\namounts[2]=6.00\n", buf.String())
	})
}
