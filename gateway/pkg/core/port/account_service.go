package port

import (
	"context"

	"github.com/isutare412/goarch/gateway/pkg/core/dto"
)

type AccountService interface {
	CreateUser(context.Context, dto.CreateUserRequest) error
	PromoteUser(context.Context, dto.PromoteUserRequest) error
	GetUserByNickname(context.Context, dto.GetUserByNicknameRequest) (dto.GetUserByNicknameResponse, error)

	CreateAdmin(context.Context, dto.CreateAdminRequest) error
	GetAdminByNickname(context.Context, dto.GetAdminByNicknameRequest) (dto.GetAdminByNicknameResponse, error)
}
