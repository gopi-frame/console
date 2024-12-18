package console

import (
	"strings"
	"time"
)

func NewTimeSliceFlag(name, shorthand, usage string, layout string, values []time.Time) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewTimeSlice(layout, values...),
	}
}

type TimeSlice struct {
	layout  string
	value   []time.Time
	changed bool
}

func NewTimeSlice(layout string, values ...time.Time) *TimeSlice {
	return &TimeSlice{
		layout: layout,
		value:  values,
	}
}

func (v *TimeSlice) Set(value string) error {
	t, err := time.Parse(v.layout, value)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []time.Time{t}
		v.changed = true
		return nil
	}
	v.value = append(v.value, t)
	return nil
}

func (v *TimeSlice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = value.Format(v.layout)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *TimeSlice) Type() string {
	return TimeSliceType
}

func (v *TimeSlice) Append(value string) error {
	t, err := time.Parse(v.layout, value)
	if err != nil {
		return err
	}
	v.value = append(v.value, t)
	return nil
}

func (v *TimeSlice) Replace(values []string) error {
	if v == nil {
		return nil
	}
	out := make([]time.Time, len(values))
	for i, value := range values {
		t, err := time.Parse(v.layout, value)
		if err != nil {
			return err
		}
		out[i] = t
	}
	v.value = out
	return nil
}

func (v *TimeSlice) GetSlice() []string {
	if v == nil {
		return nil
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = value.Format(v.layout)
	}
	return values
}

func (v *TimeSlice) Value() []time.Time {
	return v.value
}
