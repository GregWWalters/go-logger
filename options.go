package glog

import "io"

type BoolOpt uint8
type ColorOpt uint8

const (
	OFF BoolOpt = iota + 1
	ON
)

const (
	ColorLogLevel ColorOpt = iota + 1
	ColorPrefix
	ColorAll
)

type LogOptions struct {
	Level     Level
	Tags      []string
	LogTo     io.Writer
	ErrTo     io.Writer
	TimeFmt   string
	LineNum   BoolOpt
	Format    Format
	Color     ColorOpt
	Multiline BoolOpt
}
