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
- Make algor
*/
