package vhost

import (
	"context"
	"fmt"
	"log"
)

func init() {
	//log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.LUTC)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
}

var LogDebug DebugFmt = func(specifier string, v ...any) {
	if Debug {
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

var LogContextPrint ContextDefaultFmt = func(ctx context.Context, v ...any) {
	//requestId := ContextRequestId(ctx)
	u := []any{ContextRequestId(ctx)}
	log.Print(append(u, v))
}

var LogContextPrintf ContextSpecifiedFmt = func(ctx context.Context, specifier string, v ...any) {
	log.Printf(specifier, v)
}
