/**
 * In actually, this is a string. BUT to make this more
 * Typescript/TSX compiler friendly it's aliased as a
 * `JSX.Element`
 */
export type SolidNativeElement = JSX.Element;

// Define types internally

declare global {
  namespace JSX {
    interface IntrinsicElements {
      [name: string]: Record<string, unknown>;
    }
    interface Element {
      firstChild?: Element;
      parentElement?: Element;
      setProp(name: string, value: unknown): void;
      isTextElement: boolean;
      removeChild(element: Element): void;
      insertBefore(element: Element, anchor?: Element): void;
      next?: Element;
      prev?: Element;
    }

    interface ElementChildrenAttribute {
      children?: {}; // specify children name to use
    }
  }
}
