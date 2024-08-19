package snmobile

/*
TODO:
We need to create a text struct that outlines the applied text styles given a peice of text.
So a map of styles and then some text.

This needs to be generated every time a text component changes.

Use dfs to build an array of the text, and have that as something that a text component recieves

for both rendering and measurement.

Note that the SwiftUI and UIKit need to have the same font to render out.

Note that the only things that mutate the tree are setProps, insertBefore, and removeNode.

Now, anytime we mutate the tree we find the parent and reconstonstruct the array.

Also there's two text component types => the wrapper and the actual texts.

The wrappers control the styles, and the inner wrappers override the styles of the out wrappers.



Steps:
- Define base data type of text in slice/array
- Make algorithm to generate it
- essentially, its a list of text and associated styles. Essentially, we keep the logic cross platform
and reduce native implementation workload.
*/

type TextDescriptor struct {
	Text   string
	styles map[string]JSValue
}

func (t *TextDescriptor) GetStylesAsJSValue() *JSValue {
	return &JSValue{data: t.styles}
}

// Given a text component, we want to go up to its parent
// On either a:
//   - InsertBefore
//   - RemoveChild
//   - SetNodeProp
//
// TODO: Determine how to handle NodeDescriptor Yoga Nodes
// TODO: Most likely just wont have one to begin with.
// TODO: Mark node as dirty when update parent changes.
//
// - Returns the text descriptor and the parent node. DOES NOT MAKE MUTATIONS ON NODES
func generateTextDescriptor(node *NodeContainer) ([]TextDescriptor, *NodeContainer) {
	parentNode := node

	// Get top level node of text
	for {
		if parentNode.parent == nil {
			break
		}

		if !parentNode.parent.isText {
			break
		}

		parentNode = parentNode.parent
	}

	// More of a stack
	textDescriptors := generateTextDescriptorRecursively(parentNode, nil)

	return textDescriptors, parentNode
}

// Pass in parent, get text descriptor. Another benifit is that you can use channels for
// Fan out, fan in.
// Node -> []TextDescriptor
// Styles can be `nil`. If it is, it will try to use the styles from the node (node always has empty styles, so
// long as you made it with the `newNodeContainer`)
func generateTextDescriptorRecursively(node *NodeContainer, styles map[string]JSValue) []TextDescriptor {
	// Could make better by making cap == node.children length
	var textDescriptor []TextDescriptor

	// In general, we want override incoming styles. They can be nil
	var newStyles map[string]JSValue

	if styles == nil {
		newStyles = node.styleMap
	} else if node.styleMap == nil || len(node.styleMap) == 0 {
		newStyles = styles
	} else {
		newStyles = map[string]JSValue{}
		// Override them
		for k, v := range styles {
			newStyles[k] = v
		}

		for k, v := range node.styleMap {
			newStyles[k] = v
		}
	}

	if (node.children == nil || len(node.children) == 0) && node.text != "" {
		textDescriptor = []TextDescriptor{
			{
				Text:   node.text,
				styles: newStyles,
			},
		}
		return textDescriptor
	}

	textDescriptor = make([]TextDescriptor, 0, len(node.children))

	type ChannelResult struct {
		index int
		data  []TextDescriptor
	}

	resultChan := make(chan ChannelResult)
	resultsBuffer := make([][]TextDescriptor, len(node.children))

	// TODO: determine if channel is better here.
	for i, n := range node.children {
		go func(i int, n *NodeContainer) {
			resultChan <- ChannelResult{
				index: i,
				data:  generateTextDescriptorRecursively(n, newStyles),
			}
		}(i, n)
	}

	for range node.children {
		res := <-resultChan
		resultsBuffer[res.index] = res.data
	}

	close(resultChan)

	for _, d := range resultsBuffer {
		textDescriptor = append(textDescriptor, d...)
	}

	return textDescriptor
}

// Same but immutable
func generateTextDescriptorImmutable(node *NodeContainer) []TextDescriptor {
	// textDescriptor :=
	return make([]TextDescriptor, 0)
}

// Some notes on operation:

// QUEUE || EXPANDED || PROCESSED_TEXT_DESCRIPTION

// [2, 3] || [1] || []
// [4, 2, 3] || [2, 1] || []
// [2, 3]  || [2, 1] || [4]
// [2, 3]  || [2, 1] || [4] => Must revert changes of 2, back to 1. In this case its just a simple reversion but later could be more difficult.
// => General reversion just means to take styles of previous expansion of the keys of the 2, Which isn't that bad.
// We revert, pop (both on queue and expanded queue) when we know its been expanded
// [3]  || [1] || [4]
// [5, 3]  || [3, 1] || [4]
// [5, 3]  || [3, 1] || [4, 5]
// [3]  || [3, 1] || [4, 5] => If already expanded and last one, just break, we are done!
// [3]  || [1, 2, 3] || [4, 5] => ^^^ Exit

/**
<Text style={{color:"red"}}> <= 1
	<Text style={{color:"blue"}}>4</Text> <=2
	<Text style={{color:"blue"}}>5</Text> <=3
</Text>
*/

// Goal: Generate an array of TextDescriptors that determine the props.
// Need to know whether node is a text component or not so we know to dispatch it.
// determine on text node creation.
// TODO: When a text component has a parent text component, DO NOT notify the host of that component
// Child text components are a "Shadow nodes" and the host platform is only aware of the parent.
// That being said, text components should NOT have a yoga node, only the upper most actually has a relationship in the tree
// Anyways, lets continue to the algorithm

// Next, we traverse up to the top text component tree to find the parent text component.
// There are two types of text nodes: containers and actualText. Containers can have children but not text. ActualText has no children, only one parent, and text

// Our job is to determine the styles. Note that inner styles always override parent styles. Another thing is that the final
// TextDiscriptor array is only ActualText nodes.

// Algorithm:

// Find parent text node, such that the parent is not a text node. If the node is the parent, or the node has no parents, assume given node is parent.

// Whether a node is a container or not depends on whether it has children. If children, update the styles, and add to queue. Also mark that weve expanded children.
// If we know that there are children and the node has been expanded, "process it" and remove it from queue

// ? Question, how is Yoga going to handle styles? may need to look at RN for this. Naive approach is to make root node have only yoga node.
// ? But, then some styles are not going to be there. I would say it makes it a lot easier if top text node is treated like the only node.
// ? If one needs more complex layouts, wrap with regular view and use flexbox. I THINK RN does allow children to have flexbox styles,
// ? But they tend to not work as reliably compared to wrapping with views. Can experiment with RN later.
// ! ^^^ Need to figure this out before working on this more.
// TODO: Determine how RN handles text, yoga styles, and shadow text nodes (does it dispatch it to host platform?)

// Code for queue method:
/**


// Used to determine whether the node has been fully processed.
hasExpandedChildrenMap := map[*NodeContainer]struct{}{}

queue := []*NodeContainer{parentNode}

// Used to maintain styles
expandedQueue := make([]*NodeContainer, 0)

// Used for new text
stylesBuffer := map[string]JSValue{}

// bruh
dequeueQueue := func() {
	queue = slices.Delete(queue, 0, 1)
}
dequeueExpandedQueue := func() {
	expandedQueue = slices.Delete(expandedQueue, 0, 1)
}

applyStyles := func(currentNode *NodeContainer) {
	if currentNode.styleMap == nil {
		return
	}
	for key, value := range currentNode.styleMap {
		stylesBuffer[key] = value
	}
}

revertStyles := func(currentNode *NodeContainer, previousNode *NodeContainer) {
	if currentNode.styleMap == nil {
		return
	}

	if previousNode.styleMap == nil {
		for key := range currentNode.styleMap {
			delete(stylesBuffer, key)
		}
		return
	}

	for key := range currentNode.styleMap {
		prevNodeValue, exists := previousNode.styleMap[key]
		if exists {
			stylesBuffer[key] = prevNodeValue
		} else {
			delete(stylesBuffer, key)
		}
	}
}

for len(queue) != 0 {
	// If node has children, its not text.
	currentNode := queue[0]

	_, hasExpanded := hasExpandedChildrenMap[currentNode]

	// Node is last in queue and already expanded. We can go ahead and end it.
	if hasExpanded && len(queue) == 1 {
		break
	}

	// Node has already expanded, which means we need to revert its styles
	// and pop from stack
	if hasExpanded {
		// This requires the current and previous expanded node
		currentExpandedNode := currentNode
		prevExpandedNode := expandedQueue[1]
		revertStyles(currentExpandedNode, prevExpandedNode)

		// Go ahead an dequeue the expanded and queue
		dequeueExpandedQueue()
		dequeueQueue()
		continue
	}

	// Node has children and is this a container and has no
	// underlying text.
	if len(currentNode.children) != 0 {
		// dequeueQueue()
		queue = slices.Insert(queue, 0, currentNode.children...)
		// Mark that its been expanded so we don't reprocess it.
		hasExpandedChildrenMap[currentNode] = struct{}{}
		expandedQueue = slices.Insert(expandedQueue, 0, currentNode)
		applyStyles(currentNode)
		continue
	}

	// Make a new styles so that everything does point to the same buffer.
	// ? Maybe a better way to do this?
	newStylesBuffer := map[string]JSValue{}

	for key, value := range stylesBuffer {
		newStylesBuffer[key] = value
	}

	// Does not have children, which means its a text node
	// Go ahead and add it to the list, we are done.
	// TODO: Make proper text descriptor
	textDescriptors = append(textDescriptors, TextDescriptor{
		Text:   currentNode.text,
		styles: newStylesBuffer,
	})
	dequeueQueue()

}
*/
