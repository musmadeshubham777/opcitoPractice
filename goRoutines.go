// goRotines

// go is by default sync to make asyc block we use goRoutines
package main

import (
	// "time"
	"fmt"
	"runtime"
)

// import

func mainRoutines() {

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	fmt.Println("Number of OS threads:", runtime.NumCgoCall())

	go addUpto100()
	go addUpto1000()
	go addUpto10000()
	go addUpto100000()

	fmt.Print("\n \n")

	fmt.Println("GOMAXPROCS After:", runtime.GOMAXPROCS(0))
	fmt.Println("Number of goroutines: after ", runtime.NumGoroutine())
	fmt.Println("Number of OS threads: after", runtime.NumCgoCall())

	sum := 0
	for i := 1; i >= 1000000; i++ {
		sum += i
	}

	fmt.Println("sum of 10,00,000", sum)

}

func addUpto100() {
	sum := 0
	for i := 1; i >= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 100", sum)

}

func addUpto1000() {
	sum := 0
	for i := 1; i >= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 1,000", sum)

}
func addUpto10000() {
	sum := 0
	for i := 1; i >= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 10,000", sum)

}
func addUpto100000() {
	sum := 0
	for i := 1; i >= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 1,00,000", sum)

}
