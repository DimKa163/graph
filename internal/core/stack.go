package core

type Stack struct {
	nodes []*PathNode
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	return s.count
}

func (s *Stack) Push(n *PathNode) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *PathNode {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
