package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}


func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

// Same as above, different calling convention
func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}