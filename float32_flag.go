package console

import "strconv"

func NewFloat32Flag(name string, shorthand string, usage string, value float32) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewFloat32(value),
	}
}

type Float32 float32

func NewFloat32(v float32) *Float32 {
	return (*Float32)(&v)
}

func (v *Float32) Set(s string) error {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	*v = Float32(f)
	return nil
}

func (v *Float32) String() string {
	return strconv.FormatFloat(float64(*v), 'g', -1, 32)
}

func (v *Float32) Type() string {
	return Float32Type
}

func (v *Float32) Value() float32 {
	return float32(*v)
}
