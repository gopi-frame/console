package console

import "strconv"

func NewFloat64Flag(name string, shorthand string, usage string, value float64) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewFloat64(value),
	}
}

type Float64 float64

func NewFloat64(v float64) *Float64 {
	return (*Float64)(&v)
}

func (f *Float64) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*f = Float64(v)
	return nil
}

func (f *Float64) String() string {
	return strconv.FormatFloat(float64(*f), 'g', -1, 64)
}

func (f *Float64) Type() string {
	return Float64Type
}

func (f *Float64) Value() float64 {
	return float64(*f)
}
