package snmobile

import (
	"fmt"

	"gopkg.in/olebedev/go-duktape.v3"
)

// Internal function to setup the rendeder
//   - How to throw errors: https://duktape.org/api.html#concepts.9
func (s *SolidNativeMobile) registureRenderer() {
	// Push global object, we'll use it later
	s.dukContext.PushGlobalObject() // => [ globalThis ]

	s.dukContext.PushObject() // => [globalThis obj]

	// Start our functions for solid js. Honesly rn not sure how i want tos tructure the API
	// for now, just slap them in a global object called _SolidNativeRenderer

	addGoFunc := func(cbName string, cb func(ctx *duktape.Context) int) {
		// [global obj]
		s.dukContext.PushGoFunction(cb) // => [global obj func]
		s.dukContext.PutPropString(-2, cbName)
		// => [global obj]
	}

	addGoFunc("createNodeByName", func(ctx *duktape.Context) int {
		nodeType := ctx.GetString(-1)
		fmt.Println("Node type: ", nodeType)
		return 0
	})

	addGoFunc("setProp", func(ctx *duktape.Context) int {
		nodeType := ctx.GetString(-1)
		fmt.Println("Node type: ", nodeType)
		return 0
	})

	addGoFunc("insertBefore", func(ctx *duktape.Context) int {
		nodeType := ctx.GetString(-1)
		fmt.Println("Node type: ", nodeType)
		return 0
	})

	addGoFunc("isTextElement", func(ctx *duktape.Context) int {
		nodeType := ctx.GetString(-1)
		fmt.Println("Node type: ", nodeType)
		return 0
	})

	addGoFunc("removeChild", func(ctx *duktape.Context) int {
		nodeType := ctx.GetString(-1)
		fmt.Println("Node type: ", nodeType)
		return 0
	})

	addGoFunc("getParent", func(ctx *duktape.Context) int {
		nodeType := ctx.GetString(-1)
		fmt.Println("Node type: ", nodeType)
		return 0
	})

	addGoFunc("getFirstChild", func(ctx *duktape.Context) int {
		nodeType := ctx.GetString(-1)
		fmt.Println("Node type: ", nodeType)
		return 0
	})

	addGoFunc("next", func(ctx *duktape.Context) int {
		nodeType := ctx.GetString(-1)
		fmt.Println("Node type: ", nodeType)
		return 0
	})

	// Final Step, but the renderer there:
	s.dukContext.PutPropString(-2, "_SolidNativeRenderer") // => [global]
	s.dukContext.Pop()                                     // => []
}
