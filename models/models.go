package models

type Restaurant struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	WebsiteInfo string `json:"website_info"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Dish struct {
	Id           int64  `json:"id"`
	RestaurantId int64  `json:"restaurant_id"`
	Name         string `json:"name"`
	CuisineType  string `json:"cuisine_type"`
	CurrencyCode string `json:"currency_code"`
	Price        int    `json:"price"`
	PrepTime     string `json:"prep_time"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type Employee struct {
	Id           int64  `json:"id"`
	RestaurantId int64  `json:"restaurant_id"`
	Name         string `json:"name"`
	EmpType      string `json:"emp_type"`
	Free         bool   `json:"free"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type Customer struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Order struct {
	Id           int64   `json:"id"`
	RestaurantId int64   `json:"restaurant_id"`
	FoodId       []int64 `json:"food_id"`
	Paid         bool    `json:"paid"`
	Status       string  `json:"status"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type Countdown struct {
	Id          int64  `json:"id"`
	OrderId     int64  `json:"order_id"`
	CustomerId  int64  `json:"customer_id"`
	CreatedAt   string `json:"created_at"`
	PrepareTime string `json:"prepare_time"`
}

type OrderResp struct {
	Order     Order     `json:"order"`
	Customer  Customer  `json:"customer"`
	Countdown Countdown `json:"countdown"`
}
