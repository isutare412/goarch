package dto

import (
	"fmt"
	"time"

	"github.com/isutare412/goarch/gateway/ent"
	"github.com/isutare412/goarch/gateway/pkg/core/enum"
	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
)

type CreateMeetingRequest struct {
	OrganizerNickname string
	Title             string
	StartsAt          time.Time
	EndsAt            time.Time
	Description       *string
}

func (r CreateMeetingRequest) Validate() error {
	if r.OrganizerNickname == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("organizerNickname field is mandatory"),
		}
	}

	if r.Title == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("title field is mandatory"),
		}
	}

	if r.StartsAt.Before(enum.MinMeetingTime) || r.EndsAt.Before(enum.MinMeetingTime) {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoInvalidTime,
			Simple: fmt.Errorf("start/end time should be later than %v", enum.MinMeetingTime.Format(time.RFC3339)),
		}
	}
	if !r.StartsAt.Before(r.EndsAt) {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoInvalidTime,
			Simple: fmt.Errorf("startsAt should be earlier than endsAt"),
		}
	}

	if r.Description != nil && *r.Description == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("description field should not be empty"),
		}
	}
	return nil
}

type CreateMeetingResponse struct {
	MeetingID int
}

func (r CreateMeetingRequest) IntoMeeting() *ent.Meeting {
	return &ent.Meeting{
		Title:       r.Title,
		StartsAt:    r.StartsAt,
		EndsAt:      r.EndsAt,
		Description: r.Description,
	}
}

type AddParticipantsRequest struct {
	MeetingID            int
	ParticipantNicknames []string
}

func (r AddParticipantsRequest) Validate() error {
	if r.MeetingID == 0 {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("meetingId field is mandatory"),
		}
	}

	if len(r.ParticipantNicknames) == 0 {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("participants should be at least one"),
		}
	}

	seenParticipants := make(map[string]struct{}, len(r.ParticipantNicknames))
	for _, name := range r.ParticipantNicknames {
		if found, exists := seenParticipants[name]; exists {
			return pkgerr.Known{
				Errno:  pkgerr.ErrnoDuplicateValue,
				Simple: fmt.Errorf("participant(%s) exists more than once in request", found),
			}
		}
		seenParticipants[name] = struct{}{}
	}
	return nil
}
