package repository

import (
	"assignment5/models"
	"database/sql"
)

type WarehouseRepository interface {
	Create(w models.Warehouse) error
	GetAll() ([]models.Warehouse, error)
	GetByID(id int) (models.Warehouse, error)
	Update(id int, warehouse models.Warehouse) error
	Delete(id int) error
}

type warehouseRepo struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) WarehouseRepository {
	return &warehouseRepo{db}
}

func (r *warehouseRepo) Create(w models.Warehouse) error {
	_, err := r.db.Exec(`INSERT INTO warehouses (name, location) VALUES ($1, $2)`, w.Name, w.Location)
	return err
}

func (r *warehouseRepo) GetAll() ([]models.Warehouse, error) {
	rows, err := r.db.Query(`SELECT id, name, location FROM warehouses`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []models.Warehouse
	for rows.Next() {
		var w models.Warehouse
		err := rows.Scan(&w.ID, &w.Name, &w.Location)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, w)
	}
	return warehouses, nil
}

func (r *warehouseRepo) GetByID(id int) (models.Warehouse, error) {
	var warehouse models.Warehouse
	err := r.db.QueryRow(`SELECT id, name, location FROM warehouses WHERE id=$1`, id).
		Scan(&warehouse.ID, &warehouse.Name, &warehouse.Location)
	return warehouse, err
}

func (r *warehouseRepo) Update(id int, warehouse models.Warehouse) error {
	_, err := r.db.Exec(`UPDATE warehouses SET name=$1, location=$2 WHERE id=$3`, warehouse.Name, warehouse.Location, id)
	return err
}

func (r *warehouseRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM warehouses WHERE id=$1`, id)
	return err
}
