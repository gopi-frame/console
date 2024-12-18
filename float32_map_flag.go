package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewFloat32MapFlag(name, shorthand, usage string, value map[string]float32) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewFloat32Map(value),
	}
}

type Float32Map struct {
	value   map[string]float32
	changed bool
}

func NewFloat32Map(value map[string]float32) *Float32Map {
	return &Float32Map{
		value: value,
	}
}

func (m *Float32Map) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseFloat(kv[1], 32)
		if err != nil {
			return err
		}
		if !m.changed {
			m.changed = true
			m.value = make(map[string]float32)
		}
		m.value[kv[0]] = float32(v)
	}
	return nil
}

func (m *Float32Map) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%s", k, strconv.FormatFloat(float64(v), 'f', -1, 32)))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *Float32Map) Type() string {
	return Float32MapType
}

func (m *Float32Map) Value() map[string]float32 {
	return m.value
}
