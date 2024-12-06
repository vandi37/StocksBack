package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/VandiKond/vanerrors/vanstack"
)

// The errors
const ()

// The time format
const (
	FORMAT string = "15:04:05 02.01.06"
)

// Log levels
type LogLevel int

// Log level
const (
	INFO LogLevel = iota
	WARN
	ERROR
	FATAL
)

// Standard log level
var StringLogLevel = map[LogLevel]string{
	INFO:  ">>info<<:",
	WARN:  "!!warn!!:",
	ERROR: "**error**:",
	FATAL: "__fatal__:",
}

// The logger
type Logger struct {
	wInfo    io.Writer
	wWarn    io.Writer
	wError   io.Writer
	wFatal   io.Writer
	levelMap map[LogLevel]string
}

// Creates a standard logger
func NewStd(w io.Writer) Logger {
	return Logger{
		wInfo:    w,
		wWarn:    w,
		wError:   w,
		wFatal:   w,
		levelMap: StringLogLevel,
	}
}

// Logger data
type LoggerData struct {
	WInfo    io.Writer
	WWarn    io.Writer
	WError   io.Writer
	WFatal   io.Writer
	LevelMap map[LogLevel]string
}

// Creates a new logger
func New(d LoggerData) Logger {
	return Logger{
		wInfo:    d.WInfo,
		wWarn:    d.WWarn,
		wError:   d.WError,
		wFatal:   d.WFatal,
		levelMap: d.LevelMap,
	}
}

func writeln(w io.Writer, prefix string, a []any) {
	fmt.Fprintln(w, append([]any{prefix, time.Now().Format(FORMAT)}, a...)...)
}

func writef(w io.Writer, prefix string, format string, a []any) {
	fmt.Fprintf(w, "%s %s "+format+"\n", append([]any{prefix, time.Now().Format(FORMAT)}, a...)...)
}

// Prints a line
func (l Logger) Println(a ...any) {
	writeln(l.wInfo, l.levelMap[INFO], a)
}

// Prints a formatted line
func (l Logger) Printf(format string, a ...any) {
	writef(l.wInfo, l.levelMap[INFO], format, a)
}

// Prints a warn line
func (l Logger) Warnln(a ...any) {
	writeln(l.wWarn, l.levelMap[WARN], a)
}

// Prints a warn formatted line
func (l Logger) Warnf(format string, a ...any) {
	writef(l.wWarn, l.levelMap[WARN], format, a)
}

// Prints a error line
func (l Logger) Errorln(a ...any) {
	writeln(l.wError, l.levelMap[ERROR], a)
}

// Prints a error formatted line
func (l Logger) Errorf(format string, a ...any) {
	writef(l.wError, l.levelMap[ERROR], format, a)
}

// Prints a fatal line and exit
func (l Logger) Fatalln(a ...any) {
	writeln(l.wFatal, l.levelMap[FATAL], a)
	stack := vanstack.NewStack()
	stack.Fill("", 20)
	fmt.Fprintln(os.Stderr, stack)
	os.Exit(http.StatusTeapot)
}

// Prints a fatal formatted line and exit
func (l Logger) Fatalf(format string, a ...any) {
	writef(l.wFatal, l.levelMap[FATAL], format, a)
	stack := vanstack.NewStack()
	stack.Fill("", 20)
	fmt.Fprintln(os.Stderr, stack)
	os.Exit(http.StatusTeapot)
}