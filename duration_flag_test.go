package console

import (
	"testing"
	"time"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestDurationFlag(t *testing.T) {
	handler := func(input console.Input) {
		value, err := input.GetDuration("timeout")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		buf.WriteString(value.String())

	}
	flag := NewDurationFlag("timeout", "t", "timeout", time.Second*10)
	cmd := newMockCommand(handler, flag)
	kernel := NewKernel()
	kernel.AddCommand(cmd)
	kernel.SetArgs([]string{"mock"})

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run() failed")
		}
		assert.Equal(t, "10s", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "-t", "30s"})
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run() failed")
		}
		assert.Equal(t, "30s", buf.String())
	})
}
