package console

import "strconv"

func NewUint64Flag(name string, shorthand string, usage string, value uint64) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint64(value),
	}
}

type Uint64 uint64

func NewUint64(v uint64) *Uint64 {
	return (*Uint64)(&v)
}

func (v *Uint64) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return err
	}
	*v = Uint64(i)
	return nil
}

func (v *Uint64) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *Uint64) Type() string {
	return Uint64Type
}

func (v *Uint64) Value() uint64 {
	return uint64(*v)
}
