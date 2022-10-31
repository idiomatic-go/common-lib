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

	comp := ParseUrn(urn)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("NSID  : %v\n", comp.NSID)
	fmt.Printf("NSS   : %v\n", comp.NSS)
	fmt.Printf("Query : %v\n", comp.Values)

	urn = "pgxsql:sloentry?id=1001&name=test_slo"
	comp = ParseUrn(urn)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("NSID  : %v\n", comp.NSID)
	fmt.Printf("NSS   : %v\n", comp.NSS)
	fmt.Printf("Query : %v\n", comp.Values)

	//Output:
	//Urn   : pgxsql:sloentry
	//NSID  : pgxsql
	//NSS   : sloentry
	//Query : map[]
	//Urn   : pgxsql:sloentry?id=1001&name=test_slo
	//NSID  : pgxsql
	//NSS   : sloentry
	//Query : map[id:[1001] name:[test_slo]]

}
