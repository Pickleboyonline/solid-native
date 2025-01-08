package core

// Used to access arrays since gomobile can't
// expose arrays other than bytes
// TODO: Understand how memory management works between Go/Mobile
type StringArray struct {
	Values []string
}

func (e *StringArray) Length() int {
	if e.Values == nil {
		return 0
	}
	return len(e.Values)
}

func (e *StringArray) Get(index int) string {
	return e.Values[index]
}

// Used to access arrays since gomobile can't
// expose arrays other than bytes
// TODO: Understand how memory management works between Go/Mobile
type IntArray struct {
	Values []int
}

func (e *IntArray) Length() int {
	if e.Values == nil {
		return 0
	}
	return len(e.Values)
}

func (e *IntArray) Get(index int) int {
	return e.Values[index]
}
