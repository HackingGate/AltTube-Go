// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/hackinggate/alttube-go/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hackinggate/alttube-go/ent/accesstoken"
	"github.com/hackinggate/alttube-go/ent/likevideo"
	"github.com/hackinggate/alttube-go/ent/refreshtoken"
	"github.com/hackinggate/alttube-go/ent/user"
	"github.com/hackinggate/alttube-go/ent/video"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AccessToken is the client for interacting with the AccessToken builders.
	AccessToken *AccessTokenClient
	// LikeVideo is the client for interacting with the LikeVideo builders.
	LikeVideo *LikeVideoClient
	// RefreshToken is the client for interacting with the RefreshToken builders.
	RefreshToken *RefreshTokenClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Video is the client for interacting with the Video builders.
	Video *VideoClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AccessToken = NewAccessTokenClient(c.config)
	c.LikeVideo = NewLikeVideoClient(c.config)
	c.RefreshToken = NewRefreshTokenClient(c.config)
	c.User = NewUserClient(c.config)
	c.Video = NewVideoClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		AccessToken:  NewAccessTokenClient(cfg),
		LikeVideo:    NewLikeVideoClient(cfg),
		RefreshToken: NewRefreshTokenClient(cfg),
		User:         NewUserClient(cfg),
		Video:        NewVideoClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		AccessToken:  NewAccessTokenClient(cfg),
		LikeVideo:    NewLikeVideoClient(cfg),
		RefreshToken: NewRefreshTokenClient(cfg),
		User:         NewUserClient(cfg),
		Video:        NewVideoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AccessToken.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.AccessToken.Use(hooks...)
	c.LikeVideo.Use(hooks...)
	c.RefreshToken.Use(hooks...)
	c.User.Use(hooks...)
	c.Video.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.AccessToken.Intercept(interceptors...)
	c.LikeVideo.Intercept(interceptors...)
	c.RefreshToken.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
	c.Video.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AccessTokenMutation:
		return c.AccessToken.mutate(ctx, m)
	case *LikeVideoMutation:
		return c.LikeVideo.mutate(ctx, m)
	case *RefreshTokenMutation:
		return c.RefreshToken.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *VideoMutation:
		return c.Video.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AccessTokenClient is a client for the AccessToken schema.
type AccessTokenClient struct {
	config
}

// NewAccessTokenClient returns a client for the AccessToken from the given config.
func NewAccessTokenClient(c config) *AccessTokenClient {
	return &AccessTokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `accesstoken.Hooks(f(g(h())))`.
func (c *AccessTokenClient) Use(hooks ...Hook) {
	c.hooks.AccessToken = append(c.hooks.AccessToken, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `accesstoken.Intercept(f(g(h())))`.
func (c *AccessTokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.AccessToken = append(c.inters.AccessToken, interceptors...)
}

// Create returns a builder for creating a AccessToken entity.
func (c *AccessTokenClient) Create() *AccessTokenCreate {
	mutation := newAccessTokenMutation(c.config, OpCreate)
	return &AccessTokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AccessToken entities.
func (c *AccessTokenClient) CreateBulk(builders ...*AccessTokenCreate) *AccessTokenCreateBulk {
	return &AccessTokenCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AccessTokenClient) MapCreateBulk(slice any, setFunc func(*AccessTokenCreate, int)) *AccessTokenCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AccessTokenCreateBulk{err: fmt.Errorf("calling to AccessTokenClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AccessTokenCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AccessTokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AccessToken.
func (c *AccessTokenClient) Update() *AccessTokenUpdate {
	mutation := newAccessTokenMutation(c.config, OpUpdate)
	return &AccessTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccessTokenClient) UpdateOne(at *AccessToken) *AccessTokenUpdateOne {
	mutation := newAccessTokenMutation(c.config, OpUpdateOne, withAccessToken(at))
	return &AccessTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccessTokenClient) UpdateOneID(id uint) *AccessTokenUpdateOne {
	mutation := newAccessTokenMutation(c.config, OpUpdateOne, withAccessTokenID(id))
	return &AccessTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AccessToken.
func (c *AccessTokenClient) Delete() *AccessTokenDelete {
	mutation := newAccessTokenMutation(c.config, OpDelete)
	return &AccessTokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccessTokenClient) DeleteOne(at *AccessToken) *AccessTokenDeleteOne {
	return c.DeleteOneID(at.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccessTokenClient) DeleteOneID(id uint) *AccessTokenDeleteOne {
	builder := c.Delete().Where(accesstoken.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccessTokenDeleteOne{builder}
}

// Query returns a query builder for AccessToken.
func (c *AccessTokenClient) Query() *AccessTokenQuery {
	return &AccessTokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAccessToken},
		inters: c.Interceptors(),
	}
}

// Get returns a AccessToken entity by its id.
func (c *AccessTokenClient) Get(ctx context.Context, id uint) (*AccessToken, error) {
	return c.Query().Where(accesstoken.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccessTokenClient) GetX(ctx context.Context, id uint) *AccessToken {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a AccessToken.
func (c *AccessTokenClient) QueryUser(at *AccessToken) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := at.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(accesstoken.Table, accesstoken.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accesstoken.UserTable, accesstoken.UserColumn),
		)
		fromV = sqlgraph.Neighbors(at.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRefreshToken queries the refresh_token edge of a AccessToken.
func (c *AccessTokenClient) QueryRefreshToken(at *AccessToken) *RefreshTokenQuery {
	query := (&RefreshTokenClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := at.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(accesstoken.Table, accesstoken.FieldID, id),
			sqlgraph.To(refreshtoken.Table, refreshtoken.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accesstoken.RefreshTokenTable, accesstoken.RefreshTokenColumn),
		)
		fromV = sqlgraph.Neighbors(at.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccessTokenClient) Hooks() []Hook {
	return c.hooks.AccessToken
}

// Interceptors returns the client interceptors.
func (c *AccessTokenClient) Interceptors() []Interceptor {
	return c.inters.AccessToken
}

func (c *AccessTokenClient) mutate(ctx context.Context, m *AccessTokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AccessTokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AccessTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AccessTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AccessTokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown AccessToken mutation op: %q", m.Op())
	}
}

// LikeVideoClient is a client for the LikeVideo schema.
type LikeVideoClient struct {
	config
}

// NewLikeVideoClient returns a client for the LikeVideo from the given config.
func NewLikeVideoClient(c config) *LikeVideoClient {
	return &LikeVideoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `likevideo.Hooks(f(g(h())))`.
func (c *LikeVideoClient) Use(hooks ...Hook) {
	c.hooks.LikeVideo = append(c.hooks.LikeVideo, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `likevideo.Intercept(f(g(h())))`.
func (c *LikeVideoClient) Intercept(interceptors ...Interceptor) {
	c.inters.LikeVideo = append(c.inters.LikeVideo, interceptors...)
}

// Create returns a builder for creating a LikeVideo entity.
func (c *LikeVideoClient) Create() *LikeVideoCreate {
	mutation := newLikeVideoMutation(c.config, OpCreate)
	return &LikeVideoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LikeVideo entities.
func (c *LikeVideoClient) CreateBulk(builders ...*LikeVideoCreate) *LikeVideoCreateBulk {
	return &LikeVideoCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *LikeVideoClient) MapCreateBulk(slice any, setFunc func(*LikeVideoCreate, int)) *LikeVideoCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &LikeVideoCreateBulk{err: fmt.Errorf("calling to LikeVideoClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*LikeVideoCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &LikeVideoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LikeVideo.
func (c *LikeVideoClient) Update() *LikeVideoUpdate {
	mutation := newLikeVideoMutation(c.config, OpUpdate)
	return &LikeVideoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LikeVideoClient) UpdateOne(lv *LikeVideo) *LikeVideoUpdateOne {
	mutation := newLikeVideoMutation(c.config, OpUpdateOne, withLikeVideo(lv))
	return &LikeVideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LikeVideoClient) UpdateOneID(id int) *LikeVideoUpdateOne {
	mutation := newLikeVideoMutation(c.config, OpUpdateOne, withLikeVideoID(id))
	return &LikeVideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LikeVideo.
func (c *LikeVideoClient) Delete() *LikeVideoDelete {
	mutation := newLikeVideoMutation(c.config, OpDelete)
	return &LikeVideoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LikeVideoClient) DeleteOne(lv *LikeVideo) *LikeVideoDeleteOne {
	return c.DeleteOneID(lv.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LikeVideoClient) DeleteOneID(id int) *LikeVideoDeleteOne {
	builder := c.Delete().Where(likevideo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LikeVideoDeleteOne{builder}
}

// Query returns a query builder for LikeVideo.
func (c *LikeVideoClient) Query() *LikeVideoQuery {
	return &LikeVideoQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLikeVideo},
		inters: c.Interceptors(),
	}
}

// Get returns a LikeVideo entity by its id.
func (c *LikeVideoClient) Get(ctx context.Context, id int) (*LikeVideo, error) {
	return c.Query().Where(likevideo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LikeVideoClient) GetX(ctx context.Context, id int) *LikeVideo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a LikeVideo.
func (c *LikeVideoClient) QueryUser(lv *LikeVideo) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := lv.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(likevideo.Table, likevideo.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, likevideo.UserTable, likevideo.UserColumn),
		)
		fromV = sqlgraph.Neighbors(lv.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVideo queries the video edge of a LikeVideo.
func (c *LikeVideoClient) QueryVideo(lv *LikeVideo) *VideoQuery {
	query := (&VideoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := lv.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(likevideo.Table, likevideo.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, likevideo.VideoTable, likevideo.VideoColumn),
		)
		fromV = sqlgraph.Neighbors(lv.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LikeVideoClient) Hooks() []Hook {
	return c.hooks.LikeVideo
}

// Interceptors returns the client interceptors.
func (c *LikeVideoClient) Interceptors() []Interceptor {
	return c.inters.LikeVideo
}

func (c *LikeVideoClient) mutate(ctx context.Context, m *LikeVideoMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LikeVideoCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LikeVideoUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LikeVideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LikeVideoDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown LikeVideo mutation op: %q", m.Op())
	}
}

// RefreshTokenClient is a client for the RefreshToken schema.
type RefreshTokenClient struct {
	config
}

// NewRefreshTokenClient returns a client for the RefreshToken from the given config.
func NewRefreshTokenClient(c config) *RefreshTokenClient {
	return &RefreshTokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `refreshtoken.Hooks(f(g(h())))`.
func (c *RefreshTokenClient) Use(hooks ...Hook) {
	c.hooks.RefreshToken = append(c.hooks.RefreshToken, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `refreshtoken.Intercept(f(g(h())))`.
func (c *RefreshTokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.RefreshToken = append(c.inters.RefreshToken, interceptors...)
}

// Create returns a builder for creating a RefreshToken entity.
func (c *RefreshTokenClient) Create() *RefreshTokenCreate {
	mutation := newRefreshTokenMutation(c.config, OpCreate)
	return &RefreshTokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RefreshToken entities.
func (c *RefreshTokenClient) CreateBulk(builders ...*RefreshTokenCreate) *RefreshTokenCreateBulk {
	return &RefreshTokenCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RefreshTokenClient) MapCreateBulk(slice any, setFunc func(*RefreshTokenCreate, int)) *RefreshTokenCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RefreshTokenCreateBulk{err: fmt.Errorf("calling to RefreshTokenClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RefreshTokenCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RefreshTokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RefreshToken.
func (c *RefreshTokenClient) Update() *RefreshTokenUpdate {
	mutation := newRefreshTokenMutation(c.config, OpUpdate)
	return &RefreshTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RefreshTokenClient) UpdateOne(rt *RefreshToken) *RefreshTokenUpdateOne {
	mutation := newRefreshTokenMutation(c.config, OpUpdateOne, withRefreshToken(rt))
	return &RefreshTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RefreshTokenClient) UpdateOneID(id uint) *RefreshTokenUpdateOne {
	mutation := newRefreshTokenMutation(c.config, OpUpdateOne, withRefreshTokenID(id))
	return &RefreshTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RefreshToken.
func (c *RefreshTokenClient) Delete() *RefreshTokenDelete {
	mutation := newRefreshTokenMutation(c.config, OpDelete)
	return &RefreshTokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RefreshTokenClient) DeleteOne(rt *RefreshToken) *RefreshTokenDeleteOne {
	return c.DeleteOneID(rt.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RefreshTokenClient) DeleteOneID(id uint) *RefreshTokenDeleteOne {
	builder := c.Delete().Where(refreshtoken.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RefreshTokenDeleteOne{builder}
}

// Query returns a query builder for RefreshToken.
func (c *RefreshTokenClient) Query() *RefreshTokenQuery {
	return &RefreshTokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRefreshToken},
		inters: c.Interceptors(),
	}
}

// Get returns a RefreshToken entity by its id.
func (c *RefreshTokenClient) Get(ctx context.Context, id uint) (*RefreshToken, error) {
	return c.Query().Where(refreshtoken.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RefreshTokenClient) GetX(ctx context.Context, id uint) *RefreshToken {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a RefreshToken.
func (c *RefreshTokenClient) QueryUser(rt *RefreshToken) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(refreshtoken.Table, refreshtoken.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, refreshtoken.UserTable, refreshtoken.UserColumn),
		)
		fromV = sqlgraph.Neighbors(rt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAccessTokens queries the access_tokens edge of a RefreshToken.
func (c *RefreshTokenClient) QueryAccessTokens(rt *RefreshToken) *AccessTokenQuery {
	query := (&AccessTokenClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(refreshtoken.Table, refreshtoken.FieldID, id),
			sqlgraph.To(accesstoken.Table, accesstoken.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, refreshtoken.AccessTokensTable, refreshtoken.AccessTokensColumn),
		)
		fromV = sqlgraph.Neighbors(rt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RefreshTokenClient) Hooks() []Hook {
	return c.hooks.RefreshToken
}

// Interceptors returns the client interceptors.
func (c *RefreshTokenClient) Interceptors() []Interceptor {
	return c.inters.RefreshToken
}

func (c *RefreshTokenClient) mutate(ctx context.Context, m *RefreshTokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RefreshTokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RefreshTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RefreshTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RefreshTokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown RefreshToken mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccessTokens queries the access_tokens edge of a User.
func (c *UserClient) QueryAccessTokens(u *User) *AccessTokenQuery {
	query := (&AccessTokenClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(accesstoken.Table, accesstoken.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.AccessTokensTable, user.AccessTokensColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikeVideos queries the like_videos edge of a User.
func (c *UserClient) QueryLikeVideos(u *User) *LikeVideoQuery {
	query := (&LikeVideoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(likevideo.Table, likevideo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.LikeVideosTable, user.LikeVideosColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRefreshTokens queries the refresh_tokens edge of a User.
func (c *UserClient) QueryRefreshTokens(u *User) *RefreshTokenQuery {
	query := (&RefreshTokenClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(refreshtoken.Table, refreshtoken.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.RefreshTokensTable, user.RefreshTokensColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// VideoClient is a client for the Video schema.
type VideoClient struct {
	config
}

// NewVideoClient returns a client for the Video from the given config.
func NewVideoClient(c config) *VideoClient {
	return &VideoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `video.Hooks(f(g(h())))`.
func (c *VideoClient) Use(hooks ...Hook) {
	c.hooks.Video = append(c.hooks.Video, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `video.Intercept(f(g(h())))`.
func (c *VideoClient) Intercept(interceptors ...Interceptor) {
	c.inters.Video = append(c.inters.Video, interceptors...)
}

// Create returns a builder for creating a Video entity.
func (c *VideoClient) Create() *VideoCreate {
	mutation := newVideoMutation(c.config, OpCreate)
	return &VideoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Video entities.
func (c *VideoClient) CreateBulk(builders ...*VideoCreate) *VideoCreateBulk {
	return &VideoCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *VideoClient) MapCreateBulk(slice any, setFunc func(*VideoCreate, int)) *VideoCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &VideoCreateBulk{err: fmt.Errorf("calling to VideoClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*VideoCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &VideoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Video.
func (c *VideoClient) Update() *VideoUpdate {
	mutation := newVideoMutation(c.config, OpUpdate)
	return &VideoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VideoClient) UpdateOne(v *Video) *VideoUpdateOne {
	mutation := newVideoMutation(c.config, OpUpdateOne, withVideo(v))
	return &VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VideoClient) UpdateOneID(id string) *VideoUpdateOne {
	mutation := newVideoMutation(c.config, OpUpdateOne, withVideoID(id))
	return &VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Video.
func (c *VideoClient) Delete() *VideoDelete {
	mutation := newVideoMutation(c.config, OpDelete)
	return &VideoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VideoClient) DeleteOne(v *Video) *VideoDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VideoClient) DeleteOneID(id string) *VideoDeleteOne {
	builder := c.Delete().Where(video.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VideoDeleteOne{builder}
}

// Query returns a query builder for Video.
func (c *VideoClient) Query() *VideoQuery {
	return &VideoQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeVideo},
		inters: c.Interceptors(),
	}
}

// Get returns a Video entity by its id.
func (c *VideoClient) Get(ctx context.Context, id string) (*Video, error) {
	return c.Query().Where(video.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VideoClient) GetX(ctx context.Context, id string) *Video {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLikeVideos queries the like_videos edge of a Video.
func (c *VideoClient) QueryLikeVideos(v *Video) *LikeVideoQuery {
	query := (&LikeVideoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(likevideo.Table, likevideo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, video.LikeVideosTable, video.LikeVideosColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VideoClient) Hooks() []Hook {
	return c.hooks.Video
}

// Interceptors returns the client interceptors.
func (c *VideoClient) Interceptors() []Interceptor {
	return c.inters.Video
}

func (c *VideoClient) mutate(ctx context.Context, m *VideoMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&VideoCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&VideoUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&VideoDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Video mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		AccessToken, LikeVideo, RefreshToken, User, Video []ent.Hook
	}
	inters struct {
		AccessToken, LikeVideo, RefreshToken, User, Video []ent.Interceptor
	}
)
