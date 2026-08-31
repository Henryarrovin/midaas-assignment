package services

import (
	"errors"
	"midaas/data"
	"midaas/models"
)

func CreateRestaurant(r models.Restaurant) (*models.Restaurant, error) {
	data := data.Restaurant{
		Name:        r.Name,
		PhoneNumber: r.PhoneNumber,
		WebsiteInfo: r.WebsiteInfo,
	}
	resp, err := data.CreateRestaurant()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return resp, nil
}

func GetAllRestaurant() (*[]models.Restaurant, error) {
	data := data.Restaurant{}
	resp, err := data.GetAllRestaurant()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return resp, nil
}

func GetRestaurantById(r models.Restaurant) (*models.Restaurant, error) {
	data := data.Restaurant{
		Id: r.Id,
	}
	resp, err := data.GetRestaurantById()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return resp, nil
}
