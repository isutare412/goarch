package dto

import (
	"time"

	"github.com/isutare412/goarch/gateway/pkg/core/ent"
)

type User struct {
	Nickname  string    `json:"nickname" example:"redshore"`
	Email     *string   `json:"email,omitempty" example:"foo@bar.com"`
	CreatedAt time.Time `json:"created_at" example:"2022-10-30T09:04:22.799572Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-10-30T09:04:22.799572Z"`
}

func (u *User) FromUser(usr *ent.User) {
	u.Nickname = usr.Nickname
	u.Email = usr.Email
	u.CreatedAt = usr.CreateTime
	u.UpdatedAt = usr.UpdateTime
}
