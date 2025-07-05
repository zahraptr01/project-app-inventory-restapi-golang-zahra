package models

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
	RackID      int    `json:"rack_id" validate:"required"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	MinStock    int    `json:"min_stock"`
	WarehouseID int    `json:"warehouse_id" validate:"required"`
}
