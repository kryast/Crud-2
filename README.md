CRUD-2 using Go Gin Framework and Gorm

POST
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name":"Laptop", "price":15990000, "stock":10}'

GET
curl -X GET http://localhost:8080/products

curl -X GET http://localhost:8080/products/1

PUT
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Laptop Gaming", "price": 18500000, "stock": 5}'