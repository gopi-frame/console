package console

import (
	"fmt"
	"strings"
)

func NewStringMapFlag(name, shorthand, usage string, value map[string]string) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewStringMap(value),
	}
}

type StringMap struct {
	value   map[string]string
	changed bool
}

func NewStringMap(value map[string]string) *StringMap {
	return &StringMap{
		value: value,
	}
}

func (s *StringMap) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		if !s.changed {
			s.changed = true
			s.value = make(map[string]string)
		}
		s.value[kv[0]] = kv[1]
	}
	return nil
}

func (s *StringMap) String() string {
	out := make([]string, len(s.value))
	for k, v := range s.value {
		out = append(out, fmt.Sprintf("%s=%s", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (s *StringMap) Type() string {
	return StringMapType
}

func (s *StringMap) Value() map[string]string {
	return s.value
}
