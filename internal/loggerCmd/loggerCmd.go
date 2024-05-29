package loggerCmd

import (
	"bytes"
	"fmt"
	"os/exec"
)

type LoggerOpts struct {
	PollFreq uint32
	Schedule uint16
}

// LoggerCmd represents and individual logging script/job in which collections of
// are represented as groves.
type LoggerCmd struct {
	Cmd            *exec.Cmd
	pollFreq       uint32
	schedule       uint16
	enablePolling  bool
	enableSchedule bool
}

func (c *LoggerCmd) Run() (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Cmd.Stdout = &stdout
	c.Cmd.Stderr = &stderr
	err := c.Cmd.Run()
	return stdout.String(), stderr.String(), err
}

func implement(cmd *exec.Cmd, opts ...Option) (*LoggerCmd, error) {
	loggerCmd := &LoggerCmd{Cmd: cmd, enablePolling: false, enableSchedule: false, pollFreq: 0, schedule: 0}
	for _, opt := range opts {
		opt(loggerCmd)
	}
	return loggerCmd, nil
}

func New(cmd exec.Cmd, opts *LoggerOpts) (*LoggerCmd, error) {
	if (&opts.PollFreq) == nil && (&opts.Schedule) == nil {
		loggercmd, err := implement(&cmd)
		if err != nil {
			fmt.Errorf("Error creating LoggingCmd: %v", err)
			return loggercmd, err
		}
		return loggercmd, err
	}

	if (&opts.PollFreq) != nil && (&opts.Schedule) == nil {
		loggerCmd, err := implement(&cmd,
			WithPollFreq(opts.PollFreq))
		if err != nil {
			fmt.Errorf("Error creating LoggingCmd: %v", err)
			return loggerCmd, err
		}
		return loggerCmd, nil
	}
	if (&opts.Schedule) != nil && (&opts.PollFreq) == nil {
		loggerCmd, err := implement(
			&cmd,
			WithSchedule(opts.Schedule),
		)
		if err != nil {
			fmt.Errorf("Error creating LoggingCmd: %v", err)
			return loggerCmd, err
		}
		return loggerCmd, nil
	}
	loggerCmd, err := implement(
		&cmd,
		WithPollFreq(opts.PollFreq),
		WithSchedule(opts.Schedule))
	if err != nil {
		fmt.Errorf("Error creating LoggingCmd: %v", err)
		return loggerCmd, err
	}
	return loggerCmd, nil
}

type Option = func(l *LoggerCmd)

func WithPollFreq(seconds uint32) Option {
	if seconds > 86399 {
		fmt.Print("Polling value should not exceed 86399 (24 hours), instead use schedule to span multiple days.")
		return func(l *LoggerCmd) {}
	}
	return func(l *LoggerCmd) {
		l.pollFreq = seconds
		l.enablePolling = true
	}
}

func WithSchedule(days uint16) Option {
	if days > 65534 {
		fmt.Print("Maximum number of days supported is 66534")
		return func(l *LoggerCmd) {}
	}
	return func(l *LoggerCmd) {
		l.schedule = days
		l.enableSchedule = true
	}
}

// Setters
func (l *LoggerCmd) SetPollFreq(seconds uint32) Option {
	if !l.enablePolling {
		fmt.Print("Polling is disable for this command, please enable and try again.")
		return func(l *LoggerCmd) {}
	}
	if seconds > 86399 {
		fmt.Print("Polling value should not exceed 86399 (24 hours), instead use schedule to span multiple days.")
		return func(l *LoggerCmd) {}
	}
	return func(l *LoggerCmd) {
		l.pollFreq = seconds
	}
}

func (l *LoggerCmd) SetSchedule(days uint16) Option {
	if !l.enableSchedule {
		fmt.Print("Schedule is disabled for this command, please enable and try again.")
		return func(l *LoggerCmd) {}
	}
	if days > 65534 {
		fmt.Print("Maximum number of days supported is 66534")
		return func(l *LoggerCmd) {}
	}
	return func(l *LoggerCmd) {
		l.schedule = days
	}
}

func (l *LoggerCmd) SetPollingEnabled(enabled bool) Option {
	if enabled {
		return func(l *LoggerCmd) {
			l.enableSchedule = enabled
		}
	}
	return func(l *LoggerCmd) {
		l.enablePolling = enabled
		l.pollFreq = 0
	}
}

func (l *LoggerCmd) SetScheduleEnabled(enabled bool) Option {
	if enabled {
		return func(l *LoggerCmd) {
			l.enableSchedule = enabled
		}
	}
	return func(l *LoggerCmd) {
		l.enableSchedule = enabled
		l.schedule = 0
	}
}
