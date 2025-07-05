package models

type Sale struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	ItemID    int    `json:"item_id"`
	Qty       int    `json:"qty"`
	Price     int    `json:"price"`
	Total     int    `json:"total"`
}
