package console

import (
	"fmt"
	"testing"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestUint64SliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		amounts, err := input.GetUint64Slice("amounts")
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		for i, v := range amounts {
			buf.WriteString(fmt.Sprintf("amounts[%d]=%d\n", i, v))
		}

	}
	flag := NewUint64SliceFlag("amounts", "a", "amounts", []uint64{100, 200, 300})
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
		assert.Equal(t, "amounts[0]=100\namounts[1]=200\namounts[2]=300\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer func() {
			buf.Reset()
		}()
		if err := kernel.Call("mock", "--amounts", "1000", "--amounts", "2000", "--amounts", "3000"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "amounts[0]=1000\namounts[1]=2000\namounts[2]=3000\n", buf.String())
	})
}
