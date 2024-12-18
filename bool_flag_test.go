package console

import (
	"bytes"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

var buf = bytes.NewBuffer(nil)

func TestBoolFlag(t *testing.T) {
	kernel := NewKernel()
	handler := func(input console.Input) {
		value, err := input.GetBool("debug")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		if value {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}
	}
	flag := NewBoolFlag("debug", "d", "debug mode", false)
	cmd := newMockCommand(handler, flag)
	kernel.AddCommand(cmd)

	kernel.SetArgs([]string{"mock"})
	t.Run("without-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run() failed")
		}
		assert.Equal(t, "false", buf.String())
	})

	t.Run("with-flag-and-value", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "-d", "true"})
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run() failed")
		}
		assert.Equal(t, "true", buf.String())
	})

	t.Run("with-flag-but-without-value", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "-d"})
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run() failed")
		}
		assert.Equal(t, "true", buf.String())
	})
}
