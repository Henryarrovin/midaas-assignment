package handlers

import (
	"encoding/json"
	"log"
	"midaas/models"
	"midaas/services"
	"net/http"
	"strconv"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Println("employee handler --- ", err.Error())
		http.Error(w, "invalid request body ...", http.StatusBadRequest)
		return
	}

	res, err := services.CreateEmployee(employee)
	if err != nil {
		log.Println("employee handler --- ", err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	res, err := services.GetAllEmployee()
	if err != nil {
		log.Println("employee handler --- ", err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	employee := models.Employee{Id: id}
	res, err := services.GetEmployeeById(employee)
	if err != nil {
		log.Println("employee handler --- ", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
