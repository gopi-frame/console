package console

import "strconv"

func NewInt32Flag(name string, shorthand string, usage string, value int32) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt32(value),
	}
}

type Int32 int32

func NewInt32(f int32) *Int32 {
	return (*Int32)(&f)
}

func (i *Int32) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return err
	}
	*i = Int32(v)
	return nil
}

func (i *Int32) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

func (i *Int32) Type() string {
	return Int32Type
}

func (i *Int32) Value() int32 {
	return int32(*i)
}
