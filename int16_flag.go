package console

import "strconv"

func NewInt16Flag(name string, shorthand string, usage string, value int16) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt16(value),
	}
}

type Int16 int16

func NewInt16(v int16) *Int16 {
	return (*Int16)(&v)
}

func (v *Int16) Set(s string) error {
	i, err := strconv.ParseInt(s, 0, 16)
	if err != nil {
		return err
	}
	*v = Int16(i)
	return nil
}

func (v *Int16) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *Int16) Type() string {
	return Int16Type
}

func (v *Int16) Value() int16 {
	return int16(*v)
}
