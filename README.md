# logz v1.0.0
[![Go Reference](https://pkg.go.dev/badge/github.com/zomasec/logz.svg)](https://pkg.go.dev/github.com/zomasec/logz) 
[![Go Report Card](https://goreportcard.com/badge/github.com/zomasec/logz)](https://goreportcard.com/report/github.com/zomasec/logz)
[![License](https://img.shields.io/github/license/zomasec/logz)](https://github.com/zomasec/logz/blob/main/LICENSE)

logz is a simple logging package in Go with different log levels, developed by Hazem El-Sayed

## Installation

```bash
go get -u github.com/zomasec/logz
```

## Usage

### Basic Logging
```go
package main

import (
    "github.com/zomasec/logz"
)

func main() {
    // Create a new logger
    logger := logz.DefaultLogs()

    // Log messages at different levels
    logger.INFO("This is an informational message")
    logger.DEBUG("This is a debug message")
    logger.WARNING("This is a warning message")
    logger.ERROR("This is an error message")
    logger.FATAL("This is a fatal message")
}
```
## Custom Logging

```go
package main

import (
    "github.com/zomasec/logz"
)

func main() {
    // Create a custom logger
    customLogger := logz.NewLogger("CUSTOM", "This is a custom message with %d args", 42)

    // Log the custom message
    customLogger.Log()
}
```


## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author
- @ZomaSec

