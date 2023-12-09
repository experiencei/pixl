package main

import "fmt"
import "time"

type ControlMsg int

type Job struct {
	data   int
	result int
}

const (
	DoExit = iota
	ExitOk
)

func doubler(jobs, results chan Job, control chan ControlMsg) {
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
		case job := <-jobs:
			results <- Job{data: job.data, result: job.data * 2}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}