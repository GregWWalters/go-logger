package glog

import (
	"io"
	"sync"
)

// SECTION: Type

type GLogger struct {
	write     sync.Mutex
	level     Level
	std       io.Writer
	err       io.Writer
	timeFmt   string
	format    Format
	color     ColorOpt
	lineNum   bool
	multiline bool
}

// SECTION: Public Functions

func NewGLogger(opts LogOptions) *GLogger {
	newLogger := defaultLogger
	newLogger.SetOpts(opts)
	return newLogger
}

func (gl *GLogger) SetOpts(opts LogOptions) {
	if opts.Level != 0 {
		gl.level = opts.Level
	}

	if opts.LogTo != nil {
		gl.std = opts.LogTo
	}
	if opts.ErrTo != nil {
		gl.err = opts.ErrTo
	}

	if opts.TimeFmt != "" {
		gl.timeFmt = opts.TimeFmt
	}

	if opts.Format != 0 {
		gl.format = opts.Format
	}

	if opts.Color > 0 {
		gl.color = opts.Color
	}
	if opts.LineNum > 0 {
		gl.lineNum = opts.LineNum == ON
	}
	if opts.Multiline > 0 {
		gl.multiline = opts.Multiline == ON
	}
}

// SECTION: Public Methods

func (gl *GLogger) Log(l Level, args ...any) {
	gl.log(l, args...)
}

func (gl *GLogger) Logf(l Level, fmt string, args ...any) {
	gl.logf(l, fmt, args...)
}

func (gl *GLogger) Trace(args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Debug(args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Info(args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Warn(args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Error(args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Fatal(args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Tracef(fmt string, args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Debugf(fmt string, args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Infof(fmt string, args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Warnf(fmt string, args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Errorf(fmt string, args ...any) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) Fatalf(fmt string, args ...any) {
	// TODO implement me
	panic("implement me")
}

// SECTION: Private Methods

func (gl *GLogger) writeOut(l lineSerializer) {
	// TODO implement me
	panic("implement me")
}

func (gl *GLogger) log(level Level, args ...any) {
	if level < gl.level {
		return
	}
	log := line{
		level:     gl.level,
		ts:        "",
		color:     gl.color,
		lineNum:   gl.lineNum,
		args:      args,
		multiline: gl.multiline,
	}
	var toWrite []byte
	switch gl.format {
	case JSON:
		toWrite = log.toJSON()
	case TEXT:
		toWrite = []byte(log.toText())
	}
	var dest io.Writer
	if level > WARN {
		dest = gl.err
	} else {
		dest = gl.std
	}
	gl.write.Lock()
	_, _ = dest.Write(toWrite)
	gl.write.Unlock()
}

func (gl *GLogger) logf(l Level, fmt string, args ...any) {
	// TODO implement me
	panic("implement me")
}
