package meeting

import (
	"testing"

	"github.com/isutare412/goarch/gateway/ent"
	"github.com/stretchr/testify/assert"
)

func TestSearchUserNotFoundNicknames(t *testing.T) {
	type testSetInput struct {
		users     []*ent.User
		nicknames []string
	}

	type testSet struct {
		name     string
		input    testSetInput
		expected []string
	}

	var testSets = []testSet{
		{
			name: "normal-case",
			input: testSetInput{
				users: []*ent.User{
					{ID: 1, Nickname: "one"},
				},
				nicknames: []string{
					"one",
				},
			},
			expected: nil,
		},
		{
			name: "more-users-than-nicknames",
			input: testSetInput{
				users: []*ent.User{
					{ID: 1, Nickname: "one"},
					{ID: 2, Nickname: "two"},
				},
				nicknames: []string{
					"one",
				},
			},
			expected: nil,
		},
		{
			name: "less-users-than-nicknames",
			input: testSetInput{
				users: []*ent.User{
					{ID: 1, Nickname: "one"},
				},
				nicknames: []string{
					"one",
					"two",
				},
			},
			expected: []string{
				"two",
			},
		},
		{
			name: "nil-inputs",
			input: testSetInput{
				users:     nil,
				nicknames: nil,
			},
			expected: nil,
		},
	}

	for _, ts := range testSets {
		t.Run(ts.name, func(t *testing.T) {
			got := searchUserNotFoundNicknames(ts.input.users, ts.input.nicknames)
			assert.Equal(t, ts.expected, got)
		})
	}
}
