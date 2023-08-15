package glog

import (
	"errors"
	"strings"
)

type Level uint8

const (
	TRACE Level = iota + 1
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

func (l Level) String() string {
	switch l {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "INVALID"
}

// ParseString returns the Level matching the name s or returns an error
func ParseString(s string) (Level, error) {
	switch strings.ToUpper(s) {
	case "TRACE":
		return TRACE, nil
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	}
	return 0, errors.New("invalid Level \"" + s + "\"")
}

// MustParseString returns the Level matching the name s or panics
func MustParseString(s string) Level {
	l, e := ParseString(s)
	if e == nil {
		return l
	}
	panic(e)
}

func (l Level) color() color {
	switch l {
	case TRACE:
		return blue
	case DEBUG:
		return cyan
	case INFO:
		return green
	case WARN:
		return yellow
	case ERROR:
		return red
	case FATAL:
		return magenta
	default:
		return reset
	}
}
