package log

import (
    "fmt"
    "os"
)

func Println(a ...interface{}) {
    fmt.Println(a...)
}

func ErrorAndPanic(a ...interface{}) {
    panic(fmt.Sprint(a...))
}

func ErrorAndExit(a ...interface{}) {
    fmt.Println(a...)
    os.Exit(-1)
}
