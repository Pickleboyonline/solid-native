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

	// ======= Setup functions

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
		// Stack: [ nodeId key value ]
		// Get second arg type
		nodeId := ctx.GetInt(-3)
		key := ctx.GetString(-2)
		valueType := ctx.GetType(-1)

		stashKeyName := fmt.Sprintf("props:%d:%s", nodeId, key)

		// Primatives are easy since they are just a copy. However, objects are harder since
		// we dont have direct access.
		// What would be easier to to store a hashmap stash of stings in the form:
		// "prop:${nodeId}:{propName}" : "value".
		// This way the value can be accessed anywhere at anytime, and we can grab the value at
		// any time.
		// TODO: Understand how the syncing works accross threads. For now,
		// we are doing single threaded stuff so this doesnt matter.

		// Put the Stash on top:
		ctx.PushGlobalStash()               // => [ nodeId key value stash ]
		ctx.Replace(-3)                     // => [ nodeId stash value ]
		ctx.PutPropString(-2, stashKeyName) // => [ nodeId stash ]
		ctx.Pop()

		s.SetNodeProp(nodeId, key, NewJsValue(valueType, stashKeyName, s))

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
