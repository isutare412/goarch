package memory

import (
	"go.uber.org/fx"

	"github.com/isutare412/goarch/ioc/pkg/port"
)

var Module = fx.Module("memory",
	fx.Provide(
		fx.Annotate(
			NewCustomerRepository,
			fx.As(new(port.CustomerRepository)),
		),
	),
)
