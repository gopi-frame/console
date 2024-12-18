package console

import "github.com/gopi-frame/contract/console"

type groupedCommand struct {
	console.Command

	groupID string
}

func (g *groupedCommand) Group() string {
	return g.groupID
}

// WithGroup returns a new command with the specified group ID.
func WithGroup(groupID string, cmd console.Command) console.Command {
	return &groupedCommand{
		Command: cmd,
		groupID: groupID,
	}
}

type CommandGroup struct {
	id   string
	name string
	cmds []console.Command
}

func NewGroup(id, name string, cmds ...console.Command) console.Group {
	return &CommandGroup{
		id:   id,
		name: name,
		cmds: cmds,
	}
}

func (g *CommandGroup) ID() string {
	return g.id
}

func (g *CommandGroup) Name() string {
	return g.name
}

func (g *CommandGroup) Commands() []console.Command {
	return g.cmds
}

func (h *CommandGroup) AddCommand(cmd console.Command) {
	h.cmds = append(h.cmds, cmd)
}
