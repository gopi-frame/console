package console

import (
	"strconv"
)

// NewBoolFlag returns a new bool flag with the specified name, default value, and usage string.
func NewBoolFlag(name string, shorthand string, usage string, value bool) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewBool(value),
		isBool:    true,
	}
}

// Bool is a flag.Value that accepts bool value.
type Bool bool

func NewBool(v bool) *Bool {
	return (*Bool)(&v)
}

func (b *Bool) String() string {
	return strconv.FormatBool(bool(*b))
}

func (b *Bool) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*b = Bool(v)
	return nil
}

func (b *Bool) Type() string {
	return BoolType
}

func (b *Bool) Value() bool {
	return bool(*b)
}
