package repository

import (
	"assignment5/models"
	"database/sql"
)

type ItemRepository interface {
	Create(item models.Item) error
	GetAll() ([]models.Item, error)
	GetByID(id int) (models.Item, error)
	Update(id int, item models.Item) error
	Delete(id int) error
}

type itemRepo struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepo{db}
}

func (r *itemRepo) Create(item models.Item) error {
	_, err := r.db.Exec(`INSERT INTO items (name, category_id, rack_id, stock, price, min_stock, warehouse_id) VALUES ($1,$2,$3,$4,$5,$6,$7)`,
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

func (r *itemRepo) GetByID(id int) (models.Item, error) {
	var item models.Item
	err := r.db.QueryRow(`SELECT id, name, category_id, rack_id, warehouse_id, stock, min_stock, price FROM items WHERE id=$1`, id).
		Scan(&item.ID, &item.Name, &item.CategoryID, &item.RackID, &item.WarehouseID, &item.Stock, &item.MinStock, &item.Price)
	return item, err
}

func (r *itemRepo) Update(id int, item models.Item) error {
	_, err := r.db.Exec(`UPDATE items SET name=$1, category_id=$2, rack_id=$3, warehouse_id=$4, stock=$5, min_stock=$6, price=$7 WHERE id=$8`,
		item.Name, item.CategoryID, item.RackID, item.WarehouseID, item.Stock, item.MinStock, item.Price, id)
	return err
}

func (r *itemRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM items WHERE id=$1`, id)
	return err
}
