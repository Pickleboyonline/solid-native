package main

// /*
// #cgo CFLAGS: -I${SRCDIR}/yoga
// #cgo LDFLAGS: ${SRCDIR}/yoga/build/yoga/libyogacore.a -lc++
// #include <yoga/Yoga.h>
// */
// import "C"

import (
	"fmt"
	mobiles "nativecore/lib/mobile"
	"nativecore/lib/yoga"

	"gopkg.in/olebedev/go-duktape.v3"
)

// TODO: Get
func main() {

	ctx := duktape.New()
	ctx.PevalString(`2 + 4`)
	result := ctx.GetNumber(-1)
	ctx.Pop()
	mobiles.Hello()
	fmt.Println("result is:", result)

	node := yoga.NewNode()
	defer node.Free()

	node.SetHeight(10)

	node.CalculateLayout(100, 100, yoga.DirectionLTR)

	fmt.Printf("Something %v\n", node.GetLayoutHeight())

	// To prevent memory leaks, don't forget to clean up after
	// yourself when you're done using a context.
	ctx.DestroyHeap()
}
