package persistence

import (
	"context"
	"database/sql"
	"github.com/DimKa163/graph/internal/core/db"
	"github.com/DimKa163/graph/internal/domain"
	"github.com/beevik/guid"
)

const (
	GetAllWarehouse = `SELECT nrb_subwarehouse.id,
       	nrb_subwarehouse.nrb_fnrec,
       	nrb_name,
       	nrb_is_active,
       	bpm_only_stock_pickup_allowed,
       	nrb_sender_id,
       	nrb_recipient_id,
       	sc.nrb_fnrec,
       	sc.nrb_available_for_balances,
       	nw.id,
       	nw.nrb_fnrec,
       	nw.nrb_address,
       	nw.bpm_descriptor_group_name,
       	tz.id,
       	tz.code
		FROM public.nrb_subwarehouse
		JOIN nrb_subwarehouse_categories sc on sc.id=nrb_subwarehouse.nrb_category_id
		JOIN public.nrb_warehouse nw on nw.id = nrb_subwarehouse.nrb_warehouse_id
		LEFT JOIN public.time_zone tz on tz.id=nw.ask_time_zone_id
		WHERE nrb_is_active = true`
)

type WarehouseRepository struct {
	db db.QueryExecutor
}

func (w WarehouseRepository) GetAll(ctx context.Context) ([]*domain.Warehouse, error) {
	var warehouses []*domain.Warehouse
	rows, err := w.db.Query(ctx, GetAllWarehouse)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var warehouseID string
		var warehouseFnrec string
		var name string
		var isActive bool
		var onlyStockPickupAllowed bool
		var senderID sql.NullString
		var recipientID sql.NullString
		var categoryFnrec sql.NullString
		var availableForBalances bool
		var warehouseInfoID sql.NullString
		var warehouseInfoFnrec sql.NullString
		var warehouseInfoAddress sql.NullString
		var warehouseInfoDescriptorGroup sql.NullString
		var tzID sql.NullString
		var tzCode sql.NullString
		if err := rows.Scan(&warehouseID,
			&warehouseFnrec,
			&name,
			&isActive,
			&onlyStockPickupAllowed,
			&senderID,
			&recipientID,
			&categoryFnrec,
			&availableForBalances,
			&warehouseInfoID,
			&warehouseInfoFnrec,
			&warehouseInfoAddress,
			&warehouseInfoDescriptorGroup,
			&tzID,
			&tzCode); err != nil {
			return nil, err
		}
		warehouse := &domain.Warehouse{
			Name:                   name,
			Fnrec:                  warehouseFnrec,
			IsActive:               isActive,
			OnlyStockPickupAllowed: onlyStockPickupAllowed,
			Category: &domain.WarehouseCategory{
				Fnrec:               categoryFnrec.String,
				AvailableForBalance: availableForBalances,
			},
		}
		var warehouseInfo *domain.WarehouseInfo

		warehouse.ID, err = guid.ParseString(warehouseID)
		if err != nil {
			return nil, err
		}
		if senderID.Valid {
			warehouse.SenderID, err = guid.ParseString(senderID.String)
			if err != nil {
				return nil, err
			}
		}
		if recipientID.Valid {
			warehouse.RecipientID, err = guid.ParseString(recipientID.String)
			if err != nil {
				return nil, err
			}
		}
		if warehouseInfoID.Valid {
			warehouseInfo = &domain.WarehouseInfo{}
			warehouseInfo.ID, err = guid.ParseString(warehouseInfoID.String)
			if err != nil {
				return nil, err
			}
			var tz *domain.TimeZone
			if tzID.Valid {
				tz = &domain.TimeZone{}
				tz.ID, err = guid.ParseString(tzID.String)
				if err != nil {
					return nil, err
				}
				tz.Code = tzCode.String
			}
			warehouseInfo.Address = warehouseInfoAddress.String
			warehouseInfo.Fnrec = warehouseInfoFnrec.String
			warehouseInfo.DescriptorGroup = warehouseInfoDescriptorGroup.String
			warehouseInfo.TimeZone = tz
		}
		warehouse.Info = warehouseInfo

		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

func NewWarehouseRepository(db db.QueryExecutor) *WarehouseRepository {
	return &WarehouseRepository{
		db: db,
	}
}
