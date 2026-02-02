package ds

type Stack struct {
	list LinkedList
}

// add new head node
func (s *Stack) Push(value string) {
	s.list.InsertAt(0, value)
}

// remove the head
func (s *Stack) Pop() (string, bool) {
	if s.list.IsEmpty() {
		return "", false
	}
	value := s.list.Head.data
	return value, s.list.RemoveAt(0)
}
