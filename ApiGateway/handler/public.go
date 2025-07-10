package handler

import (
	"99gateway/client/listing"
	"99gateway/client/user"
	"99gateway/config"
	"99gateway/utils"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PublicHandler struct {
	ListingClient *listing.ListingClient
	UserClient    *user.UserClient
}

func NewPublicHandler(cfg *config.Config) *PublicHandler {
	return &PublicHandler{
		ListingClient: listing.NewListingClient(cfg.ListingSvcURL),
		UserClient:    user.NewUserClient(cfg.UserSvcURL),
	}
}

// GET /public-api/listings
func (h *PublicHandler) GetListings(c echo.Context) error {
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	return utils.Success(c, resp)
}

// POST /public-api/listings
func (h *PublicHandler) CreateListing(c echo.Context) error {
	userID := c.FormValue("user_id")
	price := c.FormValue("price")
	_, err := strconv.Atoi(price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid price value"})
	}
	ListingType := c.FormValue("listing_type")

	if userID == "" || price == "" || ListingType == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing fields"})
	}

	params := url.Values{}
	params.Set("user_id", userID)
	params.Set("listing_type", ListingType)
	params.Set("price", price)

	resp, err := h.ListingClient.CreateListing(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	return utils.Success(c, resp)
}

// POST /public-api/users
func (h *PublicHandler) CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "name is required"})
	}

	params := url.Values{}
	params.Set("name", name)

	resp, err := h.UserClient.CreateUser(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	return utils.Success(c, resp)
}

// GET /public-api/users/:id
func (h *PublicHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "user ID is required"})
	}

	resp, err := h.UserClient.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	return utils.Success(c, resp)
}

// GET /public-api/users
func (h *PublicHandler) GetUsers(c echo.Context) error {
	params := url.Values{}
	if p := c.QueryParam("page_num"); p != "" {
		params.Set("page_num", p)
	}
	if p := c.QueryParam("page_size"); p != "" {
		params.Set("page_size", p)
	}

	resp, err := h.UserClient.GetUsers(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	return utils.Success(c, resp)
}
