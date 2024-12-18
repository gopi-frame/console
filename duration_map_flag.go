package console

import (
	"fmt"
	"strings"
	"time"
)

func NewDurationMapFlag(name, shorthand, usage string, value map[string]time.Duration) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewDurationMap(value),
	}
}

type DurationMap struct {
	value   map[string]time.Duration
	changed bool
}

func NewDurationMap(value map[string]time.Duration) *DurationMap {
	return &DurationMap{
		value: value,
	}
}

func (d *DurationMap) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := time.ParseDuration(kv[1])
		if err != nil {
			return err
		}
		if !d.changed {
			d.changed = true
			d.value = make(map[string]time.Duration)
		}
		d.value[kv[0]] = v
	}
	return nil
}

func (d *DurationMap) String() string {
	out := make([]string, len(d.value))
	for k, v := range d.value {
		out = append(out, fmt.Sprintf("%s=%s", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (d *DurationMap) Type() string {
	return DurationMapType
}

func (d *DurationMap) Value() map[string]time.Duration {
	return d.value
}
