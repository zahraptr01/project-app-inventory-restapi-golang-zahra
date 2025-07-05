package models

type Rack struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	WarehouseID int    `json:"warehouse_id" validate:"required"`
}
