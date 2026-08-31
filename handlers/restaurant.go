package handlers

import (
	"encoding/json"
	"log"
	"midaas/models"
	"midaas/services"
	"net/http"
	"strconv"
)

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant models.Restaurant
	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
		log.Println("restaurant handler --- ", err.Error())
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	res, err := services.CreateRestaurant(restaurant)
	if err != nil {
		log.Println("restaurant handler --- ", err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetAllRestaurant(w http.ResponseWriter, r *http.Request) {
	res, err := services.GetAllRestaurant()
	if err != nil {
		log.Println("restaurant handler --- ", err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetRestaurantById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	restaurant := models.Restaurant{Id: id}
	res, err := services.GetRestaurantById(restaurant)
	if err != nil {
		log.Println("restaurant handler --- ", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
