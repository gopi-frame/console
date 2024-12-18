package console

import "strconv"

func NewUint32Flag(name string, shorthand string, usage string, value uint32) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint32(value),
	}
}

type Uint32 uint32

func NewUint32(value uint32) *Uint32 {
	return (*Uint32)(&value)
}

func (v *Uint32) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 32)
	if err != nil {
		return err
	}
	*v = Uint32(i)
	return nil
}

func (v *Uint32) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *Uint32) Type() string {
	return Uint32Type
}

func (v *Uint32) Value() uint32 {
	return uint32(*v)
}
