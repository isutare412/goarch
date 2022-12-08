package port

import (
	"context"

	"github.com/isutare412/goarch/gateway/pkg/core/ent"
)

type UserRepository interface {
	Save(context.Context, *ent.User) (*ent.User, error)
	FindByNickname(context.Context, string) (*ent.User, error)
	FindByNicknameIn(context.Context, []string) ([]*ent.User, error)
	ExistsByNickname(context.Context, string) (bool, error)
}
