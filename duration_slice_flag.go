package console

import (
	"strings"
	"time"
)

func NewDurationSliceFlag(name string, shorthand string, usage string, value []time.Duration) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewDurationSlice(value...),
	}
}

type DurationSlice struct {
	value   []time.Duration
	changed bool
}

func NewDurationSlice(values ...time.Duration) *DurationSlice {
	return &DurationSlice{
		value: values,
	}
}

func (d *DurationSlice) Set(value string) error {
	v, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	if !d.changed {
		d.value = []time.Duration{v}
		d.changed = true
		return nil
	}
	d.value = append(d.value, v)
	return nil
}

func (d *DurationSlice) String() string {
	if d == nil {
		return "[]"
	}
	out := make([]string, len(d.value))
	for i, v := range d.value {
		out[i] = v.String()
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (d *DurationSlice) Type() string {
	return DurationSliceType
}

func (d *DurationSlice) Append(value string) error {
	v, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	d.value = append(d.value, v)
	return nil
}

func (d *DurationSlice) Replace(value []string) error {
	values := make([]time.Duration, len(value))
	for i, v := range value {
		if n, err := time.ParseDuration(v); err == nil {
			values[i] = n
		} else {
			return err
		}
	}
	d.value = values
	return nil
}

func (d *DurationSlice) GetSlice() []string {
	if d == nil {
		return nil
	}
	out := make([]string, len(d.value))
	for i, v := range d.value {
		out[i] = v.String()
	}
	return out
}

func (d *DurationSlice) Value() []time.Duration {
	return d.value
}
