package core

import "github.com/DimKa163/graph/internal/shared"

type WarehouseType int

const (
	WarehouseUNRECOGNIZED WarehouseType = iota
	WarehouseFREE
	WarehouseMAIN
	WarehouseCENTER
	WarehouseMALL
)

func (w WarehouseType) String() string {
	return []string{"UNRECOGNIZED", "FREE", "MAIN", "CENTER", "MALL"}[w]
}

type Warehouse struct {
	Name                   string
	Code                   string
	Type                   WarehouseType
	AvailableRest          bool
	Address                string
	DescriptorGroup        string
	OnlyStockPickupAllowed bool
}

func MapWarehouseType(code string) WarehouseType {
	switch code {
	case shared.WarehouseCategoryFREE:

		return WarehouseFREE
	case shared.WarehouseCategoryMAIN:

		return WarehouseMAIN
	case shared.WarehouseCategoryCENTRAL:
		return WarehouseCENTER
	case shared.WarehouseCategoryMALL:
		return WarehouseMALL
	default:
		return WarehouseUNRECOGNIZED
	}
}
