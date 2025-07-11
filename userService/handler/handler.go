package handler

import (
	"99user/model"
	"99user/repository"
	"99user/utils"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo repository.Repository
}

type Handler interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
}

func NewUserHandler(repo repository.Repository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

// POST /user
func (h *UserHandler) CreateUser(c echo.Context) error {
	var input model.User
	if err := c.Bind(&input); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err)
	}

	now := time.Now().UnixMicro()
	input.CreatedAt = int(now)
	input.UpdatedAt = int(now)

	user, err := h.repo.CreateUser(input)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return utils.Response(c, http.StatusOK, "users", user)
}

// GET /user/:id
func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, errors.New("user ID is required"))
	}
	user, err := h.repo.GetUser(id)
	if err != nil {
		return err
	}
	return utils.Response(c, http.StatusOK, "users", user)
}

// GET /users
func (h *UserHandler) GetUsers(c echo.Context) error {
	pagination := model.Pagination{}
	if p := c.QueryParam("page_num"); p != "" {
		page, err := strconv.Atoi(p)
		if err != nil {
			return utils.ErrorResponse(c, http.StatusBadRequest, err)
		}

		pagination.Page = page
	}
	if p := c.QueryParam("page_size"); p != "" {
		perPage, err := strconv.Atoi(p)
		if err != nil {
			return utils.ErrorResponse(c, http.StatusBadRequest, err)
		}
		pagination.PerPage = perPage
	}

	if pagination.Page < 1 {
		pagination.Page = 1
	}
	if pagination.PerPage < 1 {
		pagination.PerPage = 10
	}

	users, err := h.repo.GetUsers(pagination)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return utils.Response(c, http.StatusOK, "users", users)
}
