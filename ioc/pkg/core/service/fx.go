package service

import (
	"go.uber.org/fx"

	"github.com/isutare412/goarch/ioc/pkg/core/service/customer"
	"github.com/isutare412/goarch/ioc/pkg/port"
)

var Module = fx.Module("service",
	fx.Provide(
		fx.Annotate(
			customer.NewService,
			fx.As(new(port.CustomerService)),
		),
	),
)
