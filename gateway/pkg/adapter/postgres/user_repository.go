package postgres

import (
	"context"
	"fmt"

	"github.com/isutare412/goarch/gateway/ent"
	"github.com/isutare412/goarch/gateway/ent/user"
	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
)

type UserRepository struct {
	client *Client
}

func NewUserRepository(client *Client) *UserRepository {
	return &UserRepository{client: client}
}

func (r *UserRepository) Save(ctx context.Context, u *ent.User) (*ent.User, error) {
	usr, err := r.txClient(ctx).User.
		Create().
		SetNickname(u.Nickname).
		SetNillableEmail(u.Email).
		Save(ctx)
	if err != nil {
		return nil, pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return usr, nil
}

func (r *UserRepository) FindByNickname(ctx context.Context, nickname string) (*ent.User, error) {
	usr, err := r.txClient(ctx).User.
		Query().
		Where(user.NicknameEQ(nickname)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, pkgerr.Known{
				Errno:  pkgerr.ErrnoNotFound,
				Origin: err,
				Simple: fmt.Errorf("user(%s) not found", nickname),
			}
		} else {
			return nil, pkgerr.Known{
				Errno:  pkgerr.ErrnoRepository,
				Origin: err,
			}
		}
	}
	return usr, nil
}

func (r *UserRepository) FindByNicknameIn(ctx context.Context, nicknames []string) ([]*ent.User, error) {
	usrs, err := r.txClient(ctx).User.
		Query().
		Where(user.NicknameIn(nicknames...)).
		All(ctx)
	if err != nil {
		return nil, pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return usrs, nil
}

func (r *UserRepository) ExistsByNickname(ctx context.Context, nickname string) (bool, error) {
	exists, err := r.txClient(ctx).User.
		Query().
		Where(user.NicknameEQ(nickname)).
		Exist(ctx)
	if err != nil {
		return false, pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return exists, nil
}

func (r *UserRepository) txClient(ctx context.Context) *ent.Client {
	if c := txFromCtx(ctx); c != nil {
		return c
	}
	return r.client.cli
}
