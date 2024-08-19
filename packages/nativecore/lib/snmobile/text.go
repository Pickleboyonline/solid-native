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
