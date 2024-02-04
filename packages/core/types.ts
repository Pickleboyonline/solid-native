import type { JSX as solidJSX } from "solid-js";
/**
 * In actually, this is a string. BUT to make this more
 * Typescript/TSX compiler friendly it's aliased as a
 * `JSX.Element`
 */
export type SolidNativeElement = JSX.Element;

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
