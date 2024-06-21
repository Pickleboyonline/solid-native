package snmobile

import (
	"fmt"
	"log"

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
		// => [ nodeType ]
		nodeType := ctx.GetString(-1)

		nodeId := s.createNode(nodeType)

		log.Printf("New Node create of type %v with id %v", nodeType, nodeId)
		ctx.PushString(nodeId) // => [ nodeType nodeId ]

		return 1 // Return top of stack
	})

	addGoFunc("setProp", func(ctx *duktape.Context) int {
		// Stack: [ nodeId key value ]
		// Get second arg type
		nodeId := ctx.GetString(-3)
		key := ctx.GetString(-2)
		valueType := ctx.GetType(-1)

		// We dont need to free this JS value from the stash
		// since it gets hased to the same key
		stashKeyName := fmt.Sprintf("props:%s:%s", nodeId, key)

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

		s.setNodeProp(nodeId, key, NewJsValue(valueType, stashKeyName, s))

		return 0
	})

	addGoFunc("insertBefore", func(ctx *duktape.Context) int {
		// => [parentId nodeId anchorId?]
		parentId := ctx.GetString(-3)
		nodeId := ctx.GetString(-2)
		anchorType := ctx.GetType(-1)
		isAnchorAString := anchorType.IsString()

		var anchorId string

		if isAnchorAString {
			anchorId = ctx.GetString(-1)
		}

		log.Printf("Node Insert called with child %v  to be inserted under parent %v", nodeId, parentId)
		s.insertBefore(parentId, nodeId, anchorId)
		log.Printf("Node %v inserted under parent %v", nodeId, parentId)
		return 0
	})

	addGoFunc("isTextElement", func(ctx *duktape.Context) int {
		// => [nodeId]
		nodeId := ctx.GetString(-1)
		isTextElement := s.hostReceiver.IsTextElement(nodeId)

		ctx.PushBoolean(isTextElement)

		return 1
	})

	addGoFunc("removeChild", func(ctx *duktape.Context) int {
		// => [parentId nodeId]
		parentId := ctx.GetString(-2)
		nodeId := ctx.GetString(-1)

		s.removeChild(parentId, nodeId)
		return 0
	})

	addGoFunc("getParent", func(ctx *duktape.Context) int {
		// => [nodeId]
		nodeId := ctx.GetString(-1)

		parentId, exists := s.getParent(nodeId)

		if !exists {
			return 0
		}

		ctx.PushString(parentId) // => [nodeId parentId]

		return 1
	})

	addGoFunc("getFirstChild", func(ctx *duktape.Context) int {
		// => [nodeId]
		nodeId := ctx.GetString(-1)

		firstChildId, exists := s.getFirstChild(nodeId)

		if !exists {
			return 0
		}

		ctx.PushString(firstChildId) // => [nodeId firstChildId]

		return 1
	})

	addGoFunc("getNextSibling", func(ctx *duktape.Context) int {
		// => [nodeId]
		nodeId := ctx.GetString(-1)

		nextSiblingId, exists := s.getNextSibling(nodeId)

		if !exists {
			return 0
		}

		ctx.PushString(nextSiblingId) // => [nodeId nextSiblingIndex]

		return 1
	})

	addGoFunc("getRootView", func(ctx *duktape.Context) int {

		ctx.PushString(s.rootNodeId) // => [nodeId rootNodeId]

		return 1
	})

	// Final Step, but the renderer there:
	s.dukContext.PutPropString(-2, "_SolidNativeRenderer") // => [global]
	s.dukContext.Pop()                                     // => []
}
