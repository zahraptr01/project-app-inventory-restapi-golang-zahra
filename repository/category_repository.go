package repository

import (
	"assignment5/models"
	"database/sql"
)

type CategoryRepository interface {
	Create(category models.Category) error
	GetAll() ([]models.Category, error)
}

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepo{db}
}

func (r *categoryRepo) Create(category models.Category) error {
	_, err := r.db.Exec(`INSERT INTO categories (name) VALUES ($1)`, category.Name)
	return err
}

func (r *categoryRepo) GetAll() ([]models.Category, error) {
	rows, err := r.db.Query(`SELECT id, name FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}
