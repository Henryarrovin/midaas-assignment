package handlers

import (
	"encoding/json"
	"log"
	"midaas/models"
	"midaas/services"
	"net/http"
	"strconv"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var model models.OrderResp
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Println("order handler --- ", err.Error())
		http.Error(w, "invalid request body ...", http.StatusBadRequest)
		return
	}

	order := models.Order{
		RestaurantId: model.Order.RestaurantId,
		FoodId:       model.Order.FoodId,
	}
	customer := models.Customer{
		Name: model.Customer.Name,
	}

	res, err := services.CreateOrder(order, customer)
	if err != nil {
		log.Println("order handler --- ", err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetCountDown(w http.ResponseWriter, r *http.Request) {
	orderId, err := strconv.ParseInt(r.PathValue("order_id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	res, err := services.GetCountDown(orderId)
	if err != nil {
		log.Println("order handler --- ", err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
