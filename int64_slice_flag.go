package console

import (
	"strconv"
	"strings"
)

func NewInt64SliceFlag(name, shorthand, usage string, value []int64) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt64Slice(value...),
	}
}

type Int64Slice struct {
	value   []int64
	changed bool
}

func NewInt64Slice(values ...int64) *Int64Slice {
	return &Int64Slice{
		value: values,
	}
}

func (v *Int64Slice) Set(value string) error {
	d, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []int64{d}
		v.changed = true
		return nil
	}
	v.value = append(v.value, d)
	return nil
}

func (v *Int64Slice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatInt(value, 10)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *Int64Slice) Type() string {
	return Int64SliceType
}

func (v *Int64Slice) Append(value string) error {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	v.value = append(v.value, i)
	return nil
}

func (v *Int64Slice) Replace(value []string) error {
	values := make([]int64, len(value))
	for i, s := range value {
		d, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		values[i] = d
	}
	v.value = values
	return nil
}

func (v *Int64Slice) GetSlice() []string {
	if v == nil {
		return nil
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatInt(value, 10)
	}
	return values
}

func (v *Int64Slice) Value() []int64 {
	return v.value
}
