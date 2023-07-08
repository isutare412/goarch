package port

import (
	"context"

	"github.com/isutare412/goarch/ioc/pkg/core/model"
)

type CustomerRepository interface {
	CreateCustomer(context.Context, *model.Customer) (*model.Customer, error)
	GetCustomer(ctx context.Context, id int) (*model.Customer, error)
}
