package logxt

import (
	"fmt"
	//"github.com/idiomatic-go/common-lib/util"
	"log"
)

var debug = false

func init() {
	//logxt.SetFlags(logxt.Ldate | logxt.Ltime | logxt.Lmicroseconds | logxt.Llongfile | logxt.LUTC)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
}

func LogDebug(v ...any) {
	if debug {
		fmt.Println(v)
	}
}

func LogDebugf(specifier string, v ...any) {
	if debug {
		fmt.Printf(specifier, v)
	}
}

var LogPanic DefaultFmt = func(v ...any) {
	log.Panic(v)
}

var LogPanicf SpecifiedFmt = func(specifier string, v ...any) {
	log.Panicf(specifier, v)
}

var LogPrint DefaultFmt = func(v ...any) {
	log.Print(v)
}

var LogPrintf SpecifiedFmt = func(specifier string, v ...any) {
	log.Printf(specifier, v)
}
