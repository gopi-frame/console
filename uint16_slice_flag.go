package console

import (
	"strconv"
	"strings"
)

func NewUint16SliceFlag(name, shorthand, usage string, value []uint16) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint16Slice(value...),
	}
}

type Uint16Slice struct {
	value   []uint16
	changed bool
}

func NewUint16Slice(values ...uint16) *Uint16Slice {
	return &Uint16Slice{
		value: values,
	}
}

func (v *Uint16Slice) Set(value string) error {
	d, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []uint16{uint16(d)}
		v.changed = true
		return nil
	}
	v.value = append(v.value, uint16(d))
	return nil
}

func (v *Uint16Slice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(uint64(value), 10)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *Uint16Slice) Type() string {
	return Uint16SliceType
}

func (v *Uint16Slice) Append(value string) error {
	d, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return err
	}
	v.value = append(v.value, uint16(d))
	return nil
}

func (v *Uint16Slice) Replace(value []string) error {
	values := make([]uint16, len(value))
	for i, value := range value {
		d, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return err
		}
		values[i] = uint16(d)
	}
	v.value = values
	return nil
}

func (v *Uint16Slice) GetSlice() []string {
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(uint64(value), 10)
	}
	return values
}

func (v *Uint16Slice) Value() []uint16 {
	return v.value
}
