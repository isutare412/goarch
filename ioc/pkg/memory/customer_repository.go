package memory

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/isutare412/goarch/ioc/pkg/core/model"
	"github.com/isutare412/goarch/ioc/pkg/log"
)

type CustomerRepository struct {
	mu             sync.Mutex
	customersByID  map[int]*model.Customer
	nextCustomerID int

	log *log.Logger
}

func NewCustomerRepository(log *log.Logger) *CustomerRepository {
	return &CustomerRepository{
		customersByID:  make(map[int]*model.Customer),
		nextCustomerID: 1,
		log:            log,
	}
}

func (repo *CustomerRepository) CreateCustomer(
	ctx context.Context,
	customer *model.Customer,
) (*model.Customer, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	copied := *customer
	copied.ID = repo.nextCustomerID
	repo.nextCustomerID++

	now := time.Now()
	copied.CreateTime = now
	copied.UpdateTime = now

	repo.customersByID[copied.ID] = &copied

	customer.ID = copied.ID
	customer.CreateTime = copied.CreateTime
	customer.UpdateTime = copied.UpdateTime
	return customer, nil
}

func (repo *CustomerRepository) GetCustomer(ctx context.Context, id int) (*model.Customer, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	customer, ok := repo.customersByID[id]
	if !ok {
		return nil, fmt.Errorf("customer not found with id(%d)", id)
	}

	copied := *customer
	return &copied, nil
}
