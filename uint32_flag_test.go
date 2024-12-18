package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestUint32Flag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetUint32("amounts")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		buf.WriteString(fmt.Sprintf("amounts = %d\n", amounts))

	}
	flag := NewUint32Flag("amounts", "a", "amounts", 1000)
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
		assert.Equal(t, "amounts = 1000\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "10000"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts = 10000\n", buf.String())
	})
}
