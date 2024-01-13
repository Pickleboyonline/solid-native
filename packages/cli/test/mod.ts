// import { createRenderer } from "solid-js/universal";

// console.log("Hello world!");

// const PROPERTIES = new Set(["className", "textContent"]);

// type Node = any;
// type TextNode = any;

// export const {
//   render,
//   effect,
//   memo,
//   createComponent,
//   createElement,
//   createTextNode,
//   insertNode,
//   insert,
//   spread,
//   setProp,
//   mergeProps,
// } = createRenderer({
//   createElement(name: string) {
//     return document.createElement(name);
//   },
//   createTextNode(value: string) {
//     return document.createTextNode(value);
//   },
//   // ! Causes re-render
//   replaceText(textNode: TextNode, value) {
//     textNode.data = value;
//   },
//   /**
//    * Any time this changes, we just call the "render"
//    * function to return a new swift ui struct.
//    *
//    * Rendering happens async,
//    */
//   // ! Causes re-render
//   setProperty(node: Node, name: string, value: any) {
//     if (name === "style") Object.assign(node.style, value);
//     else if (name.startsWith("on")) node[name.toLowerCase()] = value;
//     else if (PROPERTIES.has(name)) node[name] = value;
//     else node.setAttribute(name, value);
//   },
//   // Causes re-render
//   insertNode(parent: Node, node, anchor) {
//     parent.insertBefore(node, anchor);
//   },
//   // ! Causes re-render
//   isTextNode(node) {
//     return node.type === 3;
//   },
//   // ! Causes re-render
//   removeNode(parent, node) {
//     parent.removeChild(node);
//   },
//   // ! Causes re-render
//   getParentNode(node) {
//     return node.parentNode;
//   },
//   // ! Causes re-render
//   getFirstChild(node) {
//     return node.firstChild;
//   },
//   // ! Causes re-render
//   getNextSibling(node) {
//     return node.nextSibling;
//   },
// });

// type ICreateElementNode = {};

// // We need to have a tree like stucture to represent nodes.
// /**
//  * This is the SNUIViewRegistry.
//  * The tree is a
//  */

// const createElementNode = ({}: ICreateElementNode) => {
//   /**
//    * Potenially have getters/setters
//    */
//   let firstChild = null;
//   let parentNode = null;
//   let nextSibling = null;

//   /**
//    * Quick lookup of nodes.
//    */
//   let descendants = new Set<Node>();

//   /**
//    * O(n).
//    * Could be less time if in a set
//    * @param node
//    */
//   const removeChild = (_node: any) => {
//     // Search for nodes in set
//   };

//   const render = () => {};
// };

// /**
//  * OK, so I need a way to reference and create NODES (Objects) that build swift UI.
//  * Thats taken care of by createElement and text node.
//  *
//  * Not bad, BUT we have a problem. We need a way to re-render this.
//  * This can either be done 2 ways:
//  * 1. Create my own VDOM that manually updates and manges swift UI structs
//  * 2. Create Observable objects that when mutated cause updates.
//  *
//  * The first approach is most likely better. With obersable objects, I must have references
//  * to other observable objects (????). Actually, maybe not. I'd prefer to just use Swift UI.
//  *
//  * But how, you might ask?
//  *
//  * The reason being is that I want SwiftUI to handle changes, not me. I dont know what optimizations need to be made
//  * Thus, its better to have these class wrappers that manage observable objects and connect to other structs.
//  *
//  *
//  * Structs require certain properties to be rendered, lets say their props. We can place this in
//  * an observable object.
//  *
//  * Props work just like react. We capture the component with another component and render like another other.
//  *
//  * BUT the trick comes when we have children.
//  *
//  * This is again, a prop, but a special one because it must refer to other objects.
//  */

// // We need a published prop that we can update that indicates that a change has been made.
// // The prop is an array of other references to other modules,
