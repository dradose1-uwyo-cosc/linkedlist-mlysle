package ds

type Queue struct {
	list LinkedList
}

// add tail node
func (q *Queue) Push(value string) {
	q.list.Insert(value)
}

// remove the head
func (q *Queue) Pop() (string, bool) {
	if q.list.IsEmpty() {
		return "", false
	}
	value := q.list.Head.data
	return value, q.list.RemoveAt(0)
}
