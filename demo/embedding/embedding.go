package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}
