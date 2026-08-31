package main

import (
	"log"
	"midaas/data"
	"midaas/handlers"
	"net/http"
)

func main() {
	if err := data.NewDB(); err != nil {
		log.Fatal("failed to connect to db ...")
	}
	log.Println("Connected to db ...")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/restaurant", handlers.CreateRestaurant)
	mux.HandleFunc("GET /api/restaurant", handlers.GetAllRestaurant)
	mux.HandleFunc("GET /api/restaurant/{id}", handlers.GetRestaurantById)

	mux.HandleFunc("POST /api/dish", handlers.CreateDish)
	mux.HandleFunc("GET /api/dish", handlers.GetAllDish)
	mux.HandleFunc("GET /api/dish/{id}", handlers.GetDishById)

	mux.HandleFunc("POST /api/employee", handlers.CreateEmployee)
	mux.HandleFunc("GET /api/employee", handlers.GetAllEmployee)
	mux.HandleFunc("GET /api/employee/{id}", handlers.GetEmployeeById)

	mux.HandleFunc("POST /api/order", handlers.CreateOrder)
	mux.HandleFunc("PATCH /api/order/{order_id}", handlers.GetCountDown)

	log.Println("connected to server at port 8080 ...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("failed to start server ...")
	}
}
