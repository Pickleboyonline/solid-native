package snmobile

import (
	"encoding/json"
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

		if ctx.IsFunction(-1) {
			log.Printf("Key \"%v\" for node %v was set with a function. Skipping as not supported...", key, nodeId)
			return 0
		}

		if ctx.IsUndefined(-1) {
			// We pass in an undefined thing here.
			s.setNodeProp(nodeId, key,
				&JSValue{})
		}

		encodedJson := ctx.JsonEncode(-1)

		var unmarshaledValue interface{}

		json.Unmarshal([]byte(encodedJson), &unmarshaledValue)

		log.Printf("Json of prop: %v", unmarshaledValue)

		s.setNodeProp(nodeId, key,
			&JSValue{
				data: unmarshaledValue,
			})

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
