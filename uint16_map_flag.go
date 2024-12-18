package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewUint16MapFlag(name, shorthand, usage string, value map[string]uint16) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint16Map(value),
	}
}

type Uint16Map struct {
	value   map[string]uint16
	changed bool
}

func NewUint16Map(value map[string]uint16) *Uint16Map {
	return &Uint16Map{
		value: value,
	}
}

func (m *Uint16Map) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseUint(kv[1], 10, 16)
		if err != nil {
			return err
		}
		if !m.changed {
			m.changed = true
			m.value = make(map[string]uint16)
		}
		m.value[kv[0]] = uint16(v)
	}
	return nil
}

func (m *Uint16Map) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%d", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *Uint16Map) Type() string {
	return Uint16MapType
}

func (m *Uint16Map) Value() map[string]uint16 {
	return m.value
}
