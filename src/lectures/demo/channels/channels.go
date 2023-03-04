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

func doubler(job <-chan Job, result chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOk
				return
			default:
				panic("Unhandled control message")
			}
		case job := <-job:
			result <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {
	// Taking a input and double it
	// Job channel
	// Results
	// Controllers
	jobs := make(chan Job, 50)
	results := make(chan Result, 50)
	control := make(chan ControlMsg, 50)

	go doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Println("Result: ", result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Time out")
			control <- DoExit
			<-control
			fmt.Println("Exitted Program")
			return
		}
	}
}
