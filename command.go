package console

import "github.com/gopi-frame/contract/console"

// Command is an abstract command implementation of [console.Command].
// [console.Command.Name], [console.Command.Handle] methods must be implemented.
// Rewrite other methods of [console.Command] if necessarily.
type Command struct {
	console.Command
}

func (c *Command) Group() string {
	return ""
}

// Help returns the help information of the command, default is empty string.
func (c *Command) Help() string {
	return ""
}

// Description returns the description of the command, default is empty string.
func (c *Command) Description() string {
	return ""
}

// Example returns the example of the command, default is empty string.
func (c *Command) Example() string {
	return ""
}

// Args returns the excepted arguments of the command, default is [AnyArgs].
func (c *Command) Args() PossibleArgs {
	return AnyArgs()
}

// Flags returns the flags of the command, default is nil.
func (c *Command) Flags() []console.Flag {
	return nil
}

// PersistentFlags returns the persistent flags of the command, default is nil.
func (c *Command) PersistentFlags() []console.Flag {
	return nil
}

func (c *Command) SubCommands() []console.Command {
	return nil
}
