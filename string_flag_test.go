package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestStringFlag(t *testing.T) {
	handler := func(input console.Input) {
		name, err := input.GetString("name")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		buf.WriteString(fmt.Sprintf("hello %s", name))

	}
	flag := NewStringFlag("name", "", "name", "world")
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "hello world", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock", "--name", "gopi"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "hello gopi", buf.String())
	})
}
