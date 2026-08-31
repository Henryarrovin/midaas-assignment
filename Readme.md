Documentation:
-------------

API Designed:
-------------
Restaurant APIs:
 - POST /api/restaurant
 - GET /api/restaurant
 - GET /api/restaurant/{id}

Dish APIs:
 - POST /api/dis
 - GET /api/dish
 - GET /api/dish/{id}

Employee APIs:
 - POST /api/employee
 - GET /api/employee
 - GET /api/employee/{id}

Order APIs:
 - POST /api/order
 - PATCH /api/order/{order_id}

--- PATCH order api shows the countdown of the order

EXAMPLE CURLS:

Restaurant CURLS:
curl --location 'http://localhost:8080/api/restaurant' \
--header 'Content-Type: application/json' \
--data '{
    "name": "MoonWalk",
    "phone_number": "0123456789",
    "website_info": "MoonWalk"
}'

curl --location 'http://localhost:8080/api/restaurant' \
--header 'Content-Type: application/json' \
--data ''

curl --location 'http://localhost:8080/api/restaurant/2092251129973837824' \
--header 'Content-Type: application/json' \
--data ''

DISHES CURLS:
curl --location 'http://localhost:8080/api/dish' \
--header 'Content-Type: application/json' \
--data '{
    "restaurant_id": 2092251129973837824,
    "name": "Noodles",
    "cuisine_type": "Indian",
    "currency_code": "INR",
    "price": 400,
    "prep_time": "50s"
}'

curl --location 'http://localhost:8080/api/dish' \
--header 'Content-Type: application/json'

curl --location 'http://localhost:8080/api/dish/2092592882111221760' \
--header 'Content-Type: application/json'

EMPLOYEE CURLS:
curl --location 'http://localhost:8080/api/employee' \
--header 'Content-Type: application/json' \
--data '{
    "restaurant_id": 2092251129973837824,
    "name": "employee3",
    "emp_type": "COOK"
}'

curl --location 'http://localhost:8080/api/employee' \
--header 'Content-Type: application/json' \
--data ''

curl --location 'http://localhost:8080/api/employee/2092596401241329664' \
--header 'Content-Type: application/json' \
--data ''

ORDER CURLS:
curl --location 'http://localhost:8080/api/order' \
--header 'Content-Type: application/json' \
--data '{
    "order": {
        "restaurant_id": 2092251129973837824,
        "food_id": [
            2092592882111221760,
            2092593506139770880,
            2092592882111221760,
            2092592882111221760,
            2092593506139770880,
            2092593506139770880
        ]
    },
    "customer": {
        "name": "Henry"
    }
}'

curl --location --request PATCH 'http://localhost:8080/api/order/2093004572275511296' \
--header 'Content-Type: application/json'
