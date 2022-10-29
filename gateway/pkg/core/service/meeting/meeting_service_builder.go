package meeting

import "github.com/isutare412/goarch/gateway/pkg/core/port"

type serviceBuilder struct {
	repoSession port.RepositorySession
	userRepo    port.UserRepository
	meetingRepo port.MeetingRepository
}

func ServiceBuilder() serviceBuilder { return serviceBuilder{} }

func (b serviceBuilder) WithRepositorySession(repoSession port.RepositorySession) serviceBuilder {
	b.repoSession = repoSession
	return b
}

func (b serviceBuilder) WithUserRepository(userRepo port.UserRepository) serviceBuilder {
	b.userRepo = userRepo
	return b
}

func (b serviceBuilder) WithMeetingRepository(meetingRepo port.MeetingRepository) serviceBuilder {
	b.meetingRepo = meetingRepo
	return b
}

func (b serviceBuilder) Build() *Service {
	return &Service{
		repoSession: b.repoSession,
		userRepo:    b.userRepo,
		meetingRepo: b.meetingRepo,
	}
}
