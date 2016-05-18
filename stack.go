package stack

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Push(v interface{}) {
	e := Element{value: v, next: s.top}
	s.top = &e
	s.size += 1
}

func (s *Stack) Size() int {
	return s.size
}

func (e *Element) Value() interface{} {
	if e != nil {
		return e.value
	}
	return nil
}

func (s *Stack) Pop() *Element {
	if s.size > 0 {
		var e Element = *s.top
		s.top = e.next
		s.size -= 1
		return &e
	}
	return nil
}
