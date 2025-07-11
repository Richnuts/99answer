package repository

import (
	"99user/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

type Repository interface {
	CreateUser(user model.User) (model.User, error)
	GetUser(id string) (model.User, error)
	GetUsers(pagination model.Pagination) ([]model.User, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user model.User) (model.User, error) {
	res, err := r.db.Exec(`
		INSERT INTO users (name, created_at, updated_at)
		VALUES (?, ?, ?)
	`, user.Name, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return model.User{}, err
	}
	user.ID = int(lastID)

	return user, err
}

func (r *repository) GetUser(id string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, `
		SELECT id, name, created_at, updated_at
		FROM users
		WHERE id = ?
	`, id)
	return user, err
}

func (r *repository) GetUsers(pagination model.Pagination) ([]model.User, error) {
	var users []model.User
	fmt.Println(pagination)
	err := r.db.Select(&users, `
		SELECT id, name, created_at, updated_at
		FROM users
		ORDER BY created_at DESC LIMIT ? OFFSET ?
	`, pagination.PerPage, pagination.Page)
	return users, err
}
