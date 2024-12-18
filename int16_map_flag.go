package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewInt16MapFlag(name, shorthand, usage string, value map[string]int16) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt16Map(value),
	}
}

type Int16Map struct {
	value   map[string]int16
	changed bool
}

func NewInt16Map(value map[string]int16) *Int16Map {
	return &Int16Map{
		value: value,
	}
}

func (m *Int16Map) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseInt(kv[1], 10, 16)
		if err != nil {
			return err
		}
		if !m.changed {
			m.value = make(map[string]int16)
			m.changed = true
		}
		m.value[kv[0]] = int16(v)
	}
	return nil
}

func (m *Int16Map) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%d", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *Int16Map) Type() string {
	return Int16MapType
}

func (m *Int16Map) Value() map[string]int16 {
	return m.value
}
