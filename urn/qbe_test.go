package urn

import (
	"fmt"
	"net/url"
	"strings"
)

func _ExampleNetUrlParse() {
	//urn := "urn:pgxsql:Nss.101?op=test&op2=update"
	urn := "pgxsql:sloentry:id:101"

	u, err := url.Parse(urn)
	tokens := strings.Split(u.Opaque, ":")
	if tokens != nil {
		fmt.Printf("Components : %v\n", tokens)
	}

	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("Nid  : %v\n", u.Scheme)
	fmt.Printf("Nss   : %v\n", u.Opaque)
	fmt.Printf("URI   : %v\n", u)
	fmt.Printf("Query : %v\n", u.RawQuery)

	urn = "pgxsql:qbe:id=name,value=test_slo"

	u, err = url.Parse("")
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("Nid  : %v\n", u.Scheme)
	fmt.Printf("Nss   : %v\n", u.Opaque)
	fmt.Printf("URI   : %v\n", u)
	fmt.Printf("Query : %v\n", u.RawQuery)

	//Output:
	// fail
	// Error : <nil>
	// Urn   : pgxsql:Nss.101
	// Nid  : pgxsql
	// Nss   : Nss.101

}
func ExampleParseQbeInvalid() {
	urn := "urn:jksk-invalid:id=test_slo"

	u := ParseQbe(urn)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("Error : %v\n", u.Err)

	//Output:
	// Urn   : urn:jksk-invalid:id=test_slo
	// Error : invalid QbeURN Nid : jksk-invalid

}

func ExampleParseQbe() {
	urn := "urn:qbe:id=test_slo"
	u := ParseQbe(urn)
	fmt.Printf("Urn    : %v\n", urn)
	fmt.Printf("Nid    : %v\n", u.Nid)
	fmt.Printf("Nss    : %v\n", u.Nss)
	fmt.Printf("QBE    : %v\n", u.Grid)
	fmt.Printf("Values : %v\n", u.Values)

	urn = "qbe:id=1001,name=test_slo?content-location=embedded"
	u = ParseQbe(urn)
	fmt.Printf("Urn    : %v\n", urn)
	fmt.Printf("Nid    : %v\n", u.Nid)
	fmt.Printf("Nss    : %v\n", u.Nss)
	fmt.Printf("QBE    : %v\n", u.Grid)
	fmt.Printf("Values : %v\n", u.Values)

	//Output:
	// Urn    : urn:qbe:id=test_slo
	// Nid    : qbe
	// Nss    : id=test_slo
	// QBE    : [{id test_slo}]
	// Values : map[]
	// Urn    : qbe:id=1001,name=test_slo?content-location=embedded
	// Nid    : qbe
	// Nss    : id=1001,name=test_slo
	// QBE    : [{id 1001} {name test_slo}]
	// Values : map[content-location:[embedded]]

}

func ExampleBuildQbe() {
	u := BuildQbe(true, "id", 1001)

	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Values : %v\n", u.Values)

	u = BuildQbe(false, "id", 1001)
	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Values : %v\n", u.Values)

	u = BuildQbe(true, "id", nil)
	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Values : %v\n", u.Values)

	//Output:
	// Urn    : qbe:id=1001?content-location=embedded
	// Error  : <nil>
	// Values : map[content-location:[embedded]]
	// Urn    : qbe:id=1001
	// Error  : <nil>
	// Values : map[]
	// Urn    : qbe:id=<nil>?content-location=embedded
	// Error  : <nil>
	// Values : map[content-location:[embedded]]
}

func ExampleBuildQbeMulti() {
	u := BuildQbeMulti(true, Cell{Field: "id", Criteria: 1001})

	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Values : %v\n", u.Values)

	u = BuildQbeMulti(true, Cell{Field: "", Criteria: 1001})
	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Values : %v\n", u.Values)

	/*
		u = BuildQbe("pgxsql", "id", nil)
		fmt.Printf("Urn    : %v\n", u)
		fmt.Printf("Error  : %v\n", u.Err)
		fmt.Printf("Scheme : %v\n", u.Scheme())
	*/

	//Output:
	// Urn    : qbe:id=1001?content-location=embedded
	// Error  : <nil>
	// Values : map[content-location:[embedded]]
	// Urn    : qbe:
	// Error  : invalid QbeURN, cell field is empty
	// Values : map[content-location:[embedded]]

}
