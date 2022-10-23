// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// EdgeAdmin holds the string denoting the admin edge name in mutations.
	EdgeAdmin = "admin"
	// EdgeOrganizes holds the string denoting the organizes edge name in mutations.
	EdgeOrganizes = "organizes"
	// EdgeMeetings holds the string denoting the meetings edge name in mutations.
	EdgeMeetings = "meetings"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AdminTable is the table that holds the admin relation/edge.
	AdminTable = "admins"
	// AdminInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	AdminInverseTable = "admins"
	// AdminColumn is the table column denoting the admin relation/edge.
	AdminColumn = "user_admin"
	// OrganizesTable is the table that holds the organizes relation/edge.
	OrganizesTable = "meetings"
	// OrganizesInverseTable is the table name for the Meeting entity.
	// It exists in this package in order to avoid circular dependency with the "meeting" package.
	OrganizesInverseTable = "meetings"
	// OrganizesColumn is the table column denoting the organizes relation/edge.
	OrganizesColumn = "user_organizes"
	// MeetingsTable is the table that holds the meetings relation/edge. The primary key declared below.
	MeetingsTable = "user_meetings"
	// MeetingsInverseTable is the table name for the Meeting entity.
	// It exists in this package in order to avoid circular dependency with the "meeting" package.
	MeetingsInverseTable = "meetings"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldNickname,
	FieldEmail,
}

var (
	// MeetingsPrimaryKey and MeetingsColumn2 are the table columns denoting the
	// primary key for the meetings relation (M2M).
	MeetingsPrimaryKey = []string{"user_id", "meeting_id"}
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
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)
