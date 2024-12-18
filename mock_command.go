package console

import (
	"github.com/gopi-frame/contract/console"
)

type mockCommand struct {
	Command
	flags   []console.Flag
	handler func(input console.Input)
}

func newMockCommand(handler func(input console.Input), flags ...console.Flag) console.Command {
	return &mockCommand{
		handler: handler,
		flags:   flags,
	}
}

func (m mockCommand) Signature() string {
	return "mock"
}

func (m mockCommand) Description() string {
	return "This is a mock command"
}

func (m mockCommand) AddFlags(flags ...console.Flag) {
	m.flags = append(m.flags, flags...)
}

func (m mockCommand) Flags() []console.Flag {
	return m.flags
}

func (m mockCommand) SetHandler(handler func(input console.Input)) {
	m.handler = handler
}

func (m mockCommand) Handle(input console.Input) {
	m.handler(input)
}
