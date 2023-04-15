// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/isutare412/goarch/gateway/pkg/core/ent/admin"
	"github.com/isutare412/goarch/gateway/pkg/core/ent/user"
)

// AdminCreate is the builder for creating a Admin entity.
type AdminCreate struct {
	config
	mutation *AdminMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ac *AdminCreate) SetCreateTime(t time.Time) *AdminCreate {
	ac.mutation.SetCreateTime(t)
	return ac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ac *AdminCreate) SetNillableCreateTime(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetCreateTime(*t)
	}
	return ac
}

// SetUpdateTime sets the "update_time" field.
func (ac *AdminCreate) SetUpdateTime(t time.Time) *AdminCreate {
	ac.mutation.SetUpdateTime(t)
	return ac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ac *AdminCreate) SetNillableUpdateTime(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetUpdateTime(*t)
	}
	return ac
}

// SetPhoneNumber sets the "phone_number" field.
func (ac *AdminCreate) SetPhoneNumber(s string) *AdminCreate {
	ac.mutation.SetPhoneNumber(s)
	return ac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *AdminCreate) SetUserID(id int) *AdminCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AdminCreate) SetUser(u *User) *AdminCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AdminMutation object of the builder.
func (ac *AdminCreate) Mutation() *AdminMutation {
	return ac.mutation
}

// Save creates the Admin in the database.
func (ac *AdminCreate) Save(ctx context.Context) (*Admin, error) {
	var (
		err  error
		node *Admin
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Admin)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AdminMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AdminCreate) SaveX(ctx context.Context) *Admin {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AdminCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AdminCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AdminCreate) defaults() {
	if _, ok := ac.mutation.CreateTime(); !ok {
		v := admin.DefaultCreateTime()
		ac.mutation.SetCreateTime(v)
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		v := admin.DefaultUpdateTime()
		ac.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AdminCreate) check() error {
	if _, ok := ac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Admin.create_time"`)}
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Admin.update_time"`)}
	}
	if _, ok := ac.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New(`ent: missing required field "Admin.phone_number"`)}
	}
	if v, ok := ac.mutation.PhoneNumber(); ok {
		if err := admin.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "Admin.phone_number": %w`, err)}
		}
	}
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Admin.user"`)}
	}
	return nil
}

func (ac *AdminCreate) sqlSave(ctx context.Context) (*Admin, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AdminCreate) createSpec() (*Admin, *sqlgraph.CreateSpec) {
	var (
		_node = &Admin{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: admin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admin.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.CreateTime(); ok {
		_spec.SetField(admin.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := ac.mutation.UpdateTime(); ok {
		_spec.SetField(admin.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := ac.mutation.PhoneNumber(); ok {
		_spec.SetField(admin.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   admin.UserTable,
			Columns: []string{admin.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_admin = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AdminCreateBulk is the builder for creating many Admin entities in bulk.
type AdminCreateBulk struct {
	config
	builders []*AdminCreate
}

// Save creates the Admin entities in the database.
func (acb *AdminCreateBulk) Save(ctx context.Context) ([]*Admin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Admin, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AdminCreateBulk) SaveX(ctx context.Context) []*Admin {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AdminCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AdminCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}