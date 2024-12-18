package console

import (
	"fmt"
	"testing"
	"time"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestTimeFlag(t *testing.T) {
	handler := func(input console.Input) {
		deadline, err := input.GetTime("deadline", time.DateTime)
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		buf.WriteString(fmt.Sprintf("deadline: %s\n", deadline.Format(time.DateTime)))

	}
	flag := NewTimeFlag("deadline", "d", "deadline", time.DateTime, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "deadline: 2024-01-01 00:00:00\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock", "--deadline", "2024-01-02 00:00:00"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "deadline: 2024-01-02 00:00:00\n", buf.String())
	})
}
