package log

import (
	"io"
	internal_logger "log"
	"os"
)

// Wrapper for some niceties

// Logger Levels
const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

// DEFAULT FLAG prefixes messages such as: 2015/07/25 20:05:51 log.go:118:
const DEFAULTFLAG = internal_logger.Ldate | internal_logger.Ltime | internal_logger.Lshortfile

// Logger is a wrapper around the std.log package. It allows for more
// granular levels, as well as some formatting niceties.
type Logger interface {
	Fatal(interface{})
	Error(interface{})
	Warn(interface{})
	Info(interface{})
	Debug(interface{})
	Trace(interface{})

	Level(int)
	SetFlags(int)
	SetOutput(io.Writer)
}

// Log struct
type Log struct {
	level int
	flag  int
}

// NewLogger instantiates a new logging object.
func NewLogger(output io.Writer, level int) *Log {
	log := &Log{
		level: level,
		flag:  DEFAULTFLAG,
	}

	log.SetOutput(output)
	log.SetFlags(log.flag)
	return log
}

var std = NewLogger(os.Stderr, INFO)

// Fatal is the highest level log. Results in a panic.
func (log *Log) Fatal(message ...interface{}) {
	log.Println("FATAL:", message)
	panic(message)
}

// Fatal is the highest level log. Results in a panic.
func Fatal(message ...interface{}) {
	std.Fatal(message...)
}

// Error is for conditions that are bad, but don't warrant a panic.
func (log *Log) Error(message ...interface{}) {
	if log.level <= ERROR {
		log.Println("ERROR:", message)
	}
}

// Error is for conditions that are bad, but don't warrant a panic.
func Error(message ...interface{}) {
	std.Error(message...)
}

// Warn is for those mildly bad conditions.
func (log *Log) Warn(message ...interface{}) {
	if log.level <= WARN {
		log.Println("WARN:", message)
	}
}

// Warn is for those mildly bad conditions.
func Warn(message ...interface{}) {
	std.Warn(message...)
}

// Info is for potentially noisy messages that you would like to see in production.
func (log *Log) Info(message ...interface{}) {
	if log.level <= INFO {
		log.Println("INFO:", message)
	}
}

// Info is for potentially noisy messages that you would like to see in production.
func Info(message ...interface{}) {
	std.Info(message...)
}

// Debug is for messages that can be helpful in diagnosting problems.
func (log *Log) Debug(message ...interface{}) {
	if log.level <= DEBUG {
		log.Println("DEBUG:", message)
	}
}

// Debug is for messages that can be helpful in diagnosting problems.
func Debug(message ...interface{}) {
	std.Debug(message...)
}

// Trace is for noisy messages that you inserted, because your favorite debugger is log statements.
func (log *Log) Trace(message ...interface{}) {
	if log.level <= TRACE {
		log.Println("TRACE:", message)
	}
}

// Trace is for noisy messages that you inserted, because your favorite debugger is log statements.
func Trace(message ...interface{}) {
	std.Trace(message...)
}

// Println is a compatibility function so that you can drop in this library where you used to use the std.log
func (log *Log) Println(message ...interface{}) {
	internal_logger.Println(message)
}

// Println is a compatibility function so that you can drop in this library where you used to use the std.log
func Println(message ...interface{}) {
	std.Println(message)
}

// Level sets the lowest level to log
func (log *Log) Level(level int) {
	log.level = level
}

// Level sets the lowest level to log
func Level(level int) {
	std.Level(level)
}

// SetFlags sets the standard log information that is being prefixed
func (log *Log) SetFlags(flag int) {
	internal_logger.SetFlags(flag)
}

// SetFlags sets the standard log information that is being prefixed
func SetFlags(flag int) {
	std.SetFlags(flag)
}

// SetOutput allows you to redirect the output.
func (log *Log) SetOutput(w io.Writer) {
	internal_logger.SetOutput(w)
}

// SetOutput allows you to redirect the output.
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}
