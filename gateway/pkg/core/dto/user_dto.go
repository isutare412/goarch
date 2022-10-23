package dto

import (
	"time"

	"github.com/isutare412/goarch/gateway/ent"
)

type User struct {
	Nickname   string
	Email      *string
	CreateTime time.Time
	UpdateTime time.Time
}

func (u *User) FromUser(usr *ent.User) {
	u.Nickname = usr.Nickname
	u.Email = usr.Email
	u.CreateTime = usr.CreateTime
	u.UpdateTime = usr.UpdateTime
}
