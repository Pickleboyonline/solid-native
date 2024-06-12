import { createRenderer } from "solid-js/universal";
import { SolidNativeRenderer } from "./modules/mod.ts";

type Node = {
  id: string;
};

/**
 * When the SolidJS renderer encounters a string or text, it makes a text component.
 * This is an issue if the renderer returns string based id's. Thus, wrap in an object
 * when created to avoid this issue. In the future, we can look into the views
 * exposing a JSValue Object that allows it to be directly manipulated using the same
 * JSValueBuilder Swift Class I made.
 * @param id
 * @returns
 */
const wrapNodeIdInNode = (id: string): Node => ({ id });

export const {
  render: solidRender,
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
} = createRenderer<Node>({
  createElement(nodeName) {
    const id = SolidNativeRenderer.createNodeByName(nodeName);
    return wrapNodeIdInNode(id);
  },
  createTextNode(value) {
    const node = SolidNativeRenderer.createNodeByName("sn_text");
    SolidNativeRenderer.setProp(node, "text", value);

    return wrapNodeIdInNode(node);
  },
  replaceText({ id }, value) {
    SolidNativeRenderer.setProp(id, "text", value);
  },
  setProperty({ id }, propertyName, value) {
    SolidNativeRenderer.setProp(id, propertyName, value);
  },
  insertNode({ id: parentId }, { id: nodeId }, anchor) {
    const anchorId = anchor?.id;
    SolidNativeRenderer.insertBefore(parentId, nodeId, anchorId);
  },
  isTextNode({ id }) {
    return SolidNativeRenderer.isTextElement(id);
  },
  removeNode({ id: parentId }, { id: nodeId }) {
    return SolidNativeRenderer.removeChild(parentId, nodeId);
  },
  getParentNode({ id }) {
    const parentId = SolidNativeRenderer.getParent(id);
    if (parentId) {
      return wrapNodeIdInNode(parentId);
    }
    return undefined;
  },
  getFirstChild({ id }) {
    const firstChildId = SolidNativeRenderer.getFirstChild(id);
    if (firstChildId) {
      return wrapNodeIdInNode(firstChildId);
    }
    return undefined;
  },
  getNextSibling({ id }) {
    const nextSiblingId = SolidNativeRenderer.next(id);
    if (nextSiblingId) {
      return wrapNodeIdInNode(nextSiblingId);
    }
    return undefined;
  },
});

export const render = (code: () => JSX.Element) => {
  // deno-lint-ignore ban-ts-comment
  // @ts-ignore
  solidRender(code, wrapNodeIdInNode(SolidNativeRenderer.getRootView()));
};
