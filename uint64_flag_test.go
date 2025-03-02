package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestUint64Flag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetUint64("amounts")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		buf.WriteString(fmt.Sprintf("amounts: %d\n", amounts))

	}
	flag := NewUint64Flag("amounts", "a", "amounts", 100)
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
		assert.Equal(t, "amounts: 100\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "1000"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts: 1000\n", buf.String())
	})
}
