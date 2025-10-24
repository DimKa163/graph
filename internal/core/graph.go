package core

import "github.com/beevik/guid"

type Graph struct {
	nodes     map[guid.Guid]*Node
	nodeSlice []*Node
	edgeSlice EdgeList
}

func NewGraph() *Graph {
	return &Graph{
		nodes:     make(map[guid.Guid]*Node),
		edgeSlice: make(EdgeList),
	}
}

func (g *Graph) Find(id *guid.Guid) (n *Node, ok bool) {
	n, ok = g.nodes[*id]
	return
}

func (g *Graph) Path(destination *Node) *Path {
	path := NewPath()
	queue := NewQueue(3)
	visited := make(map[guid.Guid]bool)
	subD := make(map[guid.Guid]int)
	subD[destination.ID] = 1
	queue.Push(&PathNode{
		Node: destination,
	})
	for queue.Len() > 0 {
		item := queue.Pop()
		_, ok := visited[item.ID]
		if ok {
			continue
		}
		visited[item.ID] = true
		item.Level = subD[item.ID]
		path.AddNode(item)
		// идём назад
		edges := g.edgeSlice.IncomeTo(item.Node)
		for _, edge := range edges {
			node := edge.From
			subD[node.ID] = item.Level + 1
			queue.Push(&PathNode{
				Level: item.Level,
				Node:  node,
				Next:  item,
			})
		}
	}
	return path
}

func (g *Graph) AddNode(n *Node) {
	_, ok := g.nodes[n.ID]
	if ok {
		return
	}
	g.nodes[n.ID] = n
	g.nodeSlice = append(g.nodeSlice, n)
}

func (g *Graph) AddEdge(from, to *Node, weight int64) {
	g.edgeSlice.Add(from, to, weight)
}
