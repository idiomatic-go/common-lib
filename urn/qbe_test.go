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
	//fmt.Printf("Nid   : %v\n", u.Nid)
	//fmt.Printf("Nss   : %v\n", u.Nss)
	//fmt.Printf("QBE   : %v\n", u.Grid)
	fmt.Printf("Error : %v\n", u.Err)

	//Output:
	// Urn   : urn:jksk-invalid:id=test_slo
	// Error : invalid QbeURN Nid : jksk-invalid

}

func ExampleParseQbe() {
	urn := "urn:qbe:id=test_slo"
	u := ParseQbe(urn)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("Nid   : %v\n", u.Nid)
	fmt.Printf("Nss   : %v\n", u.Nss)
	fmt.Printf("QBE   : %v\n", u.Grid)

	urn = "qbe:id=1001,name=test_slo,scheme=pgxsql"
	u = ParseQbe(urn)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("Nid   : %v\n", u.Nid)
	fmt.Printf("Nss   : %v\n", u.Nss)
	fmt.Printf("QBE   : %v\n", u.Grid)

	//Output:
	// Urn   : urn:qbe:id=test_slo
	// Nid   : qbe
	// Nss   : id=test_slo
	// QBE   : [{id test_slo}]
	// Urn   : qbe:id=1001,name=test_slo,scheme=pgxsql
	// Nid   : qbe
	// Nss   : id=1001,name=test_slo,scheme=pgxsql
	// QBE   : [{id 1001} {name test_slo} {scheme pgxsql}]

}

func ExampleBuildQbe() {
	u := BuildQbe("pgxsql", "id", 1001)

	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Scheme : %v\n", u.Scheme())

	u = BuildQbe("", "id", 1001)
	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Scheme : %v\n", u.Scheme())

	u = BuildQbe("pgxsql", "id", nil)
	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Scheme : %v\n", u.Scheme())

	//Output:
	// Urn    : qbe:scheme=pgxsql,id=1001
	// Error  : <nil>
	// Scheme : pgxsql
	// Urn    : qbe:id=1001
	// Error  : <nil>
	// Scheme :
	// Urn    : qbe:scheme=pgxsql,id=<nil>
	// Error  : <nil>
	// Scheme : pgxsql
}

func ExampleBuildQbeMulti() {
	u := BuildQbeMulti(Cell{Field: SchemeName, Criteria: "pgxsql"}, Cell{Field: "id", Criteria: 1001})

	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Scheme : %v\n", u.Scheme())

	u = BuildQbeMulti(Cell{Field: SchemeName, Criteria: nil}, Cell{Field: "id", Criteria: 1001})
	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Scheme : %v\n", u.Scheme())

	u = BuildQbeMulti(Cell{Field: SchemeName, Criteria: nil}, Cell{Field: "", Criteria: 1001})
	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)
	fmt.Printf("Scheme : %v\n", u.Scheme())

	/*
		u = BuildQbe("pgxsql", "id", nil)
		fmt.Printf("Urn    : %v\n", u)
		fmt.Printf("Error  : %v\n", u.Err)
		fmt.Printf("Scheme : %v\n", u.Scheme())
	*/

	//Output:
	// Urn    : qbe:scheme=pgxsql,id=1001
	// Error  : <nil>
	// Scheme : pgxsql
	// Urn    : qbe:scheme=<nil>,id=1001
	// Error  : <nil>
	// Scheme :
	// Urn    : qbe:scheme=<nil>
	// Error  : invalid QbeURN, cell field is empty
	// Scheme :

}
