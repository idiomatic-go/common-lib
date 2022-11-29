package fse

import (
	"context"
	"fmt"
)

type CustomerHistory struct {
	CustomerId string
	Events     []Event
}
type Event struct {
	CustomerId  string
	Transaction string
	Status      string
	CreateDate  string
	LogEventId  string
}

func ExampleEntryErrorFile() {
	ctx := ContextWithContent(nil, fsys, "resource/error/invalid-error-content.txt")
	entry := ContextContent(ctx)

	fmt.Printf("Name  : %v\n", entry.Name)
	fmt.Printf("Buf   : %v\n", string(entry.Content))
	fmt.Printf("Error : %v\n", entry.Error())

	ctx = ContextWithContent(nil, fsys, "resource/error/next_test_ERRor.txt")
	entry = ContextContent(ctx)

	fmt.Printf("Name  : %v\n", entry.Name)
	fmt.Printf("Buf   : %v\n", string(entry.Content))
	fmt.Printf("Error : %v\n", entry.Error())

	ctx = ContextWithContent(nil, fsys, "resource/error/test-name-eRr.txt")
	entry = ContextContent(ctx)

	fmt.Printf("Name  : %v\n", entry.Name)
	fmt.Printf("Buf   : %v\n", string(entry.Content))
	fmt.Printf("Error : %v\n", entry.Error())

	//Output:
	//Name  : resource/error/invalid-error-content.txt
	//Buf   : This is invalid content
	//Error : <nil>
	//Name  : resource/error/next_test_error.txt
	//Buf   : This is example 2 of an error from next_test_ERRor.txt
	//Error : This is example 2 of an error from next_test_ERRor.txt
	//Name  : resource/error/test-name-err.txt
	//Buf   : This is example 1 of an error from test-name-eRr.txt
	//Error : This is example 1 of an error from test-name-eRr.txt

}

func ExampleProcessContentError() {
	//cust := CustomerHistory{}
	_, err := ProcessContent[CustomerHistory](nil)
	fmt.Printf("Nil Context : %v\n", err)

	_, err = ProcessContent[CustomerHistory](context.Background())
	fmt.Printf("Nil Entry   : %v\n", err)

	ctx := ContextWithContent(nil, fsys, "resource/json/event-empty.json")
	_, err = ProcessContent[CustomerHistory](ctx)
	fmt.Printf("Nil Content : %v\n", err)

	//ctx = ContextWithContent(nil, fsys, "resource/json/event.json")
	//_, err = ProcessContent[CustomerHistory](ctx)
	//fmt.Printf("Nil Any     : %v\n", err)

	ctx = ContextWithContent(nil, fsys, "resource/json/event-not-json.txt")
	_, err = ProcessContent[CustomerHistory](ctx)
	fmt.Printf("Non Json    : %v\n", err)

	//Output:
	//Nil Context : fse:ProcessContent internal error : context is nil
	//Nil Entry   : fse:ProcessEntry internal error : no file system entry available
	//Nil Content : fse:ProcessEntry internal error : no content available for entry name : resource/json/event-empty.json
	//Non Json    : fse:ProcessEntry internal error : invalid content for json.Unmarshal() : resource/json/event-not-json.txt
}

func ExampleProcessContent() {
	ctx := ContextWithContent(nil, fsys, "resource/error/test-name-eRr.txt")
	_, err := ProcessContent[CustomerHistory](ctx)
	fmt.Printf("Error : %v\n", err)

	ctx = ContextWithContent(nil, fsys, "resource/json/event.json")
	cust, err0 := ProcessContent[CustomerHistory](ctx)
	fmt.Printf("Error : %v\n", err0)
	fmt.Printf("Cust  : %v\n", cust)

	ctx = ContextWithContent(nil, fsys, "resource/json/event-list.json")
	list, err1 := ProcessContent[[]CustomerHistory](ctx)
	fmt.Printf("Error : %v\n", err1)
	fmt.Printf("List  : %v\n", list)

	//Output:
	//Error : This is example 1 of an error from test-name-eRr.txt
	//Error : <nil>
	//Cust  : {C1 [{ invoice created 2022-08-25 0001} { payment created 2022-08-26 0002}]}
	//Error : <nil>
	//List  : [{C1 [{ invoice created 2022-08-25 0001} { payment created 2022-08-26 0002}]} {C0002 [{ return created 2022-10-25 99990001} { credit created 2022-11-01 0002333}]}]

}
