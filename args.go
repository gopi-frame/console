package console

import (
	"fmt"

	"github.com/gopi-frame/contract/console"
	"github.com/gopi-frame/exception"
)

// PossibleArgs is a function that returns an error if the input does not match the expected arguments.
type PossibleArgs = func(input console.Input) error

// AnyArgs never returns an error.
func AnyArgs() PossibleArgs {
	return func(input console.Input) error {
		return nil
	}
}

// NoArgs returns an error if the input has any arguments.
func NoArgs() PossibleArgs {
	return func(input console.Input) error {
		if len(input.Args()) > 0 {
			return exception.New(fmt.Sprintf("excepted no args, but got %d", len(input.Args())))
		}
		return nil
	}
}

// ExactArgs returns an error if the input has a different number of arguments.
func ExactArgs(n int) PossibleArgs {
	return func(input console.Input) error {
		if len(input.Args()) != n {
			return exception.New(fmt.Sprintf("excepted %d args, but got %d", n, len(input.Args())))
		}
		return nil
	}
}

// MinArgs returns an error if the input has fewer arguments than the expected minimum.
func MinArgs(n int) PossibleArgs {
	return func(input console.Input) error {
		if len(input.Args()) < n {
			return exception.New(fmt.Sprintf("excepted at least %d args, but got %d", n, len(input.Args())))
		}
		return nil
	}
}

// MaxArgs returns an error if the input has more arguments than the expected maximum.
func MaxArgs(n int) PossibleArgs {
	return func(input console.Input) error {
		if len(input.Args()) > n {
			return exception.New(fmt.Sprintf("excepted %d args, but got %d", n, len(input.Args())))
		}
		return nil
	}
}

// BetweenArgs returns an error if the input has a different number of arguments than the expected range.
func BetweenArgs(min, max int) PossibleArgs {
	return func(input console.Input) error {
		if len(input.Args()) < min || len(input.Args()) > max {
			return exception.New(fmt.Sprintf("excepted %d to %d args, but got %d", min, max, len(input.Args())))
		}
		return nil
	}
}
