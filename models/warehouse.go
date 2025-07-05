package models

type Warehouse struct {
	ID      int    `json:"id"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}
