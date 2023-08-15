package glog

import (
	"os"
	"time"
)

var defaultLogger = &GLogger{
	level:   DEBUG,
	std:     os.Stdout,
	err:     os.Stderr,
	timeFmt: time.RFC3339,
	format:  TEXT,
	color:   false,
	lineNum: false,
}

func SetOpts(opts LogOptions) {
	defaultLogger.SetOpts(opts)
}
