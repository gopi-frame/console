package console

import "strconv"

func NewInt8Flag(name string, shorthand string, usage string, value int8) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt8(value),
	}
}

type Int8 int8

func NewInt8(v int8) *Int8 {
	return (*Int8)(&v)
}

func (v *Int8) Set(s string) error {
	i, err := strconv.ParseInt(s, 0, 8)
	if err != nil {
		return err
	}
	*v = Int8(i)
	return nil
}

func (v *Int8) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *Int8) Type() string {
	return Int8Type
}

func (v *Int8) Value() int8 {
	return int8(*v)
}
