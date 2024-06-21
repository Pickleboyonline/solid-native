package snmobile

import "gopkg.in/olebedev/go-duktape.v3"

// Two things:
// must be able to check the type with a function call
// must be able to grab type.
type JSValue struct {
	valueType    duktape.Type
	stashKeyName string
	// Mainly used for access to ctx to retrieve stash
	solidNativeMobile *SolidNativeMobile
}

// DO NOT create from Mobile side
// Only create if it's already in stash.
// However if not in stash wont return anything
// TODO: Optimize this, if its a primative, just make a copy into the interface
func NewJsValue(
	valueType duktape.Type,
	stashKeyName string,
	solidNativeMobile *SolidNativeMobile) *JSValue {
	return &JSValue{
		valueType,
		stashKeyName,
		solidNativeMobile,
	}
}

func (v *JSValue) IsString() bool {
	return v.valueType.IsString()
}

func (v *JSValue) GetString() string {
	v.solidNativeMobile.dukContext.PushGlobalStash()                 // => [stash]
	v.solidNativeMobile.dukContext.GetPropString(-1, v.stashKeyName) // => [stash value]
	str := v.solidNativeMobile.dukContext.GetString(-1)              // => [stash value]
	v.solidNativeMobile.dukContext.Pop2()                            // => []
	return str
}

func (v *JSValue) IsNumber() bool {
	return v.valueType.IsNumber()
}

func (v *JSValue) GetNumber() float64 {
	v.solidNativeMobile.dukContext.PushGlobalStash()                 // => [stash]
	v.solidNativeMobile.dukContext.GetPropString(-1, v.stashKeyName) // => [stash value]
	num := v.solidNativeMobile.dukContext.GetNumber(-1)              // => [stash value]
	v.solidNativeMobile.dukContext.Pop2()                            // => []
	return num
}

func (v *JSValue) IsObject() bool {
	return v.valueType.IsObject()
}

// Swift/Kotlin way of recieving values
type JSValueEnumerator interface{}

func (v *JSValue) GetArrayEnumerator() JSValueEnumerator {
	return make([]byte, 0)
}

// TODO: Impliment
func (v *JSValue) GetObjectForKey(key string) *JSValue {
	return &JSValue{}
}

// This will remove the value from the stash making the
// JS value no longer accessable.
func (v *JSValue) Free() {

}
