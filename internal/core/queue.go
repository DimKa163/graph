package core

type Queue struct {
	nodes []*PathNode
	size  int
	head  int
	tail  int
	count int
}

func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*PathNode, size),
		size:  size,
	}
}

func (q *Queue) Push(n *PathNode) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*PathNode, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *Queue) Len() int {
	return q.count
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *PathNode {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}
