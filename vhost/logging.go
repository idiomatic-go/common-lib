package vhost

import (
	"context"
	"fmt"
	"log"

	"github.com/idiomatic-go/common-lib/vhost/usr"
)

func init() {
	//log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.LUTC)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
}

var LogDebug usr.DebugFmt = func(specifier string, v ...any) {
	if usr.Debug {
		fmt.Printf(specifier, v)
	}
}

var LogPanic usr.DefaultFmt = func(v ...any) {
	log.Panic(v)
}

var LogPanicf usr.SpecifiedFmt = func(specifier string, v ...any) {
	log.Panicf(specifier, v)
}

var LogPrint usr.DefaultFmt = func(v ...any) {
	log.Print(v)
}

var LogPrintf usr.SpecifiedFmt = func(specifier string, v ...any) {
	log.Printf(specifier, v)
}

var LogContextPrint usr.ContextDefaultFmt = func(ctx context.Context, v ...any) {
	//requestId := ContextRequestId(ctx)
	u := []any{ContextRequestId(ctx)}
	log.Print(append(u, v))
}

var LogContextPrintf usr.ContextSpecifiedFmt = func(ctx context.Context, specifier string, v ...any) {
	log.Printf(specifier, v)
}
