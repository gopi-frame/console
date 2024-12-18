package console

func NewStringFlag(name string, shorthand string, usage string, value string) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewString(value),
	}
}

type String string

func NewString(v string) *String {
	return (*String)(&v)
}

func (s *String) Set(v string) error {
	*s = String(v)
	return nil
}

func (s *String) String() string {
	return string(*s)
}

func (s *String) Type() string {
	return StringType
}

func (s *String) Value() string {
	return string(*s)
}
