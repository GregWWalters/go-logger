package glog

type Logger interface {
	Log(l Level, args ...any)
	Trace(args ...any)
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

type FormatLogger interface {
	Logger
	Logf(l Level, fmt string, args ...any)
	Tracef(fmt string, args ...any)
	Debugf(fmt string, args ...any)
	Infof(fmt string, args ...any)
	Warnf(fmt string, args ...any)
	Errorf(fmt string, args ...any)
	Fatalf(fmt string, args ...any)
}

type LogSetter interface {
	SetOpts(opts LogOptions)
}
