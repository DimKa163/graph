package core

import "github.com/beevik/guid"

type Node struct {
	ID guid.Guid
	*Warehouse
}
