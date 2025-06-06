package graphapi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/theopenlane/utils/ulids"

	ent "github.com/theopenlane/dbx/internal/ent/generated"
	"github.com/theopenlane/dbx/pkg/enums"
)

type GroupBuilder struct {
	client *client

	// Fields
	Name     string
	Location string
	Region   enums.Region
}

type GroupCleanup struct {
	client *client

	// Fields
	GroupID string
}

type DatabaseBuilder struct {
	client *client

	// Fields
	Name    string
	OrgID   string
	DSN     string
	GroupID string
}

type DatabaseCleanup struct {
	client *client

	// Fields
	DatabaseID string
}

// MustNew group builder is used to create groups in the database
func (g *GroupBuilder) MustNew(ctx context.Context, _ *testing.T) *ent.Group {
	if g.Name == "" {
		g.Name = gofakeit.AppName()
	}

	if g.Location == "" {
		g.Location = "den"
	}

	if g.Region == "" {
		g.Region = enums.Amer
	}

	group := g.client.db.Group.Create().
		SetName(g.Name).
		SetPrimaryLocation(g.Location).
		SetRegion(g.Region).
		SaveX(ctx)

	return group
}

// MustDelete is used to cleanup groups in the database
func (g *GroupCleanup) MustDelete(ctx context.Context, _ *testing.T) {
	g.client.db.Group.DeleteOneID(g.GroupID).ExecX(ctx)
}

// MustNew group builder is used to create databases in the database
func (d *DatabaseBuilder) MustNew(ctx context.Context, t *testing.T) *ent.Database {
	if d.Name == "" {
		d.Name = gofakeit.AppName()
	}

	if d.OrgID == "" {
		d.OrgID = ulids.New().String()
	}

	if d.DSN == "" {
		d.DSN = fmt.Sprintf("https://%s.turso.com", gofakeit.AppName())
	}

	if d.GroupID == "" {
		group := (&GroupBuilder{client: d.client}).MustNew(ctx, t)
		d.GroupID = group.ID
	}

	db := d.client.db.Database.Create().
		SetName(d.Name).
		SetOrganizationID(d.OrgID).
		SetDsn(d.DSN).
		SetGroupID(d.GroupID).
		SaveX(ctx)

	return db
}

// MustDelete is used to cleanup databases in the database
func (d *DatabaseCleanup) MustDelete(ctx context.Context, _ *testing.T) {
	d.client.db.Database.DeleteOneID(d.DatabaseID).ExecX(ctx)
}
