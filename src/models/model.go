package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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
	ID 			string `form:"id" json:"id"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
	ContentType string `json:"content_type"`
	TokenCount 	int64
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

func (c Model) Get(e echo.Context, table ITable) (ITable, error) {
	GetLogger(4).Flogger("Get called")
	db := database.GetPQDatabase()
	q := fmt.Sprintf("SELECT %s FROM content WHERE id = $1", c.Columns)
	rows, err := db.Query(q, c.ID)
	if err != nil {
		return nil, merrors.DBQueryError{}.Wrap(err).Log()
	}
	defer rows.Close()
	var t ITable
	for rows.Next() {
		t, err = table.Scan(e, rows)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (c Model) Set(e echo.Context, table ITable) error {
	GetLogger(4).Flogger("Get called")
	db := database.GetPQDatabase()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return merrors.DBConnectionError{DB: db}.Wrap(err).Log()
	}
	q := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT(id) %s",
		c.Columns,
		c.Values,
		c.Conflict,
	)
	values, err := table.Values(e)
	if err != nil {
		return err
	}
	_, err = tx.Exec(q, values...)
	if err != nil {
		tx.Rollback()
		return merrors.DBQueryError{DB: db}.New("err: %w\nq: %s", err, q)
	}
	if err := tx.Commit(); err != nil {
		return merrors.TransactionCommitError{DB: db}.Wrap(err).Log()
	}
	GetLogger(4).Flogger("set successful\nq: %s\n\nvalues: %#v\n\n", q, values)
	return nil
}

func (c Model) List(e echo.Context, table Content) ([]Content, error) {
	GetLogger(4).Flogger("Get called")
	db := database.GetPQDatabase()
	r := make([]Content, 0)
	textOut := strings.Replace(c.Columns, "id", "id::text", 1)
	q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = $1", textOut)
	rows, err := db.Query(q, table.Model.ContentType)
	if err != nil {
		return nil, merrors.SQLQueryError{Info: "model list", DB: db}.Wrap(err).Log()
	}
	defer rows.Close()
	for rows.Next() {
		itable, err := table.Scan(e, rows)
		if err != nil {
			return nil, merrors.DBContentScanError{Info: "model list", DB: db}.Wrap(err).Log()
		}
		r = append(r, itable)
	}
	return r, nil
}

func (c Model) ListBy(e echo.Context, table ITable, column string, value string) ([]ITable, error) {
	GetLogger(4).Flogger("Get called")
	db := database.GetPQDatabase()
	r := make([]ITable, 0)
	q := fmt.Sprintf("SELECT %s FROM content WHERE %s = $1", c.Columns, column)
	rows, err := db.Query(q, value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		itable, err := table.Scan(e, rows)
		if err != nil {
			return nil, err
		}
		r = append(r, itable)
	}
	return r, nil
}

