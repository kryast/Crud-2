CRUD-2 using Go Gin Framework and Gorm

POST
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name":"Laptop", "price":15990000, "stock":10}'