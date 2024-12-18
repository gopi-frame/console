package console

import "strconv"

// NewUint8Flag returns a new uint8 flag with the specified name, default value, and usage string.
func NewUint8Flag(name string, shorthand string, usage string, value uint8) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint8(value),
	}
}

// Uint8 is a flag.Value that accepts uint8 value.
type Uint8 uint8

func NewUint8(v uint8) *Uint8 {
	return (*Uint8)(&v)
}

func (v *Uint8) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 8)
	if err != nil {
		return err
	}
	*v = Uint8(i)
	return nil
}

func (v *Uint8) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *Uint8) Type() string {
	return Uint8Type
}

func (v *Uint8) Value() uint8 {
	return uint8(*v)
}
