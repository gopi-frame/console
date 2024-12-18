package console

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/gopi-frame/contract/console"
	"github.com/stretchr/testify/assert"
)

func TestTimeMapFlag(t *testing.T) {
	handler := func(input console.Input) {
		schedule, err := input.GetTimeMap("schedule", time.DateTime)
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		var keys []string
		for key := range schedule {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			buf.WriteString(fmt.Sprintf("%s: %s\n", key, schedule[key].Format(time.DateTime)))
		}

	}
	flag := NewTimeMapFlag("schedule", "s", "schedule", time.DateTime, map[string]time.Time{
		"work":  time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
		"lunch": time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
	})
	kernel := NewKernel()
	kernel.AddCommand(newMockCommand(handler, flag))

	kernel.SetOut(buf)
	t.Run("without-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "lunch: 2023-01-01 12:00:00\nwork: 2023-01-01 09:00:00\n", buf.String())
	})
	t.Run("with-flag", func(t *testing.T) {
		defer buf.Reset()
		if err := kernel.Call("mock", "--schedule", "work=2024-01-01 09:00:00,lunch=2024-01-01 12:00:00"); err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, "lunch: 2024-01-01 12:00:00\nwork: 2024-01-01 09:00:00\n", buf.String())
	})
}
