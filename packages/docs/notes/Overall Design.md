## Solid Native:
System complexity increasing lol. Need to get Yoga Layout, which covers a lot of use cases. 

In the expo module system, the have a module function that provides a definition that describes the view. 

Rendering Pipe:
1. Javascript Engine is setup for runtime (like setTimeout)
	1. Runtime exposes Swift code to run
2. Downloads JS and runs script
3. For solidJS swift must be exposed to create a view system
	1. Call this the shadow tree
	2. Each element in the tree is the:
		1. Node, which is the ID
		2. Name to create element type
		3. Properties (map of data)
		4. And parent node
	3. By the API calls, a tree must be implimented
4. SolidJS does not handle rendering, and can be though to incrimentally update the view tree
5. Every time it updates the view tree, we must recalculate the layout.
6. Layout Metrics are passed to render system (swiftUI)
7. SwiftUI handles rendering props declaratively

Now, based on that have the following spec:
- Runtime Functions in global namespace
	- Can be things like `fetch`, set Timeout, etc...
- Render system?
	- Is this a module or runtime global?
	- Render system exposes methods to manipulate Shadow VDOM
- Each node in the VDOM houses:
	- name
	- props
	- id
	- children
	- graph system
- Each VDOM type must be accessable to the render system by runtime lookup (registry)
	- Mount Phase:
		- Node is created
		- Props are defined
		- Mounted to root tree
- RENDERING:
	- Rendering is done via swiftUI in a declarative fashion. This means once the data is available and the child is mounted, rendering happens. So, data needs to be available.
	- Calculate the layout info after each node is defined.
	- Once layout is calculated, update the tree.
	- I am fairly certain that YOGA can mark when its available.
	- Ok so theres shadow layout data and layout data. 
	- So really there are two tree's => one for representing general props and one for yoga layout.
	- We need yoga layout info for each mount phase. So, this info needs to be available before mounting
	- Mutate props, and recalculate yoga layout.
	- Yoga takes care of rendering system
	- When yoga is done, we update the layout metrics for the props for that.
	- So, it looks like this:
		- SolidShadowTree change => Queue render to Yoga on bg thread => Yoga traverses tree => If change happens, update the Node layout props => SwiftUI renders it
		- Lets say that render system traverses the tree, grabbing its associated yoga node. Essentially BFS
		- Really, we just need to update the layout system.
		- For SwiftUI, we don't render anything if layout information is not available
- On the view system, how does that work?
	- Well, we know each view/node requires certain wrappers to maintain layout and border styles.
	- Each view must be registered via a definition file. Somehow this is statically linked to the system.
	- Modules and Views can be considered one, in that everything is a module, but some modules can also be used as views if they have the view definition
	- Thus, we get a module registry and a view registry.
	- The renderer must be able to see the view registry.
	- Thus, I will say that the renderer is NOT a module but a core runtime element.
		- Basically, there are modules and there are core framework.
		- The distinction is that there are some special functions properties that are associated with the core that manage how the base app works.
			- These are things like javascript runtime management
			- View registry
			- Module registry management
		- Anytime when these must be accessed, they take the form of something else.
		- HOWEVER, you could add some flexibility, in the native modules and views are allowed to have access to those elements.
		- ^^ Naw, its too complex. Also renderer needs to know the root element. I created a hacky implimentation of this. Basically, 
Ok lets say you have a native module, we look for anything that extend the class, import it and generate a file that has a class that will call the definitions of every module you make.

Now, that definition class would have to be accessable to the render system. Also, mind that the root view is part of core, but it also need to be in the render system.


So, yeah, renderer is NOT a expernal module in that sense.

To recap:

Components:
- Javascript
- Core
	- Renderer
		- Functions related to managing view system
	- JS Engine
	- Module Registry
	- View Registry
- SwiftUI
	- Declaratively responded to JS

Module registry can be access and instanciated from the definition, which can build a JSValue for that module. The class itself is made and is associated JSValue is made via the defintion. This method needs to run prior to initialization. So, essentially to create the registry we make objects of the modules. We then compute their name. If its accessed, the JSValue is made.

When the definition is made, we make its views as well. 

This happens during initialization.

The renderer is exposed via JS runtime. it knows the view definitions.

When a node is created, it needs to address the following:
- Yoga Node
- Prop management
- Yoga Layout calculations and prop drilling
- SwiftUI wrappers to deal with layout
- Wrap associated view definition for system. 
- Question => Should view be struct or class? Or just a base swiftui view?
	- In this framework, views can just be struct defs that get passed a SolidNative prop
	- Can extend the view struct
	- No need for observable object, since, thats in the view definition. The struct takes care of that.
	- Text is sepecial because child components are "shadowed" and do not exist. However, since all the info is in the props (including children refrences, we need not worry about this and can calcualte it in the same way.

Core System:
- Builds JS system and download core
- Renderer will give chain of objects to call for system (effectively the wrapper node objects that handle layout, view styles, etc...)
- Ok understanding => How to expose methods to JS Land?
	- Some methods are directly the global namespace (think console.log => print and setTimeout)
	- some methods are core and some are external modules:
		- Now, the core is interesting because some methods may want to access core library stuff.
		- Thats OK ig, just import and access the shared instance tbh
		- On the JS side, just call the function.
		- Methods are exposed with the `_ModuleName`
		- Some modules are exposed via the core library and some are via the external module system
		- What I am getting at is that everything can be in the global namespace
		- However, if you want dynamic object allocation when accessing the values, then it may be smart to allocate that in a different system (think a require)
		- Some systems are not dynamic

WAIT we can decouple the view system and the render system

## New Overview:

Data/Component Wise:
- JS Engine
	- Methods to control it
- Module registry
	- Adding a module to it requires some macros but for now its not needed
- View Registry
- Render System:
	- Can be nde of view registry.
		- Only needs to know the name and props, both serializable things
		- Calculate yoga layout after perfoming mutation
		- Can possibply be stored in a dictionary
			- Flattened view
	- Only calls upon views when mounted per se
	- ^^ Need to determine how to do this?
	- Some linking between NATIVE side of things + Rendering
- Module system:
	- Must be able to grab from namespace

System must be able to registure new modules on the fly using code gen. Theres a file local to the project that can grab all the defs and expose to runtime or something.

=> Can take native script approach

---
Above is a future update. For now, we just need yoga layouts. So i'll pass on changing it for semantics if the functionality remains the same. Dont need 1:1 copy of RN if we can make it simplier

Just be concerned with Yoga layouts for now.

