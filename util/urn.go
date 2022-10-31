package util

import (
	"net/url"
)

func ParseUrn(urn string) *UrnComponents {
	u, err := url.Parse(urn)
	if err != nil {
		return nil
	}
	if u.RawQuery == "" {
		return &UrnComponents{NSID: u.Scheme, NSS: u.Opaque}
	}
	val, _ := url.ParseQuery(u.RawQuery)
	return &UrnComponents{NSID: u.Scheme, NSS: u.Opaque, Values: val}
}
