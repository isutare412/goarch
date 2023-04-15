// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/isutare412/goarch/gateway/pkg/core/ent/meeting"
	"github.com/isutare412/goarch/gateway/pkg/core/ent/user"
)

// Meeting is the model entity for the Meeting schema.
type Meeting struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// StartsAt holds the value of the "starts_at" field.
	StartsAt time.Time `json:"starts_at,omitempty"`
	// EndsAt holds the value of the "ends_at" field.
	EndsAt time.Time `json:"ends_at,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MeetingQuery when eager-loading is set.
	Edges          MeetingEdges `json:"edges"`
	user_organizes *int
}

// MeetingEdges holds the relations/edges for other nodes in the graph.
type MeetingEdges struct {
	// Organizer holds the value of the organizer edge.
	Organizer *User `json:"organizer,omitempty"`
	// Participants holds the value of the participants edge.
	Participants []*User `json:"participants,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrganizerOrErr returns the Organizer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MeetingEdges) OrganizerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Organizer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Organizer, nil
	}
	return nil, &NotLoadedError{edge: "organizer"}
}

// ParticipantsOrErr returns the Participants value or an error if the edge
// was not loaded in eager-loading.
func (e MeetingEdges) ParticipantsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Participants, nil
	}
	return nil, &NotLoadedError{edge: "participants"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Meeting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case meeting.FieldID:
			values[i] = new(sql.NullInt64)
		case meeting.FieldTitle, meeting.FieldDescription:
			values[i] = new(sql.NullString)
		case meeting.FieldCreateTime, meeting.FieldUpdateTime, meeting.FieldStartsAt, meeting.FieldEndsAt:
			values[i] = new(sql.NullTime)
		case meeting.ForeignKeys[0]: // user_organizes
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Meeting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Meeting fields.
func (m *Meeting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case meeting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case meeting.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				m.CreateTime = value.Time
			}
		case meeting.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				m.UpdateTime = value.Time
			}
		case meeting.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				m.Title = value.String
			}
		case meeting.FieldStartsAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field starts_at", values[i])
			} else if value.Valid {
				m.StartsAt = value.Time
			}
		case meeting.FieldEndsAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ends_at", values[i])
			} else if value.Valid {
				m.EndsAt = value.Time
			}
		case meeting.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				m.Description = new(string)
				*m.Description = value.String
			}
		case meeting.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_organizes", value)
			} else if value.Valid {
				m.user_organizes = new(int)
				*m.user_organizes = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOrganizer queries the "organizer" edge of the Meeting entity.
func (m *Meeting) QueryOrganizer() *UserQuery {
	return (&MeetingClient{config: m.config}).QueryOrganizer(m)
}

// QueryParticipants queries the "participants" edge of the Meeting entity.
func (m *Meeting) QueryParticipants() *UserQuery {
	return (&MeetingClient{config: m.config}).QueryParticipants(m)
}

// Update returns a builder for updating this Meeting.
// Note that you need to call Meeting.Unwrap() before calling this method if this Meeting
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Meeting) Update() *MeetingUpdateOne {
	return (&MeetingClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Meeting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Meeting) Unwrap() *Meeting {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Meeting is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Meeting) String() string {
	var builder strings.Builder
	builder.WriteString("Meeting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("create_time=")
	builder.WriteString(m.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(m.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(m.Title)
	builder.WriteString(", ")
	builder.WriteString("starts_at=")
	builder.WriteString(m.StartsAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ends_at=")
	builder.WriteString(m.EndsAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := m.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Meetings is a parsable slice of Meeting.
type Meetings []*Meeting

func (m Meetings) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}