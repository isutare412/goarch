package port

import (
	"context"

	"github.com/isutare412/goarch/gateway/ent"
)

type UserRepository interface {
	Create(context.Context, *ent.User) error
	ExistsByNickname(context.Context, string) (bool, error)
}
