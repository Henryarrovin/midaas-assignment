package data

import (
	"database/sql"
	"errors"
	"fmt"
	"midaas/models"

	"github.com/bwmarrin/snowflake"
	"github.com/lib/pq"
)

type Order struct {
	Id           int64
	RestaurantId int64
	FoodId       []int64
	Paid         bool
	Status       string
	CreatedAt    string
	UpdatedAt    string
}

func (o *Order) CreateOrder(tx *sql.Tx) (*models.Order, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errors.New("error generating id ...")
	}

	var res Order
	id := node.Generate().Int64()
	err = tx.QueryRow(
		`INSERT INTO orders (id, restaurant_id, food_id, paid, status) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, restaurant_id, food_id, paid, status, created_at, updated_at`,
		id, o.RestaurantId, pq.Array(o.FoodId), o.Paid, "PENDING",
	).Scan(
		&res.Id, &res.RestaurantId, pq.Array(&res.FoodId),
		&res.Paid, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(err.Error())
	}

	return &models.Order{
		Id:           res.Id,
		RestaurantId: res.RestaurantId,
		FoodId:       res.FoodId,
		Paid:         res.Paid,
		Status:       res.Status,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}, nil
}

func (o *Order) GetAllOrder() (*[]models.Order, error) {
	Orders := make([]models.Order, 0)
	selectOrders, err := DB.Query("SELECT * FROM orders;")
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}

	for selectOrders.Next() {
		var res models.Order
		err := selectOrders.Scan(
			&res.Id, &res.RestaurantId, pq.Array(&res.FoodId),
			&res.Paid, &res.Status, &res.CreatedAt, &res.UpdatedAt,
		)
		if err != nil {
			return nil, errors.New("failed to scan data ...")
		}
		Orders = append(Orders, res)
	}
	return &Orders, nil
}

func (o *Order) GetOrderById() (*models.Order, error) {
	var res models.Order
	err := DB.QueryRow("SELECT * FROM orders where id = $1", o.Id).Scan(
		&res.Id, &res.RestaurantId, pq.Array(&res.FoodId),
		&res.Paid, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("order not found ...")
	}
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &res, nil
}

func (o *Order) UpdateOrderStatus() (*string, error) {
	row, err := DB.Exec("UPDATE orders SET status = $1 WHERE id = $2", "COMPLETED", o.Id)
	if err != nil {
		return nil, errors.New("error updating row ...")
	}

	r, _ := row.RowsAffected()

	res := fmt.Sprintf("Update %d row ...", r)
	return &res, nil
}
