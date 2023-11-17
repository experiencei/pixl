// run with `go run ./demo/pkg/main`
package main

import (
	"pixlProj/demo/pkg/display"
	"pixlProj/demo/pkg/msg"
)

// import (
// 	"coursecontent/demo/pkg/display"
// 	"coursecontent/demo/pkg/msg"
// )

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("calling a package function")
}
