package dto

import (
	"fmt"
	"time"

	"github.com/isutare412/goarch/gateway/ent"
	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
)

type Admin struct {
	Nickname    string
	Email       *string
	PhoneNumber string
	CreateTime  time.Time
	UpdateTime  time.Time
}

func (a *Admin) FromAdmin(adm *ent.Admin) error {
	if adm.Edges.User == nil {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEntityNotLoaded,
			Simple: fmt.Errorf("user of admin(%d) not loaded", adm.ID),
		}
	}

	a.Nickname = adm.Edges.User.Nickname
	a.Email = adm.Edges.User.Email
	a.PhoneNumber = adm.PhoneNumber
	a.CreateTime = adm.CreateTime
	a.UpdateTime = adm.UpdateTime
	return nil
}
