package console

import (
	"fmt"
	"testing"
	"time"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestTimeSliceFlag(t *testing.T) {
	handler := func(input console.Input) {
		times, err := input.GetTimeSlice("times", time.DateTime)
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		for i, t := range times {
			buf.WriteString(fmt.Sprintf("times[%d] = %s\n", i, t.Format(time.DateTime)))
		}

	}
	flag := NewTimeSliceFlag("times", "t", "times", time.DateTime, []time.Time{
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
	})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "times[0] = 2023-01-01 00:00:00\ntimes[1] = 2023-01-02 00:00:00\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock", "-t", "2023-01-03 00:00:00", "-t", "2023-01-04 00:00:00"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "times[0] = 2023-01-03 00:00:00\ntimes[1] = 2023-01-04 00:00:00\n", buf.String())
	})
}
