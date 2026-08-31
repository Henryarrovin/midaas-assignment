package services

import (
	"errors"
	"midaas/data"
	"midaas/models"
)

func CreateDish(d models.Dish) (*models.Dish, error) {
	restaurantId := d.RestaurantId
	restaurant := data.Restaurant{Id: restaurantId}
	rest, err := restaurant.GetRestaurantById()
	if err != nil {
		return nil, errors.New("restaurant not found ..." + err.Error())
	}

	data := data.Dish{
		RestaurantId: rest.Id,
		Name:         d.Name,
		CuisineType:  d.CuisineType,
		CurrencyCode: d.CurrencyCode,
		Price:        d.Price,
		PrepTime:     d.PrepTime,
	}
	resp, err := data.CreateDish()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return resp, nil
}

func GetAllDish() (*[]models.Dish, error) {
	data := data.Dish{}
	resp, err := data.GetAllDish()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return resp, nil
}

func GetDishById(d models.Dish) (*models.Dish, error) {
	data := data.Dish{
		Id: d.Id,
	}
	resp, err := data.GetDishById()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return resp, nil
}
