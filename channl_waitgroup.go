// A WaitGroup is used to wait for a collection of Goroutines to finish executing.
// The control is blocked until all Goroutines finish executing.
// Let’s say we have 3 concurrently executing Goroutines spawned from the main Goroutine.
// The main Goroutines needs to wait for the 3 other Goroutines to finish before terminating. This can be accomplished using WaitGroup.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func main(){

//     // fmt.Println("passing")
//     // fmt.Println(time.TimeOnly)
//     var mg sync.WaitGroup

//     for i :=0 ; i<=3 ;i++{
//         // fmt.Println(i)
//         mg.Add(1)
//         go process(i , &mg)
//     }

//     mg.Wait()
//     //add--done--go --wait these are all methods of waitgroup

// }

func process(i int, mg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	mg.Done()

}

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1) //“I am about to start 1 goroutine, so increase the counter by 1.”

		// Done() → decreases the counter by 1
		// Wait() → blocks until the counter becomes 0
		//always add before calling go routine

		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func mainwaitgoup() {

	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

}

//starting is capital then public and if small then private
// Worker Goroutines listen for new tasks on the jobs buffered channel. Once a task is complete, the result is written to the results buffered channel.

// One of the important uses of buffered channel is the implementation of worker pool.
// a worker pool is a collection of threads that are waiting for tasks to be assigned to them. Once they finish the task assigned, they make themselves available again for the next task.
