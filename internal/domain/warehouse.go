package domain

import (
	"context"

	"github.com/beevik/guid"
	"github.com/jackc/pgx/v5"
)

type Warehouse struct {
	ID                     *guid.Guid
	Fnrec                  string
	Name                   string
	IsActive               bool
	OnlyStockPickupAllowed bool
	SenderID               *guid.Guid
	RecipientID            *guid.Guid
	Category               *WarehouseCategory
	Info                   *WarehouseInfo
}

type WarehouseCategory struct {
	Fnrec               string
	AvailableForBalance bool
}

type WarehouseInfo struct {
	ID              *guid.Guid
	Fnrec           string
	Address         string
	DescriptorGroup string
	TimeZone        *TimeZone
}

type TimeZone struct {
	ID   *guid.Guid
	Code string
}

type WarehouseRepository interface {
	GetAll(ctx context.Context) ([]*Warehouse, error)
}

func (w *Warehouse) Scan(dest pgx.Rows) error {
	return nil
}
