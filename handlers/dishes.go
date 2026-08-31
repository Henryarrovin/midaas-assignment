package handlers

import (
	"encoding/json"
	"log"
	"midaas/models"
	"midaas/services"
	"net/http"
	"strconv"
)

func CreateDish(w http.ResponseWriter, r *http.Request) {
	var dish models.Dish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		log.Println("dish handler --- ", err.Error())
		http.Error(w, "invalid request body ...", http.StatusBadRequest)
		return
	}

	res, err := services.CreateDish(dish)
	if err != nil {
		log.Println("dish handler --- ", err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetAllDish(w http.ResponseWriter, r *http.Request) {
	res, err := services.GetAllDish()
	if err != nil {
		log.Println("dish handler --- ", err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetDishById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	dish := models.Dish{Id: id}
	res, err := services.GetDishById(dish)
	if err != nil {
		log.Println("dish handler --- ", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
