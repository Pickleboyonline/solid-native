import { SolidNativeElement } from "./types.ts";

interface TSolidNativeCore {
  /**
   * @returns elementID
   */
  createElement: (elementName: SolidNativeElement) => SolidNativeElement;
  isTextElement: (elementId: SolidNativeElement) => boolean;
  /**
   * @returns elementID
   */
  createTextElement: () => SolidNativeElement;
  removeElement: (
    parentId: SolidNativeElement,
    childId: SolidNativeElement
  ) => void;
  insertElement: (
    parentId: SolidNativeElement,
    childId: SolidNativeElement,
    anchorId?: SolidNativeElement
  ) => void;
  getParentElementId: (
    elementId: SolidNativeElement
  ) => SolidNativeElement | undefined;
  getFirstChildElementId: (
    elementId: SolidNativeElement
  ) => SolidNativeElement | undefined;
  getNextSiblingElementId: (
    elementId: SolidNativeElement
  ) => SolidNativeElement | undefined;
  setPropertyOnElement: (
    elementId: SolidNativeElement,
    propertyName: SolidNativeElement,
    value: unknown
  ) => void;
  getRootElement: () => SolidNativeElement;
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
  getNativeModule<TSolidNativeCore>("SolidNativeCore");

export const print = getNativeModule<(str: string) => void>("print");
