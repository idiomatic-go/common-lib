package util

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

/*
func _ExampleParseQbeInvalid() {
	urn := "urn:jksk-invalid:id=test_slo"

	u := Parse(urn)
	fmt.Printf("Urn   : %v\n", urn)
	fmt.Printf("Error : %v\n", u.Err)

	//Output:
	// Urn   : urn:jksk-invalid:id=test_slo
	// Error : invalid QbeURN Nid : jksk-invalid

}

func ExampleParseQbe() {
	urn := "urn:qbe:id=test_slo"
	u := Parse(urn)
	fmt.Printf("Urn    : %v\n", urn)
	fmt.Printf("Nid    : %v\n", u.Nid)
	fmt.Printf("Nss    : %v\n", u.Nss)
	fmt.Printf("QBE    : %v\n", u.QbeGrid)
	fmt.Printf("Query  : %v\n", u.RawQuery)

	urn = "qbe:id=1001,name=test_slo?content-location=embedded"
	u = Parse(urn)
	fmt.Printf("Urn    : %v\n", urn)
	fmt.Printf("Nid    : %v\n", u.Nid)
	fmt.Printf("Nss    : %v\n", u.Nss)
	fmt.Printf("QBE    : %v\n", u.QbeGrid)
	fmt.Printf("Query  : %v\n", u.RawQuery)

	//Output:
	// Urn    : urn:qbe:id=test_slo
	// Nid    : qbe
	// Nss    : id=test_slo
	// QBE    : [{id test_slo}]
	// Query  :
	// Urn    : qbe:id=1001,name=test_slo?content-location=embedded
	// Nid    : qbe
	// Nss    : id=1001,name=test_slo
	// QBE    : [{id 1001} {name test_slo}]
	// Query  : content-location=embedded

}

func ExampleBuildQbe() {
	u := Build("qbe", "id", 1001)

	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)

	u = Build("", "id", 1001)
	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)

	u = Build("qbe", "id", nil)
	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)

	//Output:
	// Urn    : qbe:id=1001
	// Error  : <nil>
	// Urn    : qbe:id=1001
	// Error  : <nil>
	// Urn    : qbe:id=<nil>
	// Error  : <nil>

}

func ExampleBuildQbeMulti() {
	u := BuildMulti("qbe", QbeCell{Field: "id", Criteria: 1001})

	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)

	u = BuildMulti("", QbeCell{Field: "", Criteria: 1001})
	fmt.Printf("Urn    : %v\n", u)
	fmt.Printf("Error  : %v\n", u.Err)

	//Output:
	// Urn    : qbe:id=1001
	// Error  : <nil>
	// Urn    : qbe:
	// Error  : invalid URN, cell field is empty

}

*/

func ExampleNewURN() {
	u := NewURN(QbeNid, Cell{Field: "id", Criteria: 1001}, Cell{Field: "name", Criteria: "test name"}, Cell{Field: "Date", Criteria: nil})

	fmt.Printf("Urn       : %v\n", u)
	fmt.Printf("Empty     : %v\n", u.IsGridEmpty())
	fmt.Printf("Cell [id] : %v\n", u.Cell("id"))

	u = NewURN(QbeNid, Cell{Field: "id", Criteria: 1001}, Cell{Field: "", Criteria: "test name"}, Cell{Field: "Date", Criteria: nil})

	fmt.Printf("Urn       : %v\n", u)
	fmt.Printf("Empty     : %v\n", u.IsGridEmpty())
	fmt.Printf("Cell [id2]: %v\n", u.Cell("id2"))

	//Output:
	//Urn       : qbe:id=1001,name=test name,Date=<nil>
	//Empty     : false
	//Cell [id] : id=1001
	//Urn       : qbe:id=1001,=test name,Date=<nil>
	//Empty     : false
	//Cell [id2]: <nil>
}
