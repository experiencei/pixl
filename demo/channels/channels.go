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