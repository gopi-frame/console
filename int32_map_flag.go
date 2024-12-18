package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewInt32MapFlag(name, shorthand, usage string, value map[string]int32) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt32Map(value),
	}
}

type Int32Map struct {
	value   map[string]int32
	changed bool
}

func NewInt32Map(value map[string]int32) *Int32Map {
	return &Int32Map{
		value: value,
	}
}

func (m *Int32Map) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseInt(kv[1], 10, 32)
		if err != nil {
			return err
		}
		if !m.changed {
			m.changed = true
			m.value = make(map[string]int32)
		}
		m.value[kv[0]] = int32(v)
	}
	return nil
}

func (m *Int32Map) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%d", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *Int32Map) Type() string {
	return Int32MapType
}

func (m *Int32Map) Value() map[string]int32 {
	return m.value
}
