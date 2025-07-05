package repository

import (
	"assignment5/models"
	"database/sql"
)

type RackRepository interface {
	Create(rack models.Rack) error
	GetAll() ([]models.Rack, error)
	GetByID(id int) (models.Rack, error)
	Update(id int, rack models.Rack) error
	Delete(id int) error
}

type rackRepo struct {
	db *sql.DB
}

func NewRackRepository(db *sql.DB) RackRepository {
	return &rackRepo{db}
}

func (r *rackRepo) Create(rack models.Rack) error {
	_, err := r.db.Exec(`INSERT INTO racks (code, warehouse_id) VALUES ($1, $2)`, rack.Code, rack.WarehouseID)
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
		err := rows.Scan(&r.ID, &r.Code, &r.WarehouseID)
		if err != nil {
			return nil, err
		}
		racks = append(racks, r)
	}
	return racks, nil
}

func (r *rackRepo) GetByID(id int) (models.Rack, error) {
	var rack models.Rack
	err := r.db.QueryRow(`SELECT id, code, warehouse_id FROM racks WHERE id=$1`, id).
		Scan(&rack.ID, &rack.Code, &rack.WarehouseID)
	return rack, err
}

func (r *rackRepo) Update(id int, rack models.Rack) error {
	_, err := r.db.Exec(`UPDATE racks SET code=$1, warehouse_id=$2 WHERE id=$3`, rack.Code, rack.WarehouseID, id)
	return err
}

func (r *rackRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM racks WHERE id=$1`, id)
	return err
}
