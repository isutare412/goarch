package dto

import (
	"time"

	"github.com/isutare412/goarch/ioc/pkg/core/model"
)

type CustomerInput struct {
	Name  string      `json:"name"`
	Birth DateOfBirth `json:"birth"`
}

func (c *CustomerInput) FromModel(mc *model.Customer) {
	var birth DateOfBirth
	birth.FromTime(mc.DateOfBirth)

	c.Name = mc.Name
	c.Birth = birth
}

func (c *CustomerInput) ToModel() *model.Customer {
	return &model.Customer{
		Name:        c.Name,
		DateOfBirth: c.Birth.ToTime(),
	}
}

type CustomerOutput struct {
	CustomerInput
	ID         int       `json:"id"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func (c *CustomerOutput) FromModel(mc *model.Customer) {
	c.CustomerInput.FromModel(mc)
	c.ID = mc.ID
	c.CreateTime = mc.CreateTime
	c.UpdateTime = mc.UpdateTime
}

type RegisterCustomerRequest struct {
	CustomerInput
}

type RegisterCustomerResponse struct {
	CustomerOutput
}

type DateOfBirth struct {
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
}

func (dob *DateOfBirth) ToTime() time.Time {
	return time.Date(dob.Year, dob.Month, dob.Day, 0, 0, 0, 0, time.Local)
}

func (dob *DateOfBirth) FromTime(t time.Time) {
	dob.Year = t.Year()
	dob.Month = t.Month()
	dob.Day = t.Day()
}

type GetCustomerResponse struct {
	CustomerOutput
}
