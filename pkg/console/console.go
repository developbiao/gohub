package console

import (
	"fmt"
	"github.com/mgutz/ansi"
	"os"
)

// Success print green success message
func Success(msg string) {
	colorOut(msg, "green")
}

// Error print red error message
func Error(msg string) {
	colorOut(msg, "red")
}

// Warning print yellow warning message
func Warning(msg string) {
	colorOut(msg, "yellow")
}

// Exit print error message and call os.Exit
func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

// ExitIf if exist error print error
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

// colorOut  ANSI code for style
func colorOut(message, color string) {
	fmt.Fprintf(os.Stdout, ansi.Color(message, color))
}
