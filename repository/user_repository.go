package repository

import (
	"assignment5/models"
	"database/sql"
)

type UserRepository interface {
	Create(user models.User) error
	GetAll() ([]models.User, error)
	GetByID(id int) (models.User, error)
	Update(id int, user models.User) error
	Delete(id int) error
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

func (r *userRepo) GetAll() ([]models.User, error) {
	rows, err := r.db.Query(`SELECT id, name, email, password, role FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepo) GetByID(id int) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(`SELECT id, name, email, password, role FROM users WHERE id=$1`, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	return user, err
}

func (r *userRepo) Update(id int, user models.User) error {
	_, err := r.db.Exec(`UPDATE users SET name=$1, email=$2, password=$3, role=$4 WHERE id=$5`, user.Name, user.Email, user.Password, user.Role, id)
	return err
}

func (r *userRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	return err
}
