package postgres

import (
	"context"
	"fmt"

	"github.com/isutare412/goarch/gateway/pkg/core/ent"
	"github.com/isutare412/goarch/gateway/pkg/core/ent/user"
	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
)

type AdminRepository struct {
	client *Client
}

func NewAdminRepository(client *Client) *AdminRepository {
	return &AdminRepository{client: client}
}

func (r *AdminRepository) Save(ctx context.Context, adm *ent.Admin, userID int) (*ent.Admin, error) {
	adm, err := r.txClient(ctx).Admin.
		Create().
		SetPhoneNumber(adm.PhoneNumber).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return nil, pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return adm, nil
}

func (r *AdminRepository) FindByNicknameWithUser(ctx context.Context, nickname string) (*ent.Admin, error) {
	adm, err := r.txClient(ctx).User.
		Query().
		Where(user.NicknameEQ(nickname)).
		QueryAdmin().
		WithUser().
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
	return adm, nil
}

func (r *AdminRepository) ExistsByNickname(ctx context.Context, nickname string) (bool, error) {
	exists, err := r.txClient(ctx).User.
		Query().
		Where(user.NicknameEQ(nickname)).
		QueryAdmin().
		Exist(ctx)
	if err != nil {
		return false, pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return exists, nil
}

func (r *AdminRepository) txClient(ctx context.Context) *ent.Client {
	if c := txFromCtx(ctx); c != nil {
		return c
	}
	return r.client.cli
}
