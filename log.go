package glog

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type lineSerializer interface {
	toText() string
	toJSON() []byte
}

type line struct {
	level     Level
	tags      []string
	ts        string
	color     ColorOpt
	lineNum   bool
	args      []any
	multiline bool
}

type fmtline struct {
	fmt string
	line
}

func unwrapErrorToString(e error, multiline bool) string {
	var errs []string
	err := e
	for true {
		if err == nil {
			break
		}
		errs = append(errs, err.Error())
		err = errors.Unwrap(err)
	}
	sep := " > "
	if multiline {
		sep = "\n> "
	}
	return strings.Join(errs, sep)
}

func unwrapErrorToJSON(e error) []byte {
	if a, ok := e.(json.Marshaler); ok {
		b, _ := a.MarshalJSON()
		return b
	}
	// TODO: unwrap
	return []byte(e.Error())
}

func argToString(arg any, multiline bool) string {
	if a, ok := arg.(string); ok {
		return a
	}
	if a, ok := arg.(fmt.Stringer); ok {
		return a.String()
	}
	if a, ok := arg.(error); ok {
		return unwrapErrorToString(a, multiline)
	}
	return fmt.Sprint(arg)
}

func argToJSON(arg any) []byte {
	if a, ok := arg.(error); ok {
		if _, ok := a.(json.Marshaler); !ok {
			return []byte(a.Error())
		}
	}
	b, _ := json.Marshal(arg)
	return b
}

func (l line) prefix() string {
	s := strings.Builder{}
	lvl := l.level.String()
	if l.color == ColorLogLevel {
		l.level.color().apply(lvl)
	}
	s.WriteString(l.level.String() + " ")
	if l.ts != "" {
		s.WriteString(time.Now().Format(l.ts) + " ")
	}
	if l.lineNum {
		s.WriteString("[LINE NUM]")
	}
	s.WriteString(": ")
	prefix := s.String()
	if l.color == ColorPrefix {
		l.level.color().apply(prefix)
	}
	return prefix
}

func (l line) toText() string {
	msg := l.prefix() + fmt.Sprint(l.args...)
	if l.color == ColorAll {
		l.level.color().apply(msg)
	}
	return msg
}

func (fl fmtline) toText() string {
	msg := fl.prefix() + fmt.Sprintf(fl.fmt, fl.args...)
	if fl.color == ColorAll {
		fl.level.color().apply(msg)
	}
	return msg
}

func (l line) toJSON() []byte {
	jsonMap := map[string]any{
		"timestamp": l.ts,
		"level":     l.level,
		"tags":      l.tags,
		"message":   fmt.Sprint(l.args...),
	}
	if l.lineNum {
		jsonMap["line"] = "[LINE NUM]"
	}
	var bs []byte
	if l.multiline {
		bs, _ = json.MarshalIndent(jsonMap, "", "\t")
	} else {
		bs, _ = json.Marshal(jsonMap)
	}
	return bs
}

func (fl fmtline) toJSON() []byte {
	jsonMap := map[string]any{
		"timestamp": fl.ts,
		"level":     fl.level,
		"tags":      fl.tags,
		"message":   fmt.Sprintf(fl.fmt, fl.args...),
	}
	if fl.lineNum {
		jsonMap["line"] = "[LINE NUM]"
	}
	var bs []byte
	if fl.multiline {
		bs, _ = json.MarshalIndent(jsonMap, "", "\t")
	} else {
		bs, _ = json.Marshal(jsonMap)
	}
	return bs
}
