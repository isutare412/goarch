package port

import (
	"context"

	"github.com/isutare412/goarch/ioc/pkg/core/dto"
)

type CustomerService interface {
	RegisterCustomer(context.Context, dto.RegisterCustomerRequest) (dto.RegisterCustomerResponse, error)
	GetCustomer(ctx context.Context, id int) (dto.GetCustomerResponse, error)
}
