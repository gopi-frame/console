package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewIntMapFlag(name, shorthand, usage string, value map[string]int) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewIntMap(value),
	}
}

type IntMap struct {
	value   map[string]int
	changed bool
}

func NewIntMap(value map[string]int) *IntMap {
	return &IntMap{
		value: value,
	}
}

func (m *IntMap) Set(value string) error {
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
			m.value = make(map[string]int)
		}
		m.value[kv[0]] = int(v)
	}
	return nil
}

func (m *IntMap) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%d", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *IntMap) Type() string {
	return IntMapType
}

func (m *IntMap) Value() map[string]int {
	return m.value
}
