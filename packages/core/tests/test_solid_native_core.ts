// Was a test for renderer, no longer in use though

// import { getCircularReplacer } from "./cycle.js";
// import { TSolidNativeCore, print } from "./solid_native_core.ts";

// const startingId = 0;
// let id = startingId;

// function generateId() {
//   id += 1;
//   return "" + id;
// }

// export type TestSolidNativeElement = {
//   id: string;
//   name: string;
//   prev?: TestSolidNativeElement;
//   parent?: TestSolidNativeElement;
//   next?: TestSolidNativeElement;
//   children: TestSolidNativeElement[];
//   props: Record<string, unknown>;
//   firstChild?: TestSolidNativeElement;
// };

// const rootElement = {
//   id: "" + startingId,
//   name: "root_view",
//   children: [],
//   props: {},
// };

// const elementMapById: Record<string, TestSolidNativeElement | undefined> = {
//   [rootElement.id]: rootElement,
// };

// const TEXT_ELEMENT_NAME = "sn_text";

// function updateElementChildren(parent: TestSolidNativeElement) {
//   const newChildren = [];

//   let nextChild = parent.firstChild;

//   while (nextChild) {
//     newChildren.push(nextChild);
//     nextChild = nextChild.next;
//   }

//   parent.children = newChildren;
// }

// const eventListeners = new Set<StateChangeHandler>();

// let change = 0;
// const emitChange = () => {
//   print(`State ${change}: `);
//   print(JSON.stringify(elementMapById, getCircularReplacer(), 4));
//   change++;
//   eventListeners.forEach((handler) => {
//     setTimeout(() => handler(elementMapById), 0);
//   });
// };

// type StateChangeHandler = (tree: typeof elementMapById) => void;

// const addEventListener = (
//   _type: "onStateChange",
//   handler: StateChangeHandler
// ) => {
//   eventListeners.add(handler);
//   return () => {
//     eventListeners.delete(handler);
//   };
// };

// export const TestSolidNativeCore: TSolidNativeCore<TestSolidNativeElement> & {
//   addEventListener: (
//     type: "onStateChange",
//     handler: StateChangeHandler
//   ) => () => void;
// } = {
//   createElementByName(name) {
//     print("createElementByName " + arguments);
//     const newElement: TestSolidNativeElement = {
//       name,
//       id: generateId(),
//       props: {},
//       children: [],
//     };
//     return newElement;
//   },
//   createTextElement() {
//     print("createTextElement " + arguments);
//     const newElement: TestSolidNativeElement = {
//       name: TEXT_ELEMENT_NAME,
//       id: generateId(),
//       props: {},
//       children: [],
//     };
//     return newElement;
//   },
//   isTextElement(element) {
//     print("isTextElement " + arguments);
//     return element.name === TEXT_ELEMENT_NAME;
//   },
//   removeElement(parent, child) {
//     print("removeElement " + arguments);
//     if (child.next && child.prev) {
//       const { prev, next } = child;
//       prev.next = next;
//       next.prev = prev;
//     } else if (child.next && !child.prev) {
//       child.next.prev = undefined;
//       parent.firstChild = child.next;
//     } else if (!child.next && child.prev) {
//       child.prev.next = undefined;
//     }
//     child.parent = undefined;
//     child.next = undefined;
//     child.prev = undefined;

//     updateElementChildren(parent);
//     elementMapById[child.id] = undefined;
//     emitChange();
//   },
//   insertElement(parent, child, anchor) {
//     print("insertElement " + arguments);
//     const parentFirstChild = parent.firstChild;

//     if (anchor) {
//       if (anchor === parentFirstChild) {
//         anchor.prev = child;
//         child.next = anchor;
//         parent.firstChild = child;
//       } else if (anchor.prev) {
//         anchor.prev.next = child;
//         child.prev = anchor.prev;
//         anchor.prev = child;
//         child.next = anchor;
//       }
//     } else if (parentFirstChild) {
//       parentFirstChild.prev = child;
//       child.next = parentFirstChild;
//       parent.firstChild = child;
//     } else {
//       parent.firstChild = child;
//     }

//     elementMapById[child.id] = child;
//     child.parent = parent;
//     updateElementChildren(parent);
//     emitChange();
//   },
//   getRootElement() {
//     print("getRootElement " + arguments);
//     return rootElement;
//   },
//   getFirstChildElementId(element) {
//     print("getFirstChildElementId " + arguments);
//     return element.firstChild;
//   },
//   getParentElementId(element) {
//     print("getParentElementId " + arguments);
//     return element.parent;
//   },
//   getNextSiblingElementId(element) {
//     print("getNextSiblingElementId " + arguments);
//     return element.next;
//   },
//   setPropertyOnElement(elementId, propertyName, value) {
//     print("setPropertyOnElement " + arguments);
//     elementId.props[propertyName] = value;
//     emitChange();
//   },
//   addEventListener,
// };

// // function getCircularReplacer() {
// //   const ancestors = [];
// //   return function (key, value) {
// //     if (typeof value !== "object" || value === null) {
// //       return value;
// //     }
// //     // `this` is the object that value is contained in,
// //     // i.e., its direct parent.
// //     while (ancestors.length > 0 && ancestors.at(-1) !== this) {
// //       ancestors.pop();
// //     }
// //     if (ancestors.includes(value)) {
// //       return "[Circular]";
// //     }
// //     ancestors.push(value);
// //     return value;
// //   };
// // }
