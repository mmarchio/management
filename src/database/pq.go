package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mmarchio/management/errors"
)

func GetPQContext(ctx context.Context) context.Context {
	if GetPQDatabase(ctx) != nil {
		return ctx
	}
	dburl := generateConnectionString()
	conn, err := sql.Open("postgres", dburl)
	if err != nil {
		panic(fmt.Errorf("database connection error: %w", merrors.DBConnectionError{}.Wrap(err)))
	}
	ctx = SetContextPQ(ctx, conn)
	ctx = SetContextPQTx(ctx, conn)
	return ctx
}

func GetPQDatabase(ctx context.Context) *sql.DB {
	v := ctx.Value(DBKey)
	if db, ok := v.(*sql.DB); ok {
		return db
	}
	return nil
}

func GetPQTx(ctx context.Context) *sql.Tx {
	v := ctx.Value(DBTXKey)
	if tx, ok := v.(*sql.Tx); ok {
		return tx
	} else {
		ctx = GetPQContext(ctx)
		return GetPQTx(ctx)
	}
	return nil
}

func SetContextPQ(ctx context.Context, conn *sql.DB) context.Context {
	ctx = context.WithValue(ctx, DBKey, conn)
	return ctx
}

func SetContextPQTx(ctx context.Context, conn *sql.DB) context.Context {
	tx, err := conn.Begin()
	if err != nil {
		panic(err)
	}
	return context.WithValue(ctx, DBTXKey, tx)
}

func GetContextPQ(ctx context.Context) (context.Context, *sql.DB) {
	return GetPQContext(ctx), GetPQDatabase(ctx)
}

func ClearDB(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, DBKey, nil)
	return ctx
} 

