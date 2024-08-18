package snmobile

// Used to access arrays since gomobile can't
// expose arrays other than bytes
// TODO: Understand how memory management works between Go/Mobile
type StringArray struct {
	strings []string
}

func (e *StringArray) Length() int {
	if e.strings == nil {
		return 0
	}
	return len(e.strings)
}

func (e *StringArray) Get(index int) string {
	return e.strings[index]
}

// Used to access arrays since gomobile can't
// expose arrays other than bytes
// TODO: Understand how memory management works between Go/Mobile
type IntArray struct {
	ints []int
}

func (e *IntArray) Length() int {
	if e.ints == nil {
		return 0
	}
	return len(e.ints)
}

func (e *IntArray) Get(index int) int {
	return e.ints[index]
}
