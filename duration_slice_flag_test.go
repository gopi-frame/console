package console

import (
	"fmt"
	"testing"
	"time"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestDurationSliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		values, err := input.GetDurationSlice("timeout")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		for i, v := range values {
			buf.WriteString(fmt.Sprintf("%d: %s\n", i, v.String()))
		}

	}
	flag := NewDurationSliceFlag("timeout", "t", "timeout", []time.Duration{
		time.Second * 10,
		time.Second * 20,
	})
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
		assert.Equal(t, "0: 10s\n1: 20s\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "-t", "30s", "-t", "40s"})
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run() failed")
		}
		assert.Equal(t, "0: 30s\n1: 40s\n", buf.String())
	})
}
