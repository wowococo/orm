package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errorLogger = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLogger  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers     = []*log.Logger{errorLogger, infoLogger}
	mu          sync.Mutex
)

// log method
var (
	Error  = errorLogger.Println
	Errorf = errorLogger.Printf
	Info   = infoLogger.Println
	Infof  = infoLogger.Printf
)

// log level
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLogger.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infoLogger.SetOutput(ioutil.Discard)
	}

}
