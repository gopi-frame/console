package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewUintMapFlag(name, shorthand, usage string, value map[string]uint) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUintMap(value),
	}
}

type UintMap struct {
	value   map[string]uint
	changed bool
}

func NewUintMap(value map[string]uint) *UintMap {
	return &UintMap{
		value: value,
	}
}

func (m *UintMap) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseUint(kv[1], 10, 32)
		if err != nil {
			return err
		}
		if !m.changed {
			m.changed = true
			m.value = make(map[string]uint)
		}
		m.value[kv[0]] = uint(v)
	}
	return nil
}

func (m *UintMap) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%d", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *UintMap) Type() string {
	return UintMapType
}

func (m *UintMap) Value() map[string]uint {
	return m.value
}
