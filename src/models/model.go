package models

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mmarchio/management/database"
	merrors "github.com/mmarchio/management/errors"
)

type Model struct {
	ID string `form:"id" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ContentType string `json:"content_type"`
	Columns 	string
	Values 		string
	Conflict 	string
	Validated   bool
	Manifest	[]string	
}

type ShallowModel struct {
	ID string `form:"id" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ContentType string `json:"content_type"`
	Columns 	string
	Values 		string
	Conflict 	string
	Validated   bool	
	Manifest    []string `json:"manifest"`
	ExpandedID  string `json:"expanded_id"`
}

func (c ShallowModel) New(id, ct *string) ShallowModel {
	if id != nil {
		c.ID = *id
	} else {
		c.ID = uuid.NewString()
	}
	if ct != nil {
		c.ContentType = *ct
	}
	c.Init()
	return c
}

func (c Model) Validate() bool {
	if c.ID == "" {
		return false
	}
	if c.CreatedAt.IsZero() || c.UpdatedAt.IsZero() {
		return false
	}
	return true
}

func (c *Model) Init() {
	c.Columns = "id, created_at, updated_at, content_type, content"
	c.Values = "$1, $2, $3, $4, $5"
	c.Conflict = "DO UPDATE SET updated_at = $3, content = $5"
}

func (c *ShallowModel) Init() {
	c.Columns = "id, created_at, updated_at, content_type, content"
	c.Values = "$1, $2, $3, $4, $5"
	c.Conflict = "DO UPDATE SET updated_at = $3, content = $5"
}

func (c Model) Get(ctx context.Context, table ITable) (ITable, error) {
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	q := fmt.Sprintf("SELECT %s FROM content WHERE id = $1", c.Columns)
	rows, err := db.Query(q, c.ID)
	var t ITable
	for rows.Next() {
		t, err = table.Scan(ctx, rows)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (c Model) Set(ctx context.Context, table ITable) error {
	ctx, tx := database.GetDBTransaction(ctx)
	q := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT(id) %s",
		c.Columns,
		c.Values,
		c.Conflict,
	)
	values, err := table.Values(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, q, values...)
	if err != nil {
		tx.Rollback(ctx)
		return fmt.Errorf("err: %w\nq: %s", err, q)
	}
	fmt.Printf("set successful\nq: %s\n\nvalues: %#v\n\n", q, values)
	return nil
}

func (c Model) List(ctx context.Context, table Content) ([]Content, error) {
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	r := make([]Content, 0)
	textOut := strings.Replace(c.Columns, "id", "id::text", 1)
	q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = $1", textOut)
	fmt.Printf("q: %s\n", q)
	rows, err := db.Query(q, table.Model.ContentType)
	if err != nil {
		fmt.Println(err)
		return nil, merrors.SQLQueryError{Info: "model list"}.Wrap(err)
	}
	for rows.Next() {
		itable, err := table.Scan(ctx, rows)
		if err != nil {
			fmt.Println(err)
			return nil, merrors.DBContentScanError{Info: "model list"}.Wrap(err)
		}
		r = append(r, itable)
	}
	return r, nil
}

func (c Model) ListBy(ctx context.Context, table ITable, column string, value string) ([]ITable, error) {
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	r := make([]ITable, 0)
	q := fmt.Sprintf("SELECT %s FROM content WHERE %s = $1", c.Columns, column)
	rows, err := db.Query(q, value)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		itable, err := table.Scan(ctx, rows)
		if err != nil {
			return nil, err
		}
		r = append(r, itable)
	}
	return r, nil
}

func FromBase64(s string) (string, error) {
	b := bytes.NewBufferString(s)
	decoder := base64.NewDecoder(base64.StdEncoding, b)
	decodedBytes := make([]byte, 10)
	var str string
	for {
		n, err := decoder.Read(decodedBytes)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			return "", merrors.Base64DecodingError{Info: s}.Wrap(err)
		}
		str += string(decodedBytes[:n])
	}
	return str, nil
}

func ToBase64(b []byte) (string, error) {
	writer := new(bytes.Buffer)
	_, err := base64.NewEncoder(base64.StdEncoding, writer).Write(b)
	if err != nil {
		return "", merrors.Base64EncodingError{Info: string(b)}.Wrap(err)
	}
	return writer.String(), nil
}

