package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/your-username/bookstore_microservice/api/handlers"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/your-username/bookstore_microservice/docs" // This is required to initialize Swagger documentation.
)

// @title Bookstore Microservice API
// @description RESTful APIs to manage a list of books in a store.
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/books", handlers.GetAllBooks)
	e.GET("/books/:id", handlers.GetBook)
	e.POST("/books", handlers.AddBook)

	// Swagger API documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
