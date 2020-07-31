package structs

type Queue struct {
	q []interface{}
}

func (q *Queue) Put(value ...interface{})       { q.q = append(q.q, value...) }
func (q *Queue) Len() int                       { return len(q.q) }
func (q *Queue) GetQueueContent() []interface{} { return q.q }

func (q *Queue) TakeOneTask() interface{} {
	task := q.q[0]
	q.q = q.q[1:]

	return task
}
