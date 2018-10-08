package logger

import (
    "log"
    "os"
)

var bookeeLog = log.New(os.Stderr, "", log.LstdFlags)

func L() *log.Logger {
    return bookeeLog
}