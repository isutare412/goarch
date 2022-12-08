package postgres

import (
	"context"

	"github.com/isutare412/goarch/gateway/pkg/core/ent"
)

type ctxKeyTransaction struct{}

func ctxWithTx(ctx context.Context, tx *ent.Client) context.Context {
	return context.WithValue(ctx, ctxKeyTransaction{}, tx)
}

func txFromCtx(ctx context.Context) *ent.Client {
	txAny := ctx.Value(ctxKeyTransaction{})
	if tx, ok := txAny.(*ent.Client); ok {
		return tx
	}
	return nil
}

type txContext struct {
	ctx      context.Context
	commit   func() error
	rollback func() error
}

func (tc *txContext) Ctx() context.Context {
	return tc.ctx
}

func (tc *txContext) Commit() error {
	return tc.commit()
}

func (tc *txContext) Rollback() error {
	return tc.rollback()
}
