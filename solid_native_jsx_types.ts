// Define types internally

declare global {
    namespace JSX {
      interface IntrinsicElements {
        [name: string]: Record<string, unknown>
      }
      interface Element {
        [name: string]: Record<string, unknown>
      }
    }
}
