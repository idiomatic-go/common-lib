package util

import "net/url"

func UrnParse(urn string) (nsid, nss string) {
	u, err := url.Parse(urn)
	if err != nil {
		return "", ""
	}
	return u.Scheme, u.Opaque
}
