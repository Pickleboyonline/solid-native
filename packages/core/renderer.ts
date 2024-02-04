import { createRenderer } from "solid-js/universal";
import { SolidNativeRenderer } from "./modules/mod.ts";

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
} = createRenderer<string>({
  createElement(nodeName) {
    return SolidNativeRenderer.createNodeByName(nodeName);
  },
  createTextNode(value) {
    const node = SolidNativeRenderer.createNodeByName("sn_text");
    SolidNativeRenderer.setProp(node, "text", value);

    return node;
  },
  replaceText(node, value) {
    SolidNativeRenderer.setProp(node, "text", value);
  },
  setProperty(node, propertyName, value) {
    SolidNativeRenderer.setProp(node, propertyName, value);
  },
  insertNode(parent, node, anchor) {
    SolidNativeRenderer.insertBefore(parent, node, anchor);
  },
  isTextNode(node) {
    return SolidNativeRenderer.isTextElement(node);
  },
  removeNode(parent, node) {
    return SolidNativeRenderer.removeChild(parent, node);
  },
  getParentNode(node) {
    return SolidNativeRenderer.getParent(node);
  },
  getFirstChild(node) {
    return SolidNativeRenderer.getFirstChild(node);
  },
  getNextSibling(node) {
    return SolidNativeRenderer.next(node);
  },
});
