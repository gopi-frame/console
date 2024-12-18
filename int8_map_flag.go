package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewInt8MapFlag(name, shorthand, usage string, value map[string]int8) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt8Map(value),
	}
}

type Int8Map struct {
	value   map[string]int8
	changed bool
}

func NewInt8Map(value map[string]int8) *Int8Map {
	return &Int8Map{
		value: value,
	}
}

func (m *Int8Map) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseInt(kv[1], 10, 8)
		if err != nil {
			return err
		}
		if !m.changed {
			m.value = make(map[string]int8)
			m.changed = true
		}
		m.value[kv[0]] = int8(v)
	}
	return nil
}

func (m *Int8Map) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%d", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *Int8Map) Type() string {
	return Int8MapType
}

func (m *Int8Map) Value() map[string]int8 {
	return m.value
}
