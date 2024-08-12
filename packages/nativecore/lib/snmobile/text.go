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

type TypeDescriptor struct {
	text   string
	styles map[string]interface{}
}

// Given a text component, we want to go up to its parent
// On either a:
//   - InsertBefore
//   - RemoveChild
//   - SetNodeProp
func updateTextComponent() {
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

}
