// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/isutare412/goarch/gateway/pkg/core/ent/meeting"
	"github.com/isutare412/goarch/gateway/pkg/core/ent/predicate"
	"github.com/isutare412/goarch/gateway/pkg/core/ent/user"
)

// MeetingQuery is the builder for querying Meeting entities.
type MeetingQuery struct {
	config
	limit            *int
	offset           *int
	unique           *bool
	order            []OrderFunc
	fields           []string
	predicates       []predicate.Meeting
	withOrganizer    *UserQuery
	withParticipants *UserQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MeetingQuery builder.
func (mq *MeetingQuery) Where(ps ...predicate.Meeting) *MeetingQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MeetingQuery) Limit(limit int) *MeetingQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MeetingQuery) Offset(offset int) *MeetingQuery {
	mq.offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MeetingQuery) Unique(unique bool) *MeetingQuery {
	mq.unique = &unique
	return mq
}

// Order adds an order step to the query.
func (mq *MeetingQuery) Order(o ...OrderFunc) *MeetingQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryOrganizer chains the current query on the "organizer" edge.
func (mq *MeetingQuery) QueryOrganizer() *UserQuery {
	query := &UserQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(meeting.Table, meeting.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, meeting.OrganizerTable, meeting.OrganizerColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParticipants chains the current query on the "participants" edge.
func (mq *MeetingQuery) QueryParticipants() *UserQuery {
	query := &UserQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(meeting.Table, meeting.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, meeting.ParticipantsTable, meeting.ParticipantsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Meeting entity from the query.
// Returns a *NotFoundError when no Meeting was found.
func (mq *MeetingQuery) First(ctx context.Context) (*Meeting, error) {
	nodes, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{meeting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MeetingQuery) FirstX(ctx context.Context) *Meeting {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Meeting ID from the query.
// Returns a *NotFoundError when no Meeting ID was found.
func (mq *MeetingQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{meeting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MeetingQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Meeting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Meeting entity is found.
// Returns a *NotFoundError when no Meeting entities are found.
func (mq *MeetingQuery) Only(ctx context.Context) (*Meeting, error) {
	nodes, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{meeting.Label}
	default:
		return nil, &NotSingularError{meeting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MeetingQuery) OnlyX(ctx context.Context) *Meeting {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Meeting ID in the query.
// Returns a *NotSingularError when more than one Meeting ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MeetingQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{meeting.Label}
	default:
		err = &NotSingularError{meeting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MeetingQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Meetings.
func (mq *MeetingQuery) All(ctx context.Context) ([]*Meeting, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MeetingQuery) AllX(ctx context.Context) []*Meeting {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Meeting IDs.
func (mq *MeetingQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(meeting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MeetingQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MeetingQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MeetingQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MeetingQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MeetingQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MeetingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MeetingQuery) Clone() *MeetingQuery {
	if mq == nil {
		return nil
	}
	return &MeetingQuery{
		config:           mq.config,
		limit:            mq.limit,
		offset:           mq.offset,
		order:            append([]OrderFunc{}, mq.order...),
		predicates:       append([]predicate.Meeting{}, mq.predicates...),
		withOrganizer:    mq.withOrganizer.Clone(),
		withParticipants: mq.withParticipants.Clone(),
		// clone intermediate query.
		sql:    mq.sql.Clone(),
		path:   mq.path,
		unique: mq.unique,
	}
}

// WithOrganizer tells the query-builder to eager-load the nodes that are connected to
// the "organizer" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MeetingQuery) WithOrganizer(opts ...func(*UserQuery)) *MeetingQuery {
	query := &UserQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withOrganizer = query
	return mq
}

// WithParticipants tells the query-builder to eager-load the nodes that are connected to
// the "participants" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MeetingQuery) WithParticipants(opts ...func(*UserQuery)) *MeetingQuery {
	query := &UserQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withParticipants = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Meeting.Query().
//		GroupBy(meeting.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MeetingQuery) GroupBy(field string, fields ...string) *MeetingGroupBy {
	grbuild := &MeetingGroupBy{config: mq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(ctx), nil
	}
	grbuild.label = meeting.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Meeting.Query().
//		Select(meeting.FieldCreateTime).
//		Scan(ctx, &v)
func (mq *MeetingQuery) Select(fields ...string) *MeetingSelect {
	mq.fields = append(mq.fields, fields...)
	selbuild := &MeetingSelect{MeetingQuery: mq}
	selbuild.label = meeting.Label
	selbuild.flds, selbuild.scan = &mq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a MeetingSelect configured with the given aggregations.
func (mq *MeetingQuery) Aggregate(fns ...AggregateFunc) *MeetingSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MeetingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mq.fields {
		if !meeting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MeetingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Meeting, error) {
	var (
		nodes       = []*Meeting{}
		withFKs     = mq.withFKs
		_spec       = mq.querySpec()
		loadedTypes = [2]bool{
			mq.withOrganizer != nil,
			mq.withParticipants != nil,
		}
	)
	if mq.withOrganizer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, meeting.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Meeting).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Meeting{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withOrganizer; query != nil {
		if err := mq.loadOrganizer(ctx, query, nodes, nil,
			func(n *Meeting, e *User) { n.Edges.Organizer = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withParticipants; query != nil {
		if err := mq.loadParticipants(ctx, query, nodes,
			func(n *Meeting) { n.Edges.Participants = []*User{} },
			func(n *Meeting, e *User) { n.Edges.Participants = append(n.Edges.Participants, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MeetingQuery) loadOrganizer(ctx context.Context, query *UserQuery, nodes []*Meeting, init func(*Meeting), assign func(*Meeting, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Meeting)
	for i := range nodes {
		if nodes[i].user_organizes == nil {
			continue
		}
		fk := *nodes[i].user_organizes
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_organizes" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MeetingQuery) loadParticipants(ctx context.Context, query *UserQuery, nodes []*Meeting, init func(*Meeting), assign func(*Meeting, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Meeting)
	nids := make(map[int]map[*Meeting]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(meeting.ParticipantsTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(meeting.ParticipantsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(meeting.ParticipantsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(meeting.ParticipantsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Meeting]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "participants" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (mq *MeetingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.fields
	if len(mq.fields) > 0 {
		_spec.Unique = mq.unique != nil && *mq.unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MeetingQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (mq *MeetingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   meeting.Table,
			Columns: meeting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: meeting.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if unique := mq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, meeting.FieldID)
		for i := range fields {
			if fields[i] != meeting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MeetingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(meeting.Table)
	columns := mq.fields
	if len(columns) == 0 {
		columns = meeting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.unique != nil && *mq.unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MeetingGroupBy is the group-by builder for Meeting entities.
type MeetingGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MeetingGroupBy) Aggregate(fns ...AggregateFunc) *MeetingGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mgb *MeetingGroupBy) Scan(ctx context.Context, v any) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

func (mgb *MeetingGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range mgb.fields {
		if !meeting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MeetingGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql.Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
		for _, f := range mgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mgb.fields...)...)
}

// MeetingSelect is the builder for selecting fields of Meeting entities.
type MeetingSelect struct {
	*MeetingQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MeetingSelect) Aggregate(fns ...AggregateFunc) *MeetingSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MeetingSelect) Scan(ctx context.Context, v any) error {
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	ms.sql = ms.MeetingQuery.sqlQuery(ctx)
	return ms.sqlScan(ctx, v)
}

func (ms *MeetingSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(ms.sql))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ms.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ms.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ms.sql.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
