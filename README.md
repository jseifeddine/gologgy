# gologgy

gologgy is a simple logging package for Go that provides color-coded log messages for different log levels.

## Installation

To install the package, run:

```sh
go get github.com/jseifeddine/gologgy
```

## Usage

You can use the logging functions provided by the `gologgy` package in your Go applications. To import and use the package, include the following import statement in your Go file:

```go
import (
    . "github.com/jseifeddine/gologgy"
)

func main() {
    Error("This is an error message")
    Success("This is a success message")
    Info("This is an informational message")
    Warning("This is a warning message")
    Debug("This is a debug message")
}
```

## Functions

- `Error(format string, args ...interface{})`: Logs an error message in red and exits the program.
- `Success(format string, args ...interface{})`: Logs a success message in green.
- `Info(format string, args ...interface{})`: Logs an informational message in blue.
- `Warning(format string, args ...interface{})`: Logs a warning message in orange (using yellow as the closest color).
- `Debug(format string, args ...interface{})`: Logs a debug message in cyan.

## License

This project is licensed under the MIT License.
