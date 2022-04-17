package glerrors

import (
	"fmt"
)

func Error(line int, message string) {
	Report(line, "", message)
}

func Report(line int, where, message string) {
	fmt.Printf("[line %d] Error %s: %s", line, where, message)
}
