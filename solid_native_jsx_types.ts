// Define types internally

declare global {
    namespace JSX {
      interface IntrinsicElements {
        [name: string]: Record<string, unknown>
      }
      interface Element {
        // children?: Element | Element[];
        // [name: string]: Record<string, unknown>
      }

      interface ElementChildrenAttribute {
        children?: {}; // specify children name to use
      }
    }
}
