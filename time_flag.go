package console

import "time"

func NewTimeFlag(name string, shorthand string, usage string, layout string, value time.Time) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewTime(value, layout),
	}
}

type Time struct {
	value  time.Time
	layout string
}

func NewTime(value time.Time, layout string) *Time {
	return &Time{
		value:  value,
		layout: layout,
	}
}

func (t *Time) Set(s string) error {
	v, err := time.Parse(t.layout, s)
	if err != nil {
		return err
	}
	*t = Time{
		value:  v,
		layout: t.layout,
	}
	return nil
}

func (t *Time) String() string {
	return t.value.Format(t.layout)
}

func (t *Time) Type() string {
	return TimeType
}

func (t *Time) Value() time.Time {
	return t.value
}
