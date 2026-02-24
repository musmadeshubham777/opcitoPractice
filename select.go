package main

import (
	"fmt"

	"time"
	// "golang.org/x/text/cases"
)

// The select statement is used to choose from multiple send/receive channel operations. The select statement blocks until one of the send/receive operations is ready.
// The syntax is similar to switch except that each of the case statements will be a channel operation. Let’s dive right into some code for better understanding.

func server1(a chan string) {
	time.Sleep(2 * time.Second)
	a <- "Server ! running"
}

func server2(a chan string) {
	time.Sleep(1 * time.Second)
	a <- "Server @ running"
}

func processing(s chan string) {
	time.Sleep(10 * time.Second)

	s <- "Passing Value"
}

func mainSelect() {

	// output1 := make(chan string)
	// output2 := make(chan string)

	// go server1(output1)
	// go server2(output2)

	// select {
	//whichever gets ready first get ececuted
	//if both gets ready at same time , then random case get executed

	// case sq := <-output1:
	// 	fmt.Println(sq)

	// case s1 := <-output2:
	// 	fmt.Println(s1)

	// default:
	// 	fmt.Println("no value received")

	// }

	oth := make(chan string)
	go processing(oth)
	for {

		time.Sleep(1 * time.Second)

		select {
		case v := <-oth:
			fmt.Println(v)
		default:
			fmt.Println("Value not recieved")
		}
	}

}
