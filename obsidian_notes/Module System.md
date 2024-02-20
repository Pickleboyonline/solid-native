Module System Needs to be made.

Needs the Core:
- Manages modules

Requirements:
- From Swifts side, all i need is some function to get the root element at any time.
- Some root element should be there. At best, this is just some empty group that has no props. OR, just a special swift view.
- Modules:
	- Modules are objects that be looked up by the JS environment. 
	- Views are special modules. 
	- OR these can be separate. 
	- This is just a wrapper so JS can talk to Swift.
- Renderer (for solidJS):
	- Is a tree structure with nodes that are swift ui components

Question:
- Should modules and views be separate?
	- Views are a wrapper for swift ui views. These have prop and state. Callbacks can be used to reference state.
	- Modules are objects that are referenced by context. 
	- Core library provides way to access these modules.
	- At the end of the day, from the JS perspective they simply access native code.
		- They simply are just code. Thus, they shouldn't require any special rendering. 
		- In fact, the rendering agent can be a module in itself.

So, we need some way to give access to JS in a unified way with modules.

What are views? Well, views themselves can adhere to this module framework since they get called on anyways. They just need special info 

However, the way these are accessed in Solid Can be challenging. I can either use a JSON serializable identifier or the object itself.

Then there comes the issue that the root view must be linked to it. This is not a problem, since its already there. 

Summing this up, we get this:
- Module system is way for Swift and JS to talk to eachother. Think i want to make a module for a camera or something
- Views, should be their own thing. They must extend some class to be a view.
- Views and modules must be accessable via the JS layer. For views, there must be a way for JS to update props and call methods on it. 
	- Now, heres the interesting thing. I can use a rendering module to do all of this. This module is registered in the core. Furthermore, it has an index of all the views in needed.

Ok, again:
- Modules are a way for JS-Native interop.
- Rendering:
	- Rendering is a module.
	- Has registeristry of native modules that is statically linked.
	- Access views via ID to set props and such.
	- Views are a native ONLY thing. These really aren't exposed to the JS side like modules are.
- JS needs some way to get modules in its context. Modules can either be lazy loaded or available at run time.
	- For now, initial instantiate at run time. 
- Core module is really only needed to set things up. It:
	- Has the JS context
	- Downloads the JS
	- Has the blank root view
	- Attaches this view to the renderer module
	- Hooks all known modules into global JS context

Ok so now we have these classes:
- Modules
- Views
- Core
Thats it!

Ok so the core has a registry of modules that it can can instanciate.
The registry is needed to lookup modules by name in the JS code.
When JS gets a native module, it's make a module.


When module gets made, the JS side must be able to get some object that has functions in it.
