package console

import (
	"fmt"
	"io"
	"time"

	"github.com/gookit/color"
	"github.com/gopi-frame/contract/console"
)

type OutputMode uint

var _ console.OutputMode = OutputMode(0)

func (o OutputMode) Has(mode console.OutputMode) bool {
	return o&(mode.(OutputMode)) != 0
}

func (o OutputMode) Append(mode console.OutputMode) console.OutputMode {
	return o | (mode.(OutputMode))
}

func (o OutputMode) Remove(mode console.OutputMode) console.OutputMode {
	return o &^ (mode.(OutputMode))
}

const (
	OutputModeNormal OutputMode = 1 << iota
	OutputModeDebug
	OutputModeSilent
	OutputModeANSI
)

type Output struct {
	writer io.Writer
	mode   console.OutputMode
}

func NewOutput(writer io.Writer, mode OutputMode) *Output {
	if writer == nil {
		writer = io.Discard
	}
	if mode == 0 {
		mode = OutputModeNormal
	}
	return &Output{
		writer: writer,
		mode:   mode,
	}
}

func (o *Output) GetMode() console.OutputMode {
	return o.mode
}

func (o *Output) WithMode(mode console.OutputMode) console.Output {
	return &Output{
		writer: o.writer,
		mode:   mode,
	}
}

func (o *Output) Write(p []byte) (n int, err error) {
	if o.mode.Has(OutputModeSilent) {
		return 0, nil
	}
	return o.writer.Write(p)
}

func (o *Output) WriteString(s string) (n int, err error) {
	return o.Write([]byte(s))
}

func (o *Output) Debug(s string) {
	if !o.mode.Has(OutputModeDebug) {
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	if o.mode.Has(OutputModeANSI) {
		timestamp = color.New(color.FgGreen, color.OpBold).Sprint(timestamp)
	}
	level := "DEBUG"
	if o.mode.Has(OutputModeANSI) {
		level = color.New(color.FgCyan, color.OpBold).Sprint(level)
	}
	_, _ = o.WriteString(fmt.Sprintf("[%s] [%s] => %s\n", timestamp, level, s))
}

func (o *Output) Debugf(format string, a ...any) {
	o.Debug(fmt.Sprintf(format, a...))
}

func (o *Output) Info(s string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	if o.mode.Has(OutputModeANSI) {
		timestamp = color.New(color.FgGreen, color.OpBold).Sprint(timestamp)
	}
	level := "INFO"
	if o.mode.Has(OutputModeANSI) {
		level = color.New(color.FgGreen, color.OpBold).Sprint(level)
	}
	_, _ = o.WriteString(fmt.Sprintf("[%s] [%s] => %s\n", timestamp, level, s))
}

func (o *Output) Infof(format string, a ...any) {
	o.Info(fmt.Sprintf(format, a...))
}

func (o *Output) Notice(s string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	if o.mode.Has(OutputModeANSI) {
		timestamp = color.New(color.FgGreen, color.OpBold).Sprint(timestamp)
	}
	level := "NOTICE"
	if o.mode.Has(OutputModeANSI) {
		level = color.New(color.FgLightCyan, color.OpBold).Sprint(level)
	}
	_, _ = o.WriteString(fmt.Sprintf("[%s] [%s] => %s\n", timestamp, level, s))
}

func (o *Output) Noticef(format string, a ...any) {
	o.Notice(fmt.Sprintf(format, a...))
}

func (o *Output) Warn(s string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	if o.mode.Has(OutputModeANSI) {
		timestamp = color.New(color.FgGreen, color.OpBold).Sprint(timestamp)
	}
	level := "WARN"
	if o.mode.Has(OutputModeANSI) {
		level = color.New(color.FgYellow, color.OpBold).Sprint(level)
	}
	color.Print()
	_, _ = o.WriteString(fmt.Sprintf("[%s] [%s] => %s\n", timestamp, level, s))
}

func (o *Output) Warnf(format string, a ...any) {
	o.Warn(fmt.Sprintf(format, a...))
}

func (o *Output) Error(s string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	if o.mode.Has(OutputModeANSI) {
		timestamp = color.New(color.FgGreen, color.OpBold).Sprint(timestamp)
	}
	level := "ERROR"
	if o.mode.Has(OutputModeANSI) {
		level = color.New(color.FgRed, color.OpBold).Sprint(level)
	}
	_, _ = o.WriteString(fmt.Sprintf("[%s] [%s] => %s\n", timestamp, level, s))
}

func (o *Output) Errorf(format string, a ...any) {
	o.Error(fmt.Sprintf(format, a...))
}

func (o *Output) Success(s string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	if o.mode.Has(OutputModeANSI) {
		timestamp = color.New(color.FgGreen, color.OpBold).Sprint(timestamp)
	}
	level := "SUCCESS"
	if o.mode.Has(OutputModeANSI) {
		level = color.New(color.FgLightGreen, color.OpBold).Sprint(level)
	}
	_, _ = o.WriteString(fmt.Sprintf("[%s] [%s] => %s\n", timestamp, level, s))
}

func (o *Output) Successf(format string, a ...any) {
	o.Success(fmt.Sprintf(format, a...))
}

func (o *Output) Fail(s string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	if o.mode.Has(OutputModeANSI) {
		timestamp = color.New(color.FgGreen, color.OpBold).Sprint(timestamp)
	}
	level := "FAILURE"
	if o.mode.Has(OutputModeANSI) {
		level = color.New(color.FgRed, color.OpBold).Sprint(level)
	}
	_, _ = o.WriteString(fmt.Sprintf("[%s] [%s] => %s\n", timestamp, level, s))
}

func (o *Output) Failf(format string, a ...any) {
	o.Fail(fmt.Sprintf(format, a...))
}
