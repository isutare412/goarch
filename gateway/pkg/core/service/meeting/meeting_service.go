package meeting

import (
	"context"
	"fmt"

	"github.com/isutare412/goarch/gateway/ent"
	"github.com/isutare412/goarch/gateway/pkg/core/dto"
	"github.com/isutare412/goarch/gateway/pkg/core/port"
	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
)

type Service struct {
	repoSession port.RepositorySession
	userRepo    port.UserRepository
	meetingRepo port.MeetingRepository
}

func (s *Service) CreateMeeting(ctx context.Context, req dto.CreateMeetingRequest) (dto.CreateMeetingResponse, error) {
	var resp dto.CreateMeetingResponse

	if err := req.Validate(); err != nil {
		return resp, fmt.Errorf("validating request: %w", err)
	}

	var meetingToCreate = req.IntoMeeting()

	err := s.repoSession.WithTx(ctx, func(ctx context.Context) error {
		organizer, err := s.userRepo.FindByNickname(ctx, req.OrganizerNickname)
		if err != nil {
			return fmt.Errorf("finding organizer(%s): %w", req.OrganizerNickname, err)
		}

		meeting, err := s.meetingRepo.Save(ctx, meetingToCreate, organizer.ID)
		if err != nil {
			return fmt.Errorf("saving meeting(%s): %w", req.Title, err)
		}
		resp.MeetingID = meeting.ID

		return nil
	})
	if err != nil {
		return resp, fmt.Errorf("during transaction: %w", err)
	}
	return resp, nil
}

func (s *Service) AddParticipants(ctx context.Context, req dto.AddParticipantsRequest) error {
	if err := req.Validate(); err != nil {
		return fmt.Errorf("validating request: %w", err)
	}

	err := s.repoSession.WithTx(ctx, func(ctx context.Context) error {
		exists, err := s.meetingRepo.ExistsByID(ctx, req.MeetingID)
		if err != nil {
			return fmt.Errorf("checking existence of meeting(%d): %w", req.MeetingID, err)
		} else if !exists {
			return pkgerr.Known{
				Errno:  pkgerr.ErrnoNotFound,
				Simple: fmt.Errorf("meeting(%d) not found", req.MeetingID),
			}
		}

		users, err := s.userRepo.FindByNicknameIn(ctx, req.ParticipantNicknames)
		if err != nil {
			return fmt.Errorf("finding users by nicknames: %w", err)
		} else if notFounds := searchUserNotFoundNicknames(users, req.ParticipantNicknames); len(notFounds) > 0 {
			return pkgerr.Known{
				Errno:  pkgerr.ErrnoNotFound,
				Simple: fmt.Errorf("users(%v) not found", notFounds),
			}
		}

		var userIDsToParticipate = make([]int, 0, len(users))
		for _, u := range users {
			userIDsToParticipate = append(userIDsToParticipate, u.ID)
		}

		participants, err := s.meetingRepo.FindParticipantsByIDAndParticipantIDs(ctx, req.MeetingID, userIDsToParticipate)
		if err != nil {
			return fmt.Errorf("checking participants existence: %w", err)
		} else if len(participants) > 0 {
			return pkgerr.Known{
				Errno: pkgerr.ErrnoValueAlreadyExists,
				Simple: fmt.Errorf("user(%s) already participating meeting(%d)",
					participants[0].Nickname, req.MeetingID),
			}
		}

		err = s.meetingRepo.AddParticipants(ctx, req.MeetingID, userIDsToParticipate)
		if err != nil {
			return fmt.Errorf("adding participants to meeting(%d): %w", req.MeetingID, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("during transaction: %w", err)
	}
	return nil
}

func searchUserNotFoundNicknames(foundUsers []*ent.User, requestedNicknames []string) (notFoundNicknames []string) {
	var seenUsers = make(map[string]*ent.User, len(foundUsers))
	for _, u := range foundUsers {
		seenUsers[u.Nickname] = u
	}

	var unseenUserNicknames []string
	for _, name := range requestedNicknames {
		if _, exists := seenUsers[name]; !exists {
			unseenUserNicknames = append(unseenUserNicknames, name)
		}
	}
	return unseenUserNicknames
}
