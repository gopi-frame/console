package console

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestStringMapFlag(t *testing.T) {
	handler := func(input console.Input) {
		attributes, err := input.GetStringMap("attributes")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		var keys []string
		for key := range attributes {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			buf.WriteString(fmt.Sprintf("attributes[%s] = %s\n", key, attributes[key]))
		}

	}
	flag := NewStringMapFlag("attributes", "a", "attributes", map[string]string{
		"name": "John",
		"age":  "30",
	})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "attributes[age] = 30\nattributes[name] = John\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock", "--attributes", "name=Jane,age=25"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "attributes[age] = 25\nattributes[name] = Jane\n", buf.String())
	})
}
