import { createRenderer } from "solid-js/universal";
import { SolidNativeElement } from "./types.ts";
import { SolidNativeCore } from "./solid_native_core.ts";

export const {
  render,
  effect,
  memo,
  createComponent,
  createElement,
  createTextNode,
  insertNode,
  insert,
  spread,
  setProp,
  mergeProps,
} = createRenderer<SolidNativeElement>({
  createElement(elementName) {
    return SolidNativeCore.createElementByName(elementName);
  },
  createTextNode(value) {
    const element = SolidNativeCore.createTextElement();
    element.setProp("text", value);
    return element;
  },
  replaceText(element, value) {
    element.setProp("text", value);
  },
  setProperty(element, propertyName, value) {
    // print(typeof value);
    element.setProp(propertyName, value);
  },
  insertNode(parent, node, anchor) {
    parent.insertBefore(node, anchor);
  },
  isTextNode(node) {
    return node.isTextElement;
  },
  removeNode(parent, node) {
    parent.removeChild(node);
  },
  getParentNode(node) {
    return node.parentElement;
  },
  getFirstChild(node) {
    return node.firstChild;
  },
  getNextSibling(node) {
    return node.next;
  },
});
