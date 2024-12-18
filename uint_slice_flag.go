package console

import (
	"strconv"
	"strings"
)

func NewUintSliceFlag(name, shorthand, usage string, value []uint) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUintSlice(value...),
	}
}

type UintSlice struct {
	value   []uint
	changed bool
}

func NewUintSlice(values ...uint) *UintSlice {
	return &UintSlice{
		value: values,
	}
}

func (v *UintSlice) Set(value string) error {
	d, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []uint{uint(d)}
		v.changed = true
		return nil
	}
	v.value = append(v.value, uint(d))
	return nil
}

func (v *UintSlice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(uint64(value), 10)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *UintSlice) Type() string {
	return UintSliceType
}

func (v *UintSlice) Append(value string) error {
	d, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}
	v.value = append(v.value, uint(d))
	return nil
}

func (v *UintSlice) Replace(value []string) error {
	values := make([]uint, len(value))
	for i, v := range value {
		d, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		values[i] = uint(d)
	}
	v.value = values
	return nil
}

func (v *UintSlice) GetSlice() []string {
	if v == nil {
		return []string{}
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(uint64(value), 10)
	}
	return values
}

func (v *UintSlice) Value() []uint {
	return v.value
}
