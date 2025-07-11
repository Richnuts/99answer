package handler

import (
	"99gateway/client/listing"
	"99gateway/client/user"
	"99gateway/config"
	"99gateway/utils"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
)

type publicHandler struct {
	ListingClient *listing.ListingClient
	UserClient    *user.UserClient
}

type Handler interface {
	CreateListing(c echo.Context) error
	CreateUser(c echo.Context) error
	GetListings(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
}

func NewPublicHandler(cfg *config.Config) Handler {
	return &publicHandler{
		ListingClient: listing.NewListingClient(cfg.ListingSvcURL),
		UserClient:    user.NewUserClient(cfg.UserSvcURL),
	}
}

// GET /public-api/listings
func (h *publicHandler) GetListings(c echo.Context) error {
	params := url.Values{}
	if p := c.QueryParam("page_num"); p != "" {
		params.Set("page_num", p)
	}
	if p := c.QueryParam("page_size"); p != "" {
		params.Set("page_size", p)
	}
	if p := c.QueryParam("user_id"); p != "" {
		params.Set("user_id", p)
	}

	resp, err := h.ListingClient.GetListings(params)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	return utils.Success(c, resp)
}

// POST /public-api/listings
func (h *publicHandler) CreateListing(c echo.Context) error {
	userID := c.FormValue("user_id")
	price := c.FormValue("price")
	_, err := strconv.Atoi(price)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, errors.New("invalid price value"))
	}
	ListingType := c.FormValue("listing_type")

	if userID == "" || price == "" || ListingType == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, errors.New("missing fields"))
	}

	params := url.Values{}
	params.Set("user_id", userID)
	params.Set("listing_type", ListingType)
	params.Set("price", price)

	resp, err := h.ListingClient.CreateListing(params)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	return utils.Success(c, resp)
}

// POST /public-api/users
func (h *publicHandler) CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, errors.New("name is required"))
	}

	resp, err := h.UserClient.CreateUser(user.CreateUserInput{Name: name})
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	return utils.Success(c, resp)
}

// GET /public-api/users/:id
func (h *publicHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, errors.New("user ID is required"))
	}

	resp, err := h.UserClient.GetUser(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	return utils.Success(c, resp)
}

// GET /public-api/users
func (h *publicHandler) GetUsers(c echo.Context) error {
	params := url.Values{}
	if p := c.QueryParam("page_num"); p != "" {
		params.Set("page_num", p)
	}
	if p := c.QueryParam("page_size"); p != "" {
		params.Set("page_size", p)
	}

	resp, err := h.UserClient.GetUsers(params)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	return utils.Success(c, resp)
}
