package port

import (
	"context"
	"database/sql"
)

type RepositorySession interface {
	BeginTx(context.Context) (Tx, error)
	BeginTxWithOption(context.Context, *sql.TxOptions) (Tx, error)
	WithTx(context.Context, func(context.Context) error) error
}

type Tx interface {
	Ctx() context.Context
	Commit() error
	Rollback() error
}
