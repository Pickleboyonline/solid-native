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
