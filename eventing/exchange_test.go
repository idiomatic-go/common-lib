package eventing

import "fmt"

// Only used for testing, no need to declare elsewhere
const (
	statusOk       = 0  // Need to sink with gRPC Ok
	statusInternal = 13 // Need to sink with gRPC Internal
)

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
	fmt.Printf("Failure    : %v\n", isStatus(exch.current(StartupEvent), statusInternal))
	fmt.Printf("Ok         : %v\n", isStatus(exch.current(StartupEvent), statusOk))

	exch.add(CreateMessage(VirtualHost, "test:package", StartupEvent, StatusInProgress, nil))

	fmt.Printf("Messages   : %v\n", exch.count(StartupEvent))
	fmt.Printf("InProgress : %v\n", isStatus(exch.current(StartupEvent), StatusInProgress))
	fmt.Printf("Failure    : %v\n", isStatus(exch.current(StartupEvent), statusInternal))
	fmt.Printf("Ok         : %v\n", isStatus(exch.current(StartupEvent), statusOk))

	exch.add(CreateMessage(VirtualHost, "test:package", StartupEvent, statusInternal, nil))

	fmt.Printf("Messages   : %v\n", exch.count(StartupEvent))
	fmt.Printf("InProgress : %v\n", isStatus(exch.current(StartupEvent), StatusInProgress))
	fmt.Printf("Failure    : %v\n", isStatus(exch.current(StartupEvent), statusInternal))
	fmt.Printf("Ok         : %v\n", isStatus(exch.current(StartupEvent), statusOk))

	exch.add(CreateMessage(VirtualHost, "test:package", StartupEvent, statusOk, nil))

	fmt.Printf("Messages   : %v\n", exch.count(StartupEvent))
	fmt.Printf("InProgress : %v\n", isStatus(exch.current(StartupEvent), StatusInProgress))
	fmt.Printf("Failure    : %v\n", isStatus(exch.current(StartupEvent), statusInternal))
	fmt.Printf("Ok         : %v\n", isStatus(exch.current(StartupEvent), statusOk))

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
