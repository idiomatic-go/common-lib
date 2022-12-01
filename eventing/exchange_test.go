package eventing

import "fmt"

func isStatus(msg *Message, status int32) bool {
	if msg == nil {
		return false
	}
	return msg.Status == status
}

func ExampleExchange() {
	exch := createExchange()

	fmt.Printf("Messages   : %v\n", exch.count(StartupEvent))
	fmt.Printf("InProgress : %v\n", isStatus(exch.current(StartupEvent), StatusInProgress))
	fmt.Printf("Failure    : %v\n", isStatus(exch.current(StartupEvent), StatusInternal))
	fmt.Printf("Ok         : %v\n", isStatus(exch.current(StartupEvent), StatusOk))

	exch.add(CreateMessage(VirtualHost, "test:package", StartupEvent, StatusInProgress, nil))

	fmt.Printf("Messages   : %v\n", exch.count(StartupEvent))
	fmt.Printf("InProgress : %v\n", isStatus(exch.current(StartupEvent), StatusInProgress))
	fmt.Printf("Failure    : %v\n", isStatus(exch.current(StartupEvent), StatusInternal))
	fmt.Printf("Ok         : %v\n", isStatus(exch.current(StartupEvent), StatusOk))

	exch.add(CreateMessage(VirtualHost, "test:package", StartupEvent, StatusInternal, nil))

	fmt.Printf("Messages   : %v\n", exch.count(StartupEvent))
	fmt.Printf("InProgress : %v\n", isStatus(exch.current(StartupEvent), StatusInProgress))
	fmt.Printf("Failure    : %v\n", isStatus(exch.current(StartupEvent), StatusInternal))
	fmt.Printf("Ok         : %v\n", isStatus(exch.current(StartupEvent), StatusOk))

	exch.add(CreateMessage(VirtualHost, "test:package", StartupEvent, StatusOk, nil))

	fmt.Printf("Messages   : %v\n", exch.count(StartupEvent))
	fmt.Printf("InProgress : %v\n", isStatus(exch.current(StartupEvent), StatusInProgress))
	fmt.Printf("Failure    : %v\n", isStatus(exch.current(StartupEvent), StatusInternal))
	fmt.Printf("Ok         : %v\n", isStatus(exch.current(StartupEvent), StatusOk))

	//Output:
	//Messages   : 0
	//InProgress : false
	//Failure    : false
	//Ok         : false
	//Messages   : 1
	//InProgress : true
	//Failure    : false
	//Ok         : false
	//Messages   : 2
	//InProgress : false
	//Failure    : true
	//Ok         : false
	//Messages   : 3
	//InProgress : false
	//Failure    : false
	//Ok         : true

}
