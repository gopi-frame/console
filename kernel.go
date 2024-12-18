package console

import (
	"os"
	"path/filepath"

	"github.com/gopi-frame/collection/list"
	"github.com/gopi-frame/contract/console"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Kernel struct {
	*cobra.Command
	commands *list.List[console.Command]
	output   console.Kernel
}

func NewKernel() *Kernel {
	name := filepath.Base(os.Args[0])
	ext := filepath.Ext(name)
	k := &Kernel{
		Command: &cobra.Command{
			Use: name[:len(name)-len(ext)],
		},
		commands: list.NewList[console.Command](),
	}
	return k
}

func (k *Kernel) SetName(name string) {
	k.Command.Use = name
}

func (k *Kernel) AddGroup(id, name string, fn func(group console.Group)) {
	group := NewGroup(id, name)
	k.Command.AddGroup(&cobra.Group{
		ID:    id,
		Title: name,
	})
	if fn != nil {
		fn(group)
	}
	for _, command := range group.Commands() {
		k.AddCommand(WithGroup(group.ID(), command))
	}
}

func (k *Kernel) AddFlag(flag console.Flag) {
	k.Command.Flags().AddFlag(buildFlag(flag))
}

func (k *Kernel) AddPersistentFlag(flag console.Flag) {
	k.Command.PersistentFlags().AddFlag(buildFlag(flag))
}

func (k *Kernel) AddCommand(command console.Command) {
	k.commands.Lock()
	defer k.commands.Unlock()
	k.commands.Push(command)
	cmd := &cobra.Command{
		Use:     command.Signature(),
		Short:   command.Description(),
		Long:    command.Help(),
		Example: command.Example(),
		GroupID: command.Group(),
		Run: func(cmd *cobra.Command, args []string) {
			input := NewInput(cmd.Context(), cmd.Flags())
			command.Handle(input)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			input := NewInput(cmd.Context(), cmd.Flags())
			return command.Args()(input)
		},
	}
	for _, flag := range command.Flags() {
		cmd.Flags().AddFlag(buildFlag(flag))
	}
	for _, flag := range command.PersistentFlags() {
		cmd.PersistentFlags().AddFlag(buildFlag(flag))
	}
	k.Command.AddCommand(cmd)
}

func (k *Kernel) Call(cmd string, args ...string) error {
	args = append([]string{cmd}, args...)
	k.SetArgs(args)
	return k.Execute()
}

func (k *Kernel) Run() error {
	return k.Execute()
}

func buildFlag(flag console.Flag) *pflag.Flag {
	f := &pflag.Flag{
		Name:      flag.Name(),
		Shorthand: flag.Shorthand(),
		Usage:     flag.Usage(),
		Value:     flag.Value(),
		DefValue:  flag.Value().String(),
		Hidden:    flag.Hidden(),
	}
	if flag.IsBool() {
		f.NoOptDefVal = "true"
	}
	return f
}
