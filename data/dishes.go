package data

import (
	"database/sql"
	"errors"
	"midaas/models"

	"github.com/bwmarrin/snowflake"
)

type Dish struct {
	Id           int64
	RestaurantId int64
	Name         string
	CuisineType  string
	CurrencyCode string
	Price        int
	PrepTime     string
	CreatedAt    string
	UpdatedAt    string
}

func (d *Dish) CreateDish() (*models.Dish, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errors.New("error generating id ...")
	}

	var res Dish
	id := node.Generate().Int64()
	err = DB.QueryRow(
		`INSERT INTO dishes (id, restaurant_id, name, cuisine_type, currency_code, price, prep_time) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, restaurant_id, name, cuisine_type, currency_code, price, prep_time, created_at, updated_at`,
		id, d.RestaurantId, d.Name, d.CuisineType, d.CurrencyCode, d.Price, d.PrepTime,
	).Scan(&res.Id, &res.RestaurantId, &res.Name, &res.CuisineType, &res.CurrencyCode, &res.Price, &res.PrepTime,
		&res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &models.Dish{
		Id:           res.Id,
		RestaurantId: res.RestaurantId,
		Name:         res.Name,
		CuisineType:  res.CuisineType,
		CurrencyCode: res.CurrencyCode,
		Price:        res.Price,
		PrepTime:     res.PrepTime,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}, nil
}

func (d *Dish) GetAllDish() (*[]models.Dish, error) {
	dishes := make([]models.Dish, 0)
	selectDishes, err := DB.Query("SELECT * FROM dishes;")
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}

	for selectDishes.Next() {
		var res models.Dish
		err := selectDishes.Scan(
			&res.Id, &res.RestaurantId, &res.Name, &res.CuisineType, &res.CurrencyCode,
			&res.Price, &res.PrepTime, &res.CreatedAt, &res.UpdatedAt,
		)
		if err != nil {
			return nil, errors.New("failed to scan data ...")
		}
		dishes = append(dishes, res)
	}
	return &dishes, nil
}

func (d *Dish) GetDishById() (*models.Dish, error) {
	var res models.Dish
	err := DB.QueryRow("SELECT * FROM dishes where id = $1", d.Id).Scan(
		&res.Id, &res.RestaurantId, &res.Name, &res.CuisineType,
		&res.CurrencyCode, &res.Price, &res.PrepTime, &res.CreatedAt, &res.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}
	return &res, nil
}
