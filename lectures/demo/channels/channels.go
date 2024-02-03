package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

func dblr(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOk
				return
			default:
				panic("unhandled error")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {
	jobs := make(chan Job, 30)
	results := make(chan Result, 30)
	control := make(chan ControlMsg)

	go dblr(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(5000 * time.Millisecond):
			fmt.Println("timed out")
			control <- DoExit
			<-control
			fmt.Println("exit")
			return

		}
	}
}

func doubler(jobs chan Job, results chan Result, control chan ControlMsg) {
	panic("unimplemented")
}

// // using a buffered channel with size 2
// channel := make(chan int, 2)

// go func() { channel <- 1 }()
// go func() { channel <- 2 }()
// // send to channel
// go func() { channel <- 3 }()
// go func() { channel <- 4 }()

// // recieve from channel
// first := <-channel
// second := <-channel
// third := <-channel
// fourth := <-channel

// fmt.Println(first, second, third, fourth)

// one := make(chan int)
// two := make(chan int)

// for {
// 	select {
// 	case o := <-one:
// 		fmt.Println("one:", o)
// 	case t := <-two:
// 		fmt.Println("two:", t)
// 	case <-time.After(5000 * time.Millisecond):
// 		fmt.Println("no data to receive")
// 		fmt.Println("timed out")
// 		return
// 	}
// }
