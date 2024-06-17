package main

import (
	"fmt"
	mobiles "nativecore/lib/mobile"

	"gopkg.in/olebedev/go-duktape.v3"
)

func main() {
	ctx := duktape.New()
	ctx.PevalString(`2 + 4`)
	result := ctx.GetNumber(-1)
	ctx.Pop()
	mobiles.Hello()
	fmt.Println("result is:", result)

	// To prevent memory leaks, don't forget to clean up after
	// yourself when you're done using a context.
	ctx.DestroyHeap()
}
