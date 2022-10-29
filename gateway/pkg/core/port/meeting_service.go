package port

import (
	"context"

	"github.com/isutare412/goarch/gateway/pkg/core/dto"
)

type MeetingService interface {
	CreateMeeting(context.Context, dto.CreateMeetingRequest) (dto.CreateMeetingResponse, error)
	AddParticipants(context.Context, dto.AddParticipantsRequest) error
}
