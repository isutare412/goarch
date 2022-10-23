package postgres

import (
	"context"

	"github.com/isutare412/goarch/gateway/ent"
	"github.com/isutare412/goarch/gateway/ent/user"
)

type UserRepository struct {
	client *Client
}

func NewUserRepository(client *Client) *UserRepository {
	return &UserRepository{client: client}
}

func (r *UserRepository) Create(ctx context.Context, u *ent.User) error {
	return r.txClient(ctx).User.
		Create().
		SetNickname(u.Nickname).
		SetNillableEmail(u.Email).
		Exec(ctx)
}

func (r *UserRepository) ExistsByNickname(ctx context.Context, nickname string) (bool, error) {
	return r.txClient(ctx).User.
		Query().
		Where(user.NicknameEQ(nickname)).
		Exist(ctx)
}

func (r *UserRepository) txClient(ctx context.Context) *ent.Client {
	if c := txFromCtx(ctx); c != nil {
		return c
	}
	return r.client.cli
}
