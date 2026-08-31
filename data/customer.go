package data

import (
	"database/sql"
	"errors"
	"fmt"
	"midaas/models"

	"github.com/bwmarrin/snowflake"
)

type Customer struct {
	Id        int64
	Name      string
	CreatedAt string
	UpdatedAt string
}

func (c *Customer) CreateCustomer(tx *sql.Tx) (*models.Customer, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errors.New("error generating id ...")
	}

	var res Customer
	id := node.Generate().Int64()
	err = tx.QueryRow(
		`INSERT INTO customers (id, name) VALUES ($1, $2)
		RETURNING id, name, created_at, updated_at`,
		id, c.Name,
	).Scan(&res.Id, &res.Name, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(err.Error())
	}

	return &models.Customer{
		Id:        res.Id,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (c *Customer) GetAllCustomer() (*[]models.Customer, error) {
	Customers := make([]models.Customer, 0)
	selectCustomers, err := DB.Query("SELECT * FROM customers;")
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}

	for selectCustomers.Next() {
		var res models.Customer
		err := selectCustomers.Scan(
			&res.Id, &res.Name, &res.CreatedAt, &res.UpdatedAt,
		)
		if err != nil {
			return nil, errors.New("failed to scan data ...")
		}
		Customers = append(Customers, res)
	}
	return &Customers, nil
}

func (c *Customer) GetCustomerById() (*models.Customer, error) {
	var res models.Customer
	err := DB.QueryRow("SELECT * FROM customers where id = $1", c.Id).Scan(
		&res.Id, &res.Name, &res.CreatedAt, &res.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("no customer found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}
	return &res, nil
}
