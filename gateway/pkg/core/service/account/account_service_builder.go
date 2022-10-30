package account

import "github.com/isutare412/goarch/gateway/pkg/core/port"

type serviceBuilder struct {
	repoSession port.RepositorySession
	userRepo    port.UserRepository
	adminRepo   port.AdminRepository
}

func NewServiceBuilder() serviceBuilder { return serviceBuilder{} }

func (b serviceBuilder) WithRepositorySession(repoSession port.RepositorySession) serviceBuilder {
	b.repoSession = repoSession
	return b
}

func (b serviceBuilder) WithUserRepository(repo port.UserRepository) serviceBuilder {
	b.userRepo = repo
	return b
}

func (b serviceBuilder) WithAdminRepository(repo port.AdminRepository) serviceBuilder {
	b.adminRepo = repo
	return b
}

func (b serviceBuilder) Build() *Service {
	return &Service{
		repoSession: b.repoSession,
		userRepo:    b.userRepo,
		adminRepo:   b.adminRepo,
	}
}
