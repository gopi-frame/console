package console

import (
	"strconv"
	"strings"
)

func NewUint8SliceFlag(name, shorthand, usage string, value []uint8) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint8Slice(value...),
	}
}

type Uint8Slice struct {
	value   []uint8
	changed bool
}

func NewUint8Slice(values ...uint8) *Uint8Slice {
	return &Uint8Slice{
		value: values,
	}
}

func (v *Uint8Slice) Set(value string) error {
	d, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []uint8{uint8(d)}
		v.changed = true
		return nil
	}
	v.value = append(v.value, uint8(d))
	return nil
}

func (v *Uint8Slice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(uint64(value), 10)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *Uint8Slice) Type() string {
	return Uint8SliceType
}

func (v *Uint8Slice) Append(value string) error {
	d, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return err
	}
	v.value = append(v.value, uint8(d))
	return nil
}

func (v *Uint8Slice) Replace(value []string) error {
	values := make([]uint8, len(value))
	for i, s := range value {
		d, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return err
		}
		values[i] = uint8(d)
	}
	v.value = values
	return nil
}

func (v *Uint8Slice) GetSlice() []string {
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(uint64(value), 10)
	}
	return values
}

func (v *Uint8Slice) Value() []uint8 {
	return v.value
}
