// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/dbx/internal/ent/generated/database"
	"github.com/theopenlane/dbx/internal/ent/generated/group"
	"github.com/theopenlane/dbx/pkg/enums"
)

// DatabaseCreate is the builder for creating a Database entity.
type DatabaseCreate struct {
	config
	mutation *DatabaseMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DatabaseCreate) SetCreatedAt(t time.Time) *DatabaseCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableCreatedAt(t *time.Time) *DatabaseCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DatabaseCreate) SetUpdatedAt(t time.Time) *DatabaseCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableUpdatedAt(t *time.Time) *DatabaseCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetCreatedBy sets the "created_by" field.
func (dc *DatabaseCreate) SetCreatedBy(s string) *DatabaseCreate {
	dc.mutation.SetCreatedBy(s)
	return dc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableCreatedBy(s *string) *DatabaseCreate {
	if s != nil {
		dc.SetCreatedBy(*s)
	}
	return dc
}

// SetUpdatedBy sets the "updated_by" field.
func (dc *DatabaseCreate) SetUpdatedBy(s string) *DatabaseCreate {
	dc.mutation.SetUpdatedBy(s)
	return dc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableUpdatedBy(s *string) *DatabaseCreate {
	if s != nil {
		dc.SetUpdatedBy(*s)
	}
	return dc
}

// SetDeletedAt sets the "deleted_at" field.
func (dc *DatabaseCreate) SetDeletedAt(t time.Time) *DatabaseCreate {
	dc.mutation.SetDeletedAt(t)
	return dc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableDeletedAt(t *time.Time) *DatabaseCreate {
	if t != nil {
		dc.SetDeletedAt(*t)
	}
	return dc
}

// SetDeletedBy sets the "deleted_by" field.
func (dc *DatabaseCreate) SetDeletedBy(s string) *DatabaseCreate {
	dc.mutation.SetDeletedBy(s)
	return dc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableDeletedBy(s *string) *DatabaseCreate {
	if s != nil {
		dc.SetDeletedBy(*s)
	}
	return dc
}

// SetOrganizationID sets the "organization_id" field.
func (dc *DatabaseCreate) SetOrganizationID(s string) *DatabaseCreate {
	dc.mutation.SetOrganizationID(s)
	return dc
}

// SetName sets the "name" field.
func (dc *DatabaseCreate) SetName(s string) *DatabaseCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetGeo sets the "geo" field.
func (dc *DatabaseCreate) SetGeo(s string) *DatabaseCreate {
	dc.mutation.SetGeo(s)
	return dc
}

// SetNillableGeo sets the "geo" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableGeo(s *string) *DatabaseCreate {
	if s != nil {
		dc.SetGeo(*s)
	}
	return dc
}

// SetDsn sets the "dsn" field.
func (dc *DatabaseCreate) SetDsn(s string) *DatabaseCreate {
	dc.mutation.SetDsn(s)
	return dc
}

// SetGroupID sets the "group_id" field.
func (dc *DatabaseCreate) SetGroupID(s string) *DatabaseCreate {
	dc.mutation.SetGroupID(s)
	return dc
}

// SetToken sets the "token" field.
func (dc *DatabaseCreate) SetToken(s string) *DatabaseCreate {
	dc.mutation.SetToken(s)
	return dc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableToken(s *string) *DatabaseCreate {
	if s != nil {
		dc.SetToken(*s)
	}
	return dc
}

// SetStatus sets the "status" field.
func (dc *DatabaseCreate) SetStatus(es enums.DatabaseStatus) *DatabaseCreate {
	dc.mutation.SetStatus(es)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableStatus(es *enums.DatabaseStatus) *DatabaseCreate {
	if es != nil {
		dc.SetStatus(*es)
	}
	return dc
}

// SetProvider sets the "provider" field.
func (dc *DatabaseCreate) SetProvider(ep enums.DatabaseProvider) *DatabaseCreate {
	dc.mutation.SetProvider(ep)
	return dc
}

// SetNillableProvider sets the "provider" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableProvider(ep *enums.DatabaseProvider) *DatabaseCreate {
	if ep != nil {
		dc.SetProvider(*ep)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DatabaseCreate) SetID(s string) *DatabaseCreate {
	dc.mutation.SetID(s)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DatabaseCreate) SetNillableID(s *string) *DatabaseCreate {
	if s != nil {
		dc.SetID(*s)
	}
	return dc
}

// SetGroup sets the "group" edge to the Group entity.
func (dc *DatabaseCreate) SetGroup(g *Group) *DatabaseCreate {
	return dc.SetGroupID(g.ID)
}

// Mutation returns the DatabaseMutation object of the builder.
func (dc *DatabaseCreate) Mutation() *DatabaseMutation {
	return dc.mutation
}

// Save creates the Database in the database.
func (dc *DatabaseCreate) Save(ctx context.Context) (*Database, error) {
	if err := dc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DatabaseCreate) SaveX(ctx context.Context) *Database {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DatabaseCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DatabaseCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DatabaseCreate) defaults() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		if database.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized database.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := database.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		if database.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized database.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := database.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.Status(); !ok {
		v := database.DefaultStatus
		dc.mutation.SetStatus(v)
	}
	if _, ok := dc.mutation.Provider(); !ok {
		v := database.DefaultProvider
		dc.mutation.SetProvider(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		if database.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized database.DefaultID (forgotten import generated/runtime?)")
		}
		v := database.DefaultID()
		dc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dc *DatabaseCreate) check() error {
	if _, ok := dc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`generated: missing required field "Database.organization_id"`)}
	}
	if v, ok := dc.mutation.OrganizationID(); ok {
		if err := database.OrganizationIDValidator(v); err != nil {
			return &ValidationError{Name: "organization_id", err: fmt.Errorf(`generated: validator failed for field "Database.organization_id": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Database.name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := database.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Database.name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Dsn(); !ok {
		return &ValidationError{Name: "dsn", err: errors.New(`generated: missing required field "Database.dsn"`)}
	}
	if v, ok := dc.mutation.Dsn(); ok {
		if err := database.DsnValidator(v); err != nil {
			return &ValidationError{Name: "dsn", err: fmt.Errorf(`generated: validator failed for field "Database.dsn": %w`, err)}
		}
	}
	if _, ok := dc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`generated: missing required field "Database.group_id"`)}
	}
	if _, ok := dc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`generated: missing required field "Database.status"`)}
	}
	if v, ok := dc.mutation.Status(); ok {
		if err := database.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Database.status": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`generated: missing required field "Database.provider"`)}
	}
	if v, ok := dc.mutation.Provider(); ok {
		if err := database.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`generated: validator failed for field "Database.provider": %w`, err)}
		}
	}
	if len(dc.mutation.GroupIDs()) == 0 {
		return &ValidationError{Name: "group", err: errors.New(`generated: missing required edge "Database.group"`)}
	}
	return nil
}

func (dc *DatabaseCreate) sqlSave(ctx context.Context) (*Database, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Database.ID type: %T", _spec.ID.Value)
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DatabaseCreate) createSpec() (*Database, *sqlgraph.CreateSpec) {
	var (
		_node = &Database{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(database.Table, sqlgraph.NewFieldSpec(database.FieldID, field.TypeString))
	)
	_spec.Schema = dc.schemaConfig.Database
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(database.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(database.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.CreatedBy(); ok {
		_spec.SetField(database.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := dc.mutation.UpdatedBy(); ok {
		_spec.SetField(database.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := dc.mutation.DeletedAt(); ok {
		_spec.SetField(database.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := dc.mutation.DeletedBy(); ok {
		_spec.SetField(database.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := dc.mutation.OrganizationID(); ok {
		_spec.SetField(database.FieldOrganizationID, field.TypeString, value)
		_node.OrganizationID = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(database.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Geo(); ok {
		_spec.SetField(database.FieldGeo, field.TypeString, value)
		_node.Geo = value
	}
	if value, ok := dc.mutation.Dsn(); ok {
		_spec.SetField(database.FieldDsn, field.TypeString, value)
		_node.Dsn = value
	}
	if value, ok := dc.mutation.Token(); ok {
		_spec.SetField(database.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(database.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := dc.mutation.Provider(); ok {
		_spec.SetField(database.FieldProvider, field.TypeEnum, value)
		_node.Provider = value
	}
	if nodes := dc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   database.GroupTable,
			Columns: []string{database.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = dc.schemaConfig.Database
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DatabaseCreateBulk is the builder for creating many Database entities in bulk.
type DatabaseCreateBulk struct {
	config
	err      error
	builders []*DatabaseCreate
}

// Save creates the Database entities in the database.
func (dcb *DatabaseCreateBulk) Save(ctx context.Context) ([]*Database, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Database, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DatabaseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DatabaseCreateBulk) SaveX(ctx context.Context) []*Database {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DatabaseCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DatabaseCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
