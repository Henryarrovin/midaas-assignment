package data

import (
	"database/sql"
	"errors"
	"midaas/models"

	"github.com/bwmarrin/snowflake"
)

type Restaurant struct {
	Id          int64
	Name        string
	PhoneNumber string
	WebsiteInfo string
	CreatedAt   string
	UpdatedAt   string
}

func (r *Restaurant) CreateRestaurant() (*models.Restaurant, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errors.New("error generating id ...")
	}

	var res Restaurant
	id := node.Generate().Int64()
	err = DB.QueryRow(
		`INSERT INTO restaurants (id, name, phone_number, website_info) VALUES ($1, $2, $3, $4)
		RETURNING id, name, phone_number, website_info, created_at, updated_at`,
		id, r.Name, r.PhoneNumber, r.WebsiteInfo,
	).Scan(&res.Id, &res.Name, &res.PhoneNumber, &res.WebsiteInfo, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, errors.New("error creating row ...")
	}

	return &models.Restaurant{
		Id:          res.Id,
		Name:        res.Name,
		PhoneNumber: res.PhoneNumber,
		WebsiteInfo: res.WebsiteInfo,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}, nil
}

func (r *Restaurant) GetAllRestaurant() (*[]models.Restaurant, error) {
	restaurants := make([]models.Restaurant, 0)
	selectRestaurants, err := DB.Query("SELECT * FROM restaurants;")
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}

	for selectRestaurants.Next() {
		var r models.Restaurant
		err := selectRestaurants.Scan(&r.Id, &r.Name, &r.PhoneNumber, &r.WebsiteInfo, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, errors.New("failed to scan data ...")
		}
		restaurants = append(restaurants, r)
	}
	return &restaurants, nil
}

func (r *Restaurant) GetRestaurantById() (*models.Restaurant, error) {
	var restaurant models.Restaurant
	err := DB.QueryRow("SELECT * FROM restaurants where id = $1", r.Id).Scan(
		&restaurant.Id, &restaurant.Name, &restaurant.PhoneNumber,
		&restaurant.WebsiteInfo, &restaurant.CreatedAt, &restaurant.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}
	return &restaurant, nil
}
