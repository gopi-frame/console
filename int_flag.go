package console

import (
	"strconv"
)

// NewIntFlag returns a new int flag with the specified name, default value, and usage string.
func NewIntFlag(name string, shorthand string, usage string, value int) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt(value),
	}
}

// Int is a flag.Value that accepts int value.
type Int int

func NewInt(v int) *Int {
	return (*Int)(&v)
}

func (i *Int) String() string {
	return strconv.Itoa(int(*i))
}

func (i *Int) Set(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*i = Int(v)
	return nil
}

func (i *Int) Type() string {
	return IntType
}

func (i *Int) Value() int {
	return int(*i)
}
