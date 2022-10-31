package util

import (
	"net/url"
)

func ParseUrn(urn string) (nsid string, nss string, args url.Values) {
	u, err := url.Parse(urn)
	if err != nil {
		return "", "", nil
	}
	if u.RawQuery == "" {
		return u.Scheme, u.Opaque, nil
	}
	val, _ := url.ParseQuery(u.RawQuery)
	return u.Scheme, u.Opaque, val
}
