package postgres

import (
	"context"

	"github.com/isutare412/goarch/gateway/ent"
	"github.com/isutare412/goarch/gateway/ent/meeting"
	"github.com/isutare412/goarch/gateway/ent/user"
	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
)

type MeetingRepository struct {
	client *Client
}

func NewMeetingRepository(client *Client) *MeetingRepository {
	return &MeetingRepository{client: client}
}

func (r *MeetingRepository) Save(ctx context.Context, mtg *ent.Meeting, orgID int) (*ent.Meeting, error) {
	mtg, err := r.txClient(ctx).Meeting.
		Create().
		SetTitle(mtg.Title).
		SetStartsAt(mtg.StartsAt).
		SetEndsAt(mtg.EndsAt).
		SetNillableDescription(mtg.Description).
		SetOrganizerID(orgID).
		Save(ctx)
	if err != nil {
		return nil, pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return mtg, nil
}

func (r *MeetingRepository) ExistsByID(ctx context.Context, mtgID int) (bool, error) {
	exists, err := r.txClient(ctx).Meeting.
		Query().
		Where(meeting.ID(mtgID)).
		Exist(ctx)
	if err != nil {
		return false, pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return exists, nil
}

func (r *MeetingRepository) AddParticipants(ctx context.Context, mtgID int, ptcIDs []int) error {
	err := r.txClient(ctx).Meeting.
		UpdateOneID(mtgID).
		AddParticipantIDs(ptcIDs...).
		Exec(ctx)
	if err != nil {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return nil
}

func (r *MeetingRepository) RemoveParticipants(ctx context.Context, mtgID int, ptcIDs []int) error {
	err := r.txClient(ctx).Meeting.
		UpdateOneID(mtgID).
		RemoveParticipantIDs(ptcIDs...).
		Exec(ctx)
	if err != nil {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return nil
}

func (r *MeetingRepository) FindParticipantsByID(ctx context.Context, mtgID int) ([]*ent.User, error) {
	usrs, err := r.txClient(ctx).Meeting.
		Query().
		Where(meeting.ID(mtgID)).
		QueryParticipants().
		All(ctx)
	if err != nil {
		return nil, pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return usrs, nil
}

func (r *MeetingRepository) FindParticipantsByIDAndParticipantIDs(
	ctx context.Context,
	mtgID int,
	ptcIDs []int,
) ([]*ent.User, error) {
	usrs, err := r.txClient(ctx).Meeting.
		Query().
		Where(meeting.ID(mtgID)).
		QueryParticipants().
		Where(user.IDIn(ptcIDs...)).
		All(ctx)
	if err != nil {
		return nil, pkgerr.Known{
			Errno:  pkgerr.ErrnoRepository,
			Origin: err,
		}
	}
	return usrs, nil
}

func (r *MeetingRepository) txClient(ctx context.Context) *ent.Client {
	if c := txFromCtx(ctx); c != nil {
		return c
	}
	return r.client.cli
}
