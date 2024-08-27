package graphapi_test

import (
	"context"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/theopenlane/dbx/internal/ent/generated"
	"github.com/theopenlane/dbx/pkg/dbxclient"
	"github.com/theopenlane/dbx/pkg/enums"
)

func (suite *GraphTestSuite) TestQueryDatabase() {
	t := suite.T()

	db := (&DatabaseBuilder{client: suite.client}).MustNew(context.Background(), t)

	testCases := []struct {
		name     string
		query    string
		expected *ent.Database
		errorMsg string
	}{
		{
			name:     "happy path database",
			query:    db.Name,
			expected: db,
		},
		{
			name:     "database not found",
			query:    "notfound",
			expected: nil,
			errorMsg: "database not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := suite.client.dbx.GetDatabase(context.Background(), tc.query)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Database)
		})
	}

	(&DatabaseCleanup{client: suite.client, DatabaseID: db.ID}).MustDelete(context.Background(), t)
	(&GroupCleanup{client: suite.client, GroupID: db.GroupID}).MustDelete(context.Background(), t)
}

func (suite *GraphTestSuite) TestListDatabases() {
	t := suite.T()

	db1 := (&DatabaseBuilder{client: suite.client}).MustNew(context.Background(), t)
	db2 := (&DatabaseBuilder{client: suite.client}).MustNew(context.Background(), t)

	t.Run("List Databases", func(t *testing.T) {
		resp, err := suite.client.dbx.GetAllDatabases(context.Background())

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Databases)
		require.Len(t, resp.Databases.Edges, 2)
	})

	(&DatabaseCleanup{client: suite.client, DatabaseID: db1.ID}).MustDelete(context.Background(), t)
	(&DatabaseCleanup{client: suite.client, DatabaseID: db2.ID}).MustDelete(context.Background(), t)
	(&GroupCleanup{client: suite.client, GroupID: db1.GroupID}).MustDelete(context.Background(), t)
	(&GroupCleanup{client: suite.client, GroupID: db2.GroupID}).MustDelete(context.Background(), t)
}

func (suite *GraphTestSuite) TestCreateDatabase() {
	t := suite.T()

	group := (&GroupBuilder{client: suite.client}).MustNew(context.Background(), t)

	testCases := []struct {
		name     string
		orgID    string
		groupID  string
		region   enums.Region
		provider *enums.DatabaseProvider
		errorMsg string
	}{
		{
			name:     "happy path, turso database",
			orgID:    "01HSCAGDJ1XZ12Y06FESH4VEC1",
			groupID:  group.ID,
			provider: &enums.Turso,
		},
		{
			name:     "happy path, turso database with region",
			orgID:    "01HSCAGDJ1XZ12Y06FESH4VEC1",
			region:   enums.Amer,
			provider: &enums.Turso,
		},
		{
			name:     "happy path, local database",
			orgID:    "01HSCAGDJ1XZ12Y06FESH4VEC2",
			groupID:  group.ID,
			provider: &enums.Local,
		},
		{
			name:     "missing group",
			orgID:    "01HSCAGDJ1XZ12Y06FESH4VEC3",
			groupID:  "notfound",
			provider: &enums.Turso,
			errorMsg: "group not found",
		},
		{
			name:     "missing org id",
			orgID:    "",
			groupID:  group.ID,
			provider: &enums.Turso,
			errorMsg: "invalid or unparsable field: organization_id",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			g := dbxclient.CreateDatabaseInput{
				OrganizationID: tc.orgID,
				Provider:       tc.provider,
				GroupID:        tc.groupID,
			}

			if tc.region != "" {
				g.Geo = lo.ToPtr(tc.region.String())
			}

			resp, err := suite.client.dbx.CreateDatabase(context.Background(), g)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateDatabase)

			assert.Contains(t, resp.CreateDatabase.Database.Name, strings.ToLower(tc.orgID))
			assert.Equal(t, *tc.provider, resp.CreateDatabase.Database.Provider)
			assert.Equal(t, tc.orgID, resp.CreateDatabase.Database.OrganizationID)

			(&DatabaseCleanup{client: suite.client, DatabaseID: resp.CreateDatabase.Database.ID}).MustDelete(context.Background(), t)
		})
	}

	(&GroupCleanup{client: suite.client, GroupID: group.ID}).MustDelete(context.Background(), t)
}

func (suite *GraphTestSuite) TestDeleteDatabase() {
	t := suite.T()

	db := (&DatabaseBuilder{client: suite.client}).MustNew(context.Background(), t)

	testCases := []struct {
		name     string
		dbName   string
		errorMsg string
	}{
		{
			name:   "happy path database",
			dbName: db.Name,
		},
		{
			name:     "db does not exist",
			dbName:   "lost-ark",
			errorMsg: "database not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			resp, err := suite.client.dbx.DeleteDatabase(context.Background(), tc.dbName)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.DeleteDatabase)

			assert.NotEmpty(t, resp.DeleteDatabase.DeletedID)
		})
	}

	(&GroupCleanup{client: suite.client, GroupID: db.GroupID}).MustDelete(context.Background(), t)
}
