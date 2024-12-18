package console

import (
	"strconv"
	"strings"
)

func NewUint64SliceFlag(name, shorthand, usage string, value []uint64) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint64Slice(value...),
	}
}

type Uint64Slice struct {
	value   []uint64
	changed bool
}

func NewUint64Slice(values ...uint64) *Uint64Slice {
	return &Uint64Slice{
		value: values,
	}
}

func (v *Uint64Slice) Set(value string) error {
	d, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []uint64{d}
		v.changed = true
		return nil
	}
	v.value = append(v.value, d)
	return nil
}

func (v *Uint64Slice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(value, 10)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *Uint64Slice) Type() string {
	return Uint64SliceType
}

func (v *Uint64Slice) Append(value string) error {
	d, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}
	v.value = append(v.value, d)
	return nil
}

func (v *Uint64Slice) Replace(value []string) error {
	values := make([]uint64, len(value))
	for i, v := range value {
		d, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		values[i] = d
	}
	v.value = values
	return nil
}

func (v *Uint64Slice) GetSlice() []string {
	if v == nil {
		return nil
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(value, 10)
	}
	return values
}

func (v *Uint64Slice) Value() []uint64 {
	return v.value
}
