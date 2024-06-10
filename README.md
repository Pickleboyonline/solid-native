# Solid Native
React Native but for Solid.JS, powered by Deno!

React Native is nice, but it has the following issues:
1. Difficult upgrades and installations
2. Difficulty creating native modules and interfacing with native modules easily
3. Almost impossible debugging with new hermes engine (and hermes engine being relatively unstable)
4. Poor monorepo support

To address these issues, this framework aims for the following:
- Improve upgrading:
	- Most likely I will borrow from NativeScript and Flutter frameworks. These projects use the main JS/Dart package as the source of truth and generates the native files as needed. This method is in contrast to React Native where you have to manually update the native files. Manually updating the files can be cumbersome and error prone, especially if you have additional packages that need the code to be modified. Updating a react native project can take over a week at times.
- Use Deno and ESbuild over react natives node and metro bundler:
	- Deno provides a simpler runtime and improved developer workflow. ESBuilds Go backend is significantly faster than the javascript based metro bundler. I do in the future want to support NodeJS as well.
- Use SolidJS over ReactJS
	- SolidJS has a smaller more performant rendering engine
- Utilize SwiftUI and Jetpack Compose for the rendering
	- SwiftUI and Jetpack Compose are the equivalent to React Native's developer facing components. Under the current framework, react native implements a custom native rendering package to make its components, significantly increasing package complexity. By utilizing SwiftUI and Jetpack Compose's APIs that already closely mimic that of react native's I can reduce the size and complexity of the Solid Native framework.
- Flutter-like developer experience
	- Flutter tends to have every package you need in their core library developed by Google. Also, Flutter has an extremely good developer experience with a working debugger. 
	- For example, I want a library like `react-native-reanimated` with `worklets` built into the SolidNative framework.
	- Expo is a good reference for what that looks like. They are striving for an improved  developer experience, but because they rely on react native developer experience can be difficult at times.

For the JS Engine, I will probably be using a V8 engine lite version with bytecode enabled for fast startup times. This approach is similar to NativeScript. I could use React Native's hermes but it is buggy and in an early stage.

## Current Example Showcase

https://github.com/Pickleboyonline/solid-native/assets/16438041/2bb1975b-24bf-4863-89cb-563aa2961116

#### How it works:
When it the iOS app launches, it downloads a JS Bundle from the developer server. The developer server uses ESbuild and various plugins to ensure that the code transpiles correctly. I implemented a SolidJS universal renderer in the JS side by calling native module methods exposed by the Swift native code. On the Swift side, I have a core, module, and view system. Core takes care of downloading the bundle, instantiating the javascript virtual machine,  and registering the available modules. Modules are any Swift code that can be called from the JS side. For example, the native renderer is a module that exposes JS methods to manipulate views. Views are Swift classes that house and control the SwiftUI views.

When the JS bundle is executed, SolidJS takes care of calling the necessary methods to make the counter appear.

## How to run:

Make sure you have Deno installed. Then, run the following command to start the bundler:
```ts
deno task start
```
The bundler should then be running on port 8080.

From there, open the iOS app in `packages/ios`. Run the app in the simulator and it should load the test app in `packages/test_app`. Note, live reload/hot reload is not supported yet, so re-run the iOS app when you make a change to the test app!

## Docs
Powered by [lume](https://lume.land/). Use the command `deno task lume:serve` to view a local instance of the docs or go to https://pickleboyonline.github.io/solid-native/. I plan on putting a roadmap there.

## TODO

- [ ] Implement React Native's Components
	- Will use the same API structure for familiarity. However, text will not have to be wrapped in a `<Text />` component since SolidJS's rendering engine will treat bare text JSX as a text component.
- [ ] Implement React Native Modules
- [ ] Implement Expo, Flutter, and any other community packages that are needed to make a base app.
- [ ] Hot Reloading and VSCode Debugging
- [ ] Create module system and developer workflow cli
- [ ] Android Version
- [ ] Improve threading:
	- The threading model should be similar to react native. The JS code should be running on a background thread. JS code needed for animations should be running on the main thread with the `worklet` pattern.
