package logger

import "log"

func Debug(args ...string) {
    log.Println(args)
}

func Trace(args ...string) {
    log.Println(args)
}

func Info(args ...string) {
    log.Println(args)
}

func Warn(args ...string) {
    log.Println(args)
}

func Error(args ...string) {
    log.Println(args)
}

func Fatal(args ...string) {
    log.Println(args)
}
