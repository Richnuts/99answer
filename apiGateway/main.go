package main

import (
	"99gateway/config"
	"99gateway/handler"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()

	publicHandler := handler.NewPublicHandler(cfg)
	// Nice to have : req-id
	e.Use(middleware.Logger())
	// Routes
	// listings
	e.GET("/public-api/listings", publicHandler.GetListings)
	e.POST("/public-api/listings", publicHandler.CreateListing)
	// Users
	e.POST("/public-api/users", publicHandler.CreateUser)
	e.GET("/public-api/users/:id", publicHandler.GetUser)
	e.GET("/public-api/users", publicHandler.GetUsers)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
