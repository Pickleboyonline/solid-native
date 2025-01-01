package snmobile

// Meant to be used on host platform to define objects that be
// called upon by JS. Can be used for modules/views, and essentially
// provides a way to access values to and from JS side.
//
// Host platform can use module registration to make object for JS to call upon.
// I would say make module first, then views, because views are just a special
// kind of module linked to the view hierarchy.
//
// On the node creation, return the module definition of the view that corresponds
// to the JS object, that will be what is created (host platform must return it)
// The host platform will imperatively construct it and return it for later usage.
// etc... etc...
type JSFunction interface {
	Call(*JSValueArray) *JSValue
}
