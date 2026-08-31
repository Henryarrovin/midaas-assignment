package services

import (
	"errors"
	"fmt"
	"log"
	"midaas/data"
	"midaas/models"
	"strconv"
	"strings"
	"time"
)

func CreateOrder(order models.Order, customer models.Customer) (*models.OrderResp, error) {
	restaurantId := order.RestaurantId
	restaurant := data.Restaurant{Id: restaurantId}
	r, err := restaurant.GetRestaurantById()
	if err != nil {
		return nil, errors.New("restaurant not found ..." + err.Error())
	}

	empl := data.Employee{
		EmpType: "COOK",
		Free:    true,
	}
	// availableCooks, err := empl.GetAllFreeCook()
	// if err != nil {
	// 	return nil, errors.New(err.Error())
	// }

	var availableCooksRes []models.Employee
	for {
		availableCooks, err := empl.GetAllFreeCook()
		if err != nil {
			return nil, errors.New(err.Error())
		}
		availableCooksRes = *availableCooks
		if len(*availableCooks) > 0 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	foodIds := order.FoodId
	fmt.Println(order.FoodId)
	cooks := make([]int64, len(availableCooksRes))
	for _, v := range foodIds {
		food := data.Dish{Id: v}
		f, err := food.GetDishById()
		if err != nil {
			return nil, errors.New("dish with id = " + strconv.FormatInt(v, 10) + " not found ..." + err.Error())
		}

		timePrep := strings.Split(f.PrepTime, ":")
		h, err := strconv.Atoi(timePrep[0])
		if err != nil {
			return nil, errors.New("parse failed" + err.Error())
		}
		m, err := strconv.Atoi(timePrep[1])
		if err != nil {
			return nil, errors.New("parse failed" + err.Error())
		}
		s, err := strconv.Atoi(timePrep[2])
		if err != nil {
			return nil, errors.New("parse failed" + err.Error())
		}
		// prepTime := (h * 60) + m + (s / 60)
		prepTime := (h * 3600) + (m * 60) + s
		cook := 0
		for i := 1; i < len(cooks); i++ {
			if cooks[i] < cooks[cook] {
				cook = i
			}
		}
		cooks[cook] += int64(prepTime)
		cookModel := data.Employee{
			Id:   (availableCooksRes)[cook].Id,
			Free: false,
		}
		c, err := cookModel.UpdateEmployeeAvailability()
		if err != nil {
			return nil, errors.New("unable to update chef status false ...")
		}
		log.Println(c)

		cookId := cookModel.Id
		duration := time.Duration(prepTime) * time.Second

		go func() {
			time.Sleep(duration)
			cookModel := data.Employee{
				Id:   cookId,
				Free: true,
			}
			c, err := cookModel.UpdateEmployeeAvailability()
			if err != nil {
				log.Println("unable to update chef status true ...")
			}
			log.Println(*c)
		}()
	}

	totalPrepTime := 0
	for _, v := range cooks {
		if v > int64(totalPrepTime) {
			totalPrepTime = int(v)
		}
	}

	tx, err := data.DB.Begin()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	createdOrder := data.Order{
		RestaurantId: r.Id,
		FoodId:       foodIds,
		Paid:         true,
		Status:       "PENDING",
	}
	fmt.Println(createdOrder)
	orderRes, err := createdOrder.CreateOrder(tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New(err.Error())
	}

	createCustomer := data.Customer{
		Name: customer.Name,
	}
	customerRes, err := createCustomer.CreateCustomer(tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New(err.Error())
	}

	createCountdown := data.Countdown{
		OrderId:     orderRes.Id,
		CustomerId:  customerRes.Id,
		PrepareTime: strconv.Itoa(totalPrepTime),
	}
	countdownRes, err := createCountdown.CreateCountdown(tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New(err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errors.New(err.Error())
	}

	return &models.OrderResp{
		Order:     *orderRes,
		Customer:  *customerRes,
		Countdown: *countdownRes,
	}, nil
}

type GetCountdownResult struct {
	OrderId       int64  `json:"order_id"`
	CustomerId    int64  `json:"customer_id"`
	RemainingTime string `json:"remaining_time"`
	Status        string `json:"status"`
}

func GetCountDown(orderId int64) (*[]GetCountdownResult, error) {
	order := data.Order{Id: orderId}
	orderRes, err := order.GetOrderById()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	countdown := data.Countdown{OrderId: orderRes.Id}
	res := []GetCountdownResult{}

	countdowns, err := countdown.GetCountdownByOrderId()
	for _, v := range *countdowns {
		if v.PrepareTime == "0" {
			order.UpdateOrderStatus()
			continue
		}

		prepTime := 0
		time := strings.Split(v.PrepareTime, ":")
		h, err := strconv.Atoi(time[0])
		if err != nil {
			return nil, errors.New("parse failed" + err.Error())
		}
		m, err := strconv.Atoi(time[1])
		if err != nil {
			return nil, errors.New("parse failed" + err.Error())
		}
		s, err := strconv.Atoi(time[2])
		if err != nil {
			return nil, errors.New("parse failed" + err.Error())
		}
		// prepTime += (h * 60) + m + (s / 60)
		prepTime += (h * 3600) + (m * 60) + s

		remainingTime, err := calculateRemainingTime(prepTime, v.CreatedAt)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		res = append(res, GetCountdownResult{
			OrderId:       v.OrderId,
			CustomerId:    v.CustomerId,
			RemainingTime: *remainingTime,
			Status:        orderRes.Status,
		})
	}

	return &res, nil
}

func calculateRemainingTime(prepTime int, createdAtString string) (*string, error) {
	fmt.Println("PREP TIME ==> ", prepTime)
	if prepTime <= 0 {
		remainingTimeRes := "00:00"
		return &remainingTimeRes, nil
	}
	// layout := "2006-01-02 15:04:05.999999"
	createdAt, err := time.Parse(time.RFC3339, createdAtString)
	if err != nil {
		return nil, errors.New("can't parse date ..." + err.Error())
	}
	fmt.Println("CReated At ==> ", createdAt)

	elapsed := int(time.Since(createdAt).Seconds())
	remaining := max(prepTime-elapsed, 0)
	fmt.Println("ELAPSED ==> ", elapsed)
	fmt.Println("REMAINING ==> ", remaining)

	minutes := remaining / 60
	seconds := remaining % 60

	remainingTimeRes := fmt.Sprintf("%02d:%02d", minutes, seconds)
	fmt.Println("remainingTimeRes ===> ", remainingTimeRes)
	return &remainingTimeRes, nil
}
