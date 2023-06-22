// TODO:
import { createRenderer } from "npm:solid-js@1.7.6/universal";
import { SolidNativeElement } from "./types.ts";
import { SolidNativeCore, print } from "./solid_native_core.ts";

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
    print('Element Name: ' + elementName)
    return SolidNativeCore.createElement(elementName);
  },
  createTextNode(value) {
    const elementId = SolidNativeCore.createTextElement();
    SolidNativeCore.setPropertyOnElement(elementId, "text", value);
    return elementId;
  },
  replaceText(elementId, value) {
    SolidNativeCore.setPropertyOnElement(elementId, "text", value);
  },
  setProperty(elementId, propertyName, value) {
    SolidNativeCore.setPropertyOnElement(elementId, propertyName, value);
  },
  insertNode(parent, node, anchor) {
    SolidNativeCore.insertElement(parent, node, anchor);
  },
  isTextNode(node) {
    return SolidNativeCore.isTextElement(node);
  },
  removeNode(parent, node) {
    SolidNativeCore.removeElement(parent, node);
  },
  getParentNode(node) {
    return SolidNativeCore.getParentElementId(node);
  },
  getFirstChild(node) {
    return SolidNativeCore.getFirstChildElementId(node);
  },
  getNextSibling(node) {
    return SolidNativeCore.getNextSiblingElementId(node);
  },
});
