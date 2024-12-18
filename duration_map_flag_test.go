package console

import (
	"sort"
	"testing"
	"time"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestDurationMapFlag(t *testing.T) {
	handler := func(input console.Input) {
		value, err := input.GetDurationMap("timeout")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		keys := make([]string, 0, len(value))
		for k := range value {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			buf.WriteString(k + " = " + value[k].String() + "\n")
		}

	}
	flag := NewDurationMapFlag("timeout", "t", "timeout", map[string]time.Duration{
		"write": time.Second * 10,
		"read":  time.Second * 20,
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
		assert.Equal(t, "read = 20s\nwrite = 10s\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "-t", "read=30s"})
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run() failed")
		}
		assert.Equal(t, "read = 30s\n", buf.String())
	})
}
