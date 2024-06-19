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
	return false
}

func (v *JSValue) GetString() string {
	return ""
}

func (v *JSValue) IsNumber() bool {
	return false
}

func (v *JSValue) GetNumber() float64 {
	return 0
}

func (v *JSValue) IsObject() bool {
	return false
}

func (v *JSValue) IsArray() bool {
	return false
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
