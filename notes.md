## Other Ideas
Because Solid JS just needs to implement some tree manipulation methods, we utilize jetpack compose
and Swift UI for the complex ui management.

I'm not sure how Jetpack Compose would work, but I'd assume it be similar to SwiftUI. For SwiftUI,
we create Element class wrappers that have props as an ObservableObject, a render function that
exports a View binded to the props, and methods/properties exposed to the JS runtime. 

This way, I essentially only need to implement the setProperty and child manipulation methods natively
and left SwiftUI take care of the DOM and rendering dirty work. There is somewhat of a VDOM, but only
in the sense that I need a class tree to hold the props that need to be mutated. All the diffing and
optimizations are handed by SwiftUI.

I haven't checked out Jetpack Compose, but I'm working on that.

Next, I need a way to call the method on JS land. Because I want the JS interface and module downloading to
be the same, I want to use Kotlin Multiplatform to share business logic. There needs to be some way for JS to
reference a module. We can either create an object and assign properties to it. Or something else.

Either way, we need to be able to look up a node, and make mutations on it. More on this later.
(The flatting of views can help here, because the node has an index which we can lookup easily)
We can give each view a view tag. 

When JS asks for a view with view tag, we give back a list of methods and properties it can call.
These come from Swift. We can use a function definition builder to assign Swift inline functions to 
Obj-C blocks and assign them to object properties.

This can be hosted from kotlin as well. We should be able to define an object with an NSDictionary of
Objective C blocks for the functions.

OK, so we have a JSObject builder. This builder constructs a JSObject Value and returns it.
A View module essentially runs this JSObjectBuilder on init and exposes it to JS land.

We need a global object, so this object gets attached to that. That represents our node. That node
is in our view registry. 


AsyncFunctions:
A bit tricky, save for later

Goals:

- [X] Create Hello world app:
    - Write in JS code in Kotlin String
    - Compile and print "w
- [ ] Create iOS app

Reference for what JSX components would look like.
https://swiftui-react-native.vercel.app/