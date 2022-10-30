package port

import (
	"context"

	"github.com/isutare412/goarch/gateway/ent"
)

type MeetingRepository interface {
	Save(ctx context.Context, mtg *ent.Meeting, orgID int) (*ent.Meeting, error)
	ExistsByID(ctx context.Context, mtgID int) (bool, error)

	AddParticipants(ctx context.Context, mtgID int, ptcIDs []int) error
	RemoveParticipants(ctx context.Context, mtgID int, ptcIDs []int) error
	FindParticipantsByID(ctx context.Context, mtgID int) ([]*ent.User, error)
	FindParticipantsByIDAndParticipantIDs(ctx context.Context, mtgID int, ptcIDs []int) ([]*ent.User, error)
}
