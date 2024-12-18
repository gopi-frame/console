package console

import (
	"fmt"
	"strings"
	"time"
)

func NewTimeMapFlag(name, shorthand, usage string, layout string, value map[string]time.Time) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewTimeMap(value, layout),
	}
}

type TimeMap struct {
	value   map[string]time.Time
	layout  string
	changed bool
}

func NewTimeMap(value map[string]time.Time, layout string) *TimeMap {
	return &TimeMap{
		value:  value,
		layout: layout,
	}
}

func (t *TimeMap) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := time.Parse(t.layout, kv[1])
		if err != nil {
			return err
		}
		if !t.changed {
			t.changed = true
			t.value = make(map[string]time.Time)
		}
		(*t).value[kv[0]] = v
	}
	return nil
}

func (t *TimeMap) String() string {
	out := make([]string, len(t.value))
	for k, v := range t.value {
		out = append(out, fmt.Sprintf("%s=%s", k, v.Format(t.layout)))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (t *TimeMap) Type() string {
	return TimeMapType
}

func (t *TimeMap) Value() map[string]time.Time {
	return t.value
}
