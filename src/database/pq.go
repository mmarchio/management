package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/mmarchio/management/config"
	"github.com/mmarchio/management/errors"
)

type DBKeyT 	int64
type DBTXKeyT 	int64

var DBKey DBKeyT = 1
var DBTXKey DBTXKeyT = 2

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

func GetPQContext(ctx context.Context) context.Context {
	conn := GetPQDatabase()
	ctx = SetContextPQ(ctx, conn)
	ctx = SetContextPQTx(ctx, conn)
	return ctx
}

func GetPQDatabase() *sql.DB {
	dburl := generateConnectionString()
	conn, err := sql.Open("postgres", dburl)
	if err != nil {
		panic(fmt.Errorf("database connection error: %w", merrors.DBConnectionError{DB: conn}.Wrap(err)))
	}
	conn.SetConnMaxIdleTime(time.Duration(1*time.Second))
	return conn
}

func GetPQTx(ctx context.Context) *sql.Tx {
	// v := ctx.Value(DBTXKey)
	// if tx, ok := v.(*sql.Tx); ok {
	// 	return tx
	// } else {
	// 	ctx = GetPQContext(ctx)
	// 	return GetPQTx(ctx)
	// }
	return nil
}

func SetContextPQ(ctx context.Context, conn *sql.DB) context.Context {
	if _, ok := ctx.Value(DBKey).(*sql.DB); ok {
		return ctx
	}
	ctx = context.WithValue(ctx, DBKey, conn)
	return ctx
}

func SetContextPQTx(ctx context.Context, conn *sql.DB) context.Context {
	if _, ok := ctx.Value(DBTXKey).(*sql.Tx); ok {
		return ctx
	}
	tx, err := conn.Begin()
	if err != nil {
		panic(err)
	}
	return context.WithValue(ctx, DBTXKey, tx)
}

func GetContextPQ(ctx context.Context) (context.Context, *sql.DB) {
	return GetPQContext(ctx), GetPQDatabase()
}

func ClearDB(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, DBKey, nil)
	return ctx
} 

func DBStats(db *sql.DB) map[string]interface{} {
	r := make(map[string]interface{})
	r["stats"] = db.Stats()
	return r
}

