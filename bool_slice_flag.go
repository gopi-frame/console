package console

import (
	"bytes"
	"encoding/csv"
	"strconv"
)

func NewBoolSliceFlag(name, shorthand, usage string, value []bool) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewBoolSlice(value...),
		hidden:    false,
		isBool:    false,
	}
}

type BoolSlice struct {
	value   []bool
	changed bool
}

func NewBoolSlice(values ...bool) *BoolSlice {
	return &BoolSlice{
		value: values,
	}
}

func (v *BoolSlice) Set(value string) error {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []bool{b}
		v.changed = true
		return nil
	}
	v.value = append(v.value, b)
	return nil
}

func (v *BoolSlice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatBool(value)
	}
	var buf = new(bytes.Buffer)
	w := csv.NewWriter(buf)
	_ = w.Write(values)
	return "[" + buf.String() + "]"
}

func (v *BoolSlice) Type() string {
	return BoolSliceType
}

func (v *BoolSlice) Append(value string) error {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	v.value = append(v.value, b)
	return nil
}

func (v *BoolSlice) Replace(values []string) error {
	out := make([]bool, len(values))
	for i, value := range values {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		out[i] = b
	}
	v.value = out
	return nil
}

func (v *BoolSlice) GetSlice() []string {
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatBool(value)
	}
	return values
}

func (v *BoolSlice) Value() []bool {
	return v.value
}
