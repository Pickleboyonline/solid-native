import { createRenderer } from "solid-js/universal";

type Node = {
  id: string;
};

/**
 * The SolidNative renderer API exposed by the Rust core
 * Bound to globalThis.solidNative
 */
interface SolidNativeAPI {
  createElement(tag: string): string;
  createTextNode(value: string): string;
  setProp(nodeId: string, key: string, value: unknown): void;
  insertBefore(parentId: string, nodeId: string, anchorId?: string): void;
  removeChild(parentId: string, nodeId: string): void;
  isTextElement(nodeId: string): boolean;
  getParent(nodeId: string): string | null;
  getFirstChild(nodeId: string): string | null;
  getNextSibling(nodeId: string): string | null;
  getRootView(): string;
}

// Access the API from globalThis (set by Rust core)
declare const solidNative: SolidNativeAPI;

/**
 * Cache wrapped nodes for reference checking in SolidJS
 */
const wrappedNodeMap = new Map<string, Node>();

/**
 * When the SolidJS renderer encounters a string or text, it makes a text component.
 * This is an issue if the renderer returns string based IDs. Thus, wrap in an object
 * when created to avoid this issue.
 */
const wrapNodeIdInNode = (id: string): Node => {
  const node = wrappedNodeMap.get(id);
  if (node) {
    return node;
  }
  const newNode = { id };
  wrappedNodeMap.set(id, newNode);
  return newNode;
};

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
    const id = solidNative.createElement(nodeName);
    return wrapNodeIdInNode(id);
  },
  createTextNode(value) {
    const node = solidNative.createElement("sn_text");
    solidNative.setProp(node, "text", value);
    return wrapNodeIdInNode(node);
  },
  replaceText({ id }, value) {
    solidNative.setProp(id, "text", value);
  },
  setProperty({ id }, propertyName, value) {
    solidNative.setProp(id, propertyName, value);
  },
  insertNode({ id: parentId }, { id: nodeId }, anchor) {
    const anchorId = anchor?.id;
    solidNative.insertBefore(parentId, nodeId, anchorId);
  },
  isTextNode({ id }) {
    return solidNative.isTextElement(id);
  },
  removeNode({ id: parentId }, { id: nodeId }) {
    return solidNative.removeChild(parentId, nodeId);
  },
  getParentNode({ id }) {
    const parentId = solidNative.getParent(id);
    if (parentId) {
      return wrapNodeIdInNode(parentId);
    }
    return undefined;
  },
  getFirstChild({ id }) {
    const firstChildId = solidNative.getFirstChild(id);
    if (firstChildId) {
      return wrapNodeIdInNode(firstChildId);
    }
    return undefined;
  },
  getNextSibling({ id }) {
    const nextSiblingId = solidNative.getNextSibling(id);
    if (nextSiblingId) {
      return wrapNodeIdInNode(nextSiblingId);
    }
    return undefined;
  },
});

export const render = (code: () => JSX.Element, rootId?: string) => {
  const root = rootId ?? solidNative.getRootView();
  // @ts-ignore - solidRender expects the root element as second argument
  solidRender(code, wrapNodeIdInNode(root));
};
