package console

import "strconv"

func NewInt64Flag(name string, shorthand string, usage string, value int64) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt64(value),
	}
}

type Int64 int64

func NewInt64(v int64) *Int64 {
	return (*Int64)(&v)
}

func (v *Int64) Set(s string) error {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}
	*v = Int64(i)
	return nil
}

func (v *Int64) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *Int64) Type() string {
	return Int64Type
}

func (v *Int64) Value() int64 {
	return int64(*v)
}
