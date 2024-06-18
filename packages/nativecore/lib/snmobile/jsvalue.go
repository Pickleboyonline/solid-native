package snmobile

import "gopkg.in/olebedev/go-duktape.v3"

// Two things:
// must be able to check the type with a function call
// must be able to grab type.
type JsValue struct {
	valueType    duktape.Type
	stashKeyName string
	// Mainly used for access to ctx to retrieve stash
	solidNativeMobile *SolidNativeMobile
}

// DO NOT create from Mobile side
// Only create if it's already in stash.
// However if not in stash wont return anything
func NewJsValue(
	valueType duktape.Type,
	stashKeyName string,
	solidNativeMobile *SolidNativeMobile) *JsValue {
	return &JsValue{
		valueType,
		stashKeyName,
		solidNativeMobile,
	}
}
