package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewInt64MapFlag(name, shorthand, usage string, value map[string]int64) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewInt64Map(value),
	}
}

type Int64Map struct {
	value   map[string]int64
	changed bool
}

func NewInt64Map(value map[string]int64) *Int64Map {
	return &Int64Map{
		value: value,
	}

}

func (m *Int64Map) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseInt(kv[1], 10, 64)
		if err != nil {
			return err
		}
		if !m.changed {
			m.changed = true
			m.value = make(map[string]int64)
		}
		m.value[kv[0]] = v
	}
	return nil
}

func (m *Int64Map) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%d", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *Int64Map) Type() string {
	return Int64MapType
}

func (m *Int64Map) Value() map[string]int64 {
	return m.value
}
