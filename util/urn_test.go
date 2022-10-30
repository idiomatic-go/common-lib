package util

import (
	"fmt"
	"net/url"
)

func ExampleNetUrlParse() {
	urn := "pgxsql:nss.101"
	u, err := url.Parse(urn)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("NSID  : %v\n", u.Scheme)
	fmt.Printf("NSS   : %v\n", u.Opaque)

	//Output:
	// Error : <nil>
	// Urn   : pgxsql:nss.101
	// NSID  : pgxsql
	// NSS   : nss.101

}

func ExampleUrnParse() {
	urn := "pgxsql:nss.101"
	nsid, nss := UrnParse(urn)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("NSID  : %v\n", nsid)
	fmt.Printf("NSS   : %v\n", nss)

	//Output:
	// Urn   : pgxsql:nss.101
	// NSID  : pgxsql
	// NSS   : nss.101

}
