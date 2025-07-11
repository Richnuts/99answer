package main

import (
	"99user/config"
	"99user/handler"
	"99user/repository"
	"log"

	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()

	// Nice to have : req-id
	e.Use(middleware.Logger())
	db := NewDatabase(cfg.Database)
	repo := repository.NewRepository(db)
	userHandler := handler.NewUserHandler(repo)
	// Route
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users/:id", userHandler.GetUser)
	e.GET("/users", userHandler.GetUsers)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}

func NewDatabase(database string) *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", database)
	if err != nil {
		log.Fatalf("%v", err)
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL
	);`)

	return db
}
