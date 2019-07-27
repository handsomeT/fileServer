package log

import "fmt"

func Error(a ...interface{}) {
	fmt.Println(a)
}

func ErrorF(format string, a ...interface{}) {
	fmt.Printf(format, a)
	fmt.Println()
}