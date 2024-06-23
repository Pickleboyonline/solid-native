---
logo: ''
---
# Solid Native

A React Native alternative but for SolidJS, powered by Deno!

Technical Overview:
- JS Engine is powered by Duktape
- JS Build System is a combination of ESBuild and SWC to bundle and target ES5.
- Meta's Yoga (also used for react native) powers the Flexbox layout system
- Instead of C++like in React Native, Go is used to have a cross platform codebase.
	- The shared Go code implements the JS engine and Yoga layout
	- On the mobile side a ViewHostReceiver is implemented as a callback mechanism to broadcast UI changes to the host environment.

## Roadmap:
- Need to implement the Text and Button component:
	- Text is a challenge because of embedded components causing the parent to change styles
	- Both the button and text need external measure functions given their props. This means I need to use UIKit to obtain the width and height of a component based on their props.
- Callback Mechanism:
	- Just have JSON serializable data be able to be sent back via an event emitter style paradigm.
	- Right now props are sent over via JSON. React Native has a new JSI, but the exact performance benefit is unclear because when exchanging primitives between the JS and Native side you always copy data. JSON's main downside is that it needs to allocate extra memory to deserialize and needs to traverse the data. Flexbuffers could be used to get JSI level performance while still maintaining the flexibility of JSON. Only issue is that the API is only available in C++, Java, and Javascript but is lacking in Go and Swift. Testing needs to be done to determine if implementing Flexbuffers is worth it.
- Implement More Views:
	- Cover most of react native style views
- Have a module system
	- Same as above
- Hot Reloading
- Build system
	- I want to be able to generate native code as function of the source and solid native version
	- NativeScript already does this.
- Debugging
	- Duktape already has a JS debugger so I just need to read the docs
- Optimize Debugging:
	- Duktape's highest supported target is ES5, which requires a large amount of polyfills. I can precompile the polyfills and load them into the engine before loading in user data.
- Threading:
	- Need to move JS side to backend. Not exactly sure how to handle this since both the Mobile and Golang have the ability to go on a different thread.
	- Could have Golang move JS to go routine and use channels when communicating the the main thread.
- Navigation
- Implement More modules
	- Have some parity with Expo modules