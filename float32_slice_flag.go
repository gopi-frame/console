package console

import (
	"strconv"
	"strings"
)

func NewFloat32SliceFlag(name, shorthand string, usage string, value []float32) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewFloat32Slice(value...),
	}
}

type Float32Slice struct {
	value   []float32
	changed bool
}

func NewFloat32Slice(values ...float32) *Float32Slice {
	return &Float32Slice{
		value: values,
	}
}

func (v *Float32Slice) Set(value string) error {
	d, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}
	if !v.changed {
		v.value = []float32{float32(d)}
		v.changed = true
		return nil
	}
	v.value = append(v.value, float32(d))
	return nil
}

func (v *Float32Slice) String() string {
	if v == nil {
		return "[]"
	}
	values := make([]string, len(v.value))
	for i, value := range v.value {
		values[i] = strconv.FormatFloat(float64(value), 'g', -1, 32)
	}
	return "[" + strings.Join(values, ",") + "]"
}

func (v *Float32Slice) Type() string {
	return Float32SliceType
}

func (v *Float32Slice) Append(value string) error {
	d, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}
	v.value = append(v.value, float32(d))
	return nil
}

func (v *Float32Slice) Replace(values []string) error {
	if v == nil {
		return nil
	}
	out := make([]float32, len(values))
	for i, value := range values {
		d, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		out[i] = float32(d)
	}
	v.value = out
	return nil
}

func (v *Float32Slice) GetSlice() []string {
	if v == nil {
		return nil
	}
	out := make([]string, len(v.value))
	for i, value := range v.value {
		out[i] = strconv.FormatFloat(float64(value), 'g', -1, 32)
	}
	return out
}

func (v *Float32Slice) Value() []float32 {
	return v.value
}
