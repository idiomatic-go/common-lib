package httpxttest

import (
	"fmt"
	"http-boost/httpxt/internal"
	"net/http"
	"net/http/httptest"
	"strings"
)

func ExampleRequest() {
	r, err := http.NewRequest(http.MethodGet, "http://www.google.com", &internal.ReaderCloser{Reader: strings.NewReader("request body content")})
	if r == nil || err != nil {
		fmt.Println("failure")
	} else {
		fmt.Printf("body: %v\n", ReadBodyString(r))
	}

	//Output:
	// body: request body content
}

func ExampleResponse() {
	r := &http.Response{Body: &internal.ReaderCloser{Reader: strings.NewReader("response body content"), Err: nil}}
	fmt.Printf("body: %v\n", ReadBodyString(r))

	//Output:
	// body: response body content
}

func ExampleNewRecorder() {
	r := httptest.NewRecorder()
	r.WriteString("new recorder content")
	r.Flush()
	fmt.Printf("body : %v\n", ReadBodyString(r))

	//Output:
	// body : new recorder content
}
