package console

import "strconv"

// NewUintFlag returns a new uint flag with the specified name, default value, and usage string.
func NewUintFlag(name string, shorthand string, usage string, value uint) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint(value),
	}
}

type Uint uint

func NewUint(v uint) *Uint {
	return (*Uint)(&v)
}

func (u *Uint) Set(v string) error {
	i, err := strconv.ParseUint(v, 0, 32)
	if err == nil {
		*u = Uint(i)
	}
	*u = Uint(uint(i))
	return nil
}

func (u *Uint) String() string {
	return strconv.FormatUint(uint64(*u), 10)
}

func (u *Uint) Type() string {
	return UintType
}

func (u *Uint) Value() uint {
	return uint(*u)
}
