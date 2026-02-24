//Go is a concurrent language and not a parallel one.
//Concurrency
// Is the capability to deal with lots of things at once(runing .tying untied shoeless and again start running)
// Download File and render WebPage
// Concurrency is an inherent part of the Go programming language.
// Concurrency is handled in Go using Goroutines and channels.

//Parallelism
// Is doing lots of things at the same time. It might sound similar to concurrency but it’s actually different.(jogging and listening at same time is parrelalism)
//used multiple CPU cores

//  Goroutines communicate using channels.
// data can be sent from one end and received from the other end using channels.
//each channel have type associated with it.No other data type data transamision is not allowed
// chan T     ==> this is channel of type T
// zero value of channel is nil. nil channels are of no use .It has to be created using make.

//Deadlock
// One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.

// package main
// func main() {
// 	ch := make(chan int)
// 	ch <- 5
// }
//in this we are only sending data on channel no go routine is not recieving data .so this raises error.
// fatal error: all goroutines are asleep - deadlock!

package main

import "fmt"

func mainChannels() {

	// data := <- a // read from channel a
	// a <- data // write to channel a
	// var a chan int
	// a := make(chan int)
	// a <- data  //writing data to the channel
	// The direction of the arrow with respect to the channel specifies whether the data is sent or received.

	// tdata := <-a  //reading data from channel
	// fmt.Println("data recieved to the channel", tdata)

	// if a == nil {
	// 	fmt.Println("Assignin to nil channel.")

	// 	a = make(chan int)

	// 	fmt.Printf("Type of a is %T", a)
	// }

	num := 589
	s_channel := make(chan int)
	c_channel := make(chan int)

	go calculateSquare(num, s_channel)
	go calculateCubes(num, c_channel)

	square, cubes := <-s_channel, <-c_channel //reading from channel

	// fmt.Println("Square:", square)
	// fmt.Println("Cubes:", cubes)
	fmt.Println("Final Output:", cubes+square)

	// uni :=make(chan <- int) //we can only send data to the channel.
	// go sendData(uni)
	// fmt.Println(<- uni) // not possible to get data from uni direction channel

	// This is where channel conversion comes into use. /
	// It is possible to convert a bidirectional channel to a send only or receive only channel
	// but not the uni cannot be converted to bi directional.

	ch := make(chan int)
	go producer(ch)

	for {
		v, ok := <-ch
		if ok {
			fmt.Println("Recieved", v, ok)
		}
	}
}

func producer(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)

}

// func sendData(ch chan <- int){
//     fmt.Println(ch)
// }

func calculateSquare(num int, s_channel chan int) {
	sum := 0
	for num != 0 {

		digit := num % 10
		sum += digit * digit
		num /= 10

	}

	s_channel <- sum

}
func calculateCubes(num int, c_channel chan int) {
	sum := 0
	for num != 0 {

		digit := num % 10
		sum += digit * digit * digit
		num /= 10

	}

	c_channel <- sum

}
