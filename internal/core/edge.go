package core

import "github.com/beevik/guid"

type Edge struct {
	From   *Node
	To     *Node
	Weight int64
}

type EdgeList map[guid.Guid][]*Edge

func (el EdgeList) Add(from, to *Node, weight int64) {
	edge := &Edge{from, to, weight}
	el.addIncomeAdd(edge)
	el.addOutcomeAdd(edge)
}

func (el EdgeList) IncomeTo(to *Node) []*Edge {
	result := make([]*Edge, 0)
	edges, ok := el[to.ID]
	if !ok {
		return result
	}
	for _, edge := range edges {
		if edge.To.ID != to.ID {
			continue
		}
		result = append(result, edge)
	}
	return result
}

func (el EdgeList) OutcomeFrom(from *Node) []*Edge {
	result := make([]*Edge, 0)
	edges, ok := el[from.ID]
	if !ok {
		return result
	}
	for _, edge := range edges {
		if edge.From.ID != from.ID {
			continue
		}
		result = append(result, edge)
	}
	return result
}

func (el EdgeList) Edges(n *Node) ([]*Edge, bool) {
	edges, ok := el[n.ID]
	if !ok {
		return nil, false
	}
	return edges, true
}
func (el EdgeList) addOutcomeAdd(edge *Edge) {
	_, ok := el[edge.From.ID]
	if !ok {
		el[edge.From.ID] = make([]*Edge, 0)
	}
	el[edge.From.ID] = append(el[edge.From.ID], edge)
}
func (el EdgeList) addIncomeAdd(edge *Edge) {
	_, ok := el[edge.To.ID]
	if !ok {
		el[edge.To.ID] = make([]*Edge, 0)
	}
	el[edge.To.ID] = append(el[edge.To.ID], edge)
}
