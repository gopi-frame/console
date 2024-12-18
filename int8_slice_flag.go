package console

import (
	"strconv"
	"strings"
)

func NewInt8SliceFlag(name, shorthand, usage string, value []int8) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt8Slice(value...),
	}
}

type Int8Slice struct {
	value   []int8
	changed bool
}

func NewInt8Slice(values ...int8) *Int8Slice {
	return &Int8Slice{
		value: values,
	}
}

func (v *Int8Slice) Set(value string) error {
	d, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []int8{int8(d)}
		v.changed = true
		return nil
	}
	v.value = append(v.value, int8(d))
	return nil
}

func (v *Int8Slice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatInt(int64(value), 10)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *Int8Slice) Type() string {
	return Int8SliceType
}

func (v *Int8Slice) Append(value string) error {
	d, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return err
	}
	v.value = append(v.value, int8(d))
	return nil
}

func (v *Int8Slice) Replace(value []string) error {
	out := make([]int8, len(value))
	for i, s := range value {
		d, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return err
		}
		out[i] = int8(d)
	}
	v.value = out
	return nil
}

func (v *Int8Slice) GetSlice() []string {
	if v == nil {
		return nil
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatInt(int64(value), 10)
	}
	return values
}

func (v *Int8Slice) Value() []int8 {
	return v.value
}
