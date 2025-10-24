package core

import "github.com/beevik/guid"

type Path struct {
	list  []*PathNode
	nodes map[guid.Guid]*PathNode
}

func NewPath() *Path {
	return &Path{
		list:  make([]*PathNode, 0),
		nodes: make(map[guid.Guid]*PathNode),
	}
}

func (path *Path) GetList() []*PathNode {
	return path.list
}

func (path *Path) AddNode(n *PathNode) {
	path.list = append(path.list, n)
	path.nodes[n.ID] = n
}

func (path *Path) Contains(nodeId *guid.Guid) bool {
	_, ok := path.nodes[*nodeId]
	return ok
}

func (path *Path) FirstNode() *PathNode {
	if len(path.list) == 0 {
		return nil
	}
	return path.list[0]
}

func (path *Path) LastNode() *PathNode {
	if len(path.list) == 0 {
		return nil
	}
	return path.list[len(path.list)-1]
}

type PathNode struct {
	Level int
	Next  *PathNode
	*Node
}
