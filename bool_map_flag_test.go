package console

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestNewBoolMapFlag(t *testing.T) {
	handler := func(input console.Input) {
		attrs, err := input.GetBoolMap("attrs")
		if !assert.NoError(t, err) {
			assert.FailNow(t, err.Error())
		}
		keys := make([]string, 0, len(attrs))
		for k := range attrs {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			buf.WriteString(fmt.Sprintf("%s = %v\n", k, attrs[k]))
		}

	}
	flag := NewBoolMapFlag("attrs", "a", "attributes", map[string]bool{})
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
		assert.Equal(t, "", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		kernel.SetArgs([]string{"mock", "-a", "a=true", "-a", "b=false"})
		if !assert.NoError(t, kernel.Execute()) {
			assert.FailNow(t, "kernel.Run() failed")
		}
		assert.Equal(t, "a = true\nb = false\n", buf.String())
	})
}
