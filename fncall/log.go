package fncall

import "fmt"

var Debug = false

func Debugf(specifier string, v ...any) {
	if Debug {
		fmt.Printf(specifier, v...)
	}
}
