package repository

import (
	"assignment5/models"
	"database/sql"
)

type SaleRepository interface {
	Create(sale models.Sale) error
	GetAll() ([]models.Sale, error)
}

type saleRepo struct {
	db *sql.DB
}

func NewSaleRepository(db *sql.DB) SaleRepository {
	return &saleRepo{db}
}

func (r *saleRepo) Create(sale models.Sale) error {
	_, err := r.db.Exec(`INSERT INTO sales (user_id, created_at, item_id, qty, price, total) VALUES ($1,$2,$3,$4,$5,$6)`,
		sale.UserID, sale.CreatedAt, sale.ItemID, sale.Qty, sale.Price, sale.Total)
	return err
}

func (r *saleRepo) GetAll() ([]models.Sale, error) {
	rows, err := r.db.Query(`SELECT id, user_id, created_at, item_id, qty, price, total FROM sales`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []models.Sale
	for rows.Next() {
		var s models.Sale
		err := rows.Scan(&s.ID, &s.UserID, &s.CreatedAt, &s.ItemID, &s.Qty, &s.Price, &s.Total)
		if err != nil {
			return nil, err
		}
		sales = append(sales, s)
	}
	return sales, nil
}
