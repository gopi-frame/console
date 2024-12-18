package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestFloat32Flag(t *testing.T) {
	handler := func(input console.Input) {
		amount, err := input.GetFloat32("amount")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		buf.WriteString(fmt.Sprintf("amount: %.2f\n", amount))

	}
	flag := NewFloat32Flag("amount", "a", "amount", 10.01)
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))
	kernel.SetArgs([]string{"mock"})

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run()")
		}
		assert.Equal(t, "amount: 10.01\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "--amount", "10.02"})
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run()")
		}
		assert.Equal(t, "amount: 10.02\n", buf.String())
	})
}
