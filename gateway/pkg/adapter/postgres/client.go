package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/isutare412/goarch/gateway/ent"
	"github.com/isutare412/goarch/gateway/pkg/config"
	"github.com/isutare412/goarch/gateway/pkg/core/port"
	_ "github.com/lib/pq"
)

type Client struct {
	cli *ent.Client
}

func NewClient(cfg config.PostgresConfig) (*Client, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Database, cfg.User, cfg.Password)
	cli, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening PostgreSQL conn: %w", err)
	}
	return &Client{cli: cli}, nil
}

func (c *Client) MigrateSchemas(ctx context.Context) error {
	if err := c.cli.Schema.Create(ctx); err != nil {
		return fmt.Errorf("migrating schemas: %w", err)
	}
	return nil
}

func (c *Client) Close(ctx context.Context) error {
	success := make(chan struct{})
	failure := make(chan error)
	go func() {
		defer close(success)

		if err := c.cli.Close(); err != nil {
			failure <- fmt.Errorf("closing PostgreSQL conn: %w", err)
			close(failure)
		}
	}()

	select {
	case <-success:
	case <-ctx.Done():
		return fmt.Errorf("timeout while closing PostgreSQL conn")
	case err := <-failure:
		return err
	}
	return nil
}

func (c *Client) BeginTx(ctx context.Context) (port.Tx, error) {
	tx, err := c.cli.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("beginning transaction: %w", err)
	}

	return &txContext{
		ctx:      ctxWithTx(ctx, tx.Client()),
		commit:   func() error { return tx.Commit() },
		rollback: func() error { return tx.Rollback() },
	}, nil
}

func (c *Client) BeginTxWithOption(ctx context.Context, opts *sql.TxOptions) (port.Tx, error) {
	tx, err := c.cli.BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("beginning transaction: %w", err)
	}

	return &txContext{
		ctx:      ctxWithTx(ctx, tx.Client()),
		commit:   func() error { return tx.Commit() },
		rollback: func() error { return tx.Rollback() },
	}, nil
}

func (c *Client) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := c.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(tx.Ctx()); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rollbacking transaction: %v: %w", rerr, err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
