package dto

import (
	"testing"
	"time"

	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
	"github.com/stretchr/testify/require"
)

func TestCreateMeetingRequest(t *testing.T) {
	var now = time.Now()
	var (
		emptyDesc    = ""
		nonEmptyDesc = "foo"
	)

	type testSet struct {
		name          string
		input         CreateMeetingRequest
		expectError   bool
		expectedErrno pkgerr.Errno
	}

	var testSets = []testSet{
		{
			name: "normal-case",
			input: CreateMeetingRequest{
				OrganizerNickname: "redshore",
				Title:             "I'm your father!!",
				StartsAt:          now,
				EndsAt:            now.Add(time.Hour),
				Description:       &nonEmptyDesc,
			},
		},
		{
			name: "empty-description",
			input: CreateMeetingRequest{
				OrganizerNickname: "redshore",
				Title:             "I'm your father!!",
				StartsAt:          now,
				EndsAt:            now.Add(time.Hour),
				Description:       &emptyDesc,
			},
			expectError:   true,
			expectedErrno: pkgerr.ErrnoEmptyField,
		},
		{
			name: "without-meeting-time",
			input: CreateMeetingRequest{
				OrganizerNickname: "redshore",
				Title:             "I'm your father!!",
				Description:       &nonEmptyDesc,
			},
			expectError:   true,
			expectedErrno: pkgerr.ErrnoInvalidTime,
		},
		{
			name: "starts-not-before-ends",
			input: CreateMeetingRequest{
				OrganizerNickname: "redshore",
				Title:             "I'm your father!!",
				StartsAt:          now,
				EndsAt:            now,
				Description:       &nonEmptyDesc,
			},
			expectError:   true,
			expectedErrno: pkgerr.ErrnoInvalidTime,
		},
	}

	for _, ts := range testSets {
		t.Run(ts.name, func(t *testing.T) {
			err := ts.input.Validate()
			if ts.expectError {
				require.Error(t, err)

				kerr := pkgerr.AsKnown(err)
				require.NotNil(t, kerr, "error must be pkgerr.Known")
				require.Equal(t, ts.expectedErrno, kerr.Errno)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddParticipantsRequest(t *testing.T) {
	type testSet struct {
		name          string
		input         AddParticipantsRequest
		expectError   bool
		expectedErrno pkgerr.Errno
	}

	var testSets = []testSet{
		{
			name: "normal-case",
			input: AddParticipantsRequest{
				MeetingID: 1,
				ParticipantNicknames: []string{
					"foo",
					"bar",
				},
			},
		},
		{
			name: "duplicate-participants",
			input: AddParticipantsRequest{
				MeetingID: 1,
				ParticipantNicknames: []string{
					"foo",
					"bar",
					"foo",
				},
			},
			expectError:   true,
			expectedErrno: pkgerr.ErrnoDuplicateValue,
		},
	}

	for _, ts := range testSets {
		t.Run(ts.name, func(t *testing.T) {
			err := ts.input.Validate()
			if ts.expectError {
				require.Error(t, err)

				kerr := pkgerr.AsKnown(err)
				require.NotNil(t, kerr, "error must be pkgerr.Known")
				require.Equal(t, ts.expectedErrno, kerr.Errno)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
