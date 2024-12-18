package console

import "time"

func NewDurationFlag(name string, shorthand string, usage string, value time.Duration) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewDuration(value),
	}
}

type Duration time.Duration

func NewDuration(value time.Duration) *Duration {
	return (*Duration)(&value)
}

func (d *Duration) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*d = Duration(v)
	return nil
}

func (d *Duration) String() string {
	return time.Duration(*d).String()
}

func (d *Duration) Type() string {
	return DurationType
}

func (d *Duration) Value() time.Duration {
	return time.Duration(*d)
}
