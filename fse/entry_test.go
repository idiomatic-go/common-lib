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

func _ExampleEntryErrorFile() {
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
	cust := CustomerHistory{}
	err := ProcessContent(nil, nil)
	fmt.Printf("Nil Context : %v\n", err)

	err = ProcessContent(context.Background(), nil)
	fmt.Printf("Nil Entry   : %v\n", err)

	ctx := ContextWithContent(nil, fsys, "resource/json/event-empty.json")
	err = ProcessContent(ctx, nil)
	fmt.Printf("Nil Content : %v\n", err)

	ctx = ContextWithContent(nil, fsys, "resource/json/event.json")
	err = ProcessContent(ctx, nil)
	fmt.Printf("Nil Any     : %v\n", err)

	ctx = ContextWithContent(nil, fsys, "resource/json/event-not-json.txt")
	err = ProcessContent(ctx, &cust)
	fmt.Printf("Non Json    : %v\n", err)

	//Output:
	//Nil Context : invalid argument : context is nil
	//Nil Entry   : no file system entry available
	//Nil Content : no content available for entry name : resource/json/event-empty.json
	//Nil Any     : invalid argument : any parameter is nil
	//Non Json    : invalid content for json.Unmarshal() : resource/json/event-not-json.txt
}

func ExampleProcessContent() {
	cust := CustomerHistory{}

	err := ProcessContent(nil, nil)
	fmt.Printf("Nil Context : %v\n", err)

	err = ProcessContent(context.Background(), nil)
	fmt.Printf("Nil Entry   : %v\n", err)

	ctx := ContextWithContent(nil, fsys, "resource/json/event-empty.json")
	err = ProcessContent(ctx, nil)
	fmt.Printf("Nil Content : %v\n", err)

	ctx = ContextWithContent(nil, fsys, "resource/json/event.json")
	err = ProcessContent(ctx, nil)
	fmt.Printf("Nil Any     : %v\n", err)

	ctx = ContextWithContent(nil, fsys, "resource/json/event-not-json.txt")
	err = ProcessContent(ctx, &cust)
	fmt.Printf("Non Json    : %v\n", err)

	//Output:
	//Nil Context : invalid argument : context is nil
	//Nil Entry   : no file system entry available
	//Nil Content : no content available for entry name : resource/json/event-empty.json
	//Nil Any     : invalid argument : any parameter is nil
	//Non Json    : invalid content for json.Unmarshal() : resource/json/event-not-json.txt
}

