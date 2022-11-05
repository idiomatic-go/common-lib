package fse

import "fmt"

func ExampleEntryError() {
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
