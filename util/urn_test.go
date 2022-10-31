package util

import (
	"fmt"
	"net/url"
	"strings"
)

func _ExampleNetUrlParse() {
	//urn := "urn:pgxsql:nss.101?op=test&op2=update"
	urn := "pgxsql:sloentry:id:101"

	u, err := url.Parse(urn)
	tokens := strings.Split(u.Opaque, ":")
	if tokens != nil {
		fmt.Printf("Components : %v\n", tokens)
	}

	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("NSID  : %v\n", u.Scheme)
	fmt.Printf("NSS   : %v\n", u.Opaque)
	fmt.Printf("URI   : %v\n", u)
	fmt.Printf("Query : %v\n", u.RawQuery)

	//Output:
	// fail
	// Error : <nil>
	// Urn   : pgxsql:nss.101
	// NSID  : pgxsql
	// NSS   : nss.101

}

func ExampleUrnParse() {
	urn := "pgxsql:sloentry"
	nsid, nss, query := UrnParse(urn)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("NSID  : %v\n", nsid)
	fmt.Printf("NSS   : %v\n", nss)
	fmt.Printf("Query : %v\n", query)

	//Output:
	// Urn   : pgxsql:nss.101
	// NSID  : pgxsql
	// NSS   : nss.101

}
