package snmobile

// Used to access arrays since gomobile can't
// expose arrays other than bytes
// TODO: Understand how memory management works between Go/Mobile
type StringArray struct {
	values []string
}

func (e *StringArray) Length() int {
	if e.values == nil {
		return 0
	}
	return len(e.values)
}

func (e *StringArray) Get(index int) string {
	return e.values[index]
}

// Used to access arrays since gomobile can't
// expose arrays other than bytes
// TODO: Understand how memory management works between Go/Mobile
type IntArray struct {
	values []int
}

func (e *IntArray) Length() int {
	if e.values == nil {
		return 0
	}
	return len(e.values)
}

func (e *IntArray) Get(index int) int {
	return e.values[index]
}

type TextDescriptorArray struct {
	values []TextDescriptor
}

func (e *TextDescriptorArray) Length() int {
	if e.values == nil {
		return 0
	}
	return len(e.values)
}

func (e *TextDescriptorArray) Get(index int) TextDescriptor {
	return e.values[index]
}
