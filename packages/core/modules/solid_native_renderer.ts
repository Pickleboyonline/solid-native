import { getNativeModule } from "solid-native/core";

type SolidNativeRenderer = {
  getRootView: () => string;
  getFirstChild: (node: string) => string | undefined;
  getParent: (node: string) => string | undefined;
  setProp: (node: string, key: string, value: unknown) => void;
  isTextElement: (node: string) => boolean;
  removeChild: (node: string, child: string) => void;
  insertBefore: (node: string, child: string, anchor?: string) => void;
  next: (node: string) => string | undefined;
  prev: (node: string) => string | undefined;
  /**
   *
   * @param name Node Type Name
   * @returns Node ID
   */
  createNodeByName: (name: string) => string;
};

export const SolidNativeRenderer =
  getNativeModule<SolidNativeRenderer>("SNRender");
