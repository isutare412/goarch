package port

import (
	"context"

	"github.com/isutare412/goarch/gateway/pkg/core/ent"
)

type AdminRepository interface {
	Save(ctx context.Context, admin *ent.Admin, userID int) (*ent.Admin, error)
	FindByNicknameWithUser(context.Context, string) (*ent.Admin, error)
	ExistsByNickname(context.Context, string) (bool, error)
}
