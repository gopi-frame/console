package console

import (
	"strconv"
	"strings"
)

func NewInt32SliceFlag(name, shorthand, usage string, value []int32) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt32Slice(value...),
	}
}

type Int32Slice struct {
	value   []int32
	changed bool
}

func NewInt32Slice(values ...int32) *Int32Slice {
	return &Int32Slice{
		value: values,
	}
}

func (v *Int32Slice) Set(value string) error {
	d, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []int32{int32(d)}
		v.changed = true
		return nil
	}
	v.value = append(v.value, int32(d))
	return nil
}

func (v *Int32Slice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatInt(int64(value), 10)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *Int32Slice) Type() string {
	return Int32SliceType
}

func (v *Int32Slice) Append(value string) error {
	i, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return err
	}
	v.value = append(v.value, int32(i))
	return nil
}

func (v *Int32Slice) Replace(value []string) error {
	values := make([]int32, len(value))
	for i, s := range value {
		d, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return err
		}
		values[i] = int32(d)
	}
	v.value = values
	return nil
}

func (v *Int32Slice) GetSlice() []string {
	if v == nil {
		return nil
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatInt(int64(value), 10)
	}
	return values
}

func (v *Int32Slice) Value() []int32 {
	return v.value
}
