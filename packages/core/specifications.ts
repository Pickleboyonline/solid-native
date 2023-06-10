/*
 * Outlines what modules we need to make.
 */

/**
 * Creates SNView based on the name.
 * It looks it up in the view registry,
 * instantiates a new one, and returns a
 * JS Object that has the needed methods.
 */
const createElement = (name: string): Element => {};

/**
 * On the native side, construct an element that exposes
 * a JS object that conforms to this interface.
 */
interface Element {
  /**
   * Updates the prop with a value
   * @param name
   * @param value
   * @returns
   */
  setProp: (name: string, value: unknown) => void;

  addChildNodeBeforeAnchorNode: (
    childNodeViewTag: string,
    anchorNodeViewTag?: Element
  ) => void;

  removeChildNode: (childNode: Element) => void;

  isTextNode: boolean;

  /**
   * Will be for parent node
   * @returns
   */
  getParentNode: () => Element | null;

  getFirstChild: () => Element | null;

  getNextSibling: () => Element | null;

  viewTag: string
}

/**
 * Will create a module that is the equivalent of a native object
 * Has props and methods
 *
 * @param name
 */
const createModule = (name: string): Module => {};

interface Module {}

/*
 * Once both these modules are built,
 * I can work on combining the APIs into the same thing
 * with a similar tooling to Expo Modules API.
 */

// TODO: Might be able to transpile minified JS code into 
// Kotlin multiplatform code
