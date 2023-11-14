package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	// Using dot notation automatically dereferences. No asterisk(*) needed.
	counter.hits += 1
	fmt.Println("Counter", counter)
}