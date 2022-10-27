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

var LogDebug DebugFmt = func(specifier string, v ...any) {
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

/*
var LogContextPrint ContextDefaultFmt = func(ctx context.Context, v ...any) {
	u := []any{util.ContextRequestId(ctx)}
	log.Print(append(u, v))
}

var LogContextPrintf ContextSpecifiedFmt = func(ctx context.Context, specifier string, v ...any) {
	log.Printf(specifier, v)
}

*/
