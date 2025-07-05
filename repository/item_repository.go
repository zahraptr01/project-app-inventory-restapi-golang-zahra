package repository

import (
	"assignment5/models"
	"database/sql"
)

type ItemRepository interface {
	Create(item models.Item) error
	GetAll() ([]models.Item, error)
}

type itemRepo struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepo{db}
}

func (r *itemRepo) Create(item models.Item) error {
	_, err := r.db.Exec(`INSERT INTO items (name, category_id, rack_id, stock, price, min_stock, warehouse_id, ) VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		item.Name, item.CategoryID, item.RackID, item.Stock, item.Price, item.MinStock, item.WarehouseID)
	return err
}

func (r *itemRepo) GetAll() ([]models.Item, error) {
	rows, err := r.db.Query(`SELECT id, name, category_id, rack_id, warehouse_id, stock, min_stock FROM items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var i models.Item
		err := rows.Scan(&i.ID, &i.Name, &i.CategoryID, &i.RackID, &i.WarehouseID, &i.Stock, &i.MinStock)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}
