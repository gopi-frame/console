package console

import (
	"strconv"
	"strings"
)

func NewUint32SliceFlag(name, shorthand, usage string, value []uint32) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint32Slice(value...),
	}
}

type Uint32Slice struct {
	value   []uint32
	changed bool
}

func NewUint32Slice(values ...uint32) *Uint32Slice {
	return &Uint32Slice{
		value: values,
	}
}

func (v *Uint32Slice) Set(value string) error {
	d, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []uint32{uint32(d)}
		v.changed = true
		return nil
	}
	v.value = append(v.value, uint32(d))
	return nil
}

func (v *Uint32Slice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(uint64(value), 10)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *Uint32Slice) Type() string {
	return Uint32SliceType
}

func (v *Uint32Slice) Append(value string) error {
	d, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return err
	}
	v.value = append(v.value, uint32(d))
	return nil
}

func (v *Uint32Slice) Replace(value []string) error {
	values := make([]uint32, len(value))
	for i, value := range value {
		d, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		values[i] = uint32(d)
	}
	v.value = values
	return nil
}

func (v *Uint32Slice) GetSlice() []string {
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatUint(uint64(value), 10)
	}
	return values
}

func (v *Uint32Slice) Value() []uint32 {
	return v.value
}
