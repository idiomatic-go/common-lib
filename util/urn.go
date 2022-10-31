package util

import (
	"net/url"
)

func UrnParse(urn string) (nsid string, nss string, query string) {
	u, err := url.Parse(urn)
	if err != nil {
		return "", "", ""
	}
	return u.Scheme, u.Opaque, u.RawQuery
}
