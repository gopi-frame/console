package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestUint32SliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetUint32Slice("amounts")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		for i, v := range amounts {
			buf.WriteString(fmt.Sprintf("amounts[%d]=%d\n", i, v))
		}

	}
	flag := NewUint32SliceFlag("amounts", "a", "amounts", []uint32{100, 200})
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
		assert.Equal(t, "amounts[0]=100\namounts[1]=200\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "1000", "--amounts", "2000"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[0]=1000\namounts[1]=2000\n", buf.String())
	})
}
