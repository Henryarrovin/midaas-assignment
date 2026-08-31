package data

import (
	"database/sql"
	"errors"
	"fmt"
	"midaas/models"

	"github.com/bwmarrin/snowflake"
)

type Countdown struct {
	Id          int64
	OrderId     int64
	CustomerId  int64
	CreatedAt   string
	PrepareTime string
}

func (c *Countdown) CreateCountdown(tx *sql.Tx) (*models.Countdown, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errors.New("error generating id ...")
	}

	var res Countdown
	id := node.Generate().Int64()
	err = tx.QueryRow(
		`INSERT INTO countdown (id, order_id, customer_id, prepare_time) VALUES ($1, $2, $3, $4)
		RETURNING id, order_id, customer_id, created_at, prepare_time`,
		id, c.OrderId, c.CustomerId, c.PrepareTime,
	).Scan(&res.Id, &res.OrderId, &res.CustomerId, &res.CreatedAt, &res.PrepareTime)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(err.Error())
	}

	return &models.Countdown{
		Id:          res.Id,
		OrderId:     res.OrderId,
		CustomerId:  res.CustomerId,
		CreatedAt:   res.CreatedAt,
		PrepareTime: res.PrepareTime,
	}, nil
}

func (c *Countdown) GetAllCountdown() (*[]models.Countdown, error) {
	countdown := make([]models.Countdown, 0)
	selectCountdown, err := DB.Query("SELECT * FROM countdown;")
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}

	for selectCountdown.Next() {
		var r models.Countdown
		err := selectCountdown.Scan(&r.Id, &r.OrderId, &r.CustomerId, &r.CreatedAt, &r.PrepareTime)
		if err != nil {
			return nil, errors.New("failed to scan data ...")
		}
		countdown = append(countdown, r)
	}
	return &countdown, nil
}

func (c *Countdown) GetCountdownById() (*models.Countdown, error) {
	var countdown models.Countdown
	err := DB.QueryRow("SELECT * FROM countdown where id = $1", c.Id).Scan(
		&countdown.Id, &countdown.OrderId, &countdown.CustomerId,
		&countdown.CreatedAt, &countdown.PrepareTime,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}
	return &countdown, nil
}

func (c *Countdown) GetCountdownByOrderId() (*[]models.Countdown, error) {
	countdown := make([]models.Countdown, 0)
	countdownRes, err := DB.Query("SELECT * FROM countdown where order_id = $1", c.OrderId)
	if err == sql.ErrNoRows {
		return nil, errors.New("countdown for the order not found ...")
	}
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if countdownRes.Next() {
		var c models.Countdown
		err := countdownRes.Scan(&c.Id, &c.OrderId, &c.CustomerId, &c.CreatedAt, &c.PrepareTime)
		if err != nil {
			return nil, errors.New("failed to scan data ...")
		}
		countdown = append(countdown, c)
	}
	return &countdown, nil
}

func (c *Countdown) UpdateCountdownTime() (*string, error) {
	row, err := DB.Exec("UPDATE countdown SET prepare_time = $1 WHERE id = $2", c.PrepareTime, c.Id)
	if err != nil {
		return nil, errors.New("error updating row ...")
	}

	r, _ := row.RowsAffected()

	res := fmt.Sprintf("Update %d row ...", r)
	return &res, nil
}
