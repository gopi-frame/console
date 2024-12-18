package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestBoolSliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		value, err := input.GetBoolSlice("debug")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		for i, v := range value {
			buf.WriteString(fmt.Sprintf("%d: %t\n", i, v))
		}

	}
	flag := NewBoolSliceFlag("debug", "d", "debug mode", []bool{})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))
	kernel.SetArgs([]string{"mock"})

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run() failed")
		}
		assert.Equal(t, "", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "-d", "true", "-d", "false"})
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run() failed")
		}
		assert.Equal(t, "0: true\n1: false\n", buf.String())
	})
}
