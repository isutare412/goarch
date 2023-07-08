package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/isutare412/goarch/ioc/pkg/log"
	"github.com/isutare412/goarch/ioc/pkg/port"
)

type customerHandler struct {
	customerService port.CustomerService
	log             *log.Logger
}

func newCustomerHandler(log *log.Logger, customerService port.CustomerService) *customerHandler {
	return &customerHandler{
		customerService: customerService,
		log:             log,
	}
}

func (h *customerHandler) router() http.Handler {
	r := chi.NewRouter()
	r.Post("/", h.registerCustomer)
	return r
}

func (h *customerHandler) registerCustomer(w http.ResponseWriter, r *http.Request) {
	h.log.S().Infof("Hello world!")
}
