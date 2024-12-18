package console

import "strings"

func NewStringSliceFlag(name, shorthand, usage string, value []string) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewStringSlice(value...),
	}
}

type StringSlice struct {
	value   []string
	changed bool
}

func NewStringSlice(values ...string) *StringSlice {
	return &StringSlice{
		value: values,
	}
}

func (v *StringSlice) Set(value string) error {
	if !v.changed {
		v.value = []string{value}
		v.changed = true
		return nil
	}
	v.value = append(v.value, value)
	return nil
}

func (v *StringSlice) String() string {
	if v == nil {
		return "[]"
	}
	return "[" + strings.Join(v.value, ",") + "]"
}

func (v *StringSlice) Type() string {
	return StringSliceType
}

func (v *StringSlice) Append(value string) error {
	v.value = append(v.value, value)
	return nil
}

func (v *StringSlice) Replace(value []string) error {
	v.value = value
	return nil
}

func (v *StringSlice) GetSlice() []string {
	return v.value
}

func (v *StringSlice) Value() []string {
	return v.value
}
