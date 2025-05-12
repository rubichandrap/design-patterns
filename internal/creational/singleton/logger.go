
package singleton

import (
    "fmt"
    "log"
    "os"
    "sync"
)

type Logger struct {
    logger *log.Logger
}

var instance *Logger
var once sync.Once

func GetInstance() *Logger {
    once.Do(func() {
        instance = &Logger{
            logger: log.New(os.Stdout, "", log.LstdFlags),
        }
    })
    return instance
}

func (l *Logger) Log(msg string) {
    l.logger.Println("[LOG]", msg)
}

func (l *Logger) Error(msg string) {
    l.logger.Println("[ERROR]", msg)
}
