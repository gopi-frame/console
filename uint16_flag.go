package console

import "strconv"

// NewUint16Flag returns a new uint16 flag with the specified name, default value, and usage string.
func NewUint16Flag(name string, shorthand string, usage string, value uint16) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint16(value),
	}
}

// Uint16 is a flag.Value that accepts uint16 value.
type Uint16 uint16

func NewUint16(v uint16) *Uint16 {
	return (*Uint16)(&v)
}

func (v *Uint16) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 16)
	if err != nil {
		return err
	}
	*v = Uint16(i)
	return nil
}

func (v *Uint16) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *Uint16) Type() string {
	return Uint16Type
}

func (v *Uint16) Value() uint16 {
	return uint16(*v)
}
