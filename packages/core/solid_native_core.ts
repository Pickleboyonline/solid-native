import { SolidNativeElement, SolidNativeNode } from "./types.ts";

export interface TSolidNativeCore<Node> {
  /**
   * @returns elementID
   */
  createElementByName: (elementName: string) => Node;
  // isTextElement: (elementId: Node) => boolean;
  /**
   * @returns elementID
   */
  createTextElement: () => Node;
  // removeElement: (
  //   parentId: Node,
  //   childId: Node
  // ) => void;
  // insertElement: (
  //   parentId: Node,
  //   childId: Node,
  //   anchorId?: Node
  // ) => void;
  // getParentElementId: (
  //   elementId: Node
  // ) => Node | undefined;
  // getFirstChildElementId: (
  //   elementId: Node
  // ) => Node | undefined;
  // getNextSiblingElementId: (
  //   elementId: Node
  // ) => Node | undefined;
  // setPropertyOnElement: (
  //   elementId: Node,
  //   propertyName: string,
  //   value: unknown
  // ) => void;
  getRootElement: () => Node;
}

/**
 * Returns module from Solid Native module manager
 */
export function getNativeModule<ModuleType>(moduleName: string): ModuleType {
  // deno-lint-ignore no-explicit-any
  const mod = (globalThis as any)._getNativeModule(moduleName);

  return mod;
}

/**
 * @deprecated
 */
export const SolidNativeCore = {};
//   getNativeModule<TSolidNativeCore<SolidNativeNode>>("SolidNativeCore");

// deno-lint-ignore no-explicit-any
export const print = (globalThis as any)._print as (str: string) => void;
