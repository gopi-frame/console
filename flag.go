package console

import (
	"github.com/gopi-frame/contract/console"
)

type Flag struct {
	name      string
	shorthand string
	usage     string
	value     console.Value
	hidden    bool
	isBool    bool
}

func (f *Flag) SetUsage(usage string) {
	f.usage = usage
}

func (f *Flag) SetValue(value console.Value) {
	f.value = value
}

func (f *Flag) SetHidden(hidden bool) {
	f.hidden = hidden
}

func (f *Flag) Name() string {
	return f.name
}

func (f *Flag) Shorthand() string {
	return f.shorthand
}

func (f *Flag) Usage() string {
	return f.usage
}

func (f *Flag) Type() string {
	return f.value.Type()
}

func (f *Flag) Value() console.Value {
	return f.value
}

func (f *Flag) Hidden() bool {
	return f.hidden
}

func (f *Flag) IsBool() bool {
	return f.isBool
}
