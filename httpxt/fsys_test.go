package httpxt

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"
	//"io/fs"
)

//go:embed resource/*
var fsysContent embed.FS

func init() {
	Mount(fsysContent)
}

func ExampleReader() {
	buf, err := fsysContent.ReadFile("resource/readme.txt")
	if err != nil {
		fmt.Printf("failure : [%v]\n", err)
	} else {
		fmt.Println(string(buf))
	}

	//Output:
	// The resource folder is an example of an embedded resource. The Go language supports embedded resources, just
	// like a Windows program, that are loaded into application types at application startup
	//
	// In this case, the entire resource directory will be accessible as a file system in the application. So, this
	// is a way to mount an in-memory file system.
	//
	// Here is the link to the package : https://pkg.go.dev/embed
	//
	// Be careful on formatting if returning a http response. There needs to be blank lines between the header and body.
	// Look at http-503.txt, which ends with a blank line and no body. The last blank line is required
}

func ExampleDirectory() {
	var dirs []fs.DirEntry
	var err error

	dirs, err = fsysContent.ReadDir("resource")
	if err != nil {
		fmt.Println("failure")
	} else {
		if len(dirs) > 0 {
			for _, info := range dirs {
				fmt.Printf("%v\n", info.Name())
			}
		}
	}

	//Output:
	// http
	// json
	// readme.txt
	// text
}

func ExampleHttp504Response() {
	buf, err := fsysContent.ReadFile("resource/http/http-504.txt")
	if err != nil {
		fmt.Println("failure")
	} else {
		s := string(buf)
		fmt.Printf("Echo : %v\n", strings.TrimSpace(s))
		resp, err0 := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
		if err0 != nil {
			fmt.Println(err0)
		}
		fmt.Printf("Response : %v\n", resp != nil)
		fmt.Printf("Status Code : %v\n", resp.StatusCode)
	}

	//Output:
	// Echo : HTTP/1.2 504 GATEWAY TIMEOUT
	// Response : true
	// Status Code : 504
}

func ExampleHtmlResponse() {
	buf, err := fsysContent.ReadFile("resource/http/html-response.html")
	if err != nil {
		fmt.Println("failure")
	} else {
		resp, _ := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
		fmt.Printf("Response : %v\n", resp != nil)
		fmt.Printf("Status Code : %v\n", resp.StatusCode)
		fmt.Printf("Date : %v\n", resp.Header.Get("Date"))
		fmt.Printf("Server : %v\n", resp.Header.Get("Server"))
		fmt.Printf("Content-Type : %v\n", resp.Header.Get("Content-Type"))
		fmt.Printf("Connection : %v\n", resp.Header.Get("Connection"))
		defer resp.Body.Close()
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("failure\n")
		} else {
			fmt.Printf("Body : %v\n", string(buf))
		}
	}

	//Output:
	// Response : true
	// Status Code : 200
	// Date : Mon, 27 Jul 2009 12:28:53 GMT
	// Server : Apache/2.2.14 (Win32)
	// Content-Type : text/html
	// Connection : Closed
	// Body : <html>
	// <body>
	// <h1>Hello, World!</h1>
	// </body>
	// </html>
}
