package console

import (
	"fmt"
	"strconv"
	"strings"
)

func NewBoolMapFlag(name, shorthand, usage string, value map[string]bool) *Flag {
	return &Flag{
		name:      name,
		shorthand: shorthand,
		usage:     usage,
		value:     NewBoolMap(value),
	}
}

type BoolMap struct {
	value   map[string]bool
	changed bool
}

func NewBoolMap(value map[string]bool) *BoolMap {
	return &BoolMap{
		value: value,
	}
}

func (b *BoolMap) Set(value string) error {
	ss := strings.Split(value, ",")
	for _, p := range ss {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", value)
		}
		v, err := strconv.ParseBool(kv[1])
		if err != nil {
			return err
		}
		if !b.changed {
			b.changed = true
			b.value = make(map[string]bool)
		}
		b.value[kv[0]] = v
	}
	return nil
}

func (b *BoolMap) String() string {
	out := make([]string, len(b.value))
	for k, v := range b.value {
		out = append(out, fmt.Sprintf("%s=%t", k, v))
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (b *BoolMap) Type() string {
	return BoolMapType
}

func (b *BoolMap) Value() map[string]bool {
	return b.value
}
