package repository

import (
	"assignment5/models"
	"database/sql"
)

type RackRepository interface {
	Create(rack models.Rack) error
	GetAll() ([]models.Rack, error)
}

type rackRepo struct {
	db *sql.DB
}

func NewRackRepository(db *sql.DB) RackRepository {
	return &rackRepo{db}
}

func (r *rackRepo) Create(rack models.Rack) error {
	_, err := r.db.Exec(`INSERT INTO racks (name, warehouse_id) VALUES ($1, $2)`, rack.Name, rack.WarehouseID)
	return err
}

func (r *rackRepo) GetAll() ([]models.Rack, error) {
	rows, err := r.db.Query(`SELECT id, code, warehouse_id FROM racks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var racks []models.Rack
	for rows.Next() {
		var r models.Rack
		err := rows.Scan(&r.ID, &r.Name, &r.WarehouseID)
		if err != nil {
			return nil, err
		}
		racks = append(racks, r)
	}
	return racks, nil
}
