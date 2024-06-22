package snmobile

// Two things:
// must be able to check the type with a function call
// must be able to grab type.
type JSValue struct {
	// Anything
	data interface{}
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

// TODO: Impliment
func (v *JSValue) GetJSValueForKey(key string) *JSValue {
	obj := v.data.(map[string]interface{})
	return &JSValue{data: obj[key]}
}
