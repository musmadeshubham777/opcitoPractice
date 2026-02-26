package main

import (
	"fmt"
	"time"
)

func totalTime(start time.Time) {
	fmt.Printf("Total time taken %f seconds", time.Since(start).Seconds())
}

func test() {

	start := time.Now()
	defer totalTime(start)

	time.Sleep(2 * time.Second)
	fmt.Println("Sleep complete")

}

func main() {
	test()
}

//rmg full form in hr ips  oplnnovates
//When a function has multiple defer calls, they are pushed to a stack and executed in Last In First Out (LIFO) order.
//defer delays execution until sorrounding code gets executed
//
