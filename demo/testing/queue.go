package queue

type Queue struct {
	items    []int
	capacity int
}


func New(capacity int) Queue {
	return Queue{make([]int, 0, capacity), capacity}
}

func (q *Queue) Append(item int) bool {
	if len(q.items) == q.capacity {
		return false
	}
	q.items = append(q.items, item)
	return true
}