package account

import (
	"context"
	"fmt"

	"github.com/isutare412/goarch/gateway/pkg/core/dto"
	"github.com/isutare412/goarch/gateway/pkg/core/port"
	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
)

type Service struct {
	repoSession port.RepositorySession
	userRepo    port.UserRepository
	adminRepo   port.AdminRepository
}

func (s *Service) CreateUser(ctx context.Context, req dto.CreateUserRequest) error {
	if err := req.Validate(); err != nil {
		return fmt.Errorf("validating request: %w", err)
	}

	var userToCreate = req.IntoUser()

	err := s.repoSession.WithTx(ctx, func(ctx context.Context) error {
		userExists, err := s.userRepo.ExistsByNickname(ctx, userToCreate.Nickname)
		if err != nil {
			return fmt.Errorf("checking existence of user(%s): %w", userToCreate.Nickname, err)
		} else if userExists {
			return pkgerr.Known{
				Errno:  pkgerr.ErrnoValueAlreadyExists,
				Simple: fmt.Errorf("user(%s) already exists", userToCreate.Nickname),
			}
		}

		_, err = s.userRepo.Save(ctx, userToCreate)
		if err != nil {
			return fmt.Errorf("saving user(%s): %w", userToCreate.Nickname, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("during transaction: %w", err)
	}
	return nil
}

func (s *Service) PromoteUser(ctx context.Context, req dto.PromoteUserRequest) error {
	if err := req.Validate(); err != nil {
		return fmt.Errorf("validating request: %w", err)
	}

	var adminToCreate = req.IntoAdmin()

	err := s.repoSession.WithTx(ctx, func(ctx context.Context) error {
		foundUser, err := s.userRepo.FindByNickname(ctx, req.Nickname)
		if err != nil {
			return fmt.Errorf("finding user(%s): %w", req.Nickname, err)
		}

		adminExists, err := s.adminRepo.ExistsByNickname(ctx, foundUser.Nickname)
		if err != nil {
			return fmt.Errorf("checking existence of admin(%s): %w", foundUser.Nickname, err)
		} else if adminExists {
			return pkgerr.Known{
				Errno:  pkgerr.ErrnoValueAlreadyExists,
				Simple: fmt.Errorf("admin(%s) already exists", foundUser.Nickname),
			}
		}

		_, err = s.adminRepo.Save(ctx, adminToCreate, foundUser.ID)
		if err != nil {
			return fmt.Errorf("saving admin(%s): %w", req.Nickname, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("during transaction: %w", err)
	}
	return nil
}

func (s *Service) GetUserByNickname(
	ctx context.Context,
	req dto.GetUserByNicknameRequest,
) (dto.GetUserByNicknameResponse, error) {
	var resp dto.GetUserByNicknameResponse

	if err := req.Validate(); err != nil {
		return resp, fmt.Errorf("validating request: %w", err)
	}

	foundUser, err := s.userRepo.FindByNickname(ctx, req.Nickname)
	if err != nil {
		return resp, fmt.Errorf("finding user(%s): %w", req.Nickname, err)
	}

	resp.FromUser(foundUser)
	return resp, nil
}

func (s *Service) CreateAdmin(ctx context.Context, req dto.CreateAdminRequest) error {
	if err := req.Validate(); err != nil {
		return fmt.Errorf("validating request: %w", err)
	}

	var userToCreate = req.IntoUser()
	var adminToCreate = req.IntoAdmin()

	err := s.repoSession.WithTx(ctx, func(ctx context.Context) error {
		userExists, err := s.userRepo.ExistsByNickname(ctx, userToCreate.Nickname)
		if err != nil {
			return fmt.Errorf("checking existence of user(%s): %w", userToCreate.Nickname, err)
		} else if userExists {
			return pkgerr.Known{
				Errno:  pkgerr.ErrnoValueAlreadyExists,
				Simple: fmt.Errorf("user(%s) already exists", userToCreate.Nickname),
			}
		}

		createdUser, err := s.userRepo.Save(ctx, userToCreate)
		if err != nil {
			return fmt.Errorf("saving user(%s): %w", userToCreate.Nickname, err)
		}

		_, err = s.adminRepo.Save(ctx, adminToCreate, createdUser.ID)
		if err != nil {
			return fmt.Errorf("saving admin(%s): %w", createdUser.Nickname, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("during transaction: %w", err)
	}
	return nil
}

func (s *Service) GetAdminByNickname(
	ctx context.Context,
	req dto.GetAdminByNicknameRequest,
) (dto.GetAdminByNicknameResponse, error) {
	var resp dto.GetAdminByNicknameResponse

	if err := req.Validate(); err != nil {
		return resp, fmt.Errorf("validating request: %w", err)
	}

	admin, err := s.adminRepo.FindByNicknameWithUser(ctx, req.Nickname)
	if err != nil {
		return resp, fmt.Errorf("finding admin by nickname(%s): %w", req.Nickname, err)
	}

	if err := resp.FromAdmin(admin); err != nil {
		return resp, fmt.Errorf("converting admin entity to response: %w", err)
	}
	return resp, nil
}
