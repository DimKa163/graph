package appcontext

import (
	"context"
	"sync"

	"github.com/DimKa163/graph/internal/core"
)

type GraphContext struct {
	graph *core.Graph
	mutex *sync.RWMutex
}

func NewGraphContext() *GraphContext {
	return &GraphContext{
		mutex: &sync.RWMutex{},
	}
}

func (gc *GraphContext) Get(ctx context.Context) (*core.Graph, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		gc.mutex.RLock()
		defer gc.mutex.RUnlock()
		return gc.graph, nil
	}

}

func (gc *GraphContext) Update(graph *core.Graph) {
	gc.mutex.Lock()
	defer gc.mutex.Unlock()
	gc.graph = graph
}
