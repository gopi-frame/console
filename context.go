package console

import (
	"context"
	"os"

	"github.com/gopi-frame/contract/console"
)

type contextKey string

const (
	contextKeyOutput contextKey = "output"
)

func WithOutput(ctx context.Context, output console.Output) context.Context {
	return context.WithValue(ctx, contextKeyOutput, output)
}

func GetOutput(ctx context.Context) console.Output {
	output, ok := ctx.Value(contextKeyOutput).(console.Output)
	if !ok {
		return NewOutput(os.Stderr, OutputModeNormal|OutputModeANSI)
	}
	return output
}
