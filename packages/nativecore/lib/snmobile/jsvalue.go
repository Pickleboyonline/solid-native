package snmobile

// Two things:
// must be able to check the type with a function call
// must be able to grab type.
// TODO: Make iterable (return JSValueArray)
type JSValue struct {
	// Anything
	data interface{}
}

// Easier way to make a JSValue
// Not accessible from host platform (if Swift/ObjC or Kotlin/Java)
func NewJSValue(data any) JSValue {
	return JSValue{data}
}

func (v *JSValue) IsString() bool {
	_, ok := v.data.(string)
	return ok
}

func (v *JSValue) GetString() string {
	str := v.data.(string)
	return str
}

func (v *JSValue) IsNumber() bool {
	switch v.data.(type) {
	case int:
		return true
	case float32:
		return true
	case float64:
		return true
	default:
		return false
	}
}

func (v *JSValue) GetNumber() float64 {
	switch v.data.(type) {
	case int:
		return float64(v.data.(int))
	case float32:
		return float64(v.data.(float32))
	case float64:
		return v.data.(float64)
	default:
		return 0
	}
}

func (v *JSValue) IsObject() bool {
	_, ok := v.data.(map[string]interface{})
	return ok
}

func (v *JSValue) GetJSValueForKey(key string) *JSValue {
	obj, ok := v.data.(map[string]interface{})

	if !ok {
		return &JSValue{}
	}

	return &JSValue{data: obj[key]}
}

func (v *JSValue) GetObjectKeys() *StringArray {
	m, ok := v.data.(map[string]interface{})

	if !ok {
		return &StringArray{}
	}

	keys := []string{}

	for k := range m {
		keys = append(keys, k)
	}

	return &StringArray{values: keys}
}
