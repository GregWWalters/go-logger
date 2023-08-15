package glog

// MARK: Type

type color uint8

// MARK: Variables

// terminal defaults
var (
	reset color = 0
)

// standard colors
var (
	grey    color = 37
	gray    color = grey
	black   color = 30
	red     color = 31
	green   color = 32
	blue    color = 34
	cyan    color = 36
	magenta color = 35
	yellow  color = 33
)

// bright colors
var (
	brightGrey    color = grey + 60
	brightGray    color = brightGrey
	white         color = brightGrey
	brightBlack   color = black + 60
	brightRed     color = red + 60
	brightGreen   color = green + 60
	brightBlue    color = blue + 60
	brightCyan    color = cyan + 60
	brightMagenta color = magenta + 60
	brightYellow  color = yellow + 60
)

// MARK: private functions

// apply returns the string wrapped in the codes to apply the foreground
// color and reset
func (c color) apply(s string) string {
	return "\u001b[" + string(c) + "m" + string(s) + "\u001b[0m"
}
