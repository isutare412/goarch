package model

import "time"

type Customer struct {
	ID          int
	CreateTime  time.Time
	UpdateTime  time.Time
	Name        string
	DateOfBirth time.Time
}
