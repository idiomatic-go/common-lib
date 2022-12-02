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

func Debug(v ...any) {
	if debug {
		fmt.Println(v...)
	}
}

func Debugf(specifier string, v ...any) {
	if debug {
		fmt.Printf(specifier, v...)
	}
}

func Panic(v ...any) {
	log.Panic(v...)
}

func Panicf(specifier string, v ...any) {
	log.Panicf(specifier, v...)
}

func Print(v ...any) {
	log.Print(v...)
}

func Println(v ...any) {
	log.Println(v...)
}

func Printf(specifier string, v ...any) {
	log.Printf(specifier, v...)
}

/*
func LogContextPrint(ctx context.Context, v ...any) {
	u := []any{ContextRequestId(ctx)}
	Printf("%v\n", append(u, v))
}

func LogContextPrintf(ctx context.Context, specifier string, v ...any) {
	Printf(specifier, v)
}


*/
