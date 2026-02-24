// Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.
// capacity in the above syntax should be greater than 0 for a channel to have a buffer.

// Length vs Capacity
// The capacity of a buffered channel is the number of values that the channel can hold. This is the value we specify when creating the buffered channel using the make function.

// The length of the buffered channel is the number of elements currently queued in it.

package main

import (
	"fmt"
	//    "time"
)

func mainBufferedChannel() {
	// fmt.Println("Pass")

	// bc :=make(chan string , 2)

	// bc <- "shubham"
	// bc <- "musmade"
	// display(bc)

	// bc := make(chan int , 3)
	// fmt.Println(bc)
	// write(bc)
	// time.Sleep(2 * time.Second)

	// for  {
	//     v , ok := <- bc
	//     if ok {
	//          fmt.Println(v)
	//     }
	// }

	// for v := range bc {
	// fmt.Println(v)
	// }

	ch1 := make(chan string, 6)

	ch1 <- "Shu"
	ch1 <- "Nik"
	ch1 <- "Sid"
	ch1 <- "Pra"
	ch1 <- "Tan"
	ch1 <- "san"
	close(ch1) //if i dont close here it will raise error.

	//cannot use index witth channels

	// for v := range ch1 {
	//     fmt.Println(v)
	// }

	// for i := 0; i <= cap(ch1); i++ {
	//         fmt.Println(<-ch1)
	// }

	// for  {
	//     v , ok := <- ch1
	//     if ok {
	//          fmt.Println(v)
	//     }
	// }

	fmt.Println("Capacity of channel", cap(ch1))
	fmt.Println("Length of channel", len(ch1))

}

// func write(ch chan int){
//     for i := 0; i < 5; i++ {
//         ch <- i
//     }
//     close(ch)

// }
