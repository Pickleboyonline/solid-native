import { JSX } from "solid-js/jsx-runtime";
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

// Helper to safely get id from a node
const getNodeId = (node: unknown, context: string): string => {
  if (!node || typeof node !== 'object' || !('id' in node)) {
    throw new Error(`${context}: invalid node ${JSON.stringify(node)}`);
  }
  const id = (node as Node).id;
  if (id === undefined || id === null) {
    throw new Error(`${context}: node.id is ${id}, node: ${JSON.stringify(node)}`);
  }
  return id;
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
    if (nodeName === undefined || nodeName === null) {
      throw new Error(`createElement: nodeName is ${nodeName}`);
    }
    try {
      const id = solidNative.createElement(nodeName);
      return wrapNodeIdInNode(id);
    } catch (e) {
      throw new Error(`createElement failed for nodeName="${nodeName}": ${e}`);
    }
  },
  createTextNode(value) {
    if (value === undefined) {
      throw new Error(`createTextNode: value is undefined`);
    }
    try {
      const node = solidNative.createElement("sn_text");
      solidNative.setProp(node, "text", String(value));
      return wrapNodeIdInNode(node);
    } catch (e) {
      throw new Error(`createTextNode failed for value="${value}": ${e}`);
    }
  },
  replaceText(node, value) {
    const id = getNodeId(node, "replaceText");
    // Ensure value is a string
    solidNative.setProp(id, "text", value === undefined ? "" : String(value));
  },
  setProperty(node, propertyName, value) {
    const id = getNodeId(node, "setProperty");
    // Don't call setProp with undefined property names
    if (propertyName === undefined || propertyName === null) {
      return;
    }
    solidNative.setProp(id, propertyName, value);
  },
  insertNode(parent, node, anchor) {
    const parentId = getNodeId(parent, "insertNode.parent");
    const nodeId = getNodeId(node, "insertNode.node");
    // Only pass anchorId if anchor exists - rquickjs Opt<String> doesn't handle explicit undefined
    if (anchor) {
      const anchorId = getNodeId(anchor, "insertNode.anchor");
      solidNative.insertBefore(parentId, nodeId, anchorId);
    } else {
      solidNative.insertBefore(parentId, nodeId);
    }
  },
  isTextNode(node) {
    // SolidJS may pass null/undefined during cleanup - return false for invalid nodes
    if (!node || typeof node !== 'object' || !('id' in node)) {
      return false;
    }
    const id = (node as Node).id;
    if (id === undefined || id === null) {
      return false;
    }
    return solidNative.isTextElement(id);
  },
  removeNode(parent, node) {
    const parentId = getNodeId(parent, "removeNode.parent");
    const nodeId = getNodeId(node, "removeNode.node");
    return solidNative.removeChild(parentId, nodeId);
  },
  getParentNode(node) {
    const id = getNodeId(node, "getParentNode");
    const parentId = solidNative.getParent(id);
    // Empty string means not found (rquickjs doesn't handle Option<String>)
    if (parentId && parentId.length > 0) {
      return wrapNodeIdInNode(parentId);
    }
    return undefined;
  },
  getFirstChild(node) {
    const id = getNodeId(node, "getFirstChild");
    const firstChildId = solidNative.getFirstChild(id);
    // Empty string means not found
    if (firstChildId && firstChildId.length > 0) {
      return wrapNodeIdInNode(firstChildId);
    }
    return undefined;
  },
  getNextSibling(node) {
    const id = getNodeId(node, "getNextSibling");
    const nextSiblingId = solidNative.getNextSibling(id);
    // Empty string means not found
    if (nextSiblingId && nextSiblingId.length > 0) {
      return wrapNodeIdInNode(nextSiblingId);
    }
    return undefined;
  },
});

export const render = (code: () => JSX.Element, rootId?: string) => {
  // Get root from API (returns empty string if not set)
  const rootFromApi = solidNative.getRootView();

  // Check for empty string (means root not set)
  if (!rootFromApi || rootFromApi.length === 0) {
    throw new Error(`render: solidNative.getRootView() returned empty string. Root node not set. Call createRoot() from the host platform first.`);
  }

  const root = rootId ?? rootFromApi;
  const wrappedRoot = wrapNodeIdInNode(root);

  // @ts-ignore - solidRender expects the root element as second argument
  solidRender(code, wrappedRoot);
};
