package repository

import (
	"assignment5/models"
	"database/sql"
)

type WarehouseRepository interface {
	Create(w models.Warehouse) error
	GetAll() ([]models.Warehouse, error)
}

type warehouseRepo struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) WarehouseRepository {
	return &warehouseRepo{db}
}

func (r *warehouseRepo) Create(w models.Warehouse) error {
	_, err := r.db.Exec(`INSERT INTO warehouses (name, address) VALUES ($1, $2)`, w.Name, w.Address)
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
		err := rows.Scan(&w.ID, &w.Name, &w.Address)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, w)
	}
	return warehouses, nil
}
