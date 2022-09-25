package httpxt

import (
	"net/http"
)

// See https://tools.ietf.org/html/rfc6265 for details of each of the fields of the above cookie.

func ReadCookies(req *http.Request) map[string]*http.Cookie {
	if req == nil {
		return nil
	}
	jar := make(map[string]*http.Cookie)
	for _, c := range req.Cookies() {
		jar[c.Name] = c
	}
	return jar
}
