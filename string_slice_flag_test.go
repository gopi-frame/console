package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestStringSliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		tags, err := input.GetStringSlice("tags")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		for i, tag := range tags {
			buf.WriteString(fmt.Sprintf("tags[%d]: %s\n", i, tag))
		}

	}
	flag := NewStringSliceFlag("tags", "t", "tags", []string{"tag1", "tag2"})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "tags[0]: tag1\ntags[1]: tag2\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock", "--tags", "tag1", "--tags", "tag2", "--tags", "tag3", "--tags", "tag4"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "tags[0]: tag1\ntags[1]: tag2\ntags[2]: tag3\ntags[3]: tag4\n", buf.String())
	})
}
