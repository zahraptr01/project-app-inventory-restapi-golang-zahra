package models

type Rack struct {
	ID          int    `json:"id"`
	Code        string `json:"code" validate:"required"`
	WarehouseID int    `json:"warehouse_id" validate:"required"`
}
