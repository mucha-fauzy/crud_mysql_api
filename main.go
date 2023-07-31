package main

import (
	"crud_mysql_api/infras"
	"crud_mysql_api/internal/repository"
	"crud_mysql_api/internal/services"
	"crud_mysql_api/transport/routes"
	"fmt"
	"net/http"
)

func main() {
	// Create a new database connection
	db := infras.ProvideConn()

	// Initialize the repository with the database connection
	repo := repository.ProvideRepo(&db)

	// Initialize the service with the repository
	svc := services.ProvideService(repo)

	// Initialize the router with the service
	r := routes.NewRouter(svc)

	// Set up the HTTP server with the router and start listening
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r.SetupRoutes())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
