package data

import (
	"database/sql"
	"errors"
	"fmt"
	"midaas/models"

	"github.com/bwmarrin/snowflake"
)

type Employee struct {
	Id           int64
	RestaurantId int64
	Name         string
	EmpType      string
	Free         bool
	CreatedAt    string
	UpdatedAt    string
}

func (e *Employee) CreateEmployee() (*models.Employee, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, errors.New("error generating id ...")
	}

	var res Employee
	id := node.Generate().Int64()
	err = DB.QueryRow(
		`INSERT INTO employees (id, restaurant_id, name, emp_type, free) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, restaurant_id, name, emp_type, free, created_at, updated_at`,
		id, e.RestaurantId, e.Name, e.EmpType, true,
	).Scan(&res.Id, &res.RestaurantId, &res.Name, &res.EmpType, &res.Free, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &models.Employee{
		Id:           res.Id,
		RestaurantId: res.RestaurantId,
		Name:         res.Name,
		EmpType:      res.EmpType,
		Free:         res.Free,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}, nil
}

func (e *Employee) GetAllEmployee() (*[]models.Employee, error) {
	Employees := make([]models.Employee, 0)
	selectEmployees, err := DB.Query("SELECT * FROM employees;")
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}

	for selectEmployees.Next() {
		var res models.Employee
		err := selectEmployees.Scan(
			&res.Id, &res.RestaurantId, &res.Name, &res.EmpType, &res.Free, &res.CreatedAt, &res.UpdatedAt,
		)
		if err != nil {
			return nil, errors.New("failed to scan data ...")
		}
		Employees = append(Employees, res)
	}
	return &Employees, nil
}

func (e *Employee) GetEmployeeById() (*models.Employee, error) {
	var res models.Employee
	err := DB.QueryRow("SELECT * FROM employees where id = $1", e.Id).Scan(
		&res.Id, &res.RestaurantId, &res.Name, &res.EmpType, &res.Free, &res.CreatedAt, &res.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}
	return &res, nil
}

func (e *Employee) GetFreeEmployeeByIdAndType() (*models.Employee, error) {
	var res models.Employee
	err := DB.QueryRow("SELECT * FROM employees where id = $1 AND emp_type = $2 AND free = true", e.Id, e.EmpType).Scan(
		&res.Id, &res.RestaurantId, &res.Name, &res.EmpType, &res.Free, &res.CreatedAt, &res.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}
	return &res, nil
}

func (e *Employee) GetAllFreeCook() (*[]models.Employee, error) {
	Employees := make([]models.Employee, 0)
	selectEmployees, err := DB.Query("SELECT * FROM employees where emp_type = 'COOK' AND free = true;")
	if err == sql.ErrNoRows {
		return nil, errors.New("no data found ...")
	}
	if err != nil {
		return nil, errors.New("error querying data ...")
	}

	for selectEmployees.Next() {
		var res models.Employee
		err := selectEmployees.Scan(
			&res.Id, &res.RestaurantId, &res.Name, &res.EmpType, &res.Free, &res.CreatedAt, &res.UpdatedAt,
		)
		if err != nil {
			return nil, errors.New("failed to scan data ...")
		}
		Employees = append(Employees, res)
	}
	return &Employees, nil
}

func (e *Employee) UpdateEmployeeAvailability() (*string, error) {
	row, err := DB.Exec("UPDATE employees SET free = $1 WHERE id = $2", e.Free, e.Id)
	if err != nil {
		return nil, errors.New("error updating row ...")
	}

	r, _ := row.RowsAffected()

	res := fmt.Sprintf("Update %d row ...", r)
	return &res, nil
}
