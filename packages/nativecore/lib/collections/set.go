package collections

// Got it from here:
// https://gist.github.com/bgadrian/cb8b9344d9c66571ef331a14eb7a2e80
// Used to mark nodes as dirty
type Set struct {
	list map[int]struct{} //empty structs occupy 0 memory
}

func (s *Set) Has(v int) bool {
	_, ok := s.list[v]
	return ok
}

func (s *Set) Add(v int) {
	s.list[v] = struct{}{}
}

func (s *Set) Remove(v int) {
	delete(s.list, v)
}

func (s *Set) Clear() {
	s.list = make(map[int]struct{})
}

func (s *Set) Size() int {
	return len(s.list)
}

func NewSet() *Set {
	s := &Set{}
	s.list = make(map[int]struct{})
	return s
}
