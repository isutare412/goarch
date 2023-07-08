package model

import "time"

type Order struct {
	ID         int
	CustomerID int
	CreateTime time.Time
}

type OrderItem struct {
	OrderID  int
	ItemName string
	Count    int
}
