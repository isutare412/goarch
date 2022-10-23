// Code generated by ent, DO NOT EDIT.

package meeting

import (
	"time"
)

const (
	// Label holds the string label denoting the meeting type in the database.
	Label = "meeting"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldStartsAt holds the string denoting the starts_at field in the database.
	FieldStartsAt = "starts_at"
	// FieldEndsAt holds the string denoting the ends_at field in the database.
	FieldEndsAt = "ends_at"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeOrganizer holds the string denoting the organizer edge name in mutations.
	EdgeOrganizer = "organizer"
	// Table holds the table name of the meeting in the database.
	Table = "meetings"
	// OrganizerTable is the table that holds the organizer relation/edge. The primary key declared below.
	OrganizerTable = "user_meetings"
	// OrganizerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OrganizerInverseTable = "users"
)

// Columns holds all SQL columns for meeting fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldTitle,
	FieldStartsAt,
	FieldEndsAt,
	FieldDescription,
}

var (
	// OrganizerPrimaryKey and OrganizerColumn2 are the table columns denoting the
	// primary key for the organizer relation (M2M).
	OrganizerPrimaryKey = []string{"user_id", "meeting_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
)