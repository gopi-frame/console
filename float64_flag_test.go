package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestFloat64Flag(t *testing.T) {
	handler := func(input console.Input) {
		amount, err := input.GetFloat64("amount")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		buf.WriteString(fmt.Sprintf("amount: %.2f\n", amount))

	}
	flag := NewFloat64Flag("amount", "a", "amount", 10.0)
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock"})
		if err := kernel.Execute(); !assert.NoError(t, err) {
			assert.FailNow(t, "kernel.Run() should not return error")
		}
		assert.Equal(t, "amount: 10.00\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "--amount", "100.0"})
		if err := kernel.Execute(); !assert.NoError(t, err) {
			assert.FailNow(t, "kernel.Run() should not return error")
		}
		assert.Equal(t, "amount: 100.00\n", buf.String())
	})
}
