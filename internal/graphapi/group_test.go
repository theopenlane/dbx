package graphapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/theopenlane/dbx/internal/ent/generated"
	"github.com/theopenlane/dbx/pkg/dbxclient"
)

func (suite *GraphTestSuite) TestQueryGroup() {
	t := suite.T()

	group := (&GroupBuilder{client: suite.client}).MustNew(context.Background(), t)

	testCases := []struct {
		name     string
		query    string
		expected *ent.Group
		errorMsg string
	}{
		{
			name:     "happy path group",
			query:    group.Name,
			expected: group,
		},
		{
			name:     "group not found",
			query:    "notfound",
			expected: nil,
			errorMsg: "group not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := suite.client.dbx.GetGroup(context.Background(), tc.query)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Group)
		})
	}

	(&GroupCleanup{client: suite.client, GroupID: group.ID}).MustDelete(context.Background(), t)
}

func (suite *GraphTestSuite) TestListGroups() {
	t := suite.T()

	group1 := (&GroupBuilder{client: suite.client}).MustNew(context.Background(), t)
	group2 := (&GroupBuilder{client: suite.client}).MustNew(context.Background(), t)

	t.Run("List Groups", func(t *testing.T) {
		resp, err := suite.client.dbx.GetAllGroups(context.Background())

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Groups)

		assert.Len(t, resp.Groups.Edges, 2)

		group1Found := false
		group2Found := false

		for _, g := range resp.Groups.Edges {
			if g.Node.Name == group1.Name {
				group1Found = true
			} else if g.Node.Name == group2.Name {
				group2Found = true
			}
		}

		assert.True(t, group1Found)
		assert.True(t, group2Found)
	})

	(&GroupCleanup{client: suite.client, GroupID: group1.ID}).MustDelete(context.Background(), t)
	(&GroupCleanup{client: suite.client, GroupID: group2.ID}).MustDelete(context.Background(), t)
}

func (suite *GraphTestSuite) TestCreateGroup() {
	t := suite.T()

	groupIDs := []string{}

	testCases := []struct {
		name      string
		groupName string
		loc       string
		errorMsg  string
	}{
		{
			name:      "happy path group",
			groupName: "indiana-jones",
			loc:       "den",
		},
		{
			name:      "group already exists",
			groupName: "indiana-jones",
			loc:       "den",
			errorMsg:  "constraint failed",
		},
		{
			name:      "empty group name",
			groupName: "",
			loc:       "den",
			errorMsg:  "invalid or unparsable field: name",
		},
		{
			name:      "empty location",
			groupName: "lost-ark",
			loc:       "",
			errorMsg:  "invalid or unparsable field: primary_location",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			g := dbxclient.CreateGroupInput{
				Name:            tc.groupName,
				PrimaryLocation: tc.loc,
			}

			resp, err := suite.client.dbx.CreateGroup(context.Background(), g)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateGroup)

			assert.Equal(t, tc.groupName, resp.CreateGroup.Group.Name)

			groupIDs = append(groupIDs, resp.CreateGroup.Group.ID)
		})
	}

	// Cleanup groups
	for _, id := range groupIDs {
		(&GroupCleanup{client: suite.client, GroupID: id}).MustDelete(context.Background(), t)
	}
}

func (suite *GraphTestSuite) TestDeleteGroup() {
	t := suite.T()

	group := (&GroupBuilder{client: suite.client}).MustNew(context.Background(), t)

	testCases := []struct {
		name      string
		groupName string
		errorMsg  string
	}{
		{
			name:      "happy path group",
			groupName: group.Name,
		},
		{
			name:      "group does not exist",
			groupName: "raiders",
			errorMsg:  "group not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			resp, err := suite.client.dbx.DeleteGroup(context.Background(), tc.groupName)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.DeleteGroup)

			assert.NotEmpty(t, resp.DeleteGroup.DeletedID)
		})
	}
}
