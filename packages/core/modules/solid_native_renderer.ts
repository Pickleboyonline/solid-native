import { getNativeModule } from "solid-native/core";

type SolidNativeRenderer = {
  getRootView: () => number;
  getFirstChild: (node: number) => number | undefined;
  getParent: (node: number) => number | undefined;
  setProp: (node: number, key: string, value: unknown) => void;
  isTextElement: (node: number) => boolean;
  removeChild: (node: number, child: number) => void;
  insertBefore: (node: number, child: number, anchor?: number) => void;
  getNextSibling: (node: number) => number | undefined;
  /**
   * @param name Node Type Name
   * @returns Node ID
   */
  createNodeByName: (name: string) => number;
};

export const SolidNativeRenderer = getNativeModule<SolidNativeRenderer>(
  "_SolidNativeRenderer",
);
