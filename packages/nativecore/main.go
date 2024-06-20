package main

// /*
// #cgo CFLAGS: -I${SRCDIR}/yoga
// #cgo LDFLAGS: ${SRCDIR}/yoga/build/yoga/libyogacore.a -lc++
// #include <yoga/Yoga.h>
// */
// import "C"

import (
	mobiles "nativecore/lib/snmobile"
)

// TODO: Get
func main() {
	m := mobiles.NewSolidNativeMobile(nil)
	defer m.FreeMemory()

	// m.RunJs()

	// ctx := duktape.New()
	// ctx.PevalString(`2 + 4`)
	// result := ctx.GetNumber(-1)
	// ctx.Pop()
	// mobiles.Hello()
	// fmt.Println("result is:", result)

	// node := yoga.NewNode()
	// defer node.Free()

	// node.SetHeight(10)

	// node.CalculateLayout(100, 100, yoga.DirectionLTR)

	// fmt.Printf("Something %v\n", node.GetLayoutHeight())

	// // To prevent memory leaks, don't forget to clean up after
	// // yourself when you're done using a context.
	// ctx.DestroyHeap()
}
