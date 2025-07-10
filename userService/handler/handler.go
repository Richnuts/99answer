package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo *sqlx.DB
}

type Handler interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
}

func NewUserHandler(db *sqlx.DB) *UserHandler {
	return &UserHandler{
		repo: db,
	}
}

// POST /user
func (h *UserHandler) CreateUser(c echo.Context) error {
	return nil
}

// GET /user/:id
func (h *UserHandler) GetUser(c echo.Context) error {
	return nil
}

// GET /users
func (h *UserHandler) GetUsers(c echo.Context) error {
	return nil
}
