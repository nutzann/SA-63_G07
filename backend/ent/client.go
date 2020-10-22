// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/nutzann/app/ent/migrate"

	"github.com/nutzann/app/ent/area"
	"github.com/nutzann/app/ent/disease"
	"github.com/nutzann/app/ent/employee"
	"github.com/nutzann/app/ent/level"
	"github.com/nutzann/app/ent/statistic"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Area is the client for interacting with the Area builders.
	Area *AreaClient
	// Disease is the client for interacting with the Disease builders.
	Disease *DiseaseClient
	// Employee is the client for interacting with the Employee builders.
	Employee *EmployeeClient
	// Level is the client for interacting with the Level builders.
	Level *LevelClient
	// Statistic is the client for interacting with the Statistic builders.
	Statistic *StatisticClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Area = NewAreaClient(c.config)
	c.Disease = NewDiseaseClient(c.config)
	c.Employee = NewEmployeeClient(c.config)
	c.Level = NewLevelClient(c.config)
	c.Statistic = NewStatisticClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Area:      NewAreaClient(cfg),
		Disease:   NewDiseaseClient(cfg),
		Employee:  NewEmployeeClient(cfg),
		Level:     NewLevelClient(cfg),
		Statistic: NewStatisticClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:    cfg,
		Area:      NewAreaClient(cfg),
		Disease:   NewDiseaseClient(cfg),
		Employee:  NewEmployeeClient(cfg),
		Level:     NewLevelClient(cfg),
		Statistic: NewStatisticClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Area.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Area.Use(hooks...)
	c.Disease.Use(hooks...)
	c.Employee.Use(hooks...)
	c.Level.Use(hooks...)
	c.Statistic.Use(hooks...)
}

// AreaClient is a client for the Area schema.
type AreaClient struct {
	config
}

// NewAreaClient returns a client for the Area from the given config.
func NewAreaClient(c config) *AreaClient {
	return &AreaClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `area.Hooks(f(g(h())))`.
func (c *AreaClient) Use(hooks ...Hook) {
	c.hooks.Area = append(c.hooks.Area, hooks...)
}

// Create returns a create builder for Area.
func (c *AreaClient) Create() *AreaCreate {
	mutation := newAreaMutation(c.config, OpCreate)
	return &AreaCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Area.
func (c *AreaClient) Update() *AreaUpdate {
	mutation := newAreaMutation(c.config, OpUpdate)
	return &AreaUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AreaClient) UpdateOne(a *Area) *AreaUpdateOne {
	mutation := newAreaMutation(c.config, OpUpdateOne, withArea(a))
	return &AreaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AreaClient) UpdateOneID(id int) *AreaUpdateOne {
	mutation := newAreaMutation(c.config, OpUpdateOne, withAreaID(id))
	return &AreaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Area.
func (c *AreaClient) Delete() *AreaDelete {
	mutation := newAreaMutation(c.config, OpDelete)
	return &AreaDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AreaClient) DeleteOne(a *Area) *AreaDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AreaClient) DeleteOneID(id int) *AreaDeleteOne {
	builder := c.Delete().Where(area.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AreaDeleteOne{builder}
}

// Create returns a query builder for Area.
func (c *AreaClient) Query() *AreaQuery {
	return &AreaQuery{config: c.config}
}

// Get returns a Area entity by its id.
func (c *AreaClient) Get(ctx context.Context, id int) (*Area, error) {
	return c.Query().Where(area.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AreaClient) GetX(ctx context.Context, id int) *Area {
	a, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return a
}

// QueryDisease queries the disease edge of a Area.
func (c *AreaClient) QueryDisease(a *Area) *DiseaseQuery {
	query := &DiseaseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(area.Table, area.FieldID, id),
			sqlgraph.To(disease.Table, disease.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, area.DiseaseTable, area.DiseaseColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStatistic queries the statistic edge of a Area.
func (c *AreaClient) QueryStatistic(a *Area) *StatisticQuery {
	query := &StatisticQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(area.Table, area.FieldID, id),
			sqlgraph.To(statistic.Table, statistic.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, area.StatisticTable, area.StatisticColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLevel queries the level edge of a Area.
func (c *AreaClient) QueryLevel(a *Area) *LevelQuery {
	query := &LevelQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(area.Table, area.FieldID, id),
			sqlgraph.To(level.Table, level.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, area.LevelTable, area.LevelColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEmployee queries the employee edge of a Area.
func (c *AreaClient) QueryEmployee(a *Area) *EmployeeQuery {
	query := &EmployeeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(area.Table, area.FieldID, id),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, area.EmployeeTable, area.EmployeeColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AreaClient) Hooks() []Hook {
	return c.hooks.Area
}

// DiseaseClient is a client for the Disease schema.
type DiseaseClient struct {
	config
}

// NewDiseaseClient returns a client for the Disease from the given config.
func NewDiseaseClient(c config) *DiseaseClient {
	return &DiseaseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `disease.Hooks(f(g(h())))`.
func (c *DiseaseClient) Use(hooks ...Hook) {
	c.hooks.Disease = append(c.hooks.Disease, hooks...)
}

// Create returns a create builder for Disease.
func (c *DiseaseClient) Create() *DiseaseCreate {
	mutation := newDiseaseMutation(c.config, OpCreate)
	return &DiseaseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Disease.
func (c *DiseaseClient) Update() *DiseaseUpdate {
	mutation := newDiseaseMutation(c.config, OpUpdate)
	return &DiseaseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DiseaseClient) UpdateOne(d *Disease) *DiseaseUpdateOne {
	mutation := newDiseaseMutation(c.config, OpUpdateOne, withDisease(d))
	return &DiseaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DiseaseClient) UpdateOneID(id int) *DiseaseUpdateOne {
	mutation := newDiseaseMutation(c.config, OpUpdateOne, withDiseaseID(id))
	return &DiseaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Disease.
func (c *DiseaseClient) Delete() *DiseaseDelete {
	mutation := newDiseaseMutation(c.config, OpDelete)
	return &DiseaseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DiseaseClient) DeleteOne(d *Disease) *DiseaseDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DiseaseClient) DeleteOneID(id int) *DiseaseDeleteOne {
	builder := c.Delete().Where(disease.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DiseaseDeleteOne{builder}
}

// Create returns a query builder for Disease.
func (c *DiseaseClient) Query() *DiseaseQuery {
	return &DiseaseQuery{config: c.config}
}

// Get returns a Disease entity by its id.
func (c *DiseaseClient) Get(ctx context.Context, id int) (*Disease, error) {
	return c.Query().Where(disease.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DiseaseClient) GetX(ctx context.Context, id int) *Disease {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryArea queries the area edge of a Disease.
func (c *DiseaseClient) QueryArea(d *Disease) *AreaQuery {
	query := &AreaQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(disease.Table, disease.FieldID, id),
			sqlgraph.To(area.Table, area.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, disease.AreaTable, disease.AreaColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DiseaseClient) Hooks() []Hook {
	return c.hooks.Disease
}

// EmployeeClient is a client for the Employee schema.
type EmployeeClient struct {
	config
}

// NewEmployeeClient returns a client for the Employee from the given config.
func NewEmployeeClient(c config) *EmployeeClient {
	return &EmployeeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `employee.Hooks(f(g(h())))`.
func (c *EmployeeClient) Use(hooks ...Hook) {
	c.hooks.Employee = append(c.hooks.Employee, hooks...)
}

// Create returns a create builder for Employee.
func (c *EmployeeClient) Create() *EmployeeCreate {
	mutation := newEmployeeMutation(c.config, OpCreate)
	return &EmployeeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Employee.
func (c *EmployeeClient) Update() *EmployeeUpdate {
	mutation := newEmployeeMutation(c.config, OpUpdate)
	return &EmployeeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmployeeClient) UpdateOne(e *Employee) *EmployeeUpdateOne {
	mutation := newEmployeeMutation(c.config, OpUpdateOne, withEmployee(e))
	return &EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmployeeClient) UpdateOneID(id int) *EmployeeUpdateOne {
	mutation := newEmployeeMutation(c.config, OpUpdateOne, withEmployeeID(id))
	return &EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Employee.
func (c *EmployeeClient) Delete() *EmployeeDelete {
	mutation := newEmployeeMutation(c.config, OpDelete)
	return &EmployeeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EmployeeClient) DeleteOne(e *Employee) *EmployeeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EmployeeClient) DeleteOneID(id int) *EmployeeDeleteOne {
	builder := c.Delete().Where(employee.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmployeeDeleteOne{builder}
}

// Create returns a query builder for Employee.
func (c *EmployeeClient) Query() *EmployeeQuery {
	return &EmployeeQuery{config: c.config}
}

// Get returns a Employee entity by its id.
func (c *EmployeeClient) Get(ctx context.Context, id int) (*Employee, error) {
	return c.Query().Where(employee.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmployeeClient) GetX(ctx context.Context, id int) *Employee {
	e, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return e
}

// QueryArea queries the area edge of a Employee.
func (c *EmployeeClient) QueryArea(e *Employee) *AreaQuery {
	query := &AreaQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, id),
			sqlgraph.To(area.Table, area.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, employee.AreaTable, employee.AreaColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EmployeeClient) Hooks() []Hook {
	return c.hooks.Employee
}

// LevelClient is a client for the Level schema.
type LevelClient struct {
	config
}

// NewLevelClient returns a client for the Level from the given config.
func NewLevelClient(c config) *LevelClient {
	return &LevelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `level.Hooks(f(g(h())))`.
func (c *LevelClient) Use(hooks ...Hook) {
	c.hooks.Level = append(c.hooks.Level, hooks...)
}

// Create returns a create builder for Level.
func (c *LevelClient) Create() *LevelCreate {
	mutation := newLevelMutation(c.config, OpCreate)
	return &LevelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Level.
func (c *LevelClient) Update() *LevelUpdate {
	mutation := newLevelMutation(c.config, OpUpdate)
	return &LevelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LevelClient) UpdateOne(l *Level) *LevelUpdateOne {
	mutation := newLevelMutation(c.config, OpUpdateOne, withLevel(l))
	return &LevelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LevelClient) UpdateOneID(id int) *LevelUpdateOne {
	mutation := newLevelMutation(c.config, OpUpdateOne, withLevelID(id))
	return &LevelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Level.
func (c *LevelClient) Delete() *LevelDelete {
	mutation := newLevelMutation(c.config, OpDelete)
	return &LevelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *LevelClient) DeleteOne(l *Level) *LevelDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *LevelClient) DeleteOneID(id int) *LevelDeleteOne {
	builder := c.Delete().Where(level.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LevelDeleteOne{builder}
}

// Create returns a query builder for Level.
func (c *LevelClient) Query() *LevelQuery {
	return &LevelQuery{config: c.config}
}

// Get returns a Level entity by its id.
func (c *LevelClient) Get(ctx context.Context, id int) (*Level, error) {
	return c.Query().Where(level.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LevelClient) GetX(ctx context.Context, id int) *Level {
	l, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return l
}

// QueryArea queries the area edge of a Level.
func (c *LevelClient) QueryArea(l *Level) *AreaQuery {
	query := &AreaQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(level.Table, level.FieldID, id),
			sqlgraph.To(area.Table, area.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, level.AreaTable, level.AreaColumn),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LevelClient) Hooks() []Hook {
	return c.hooks.Level
}

// StatisticClient is a client for the Statistic schema.
type StatisticClient struct {
	config
}

// NewStatisticClient returns a client for the Statistic from the given config.
func NewStatisticClient(c config) *StatisticClient {
	return &StatisticClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `statistic.Hooks(f(g(h())))`.
func (c *StatisticClient) Use(hooks ...Hook) {
	c.hooks.Statistic = append(c.hooks.Statistic, hooks...)
}

// Create returns a create builder for Statistic.
func (c *StatisticClient) Create() *StatisticCreate {
	mutation := newStatisticMutation(c.config, OpCreate)
	return &StatisticCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Statistic.
func (c *StatisticClient) Update() *StatisticUpdate {
	mutation := newStatisticMutation(c.config, OpUpdate)
	return &StatisticUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatisticClient) UpdateOne(s *Statistic) *StatisticUpdateOne {
	mutation := newStatisticMutation(c.config, OpUpdateOne, withStatistic(s))
	return &StatisticUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatisticClient) UpdateOneID(id int) *StatisticUpdateOne {
	mutation := newStatisticMutation(c.config, OpUpdateOne, withStatisticID(id))
	return &StatisticUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Statistic.
func (c *StatisticClient) Delete() *StatisticDelete {
	mutation := newStatisticMutation(c.config, OpDelete)
	return &StatisticDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StatisticClient) DeleteOne(s *Statistic) *StatisticDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StatisticClient) DeleteOneID(id int) *StatisticDeleteOne {
	builder := c.Delete().Where(statistic.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatisticDeleteOne{builder}
}

// Create returns a query builder for Statistic.
func (c *StatisticClient) Query() *StatisticQuery {
	return &StatisticQuery{config: c.config}
}

// Get returns a Statistic entity by its id.
func (c *StatisticClient) Get(ctx context.Context, id int) (*Statistic, error) {
	return c.Query().Where(statistic.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatisticClient) GetX(ctx context.Context, id int) *Statistic {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QueryArea queries the area edge of a Statistic.
func (c *StatisticClient) QueryArea(s *Statistic) *AreaQuery {
	query := &AreaQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(statistic.Table, statistic.FieldID, id),
			sqlgraph.To(area.Table, area.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, statistic.AreaTable, statistic.AreaColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StatisticClient) Hooks() []Hook {
	return c.hooks.Statistic
}
