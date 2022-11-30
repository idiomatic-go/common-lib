package vhost

import (
	"fmt"
	"sync"
	"time"
)

type Request func() int

var evens = []int{2, 4, 6, 8, 10, 12, 14, 16}
var odds = []int{1, 3, 5, 7, 9, 11, 13, 15}
var odd3 = []int{1, 3, 5}
var even3 = []int{2, 4, 6}

func wait(name string, ms int64) {
	if len(name) > 0 {
		fmt.Printf("%v : waiting %v seconds\n", name, ms)
	}
	time.Sleep(time.Second * time.Duration(ms))
}

func ExampleMain() {

	//fmt.Println("\nTimeoutOnly : start")
	//go example_TimeoutOnly()
	//wait("TimeoutOnly", 10)

	//fmt.Println("\nTimeoutOnlyClose : start")
	//go example_TimeoutOnlyClose()
	//wait("TimeoutOnlyClose", 10)

	//fmt.Println("\nBatchEvens : start")
	//go example_BatchEvens()
	//wait("BatchEvens", 10)

	// No timeouts
	//fmt.Println("\nBatchEvensClose : start")
	//go example_BatchEvensClose()
	//wait("BatchEvensClose", 10)

	//fmt.Println("\nBatchEvensMultipleSenders : start")
	//go example_BatchEvensMultipleSenders()
	//wait("BatchEvensMultipleSenders", 10)

	//fmt.Println("\nBatchWithTimeout : start")
	//go example_BatchWithTimeout()
	//wait("BatchWithTimeout", 10)

	//fmt.Println("\nBatchWithTimeout_Timeout : start")
	//go example_BatchWithTimeout_Timeout()
	//wait("BatchWithTimeout_Timeout", 10)

	//fmt.Println("\nBatchWithTimeout_NoTimeout : start")
	//go example_BatchWithTimeout_NoTimeout()
	//wait("BatchWithTimeout_NoTimeout", 10)

	//fmt.Println("\nPollingNilRequest : start")
	//go example_PollingNilRequest()
	//wait("PollingNilRequest", 10)

	fmt.Println("\nPollingRequest : start")
	go example_PollingRequest()
	wait("PollingRequest", 10)

	// Uncomment the /* */ to run this example
	/*
		//Output:
		// fail
	*/
}

func example_TimeoutOnly() {
	c := make(chan int)

	go receive(c)
}

func example_TimeoutOnlyClose() {
	c := make(chan int)

	go receive(c)
	wait("TimeoutOnlyClose", 3)
	close(c)
}

func example_BatchEvens() {
	c := make(chan int)

	go receive(c)
	send(c, evens)
}

func example_BatchEvensClose() {
	c := make(chan int)

	go receive(c)
	send(c, evens)
	close(c)
}

func example_BatchEvensMultipleSenders() {
	c := make(chan int)

	go receive(c)
	go send(c, evens)
	go send(c, odds)
}

func example_BatchWithTimeout() {
	c := make(chan int)

	go receive(c)
	go send(c, even3)
}

func example_BatchWithTimeout_Timeout() {
	c := make(chan int)

	go receive(c)
	go send(c, even3)
	wait("", 2)
	go send(c, odd3)
}

func example_BatchWithTimeout_NoTimeout() {
	c := make(chan int)

	go receive(c)
	go send(c, even3)
	wait("", 2)
	go send(c, odds)
}

func receive(c chan int) {
	var batch = new([2]int)
	var mu sync.Mutex
	var i = 0

	for {
		select {
		case msg, open := <-c:
			// Exit on a closed channel
			if !open {
				fmt.Println("receive : exit")
				return
			}
			fmt.Println("received  : ", msg)
			mu.Lock()
			batch[i] = msg
			i++
			if i == 2 {
				go processBatch(batch)
				i = 0
				batch = new([2]int)
			}
			mu.Unlock()
		case <-time.After(time.Millisecond * 500):
			fmt.Println("received  : timeout")
			mu.Lock()
			if i > 0 {
				go processBatch(batch)
				i = 0
				batch = new([2]int)
			}
			mu.Unlock()
		}
	}
}

func processBatch(batch *[2]int) {
	for _, v := range *batch {
		if v > 0 {
			fmt.Println("processed : ", v)
		}
	}
}

func send(c chan int, items []int) {
	for _, v := range items {
		c <- v
	}
}

func example_PollingNilRequest() {
	c := make(chan int)
	stop := make(chan struct{}, 1)
	go receivePoll(c)
	go poll(c, time.Millisecond*500, stop, time.Second*1, nil)
}

func example_PollingRequest() {
	c := make(chan int)
	stop := make(chan struct{}, 1)
	var count int

	go receivePoll(c)
	go poll(c, time.Millisecond*500, stop, time.Second*1, func() int {
		count++
		return count
	})
	wait("", 4)
	close(stop)
}

func poll(data chan int, dataPolling time.Duration, stop chan struct{}, stopPolling time.Duration, req Request) {
	stopTick := time.Tick(dataPolling)
	pollTick := time.Tick(stopPolling)

	for {
		// Using seperate select statements because if the both of the ticks ocurred at the same time, one tick would be
		// ignored possibly leading to starvation of the close message.
		// Go documentation states that the selection of a case statement is indeterminate when both cases are asserted at
		// the same time.
		if stop != nil {
			select {
			case <-stopTick:
				select {
				case <-stop:
					fmt.Println("polling : closed")
					// Cleanup t ?
					return
				default:
				}
			default:
			}
		}
		select {
		case <-pollTick:
			if req != nil {
				fmt.Println("polling : send")
				data <- req()
			}
		default:
		}
	}
}

func receivePoll(c chan int) {
	for {
		select {
		case msg, open := <-c:
			// Exit on a closed channel
			if !open {
				fmt.Println("receivePoll : exit")
				return
			}
			fmt.Println("received  : ", msg)
		}
	}
}
