package dto

import (
	"fmt"

	"github.com/isutare412/goarch/gateway/pkg/core/ent"
	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
)

type CreateUserRequest struct {
	Nickname string  `json:"nickname" example:"redshore"`
	Email    *string `json:"email" example:"foo@bar.com"`
}

func (r CreateUserRequest) Validate() error {
	if r.Nickname == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("nickname field is mandatory"),
		}
	}
	if r.Email != nil && *r.Email == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("email field should not be empty"),
		}
	}
	return nil
}

func (r CreateUserRequest) IntoUser() *ent.User {
	return &ent.User{
		Nickname: r.Nickname,
		Email:    r.Email,
	}
}

type PromoteUserRequest struct {
	Nickname    string
	PhoneNumber string
}

func (r PromoteUserRequest) Validate() error {
	if r.Nickname == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("nickname field is mandatory"),
		}
	}

	if r.PhoneNumber == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("phoneNumber field is mandatory"),
		}
	}
	return nil
}

func (r PromoteUserRequest) IntoAdmin() *ent.Admin {
	return &ent.Admin{
		PhoneNumber: r.PhoneNumber,
	}
}

type GetUserByNicknameRequest struct {
	Nickname string `json:"nickname" example:"redshore"`
}

func (r GetUserByNicknameRequest) Validate() error {
	if r.Nickname == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("nickname field is mandatory"),
		}
	}
	return nil
}

type GetUserByNicknameResponse struct {
	User
}

type CreateAdminRequest struct {
	Nickname    string
	Email       *string
	PhoneNumber string
}

func (r CreateAdminRequest) Validate() error {
	if r.Nickname == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("nickname field is mandatory"),
		}
	}
	if r.Email != nil && *r.Email == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("email field should not be empty"),
		}
	}
	if r.PhoneNumber == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("phoneNumber field is mandatory"),
		}
	}
	return nil
}

func (r CreateAdminRequest) IntoUser() *ent.User {
	return &ent.User{
		Nickname: r.Nickname,
		Email:    r.Email,
	}
}

func (r CreateAdminRequest) IntoAdmin() *ent.Admin {
	return &ent.Admin{
		PhoneNumber: r.PhoneNumber,
	}
}

type GetAdminByNicknameRequest struct {
	Nickname string
}

func (r GetAdminByNicknameRequest) Validate() error {
	if r.Nickname == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("nickname field is mandatory"),
		}
	}
	return nil
}

type GetAdminByNicknameResponse struct {
	Admin
}
