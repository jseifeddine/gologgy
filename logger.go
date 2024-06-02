package gologgy

import (
    "fmt"
    "log"
    "os"
)

// Error logs an error message with a red color and exits the program
func Error(format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    fmt.Printf("\x1b[31;1m%s\x1b[0m\n", message)
    os.Exit(1)
}

// Success logs a success message with a green color
func Success(format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    log.Printf("\x1b[32;1m%s\x1b[0m\n", message)
}

// Info logs an informational message with a blue color
func Info(format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    log.Printf("\x1b[34;1m%s\x1b[0m\n", message)
}

// Warning logs a warning message with an orange color (using yellow as closest color)
func Warning(format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    log.Printf("\x1b[33;1m%s\x1b[0m\n", message)
}

