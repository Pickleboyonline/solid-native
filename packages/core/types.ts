import type { JSX as solidJSX } from "solid-js";
/**
 * In actually, this is a string. BUT to make this more
 * Typescript/TSX compiler friendly it's aliased as a
 * `JSX.Element`
 */
export type SolidNativeElement = JSX.Element;

export type SolidNativeNode = {
  firstChild?: SolidNativeNode;
  parentElement?: SolidNativeNode;
  setProp(name: string, value: unknown): void;
  isTextElement: boolean;
  removeChild(element: SolidNativeNode): void;
  insertBefore(element: SolidNativeNode, anchor?: SolidNativeNode): void;
  next?: SolidNativeNode;
  prev?: SolidNativeNode;
};
// Define types internally

declare global {
  namespace JSX {
    interface IntrinsicElements {
      [name: string]: Record<string, unknown>;
    }
    type Element = solidJSX.Element;

    interface ElementChildrenAttribute {
      children?: unknown; // specify children name to use
    }
  }
}
