package util

import "fmt"

var intEqual IsEqual = func(key, val any) bool {
	if key == nil || val == nil {
		return false
	}
	k, ok := key.(int)
	if !ok {
		return false
	}
	v, ok1 := val.(int)
	if !ok1 {
		return false
	}
	return k == v
}

func ExampleAnyTableInt() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t := CreateAnyTable(intEqual)
	fmt.Printf("Count : %v\n", t.Count())
	for _, v := range data {
		t.Put(v)
	}
	fmt.Printf("Count : %v\n", t.Count())

	v, ok := t.Get(1)
	fmt.Printf("Get [1] : %v %v\n", v, ok)

	v, ok = t.Get(10)
	fmt.Printf("Get [10] : %v %v\n", v, ok)

	v, ok = t.Get(7)
	fmt.Printf("Get [7] : %v %v\n", v, ok)

	ok = t.Delete(7)
	fmt.Printf("Delete [7] : %v\n", ok)

	v, ok = t.Get(7)
	fmt.Printf("Get [7] : %v %v\n", v, ok)

	fmt.Printf("Count : %v\n", t.Count())

	//Output:
	// Count : 0
	// Count : 10
	// Get [1] : 1 true
	// Get [10] : <nil> false
	// Get [7] : 7 true
	// Delete [7] : true
	// Get [7] : <nil> false
	// Count : 9

}

//“The use of COBOL cripples the mind; its teaching should, therefore, be regarded as a criminal offense.”
//― Edsger W. Dijkstra
var stringEqual IsEqual = func(key, val any) bool {
	if key == nil || val == nil {
		return false
	}
	k, ok := key.(string)
	if !ok {
		return false
	}
	v, ok1 := val.(string)
	if !ok1 {
		return false
	}
	return k == v
}

func ExampleAnyTableString() {
	data := []string{"The", "use", "of", "COBOL", "cripples", "the", "mind;", "its", "teaching", "should,", "therefore,", "be", "regarded", "as", "a", "criminal", "offense", "Edsger", "W.", "Dijkstra"}
	t := CreateAnyTable(stringEqual)
	fmt.Printf("Count : %v\n", t.Count())
	for _, v := range data {
		t.Put(v)
	}
	fmt.Printf("Count : %v\n", t.Count())

	v, ok := t.Get("COBOL")
	fmt.Printf("Get [COBOL] : %v %v\n", v, ok)

	v, ok = t.Get("teaching")
	fmt.Printf("Get [teaching] : %v %v\n", v, ok)

	v, ok = t.Get("criminal")
	fmt.Printf("Get [criminal] : %v %v\n", v, ok)

	ok = t.Delete("criminal")
	fmt.Printf("Delete [criminal] : %v\n", ok)

	v, ok = t.Get("criminal")
	fmt.Printf("Get [criminal] : %v %v\n", v, ok)

	fmt.Printf("Count : %v\n", t.Count())

	//Output:
	// Count : 0
	// Count : 20
	// Get [COBOL] : COBOL true
	// Get [teaching] : teaching true
	// Get [criminal] : criminal true
	// Delete [criminal] : true
	// Get [criminal] : <nil> false
	// Count : 19

}

type lookup struct {
	Name  string
	Email string
}

var lookupEqual IsEqual = func(key, val any) bool {
	if key == nil || val == nil {
		return false
	}
	k, ok := key.(string)
	if !ok {
		return false
	}
	v, ok1 := val.(lookup)
	if !ok1 {
		return false
	}
	return k == v.Name
}

func ExampleAnyTableLookup() {
	data := []lookup{lookup{"dave", "test@google.com"}, lookup{Name: "bill", Email: "billg@msn.com"}, lookup{Name: "zuck", Email: "king@facebook.com"}}
	t := CreateAnyTable(lookupEqual)
	fmt.Printf("Count : %v\n", t.Count())
	for _, v := range data {
		t.Put(v)
	}
	fmt.Printf("Count : %v\n", t.Count())

	v, ok := t.Get("bill")
	fmt.Printf("Get [bill] : %v %v\n", v, ok)

	v, ok = t.Get("zuck")
	fmt.Printf("Get [zuck] : %v %v\n", v, ok)

	v, ok = t.Get("dave")
	fmt.Printf("Get [dave] : %v %v\n", v, ok)

	ok = t.Delete("dave")
	fmt.Printf("Delete [dave] : %v\n", ok)

	v, ok = t.Get("dave")
	fmt.Printf("Get [dave] : %v %v\n", v, ok)

	fmt.Printf("Count : %v\n", t.Count())

	//Output:
	// Count : 0
	// Count : 3
	// Get [bill] : {bill billg@msn.com} true
	// Get [zuck] : {zuck king@facebook.com} true
	// Get [dave] : {dave test@google.com} true
	// Delete [dave] : true
	// Get [dave] : <nil> false
	// Count : 2

}
