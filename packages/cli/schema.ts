const solidNativeData = {
  elementsById: {
    1: {
      id: 1,
      viewType: "button",
      next: 4,
      prev: undefined,
      isTextElement: false,
      props: {
        text: "dsadasdas",
        children: [2, 3],
      },
    },
    2: {
      id: 2,
      children: [],
    },
    3: {
      id: 3,
      children: [],
    },
    4: {
      id: 4,
      children: [],
    },
  },
  viewsByType: {
    // type: Type.self
  },
  propsByElementId: {
    // Mutable observable objects that shadow
    // immutable element props. This way, they are updated.
  }
};

// Elements
const SolidNativeCore = {
  createElement: (type: string) => {
    // Add element to elements byID
    //
    // return element ID
  },
  replaceText: (elementId: string, text: string) => {},
  createTextElement: (text: string) => {
    SolidNativeCore.createElement("text");
  },
  setProperty: (elementId: string, propertyName: string, value: any) => {
    // Grab element
    // update it
    // update data
  },
  insertNode: (elementId: string) => {

  }
};



/**
 * How it works:
 * For the most part, all data is represented in a normalized tree like so with a 
 * static class that can maniplate said chain.
 * 
 * The only thing SwiftUI needs is to be updated from the props values change.
 * SwiftUI does this with an observable object that is instantiated in some registry
 * separate from our element nodes. When the element nodes change, they update the
 * corresponding prop object in the registry that was used to build the.
 * 
 * To build it, all there needs to be a static method that takes in the node and subsquently
 * builds the SwiftUI struct needed. Only issue is that we may not be able to dynmically
 * build structs, only classes.
 */

/**
 * Swift UI needs a root SwiftUI view. We need to dynamically registure and load views.
 * SwiftUI Views are structs. They can be updated externally through observable objects.
 * 
 */