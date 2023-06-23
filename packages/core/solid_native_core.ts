import { SolidNativeElement } from "./types.ts";

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
 * Returns typed module from global object.
 *
 * When exporting to via JavascriptCore on Swift,
 * ensure the module starts with a "_" to avoid naming
 * conflicts.
 *
 * When getting the module from this function, you can omit the
 * starting "_".
 */
export function getNativeModule<ModuleType>(moduleName: string): ModuleType {
  const name = (() => {
    if (moduleName.charAt(0) === "_") {
      return moduleName;
    }
    return "_" + moduleName;
  })();

  // deno-lint-ignore no-explicit-any
  return (globalThis as any)[name] as ModuleType;
}

export const SolidNativeCore =
  getNativeModule<TSolidNativeCore<SolidNativeElement>>("SolidNativeCore");

export const print = getNativeModule<(str: string) => void>("print");
