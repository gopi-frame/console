package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewUint32MapFlag(name, shorthand, usage string, value map[string]uint32) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint32Map(value),
	}
}

type Uint32Map struct {
	value   map[string]uint32
	changed bool
}

func NewUint32Map(value map[string]uint32) *Uint32Map {
	return &Uint32Map{
		value: value,
	}
}

func (m *Uint32Map) Set(value string) error {
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
			m.value = make(map[string]uint32)
		}
		m.value[kv[0]] = uint32(v)
	}
	return nil
}

func (m *Uint32Map) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%d", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *Uint32Map) Type() string {
	return Uint32MapType
}

func (m *Uint32Map) Value() map[string]uint32 {
	return m.value
}
