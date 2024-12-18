package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewUint8MapFlag(name, shorthand, usage string, value map[string]uint8) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint8Map(value),
	}
}

type Uint8Map struct {
	value   map[string]uint8
	changed bool
}

func NewUint8Map(value map[string]uint8) *Uint8Map {
	return &Uint8Map{
		value: value,
	}
}

func (m *Uint8Map) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseUint(kv[1], 10, 8)
		if err != nil {
			return err
		}
		if !m.changed {
			m.changed = true
			m.value = make(map[string]uint8)
		}
		m.value[kv[0]] = uint8(v)
	}
	return nil
}

func (m *Uint8Map) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%d", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *Uint8Map) Type() string {
	return Uint8MapType
}

func (m *Uint8Map) Value() map[string]uint8 {
	return m.value
}
