import xcode from "npm:xcode@3.0.1";
// Source: https://github.dev/apache/cordova-node-xcode
const path = import.meta
  .resolve("./../new_ios/SolidNative/SolidNative.xcodeproj/project.pbxproj")
  .replace("file://", "");
// Another xcode manager.
// import { Xcode } from "npm:pbxproj-dom@1.2.0/xcode.js";
// Useful for execution:
// https://www.npmjs.com/package/ioslib

const proj = xcode.project(path);

// const getMethods = (obj) =>
//   Object.getOwnPropertyNames(obj).filter(
//     (item) => typeof obj[item] === "function"
//   );

proj.parse(() => {
  console.log(xcode.prototype);
});

// console.log(xcode.document.targets);
