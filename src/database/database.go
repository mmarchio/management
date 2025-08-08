package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/mmarchio/management/config"
	"github.com/mmarchio/management/errors"
)

type DBKeyT 	int64
type DBTXKeyT 	int64

var DBKey DBKeyT = 1
var DBTXKey DBTXKeyT = 2

func GetDatabaseCtx() context.Context {
	ctx := context.Background()
	ctx, _ = GetDatabase(ctx)
	return ctx
}

func GetDatabase(ctx context.Context) (context.Context, *pgx.Conn) {
	dburl := generateConnectionString()
	conn, err := pgx.Connect(ctx, dburl)
	if err != nil {
		panic(fmt.Errorf("database connection error: %w", merrors.DBConnectionError{}.Wrap(err)))
	}
	ctx = SetContextDB(ctx, conn)
	return ctx, conn
}

func SetContextDB(ctx context.Context, db *pgx.Conn) context.Context {
	return context.WithValue(ctx, DBKey, db)
}

func GetContextDB(ctx context.Context) (context.Context, *pgx.Conn) {
	dbi := ctx.Value(DBKey)
	var db *pgx.Conn
	var ok bool
	if db, ok = dbi.(*pgx.Conn); ok {
		return ctx, db
	}
	ctx, db = GetDatabase(ctx)	
	return ctx, db
}

func SetContextTX(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, DBTXKey, tx)
}

func GetDBTransaction(ctx context.Context) (context.Context, pgx.Tx) {
	var err error
	var tx pgx.Tx
	var db *pgx.Conn
	var ndb *pgx.Conn
	txi := ctx.Value(DBTXKey)
	if txp, ok := txi.(pgx.Tx); ok {
		tx = txp
	}
	if tx == nil {
		dbi := ctx.Value(DBKey)
		if dbp, ok := dbi.(pgx.Conn); ok {
			db = &dbp
		}
		var ntx pgx.Tx
		if db == nil {
			ctx, ndb = GetDatabase(ctx)
			ctx = SetContextDB(ctx, ndb)
			ntx, err = ndb.Begin(ctx)
			if err != nil {
				panic(err)
			}
			ctx = SetContextTX(ctx, ntx)
			return ctx, ntx
		} else {
			ntx, err = db.Begin(ctx)
			if err != nil {
				panic(err)
			}
			ctx = SetContextTX(ctx, ntx)
			return ctx, ntx
		}
	}
	return ctx, tx
}

func generateConnectionString() string {
	//postgres://<username>:<password>@<host>:<port>/<database_name>?<options>
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s", 
		config.DBUser, 
		config.DBPass, 
		config.DBHost, 
		config.DBPort, 
		config.DBName, 
		config.DBOptions,
	)
}
