package console

import (
	"strconv"
	"strings"
)

func NewIntSliceFlag(name, shorthand, usage string, values []int) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewIntSlice(values...),
	}
}

type IntSlice struct {
	value   []int
	changed bool
}

func NewIntSlice(values ...int) *IntSlice {
	return &IntSlice{
		value: values,
	}
}

func (v *IntSlice) Set(value string) error {
	d, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []int{d}
		v.changed = true
		return nil
	}
	v.value = append(v.value, d)
	return nil
}

func (v *IntSlice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(uint64(value), 10)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *IntSlice) Type() string {
	return IntSliceType
}

func (v *IntSlice) Append(value string) error {
	d, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	v.value = append(v.value, d)
	return nil
}

func (v *IntSlice) Replace(val []string) error {
	out := make([]int, len(val))
	for i, d := range val {
		var err error
		out[i], err = strconv.Atoi(d)
		if err != nil {
			return err
		}
	}
	v.value = out
	return nil
}

func (v *IntSlice) GetSlice() []string {
	out := make([]string, len(v.value))
	for i, d := range v.value {
		out[i] = strconv.Itoa(d)
	}
	return out
}

func (v *IntSlice) Value() []int {
	return v.value
}
