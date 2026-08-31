package services

import (
	"errors"
	"midaas/data"
	"midaas/models"
)

func CreateEmployee(e models.Employee) (*models.Employee, error) {
	restaurantId := e.RestaurantId
	restaurant := data.Restaurant{Id: restaurantId}
	rest, err := restaurant.GetRestaurantById()
	if err != nil {
		return nil, errors.New("restaurant not found ..." + err.Error())
	}

	data := data.Employee{
		RestaurantId: rest.Id,
		Name:         e.Name,
		EmpType:      e.EmpType,
	}
	resp, err := data.CreateEmployee()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return resp, nil
}

func GetAllEmployee() (*[]models.Employee, error) {
	data := data.Employee{}
	resp, err := data.GetAllEmployee()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return resp, nil
}

func GetEmployeeById(e models.Employee) (*models.Employee, error) {
	data := data.Employee{
		Id: e.Id,
	}
	resp, err := data.GetEmployeeById()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return resp, nil
}
