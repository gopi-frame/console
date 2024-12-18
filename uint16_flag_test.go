package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestUint16Flag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetUint16("amounts")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		buf.WriteString(fmt.Sprintf("amounts = %d\n", amounts))

	}
	flag := NewUint16Flag("amounts", "a", "amounts", 100)
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
		assert.Equal(t, "amounts = 100\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "200"); !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts = 200\n", buf.String())
	})
}
