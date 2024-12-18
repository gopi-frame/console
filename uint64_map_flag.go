package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewUint64MapFlag(name, shorthand, usage string, value map[string]uint64) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewUint64Map(value),
	}
}

type Uint64Map struct {
	value   map[string]uint64
	changed bool
}

func NewUint64Map(value map[string]uint64) *Uint64Map {
	return &Uint64Map{
		value: value,
	}
}

func (m *Uint64Map) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseUint(kv[1], 10, 64)
		if err != nil {
			return err
		}
		if !m.changed {
			m.changed = true
			m.value = make(map[string]uint64)
		}
		m.value[kv[0]] = v
	}
	return nil
}

func (m *Uint64Map) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%d", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *Uint64Map) Type() string {
	return Uint64MapType
}

func (m *Uint64Map) Value() map[string]uint64 {
	return m.value
}
