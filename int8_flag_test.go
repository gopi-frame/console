package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestInt8Flag(t *testing.T) {
	handler := func(input console.Input) {
		amount, err := input.GetInt8("amount")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		buf.WriteString(fmt.Sprintf("amount = %d\n", amount))

	}
	flag := NewInt8Flag("amount", "a", "amount", 10)
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
		assert.Equal(t, "amount = 10\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amount", "11"); !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amount = 11\n", buf.String())
	})
}
