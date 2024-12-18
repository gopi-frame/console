package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewFloat64MapFlag(name, shorthand, usage string, value map[string]float64) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewFloat64Map(value),
	}
}

type Float64Map struct {
	value   map[string]float64
	changed bool
}

func NewFloat64Map(value map[string]float64) *Float64Map {
	return &Float64Map{
		value: value,
	}
}

func (m *Float64Map) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseFloat(kv[1], 64)
		if err != nil {
			return err
		}
		if !m.changed {
			m.changed = true
			m.value = make(map[string]float64)
		}
		m.value[kv[0]] = v
	}
	return nil
}

func (m *Float64Map) String() string {
	out := make([]string, len(m.value))
	for k, v := range m.value {
		out = append(out, fmt.Sprintf("%s=%s", k, strconv.FormatFloat(v, 'f', -1, 64)))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (m *Float64Map) Type() string {
	return Float64MapType
}

func (m *Float64Map) Value() map[string]float64 {
	return m.value
}
