package httpxt

import (
	"embed"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
	"io"
	"net/http"
	"strings"
)

//go:embed resource/*
var content embed.FS

func isInvalidArgument(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), "invalid argument:")
}

func ExampleDoNilRequest() {
	_, err := Do(nil)
	fmt.Println(isInvalidArgument(err))
	fmt.Println(err)

	//Output:
	// true
	// invalid argument: Request is nil
}

func ExampleDoHttpError() {
	req, _ := http.NewRequest(http.MethodGet, "echo://www.somestupidname.com?httpError=true", nil)
	resp, err := Do(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)

	//Output:
	// Error : http: connection has been hijacked
	// Response : false

}

func ExampleDoIOError() {
	req, _ := http.NewRequest(http.MethodGet, "echo://www.somestupidname.com?ioError=true", nil)
	resp, err := Do(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Body : %v\n", resp.Body != nil)

	defer resp.Body.Close()
	buf, ioerr := io.ReadAll(resp.Body)
	fmt.Printf("IOError : %v\n", ioerr)
	fmt.Printf("Body : %v\n", string(buf))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 200
	// Body : true
	// IOError : unexpected EOF
	// Body :
}

func ExampleDoSuccess() {
	var uri = "echo://www.somestupidname.com"
	uri += "?content-type=text/html"
	uri += "&content-length=1234"
	uri += "&body=<html><body><h1>Hello, World</h1></body></html>"
	req, err0 := http.NewRequest(http.MethodGet, uri, nil)
	if err0 != nil {
		fmt.Println("failure")
	}
	resp, err := Do(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Content-Type : %v\n", resp.Header.Get("content-type"))
	fmt.Printf("Content-Length : %v\n", resp.Header.Get("content-length"))
	fmt.Printf("Body : %v\n", resp.Body != nil)

	defer resp.Body.Close()
	buf, ioerr := io.ReadAll(resp.Body)
	fmt.Printf("IOError : %v\n", ioerr)
	fmt.Printf("Body : %v\n", string(buf))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 200
	// Content-Type : text/html
	// Content-Length : 1234
	// Body : true
	// IOError : <nil>
	// Body : <html><body><h1>Hello, World</h1></body></html>

}

func ExampleDoEmbeddedContent() {
	var uri = "fse://google.com?search=query by example" //&content-location=embedded"
	req, err0 := http.NewRequestWithContext(fse.ContextWithContent(nil, content, "resource/http/http-503.txt"), http.MethodGet, uri, nil)
	if err0 != nil {
		fmt.Println("failure")
	}
	resp, err := Do(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Content-Type : %v\n", resp.Header.Get("content-type"))
	//fmt.Printf("Content-Length : %v\n", resp.Header.Get("content-length"))
	fmt.Printf("Body : %v\n", resp.Body != nil)

	defer resp.Body.Close()
	_, ioerr := io.ReadAll(resp.Body)
	fmt.Printf("IOError : %v\n", ioerr)
	//fmt.Printf("Body : %v\n", string(buf))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 503
	// Content-Type : text/html
	// Body : true
	// IOError : <nil>

}

func ExampleEchoNil() {
	_, err := createEchoResponse(nil)
	fmt.Println(isInvalidArgument(err))
	fmt.Println(err)

	//Output:
	// true
	// invalid argument: Request is nil
}

func ExampleEchoNoArgs() {
	req, _ := http.NewRequest("", "http://www.somestupidname.com", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Request : %v\n", req != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)

	//Output:
	// Error : <nil>
	// Request : true
	// Status Code : 200

}

func ExampleEchoHttpError() {
	req, _ := http.NewRequest("", "http://somestupidname.com?httpError=true", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)

	//Output:
	// Error : http: connection has been hijacked
	// Response : false
}

func ExampleEchoHttp404() {
	req, _ := http.NewRequest("", "http://somestupidname.com?status=404", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 404

}

func ExampleEchoHeaders() {
	req, _ := http.NewRequest("", "http://somestupidname.com?content-type=application/json&content-length=1234", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Content-Type : %v\n", resp.Header.Get("content-type"))
	fmt.Printf("Content-Length : %v\n", resp.Header.Get("content-length"))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 200
	// Content-Type : application/json
	// Content-Length : 1234

}

func ExampleEchoBody() {
	req, _ := http.NewRequest(http.MethodGet, "http://somestupidname.com?body=this is body content", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Body : %v\n", resp.Body != nil)

	defer resp.Body.Close()
	buf, ioerr := io.ReadAll(resp.Body)
	fmt.Printf("IOError : %v\n", ioerr)
	fmt.Printf("Body : %v\n", string(buf))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 200
	// Body : true
	// IOError : <nil>
	// Body : this is body content

}

func ExampleEchoIOError() {
	req, _ := http.NewRequest(http.MethodGet, "http://somestupidname.com?ioError=true", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Body : %v\n", resp.Body != nil)

	defer resp.Body.Close()
	buf, ioerr := io.ReadAll(resp.Body)
	fmt.Printf("IOError : %v\n", ioerr)
	fmt.Printf("Body : %v\n", string(buf))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 200
	// Body : true
	// IOError : unexpected EOF
	// Body :

}
