package repository

import (
	"assignment5/models"
	"database/sql"
)

type UserRepository interface {
	Create(user models.User) error
	GetByEmail(email string) (models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(user models.User) error {
	_, err := r.db.Exec(`INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)`, user.Name, user.Email, user.Password, user.Role)
	return err
}

func (r *userRepo) GetByEmail(email string) (models.User, error) {
	row := r.db.QueryRow(`SELECT id, name, email, password, role FROM users WHERE email=$1`, email)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	return user, err
}
