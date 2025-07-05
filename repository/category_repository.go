package repository

import (
	"assignment5/models"
	"database/sql"
)

type CategoryRepository interface {
	Create(category models.Category) error
	GetAll() ([]models.Category, error)
	GetByID(id int) (models.Category, error)
	Update(id int, category models.Category) error
	Delete(id int) error
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

func (r *categoryRepo) GetByID(id int) (models.Category, error) {
	var category models.Category
	err := r.db.QueryRow(`SELECT id, name FROM categories WHERE id=$1`, id).
		Scan(&category.ID, &category.Name)
	return category, err
}

func (r *categoryRepo) Update(id int, category models.Category) error {
	_, err := r.db.Exec(`UPDATE categories SET name=$1 WHERE id=$2`, category.Name, id)
	return err
}

func (r *categoryRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM categories WHERE id=$1`, id)
	return err
}
