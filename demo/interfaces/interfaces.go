package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string