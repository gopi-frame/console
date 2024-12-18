package console

import (
	"strconv"
	"strings"
)

func NewInt16SliceFlag(name, shorthand, usage string, value []int16) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt16Slice(value...),
	}
}

type Int16Slice struct {
	value   []int16
	changed bool
}

func NewInt16Slice(values ...int16) *Int16Slice {
	return &Int16Slice{
		value: values,
	}
}

func (v *Int16Slice) Set(value string) error {
	d, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []int16{int16(d)}
		v.changed = true
		return nil
	}
	v.value = append(v.value, int16(d))
	return nil
}

func (v *Int16Slice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatInt(int64(value), 10)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *Int16Slice) Type() string {
	return Int16SliceType
}

func (v *Int16Slice) Append(value string) error {
	i, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return err
	}
	v.value = append(v.value, int16(i))
	return nil
}

func (v *Int16Slice) Replace(value []string) error {
	values := make([]int16, len(value))
	for i, s := range value {
		d, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return err
		}
		values[i] = int16(d)
	}
	v.value = values
	return nil
}

func (v *Int16Slice) GetSlice() []string {
	if v == nil {
		return []string{}
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatInt(int64(value), 10)
	}
	return values
}

func (v *Int16Slice) Value() []int16 {
	return v.value
}
