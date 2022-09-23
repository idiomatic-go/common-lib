package util

import "fmt"

type versionedStruct struct {
	vers  string
	count int
}

var getVersion = func(a any) string {
	if a == nil {
		return ""
	}
	if data, ok := a.(*versionedStruct); ok {
		return data.vers
	}
	return ""
}

var getEntity = func(a any) *versionedStruct {
	if a == nil {
		return nil
	}
	if data, ok := a.(*versionedStruct); ok {
		return data
	}
	return nil
}

func ExampleNoEntity() {
	e := CreateVersionedEntity(nil, getVersion)
	fmt.Printf("IsNewVersion [empty] : %v\n", e.IsNewVersion(""))
	fmt.Printf("IsNewVersion [123] : %v\n", e.IsNewVersion("123"))
	fmt.Printf("IsNewVersion [1.2.3] : %v\n", e.IsNewVersion("1.2.3"))
	if e.Get() == nil {
		fmt.Println("Get : valid")
	} else {
		fmt.Println("failure")
	}

	//Output:
	// IsNewVersion [empty] : false
	// IsNewVersion [123] : false
	// IsNewVersion [1.2.3] : false
	// Get : valid
}

func ExampleNoGetFn() {
	s := versionedStruct{vers: "1.2.3", count: 1}
	e := CreateVersionedEntity(&s, nil)
	fmt.Printf("IsNewVersion [empty] : %v\n", e.IsNewVersion(""))
	fmt.Printf("IsNewVersion [123] : %v\n", e.IsNewVersion("123"))
	fmt.Printf("IsNewVersion [1.2.3] : %v\n", e.IsNewVersion("1.2.3"))

	//Output:
	// IsNewVersion [empty] : false
	// IsNewVersion [123] : true
	// IsNewVersion [1.2.3] : true
}

func ExampleValidEntity() {
	s := versionedStruct{vers: "1.2.3", count: 1}
	e := CreateVersionedEntity(&s, getVersion)
	fmt.Printf("IsNewVersion [empty] : %v\n", e.IsNewVersion(""))
	fmt.Printf("IsNewVersion [123] : %v\n", e.IsNewVersion("123"))
	fmt.Printf("IsNewVersion [1.2.3] : %v\n", e.IsNewVersion("1.2.3"))
	fmt.Printf("Version : %v\n", e.GetVersion())
	entity := getEntity(e.Get())
	if entity != nil {
		fmt.Printf("Entity : [%v]", entity)
	} else {
		fmt.Println("Entity : invalid")
	}

	//Output:
	// IsNewVersion [empty] : true
	// IsNewVersion [123] : true
	// IsNewVersion [1.2.3] : false
	// Version : 1.2.3
	// Get : valid
}

func _ExampleChangeEntity() {
	s := versionedStruct{vers: "1.2.3", count: 1}
	e := CreateVersionedEntity(&s, getVersion)
	fmt.Printf("Version : %v\n", e.GetVersion())
	fmt.Printf("IsNewVersion [1.2.3] : %v\n", e.IsNewVersion("1.2.3"))
	entity := getEntity(e.Get())
	if entity != nil {
		fmt.Printf("Entity : [%v]", entity)
	} else {
		fmt.Println("Entity : invalid")
	}

	//Output:
	// IsNewVersion [empty] : true
	// IsNewVersion [123] : true
	// IsNewVersion [1.2.3] : false
	// Version : 1.2.3
	// Get : valid
}
