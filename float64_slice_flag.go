package console

import (
	"strconv"
	"strings"
)

func NewFloat64SliceFlag(name string, shorthand string, usage string, value []float64) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		value:     NewFloat64Slice(value...),
		usage:     usage,
	}
}

type Float64Slice struct {
	value   []float64
	changed bool
}

func NewFloat64Slice(values ...float64) *Float64Slice {
	return &Float64Slice{
		value: values,
	}
}

func (v *Float64Slice) Set(value string) error {
	d, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []float64{d}
		v.changed = true
		return nil
	}
	v.value = append(v.value, d)
	return nil
}

func (v *Float64Slice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatFloat(value, 'g', -1, 64)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *Float64Slice) Type() string {
	return Float64SliceType
}

func (v *Float64Slice) Append(value string) error {
	d, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	v.value = append(v.value, d)
	return nil
}

func (v *Float64Slice) Replace(value []string) error {
	values := make([]float64, len(value))
	for i, s := range value {
		d, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		values[i] = d
	}
	v.value = values
	return nil
}

func (v *Float64Slice) GetSlice() []string {
	if v == nil {
		return nil
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatFloat(value, 'g', -1, 64)
	}
	return values
}

func (v *Float64Slice) Value() []float64 {
	return v.value
}
