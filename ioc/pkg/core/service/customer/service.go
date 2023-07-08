package customer

import (
	"context"
	"fmt"

	"github.com/isutare412/goarch/ioc/pkg/core/dto"
	"github.com/isutare412/goarch/ioc/pkg/log"
	"github.com/isutare412/goarch/ioc/pkg/port"
)

type Service struct {
	customerRepo port.CustomerRepository
	log          *log.Logger
}

func NewService(customerRepo port.CustomerRepository, log *log.Logger) *Service {
	return &Service{
		customerRepo: customerRepo,
		log:          log,
	}
}

func (s *Service) RegisterCustomer(
	ctx context.Context,
	req dto.RegisterCustomerRequest,
) (dto.RegisterCustomerResponse, error) {
	customerRequested := req.ToModel()
	customerCreated, err := s.customerRepo.CreateCustomer(ctx, customerRequested)
	if err != nil {
		return dto.RegisterCustomerResponse{}, fmt.Errorf("creating customer: %w", err)
	}

	var customerDTO dto.CustomerOutput
	customerDTO.FromModel(customerCreated)
	return dto.RegisterCustomerResponse{CustomerOutput: customerDTO}, nil
}

func (s *Service) GetCustomer(ctx context.Context, id int) (dto.GetCustomerResponse, error) {
	customer, err := s.customerRepo.GetCustomer(ctx, id)
	if err != nil {
		return dto.GetCustomerResponse{}, fmt.Errorf("getting customer: %w", err)
	}

	var customerDTO dto.CustomerOutput
	customerDTO.FromModel(customer)
	return dto.GetCustomerResponse{CustomerOutput: customerDTO}, nil
}
