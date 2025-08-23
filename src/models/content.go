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

type Content struct {
	Model
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ContentType string `json:"content_type"`
	Content string `json:"content"`
}

type ShallowContent struct {
	ShallowModel
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ContentType string `json:"content_type"`
	Content string `json:"content"`
}

func NewShallowContent(id *string) ShallowContent {
	c := ShallowContent{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_content"
	return c
}

func (c Content) ShallowGetIn(e echo.Context) ([]ShallowContent, error) {
	GetLogger(4).Flogger("ShallowGetIn called")
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	q := fmt.Sprintf("SELECT %s FROM content WHERE id IN ('%s'::uuid)", c.Columns, strings.Join(c.Model.Manifest, "'::uuid, '"))
	rows, err := db.Query(q)
	defer rows.Close()
	if err != nil {
		return nil, merrors.ContentGetError{Info: q, Package: "models", Struct: "Content", Function: "ShallowGetIn", DB:db}.Wrap(err).Log()
	}
	ta := make([]ShallowContent, 0)
	for rows.Next() {
		t := ShallowContent{}
		t, err = t.Scan(e, rows)
		if err != nil {
			return nil, err
		}
		ta = append(ta, t)
	}
	return ta, nil
}

func (c Content) GetIn(e echo.Context) ([]Content, error) {
	GetLogger(4).Flogger("GetIn called")
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	q := fmt.Sprintf("SELECT %s FROM content WHERE id IN ('%s'::uuid)", c.Columns, strings.Join(c.Model.Manifest, "'::uuid, '"))
	rows, err := db.Query(q)
	defer rows.Close()
	if err != nil {
		return nil, merrors.ContentGetError{Info: q, DB:db}.Wrap(err).Log()
	}
	ta := make([]Content, 0)
	for rows.Next() {
		t, err := c.Scan(e, rows)
		if err != nil {
			return nil, err
		}
		ta = append(ta, t)
	}
	return ta, nil
}

func (c Content) Check(e echo.Context, id string) (int64, error) {
	db := database.GetPQDatabase()
	defer db.Close()
	countQ := fmt.Sprintf("SELECT COUNT(id) FROM content WHERE id = '%s'", id)
	rows, err := db.Query(countQ)
	if err != nil {
		return 0, merrors.DBQueryError{DB: db}.Wrap(err).Log()
	}
	var count interface{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&count)
	}

	if r, ok := count.(int64); ok {
		return r, nil
	}
	return 0, nil
}

func (c *Content) Get(e echo.Context) error {
	c.Init()
	GetLogger(4).Flogger("Content Get called for %s", c.Model.ID)
	db := database.GetPQDatabase()
	defer db.Close()
	var id string
	if c.Model.ID != "" {
		id = c.Model.ID
	}
	if c.ID != "" {
		id = c.ID
	}
	check, err := c.Check(e, id)
	if err != nil {
		return merrors.DBQueryError{DB: db}.Wrap(err).Log()
	}
	if check > 0 {
		q := fmt.Sprintf("SELECT %s FROM content WHERE id = '%s'::uuid OR content @> '{\"id\":\"%s\"}'", c.Columns, id, id)
		rows, err := db.Query(q)
		if err != nil {
			return merrors.ContentGetError{Info: fmt.Sprintf("id: %s, q: %s", id, q), DB:db}.Wrap(err).Log()
		}
		defer rows.Close()
		var t Content
		for rows.Next() {
			t, err = c.Scan(e, rows)
			if err != nil {
				return merrors.DBContentScanError{Info: fmt.Sprintf("id: %s, q: %s", id, q), DB:db}.Wrap(err).Log()
			}
			*c = t
		}
	}
	return nil
}

func (c *ShallowContent) Get(e echo.Context) error {
	GetLogger(4).Flogger("Get called")
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	var id string
	if c.ShallowModel.ID != "" {
		id = c.ShallowModel.ID
	}
	if c.ID != "" {
		id = c.ID
	}
	q := fmt.Sprintf("SELECT %s FROM content WHERE id = '%s'::uuid", c.Columns, id)
	rows, err := db.Query(q)
	if err != nil {
		
		return merrors.ContentGetError{Info: fmt.Sprintf("id: %s, q: %s", id, q), DB:db}.Wrap(err).Log()
	}
	defer rows.Close()
	var t ShallowContent
	for rows.Next() {
		t, err = c.Scan(e, rows)
		if err != nil {
			
			return merrors.DBContentScanError{Info: fmt.Sprintf("id: %s, q: %s", id, q), DB:db}.Wrap(err).Log()
		}
		*c = t
	}
	return nil
}

func (c *Content) FindBy(e echo.Context, key, value string) error {
	GetLogger(4).Flogger("FindBy called")
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	q := fmt.Sprintf("SELECT %s FROM content WHERE content @> '{\"%s\":\"%s\"}'", c.Columns, key, value)
	rows, err := db.Query(q)
	if err != nil {
		
		return merrors.ContentFindByError{Info: fmt.Sprintf("id: %s, q: %s", c.Model.ID, q), DB: db}.Wrap(err).Log()
	}
	defer rows.Close()
	var t Content
	ctr := 0
	for rows.Next() {
		ctr++
		t, err = c.Scan(e, rows)
		if err != nil {
			
			return merrors.DBContentScanError{Info: fmt.Sprintf("id: %s, q: %s", c.Model.ID, q), DB: db}.Wrap(err).Log()
		}
	}
	if ctr == 0 {
		
		return merrors.NilContentError{Info: q, Package: "models", Struct: "Content", Function: "FindBy", Code: 404}
	}
	if t.Content == "" {
		
		return merrors.NilContentError{Package: "models", Struct: "Content", Function: "FindBy", Code: 404, DB:db}.Wrap(err).Log().BubbleCode()
	}
	*c = t
	return nil
}

func (c *ShallowContent) FindBy(e echo.Context, key, value string) error {
	GetLogger(4).Flogger("FindBy called")
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	q := fmt.Sprintf("SELECT %s FROM content WHERE content @> '{\"%s\":\"%s\"}'", c.Columns, key, value)
	rows, err := db.Query(q)
	if err != nil {
		
		return merrors.ContentFindByError{Info: fmt.Sprintf("id: %s, q: %s", c.ShallowModel.ID, q), DB: db}.Wrap(err).Log()
	}
	defer rows.Close()
	var t ShallowContent
	ctr := 0
	for rows.Next() {
		ctr++
		t, err = c.Scan(e, rows)
		if err != nil {
			
			return merrors.DBContentScanError{Info: fmt.Sprintf("id: %s, q: %s", c.ShallowModel.ID, q), DB: db}.Wrap(err).Log()
		}
	}
	if ctr == 0 {
		
		return merrors.NilContentError{Info: q, Package: "models", Struct: "Content", Function: "FindBy", Code: 404}
	}
	if t.Content == "" {
		
		return merrors.NilContentError{Package: "models", Struct: "Content", Function: "FindBy", Code: 404, DB:db}.Wrap(err).Log().BubbleCode()
	}
	*c = t
	return nil
}

func (c Content) Set(e echo.Context, update bool) error {
	GetLogger(3).Flogger("model.content.set id: %s", c.Model.ID)
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	tx, err := database.GetPQDatabase().Begin()
	if err != nil {
		return merrors.DBConnectionError{DB: db}.Wrap(err).Log()
	}
	q := ""
	check, err := c.Check(e, c.Model.ID)
	if err != nil {
		return merrors.DBConnectionError{DB: db}.Wrap(err).Log()
	}

	if update && check > int64(0) {
		q = fmt.Sprintf("UPDATE content SET updated_at = '%s', content = '%s' WHERE id = '%s' RETURNING id", time.Now().Format(time.RFC3339), c.Content, c.Model.ID)
		_, err = tx.Exec(q)
		if err != nil {
			tx.Rollback()
			GetLogger(1).Flogger("rollback: %s", err.Error())
			return merrors.SQLQueryError{Info: fmt.Sprintf("q: %s, values: %#v", q, c.Values()[0]), DB: db}.Wrap(err).Log()
		}
	} else {
		q = fmt.Sprintf("INSERT INTO content (%s) VALUES (%s) ON CONFLICT(id) %s RETURNING id", c.Model.Columns, c.Model.Values, c.Model.Conflict)
		_, err = tx.Exec(q, c.Values()...)
		if err != nil {
			tx.Rollback()
			return merrors.SQLQueryError{Info: fmt.Sprintf("q: %s, values: %#v", q, c.Values()[0]), DB: db}.Wrap(err).Log()
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		
		return merrors.TransactionCommitError{Info: "content set", DB: db}.Wrap(err).Log()
	}
	return nil
}

func betterEscapes(s string) string {
	s = strings.ReplaceAll(s, "\\\\\"", "\"")
	s = strings.ReplaceAll(s, "\\\"", "\"")
	s = strings.ReplaceAll(s, "'", "\\'")
	return s
}

func (c ShallowContent) Set(e echo.Context, update bool) error {
	GetLogger(4).Flogger("Set called")
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		
		return merrors.DBConnectionError{}.Wrap(err).Log()
	}
	q := ""
	if update {
		q = fmt.Sprintf("UPDATE content SET updated_at = '%s', content = '%s' WHERE id = '%s' RETURNING id", time.Now().Format(time.RFC3339), c.Content, c.ID)
		_, err = tx.Exec(q)
		if err != nil {
			tx.Rollback()
			
			return merrors.SQLQueryError{Info: fmt.Sprintf("q: %s, values: %#v", q, c.Values()[0]), DB: db}.Wrap(err).Log()
		}
	} else {
		q = fmt.Sprintf("INSERT INTO content (%s) VALUES (%s) ON CONFLICT(id) %s RETURNING id", c.ShallowModel.Columns, c.ShallowModel.Values, c.ShallowModel.Conflict)
		_, err = tx.Exec(q, c.Values()...)
		if err != nil {
			tx.Rollback()
			
			return merrors.SQLQueryError{Info: fmt.Sprintf("q: %s, values: %#v", q, c.Values()[0]), DB: db}.Wrap(err).Log()
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		
		return merrors.TransactionCommitError{Info: "content set", DB: db}.Wrap(err).Log()
	}
	return nil
}

func (c Content) List(e echo.Context) ([]Content, error) {
	GetLogger(4).Flogger("models.Content.List called")
	GetLogger(4).Flogger("model.Content.ContentType: %s", c.Model.ContentType)
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	if db == nil {
		return nil, merrors.DBConnectionError{DB: db}.New("db is nil").Log()
	}

	textOut := strings.Replace(c.Model.Columns, "e, content", "e, content::text", 1)
	q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = $1", textOut)
	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, merrors.DBPrepareStatementError{Info: q, DB:db}.Wrap(err).Log()
	}
	rows, err := stmt.Query(c.Model.ContentType)
	if err != nil {
		return nil, merrors.DBStatementQueryQueryError{Info: q, DB:db}.Wrap(err).Log()
	}
	defer rows.Close()
	ids := make([]string, 0)
	r := make([]Content, 0)
	for rows.Next() {
		content, err := c.Scan(e, rows)
		ids = append(ids, content.Model.ID)
		r = append(r, content)
		if err != nil {
			return nil, merrors.DBContentScanError{DB:db}.Wrap(err).Log()
		}
		if rows.Err() != nil {
			return nil, merrors.DBConnectionError{DB:db}.Wrap(err).Log()
		}
	}
	if rows.Err() != nil {
		return nil, merrors.DBContentScanError{DB:db}.Wrap(err).Log()
	}
	return r, nil
}

func (c ShallowContent) List(e echo.Context) ([]ShallowContent, error) {
	GetLogger(4).Flogger("List called")
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	if db == nil {
		return nil, merrors.DBConnectionError{DB: db}.New("db is nil").Log()
	}

	textOut := strings.Replace(c.ShallowModel.Columns, "e, content", "e, content::text", 1)
	q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = $1", textOut)
	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, merrors.DBPrepareStatementError{Info: q, DB:db}.Wrap(err).Log()
	}
	rows, err := stmt.Query(c.ShallowModel.ContentType)
	if err != nil {
		return nil, merrors.DBStatementQueryQueryError{Info: q, DB:db}.Wrap(err).Log()
	}
	defer rows.Close()
	r := make([]ShallowContent, 0)
	for rows.Next() {
		content, err := c.Scan(e, rows)
		r = append(r, content)
		if err != nil {
			return nil, merrors.DBContentScanError{DB:db}.Wrap(err).Log()
		}
		if rows.Err() != nil {
			return nil, merrors.DBConnectionError{DB:db}.Wrap(err).Log()
		}
	}
	if rows.Err() != nil {
		return nil, merrors.DBContentScanError{DB:db}.Wrap(err).Log()
	}
	rows.Close()
	return r, nil
}

func (c Content) ListBy(e echo.Context, key, value interface{}) ([]Content, error) {
	GetLogger(4).Flogger("ListBy called")
	db := database.GetPQDatabase()
	defer db.Close()
	c.Init()
	q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = $1 AND content @> '{\"%s\":\"%v\"}'", c.Model.Columns, key, value)
	rows, err := db.Query(q, c.Model.ContentType)
	if err != nil {
		return nil, merrors.DBQueryError{Info: q, Package: "models", Struct: "Content", Function: "ListBy", DB: db}.Wrap(err).Log()
	}
	defer rows.Close()
	r := make([]Content, 0)
	for rows.Next() {
		content, err := c.Scan(e, rows)
		if err != nil {
			return nil, merrors.DBContentScanError{Info: q, Package: "models", Struct: "Content", Function: "ListBy", DB: db}.Wrap(err).Log()
		}
		r = append(r, content)
	}
	return r, nil
}

func (c ShallowContent) ListBy(e echo.Context, key, value interface{}) ([]ShallowContent, error) {
	GetLogger(4).Flogger("ListBy called")
	db := database.GetPQDatabase()
	defer db.Close()
	c.Init()
	q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = $1 AND content @> '{\"%s\":\"%v\"}'", c.ShallowModel.Columns, key, value)
	rows, err := db.Query(q, c.ShallowModel.ContentType)
	if err != nil {
		return nil, merrors.DBQueryError{Info: q, Package: "models", Struct: "Content", Function: "ListBy", DB: db}.Wrap(err).Log()
	}
	defer rows.Close()
	r := make([]ShallowContent, 0)
	for rows.Next() {
		content, err := c.Scan(e, rows)
		if err != nil {
			return nil, merrors.DBContentScanError{Info: q, Package: "models", Struct: "Content", Function: "ListBy", DB: db}.Wrap(err).Log()
		}
		r = append(r, content)
	}
	return r, nil
}

func (c Content) Delete(e echo.Context) error {
	GetLogger(4).Flogger("Content Delete called")
	GetLogger(4).Flogger("content.id: %s", c.Model.ID)
	db := database.GetPQDatabase()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return merrors.DBConnectionError{}.Wrap(err).Log()
	}
	q := fmt.Sprintf("DELETE FROM content WHERE id = $1 OR content @> '{\"id\":\"%s\"}'", c.Model.ID)
	_, err = tx.Exec(q, c.Model.ID)
	if err != nil {
		tx.Rollback()
		return merrors.SQLDeleteErorr{Info: q, Package: "models", Struct: "Content", Function: "Delete", DB: db}.Wrap(err).Log()
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return merrors.TransactionCommitError{Package: "models", Struct: "Content", Function: "Delete", DB: db}.Wrap(err).Log()
	}
	
	return nil
}

func (c Content) CustomQuery(e echo.Context, write bool, q string, vars ...any) ([]Content, error) {
	GetLogger(4).Flogger("CustomQuery called")
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	if write {
		tx, err := db.Begin()
		if err != nil {
			return nil, merrors.DBConnectionError{DB: db}.Wrap(err).Log()
		}
		vals := make([]any, 0)
		vals = append(vals, c.Model.Columns)
		vals = append(vals, c.Model.Values)
		vals = append(vals, c.Model.Conflict)
		vals = append(vals, vars)
		_, err = tx.Exec(q, vals...)
		if err != nil {
			if err = tx.Rollback(); err != nil {
				panic(fmt.Errorf("rollback error: %w", err))
			}
			if err = tx.Commit(); err != nil {
				panic(fmt.Errorf("commit error: %w", err))
			}
			return nil, merrors.SQLQueryError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery", DB: db}.Wrap(err).Log()
		}
		if err = tx.Commit(); err != nil {
			panic(fmt.Errorf("commit error: %w", err))
		}
		return nil, nil
	}
	rows, err := db.Query(q)
	if err != nil {
		return nil, merrors.DBQueryError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery", DB: db}.Wrap(err).Log()			
	}
	defer rows.Close()
	ctr := 0
	r := make([]Content, 0)
	for rows.Next() {
		ctr++
		content, err := c.Scan(e, rows)
		if err != nil {
			return nil, merrors.DBContentScanError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery", DB: db}.Wrap(err).Log()
		}
		r = append(r, content)
	}
	if rows.Err() != nil {
			return nil, merrors.DBContentScanError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery", DB: db}.Wrap(rows.Err())
	}
	if ctr == 0 {
		return nil, merrors.NilContentError{Info: q, Package: "models", Struct: "Content", Function: "FindBy", Code: 404, DB:db}.Wrap(err).Log().BubbleCode()
	}
	for _, t := range r {
		if t.Content == "" {
			return nil, merrors.NilContentError{Package: "models", Struct: "Content", Function: "FindBy", Code: 500, DB:db}.Wrap(err).Log()
		}
	}
	return r, nil
}

func (c ShallowContent) CustomQuery(e echo.Context, write bool, q string, vars ...any) ([]ShallowContent, error) {
	GetLogger(4).Flogger("CustomQuery called")
	c.Init()
	db := database.GetPQDatabase()
	defer db.Close()
	if write {
		tx, err := db.Begin()
		if err != nil {
			return nil, merrors.DBConnectionError{DB: db}.Wrap(err).Log()
		}
		vals := make([]any, 0)
		vals = append(vals, c.ShallowModel.Columns)
		vals = append(vals, c.ShallowModel.Values)
		vals = append(vals, c.ShallowModel.Conflict)
		vals = append(vals, vars)
		_, err = tx.Exec(q, vals...)
		if err != nil {
			if err = tx.Rollback(); err != nil {
				panic(fmt.Errorf("rollback error: %w", err))
			}
			if err = tx.Commit(); err != nil {
				panic(fmt.Errorf("commit error: %w", err))
			}
			return nil, merrors.SQLQueryError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery", DB: db}.Wrap(err).Log()
		}
		if err = tx.Commit(); err != nil {
			panic(fmt.Errorf("commit error: %w", err))
		}
		return nil, nil
	}
	// vals = append(vals, c.Model.Columns)
	// vals = append(vals, vars)
	// TODO: fix error when another use case arises
	// rows, err := db.Query(q, vals...)
	rows, err := db.Query(q)
	if err != nil {
		return nil, merrors.DBQueryError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery", DB: db}.Wrap(err).Log()			
	}
	defer rows.Close()
	ctr := 0
	r := make([]ShallowContent, 0)
	for rows.Next() {
		ctr++
		content, err := c.Scan(e, rows)
		if err != nil {
			return nil, merrors.DBContentScanError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery", DB: db}.Wrap(err).Log()
		}
		r = append(r, content)
	}
	if rows.Err() != nil {
			return nil, merrors.DBContentScanError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery", DB: db}.Wrap(rows.Err())
	}
	if ctr == 0 {
		return nil, merrors.NilContentError{Info: q, Package: "models", Struct: "Content", Function: "FindBy", Code: 404, DB:db}.Wrap(err).Log().BubbleCode()
	}
	for _, t := range r {
		if t.Content == "" {
			return nil, merrors.NilContentError{Package: "models", Struct: "Content", Function: "FindBy", Code: 500, DB:db}.Wrap(err).Log()
		}
	}
	return r, nil
}

func (c Content) Scan(e echo.Context, rows Scannable) (Content, error) {
	err := rows.Scan(&c.Model.ID, &c.Model.CreatedAt, &c.Model.UpdatedAt, &c.Model.ContentType, &c.Content)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c ShallowContent) Scan(e echo.Context, rows Scannable) (ShallowContent, error) {
	GetLogger(4).Flogger("Scan called")
	err := rows.Scan(&c.ShallowModel.ID, &c.ShallowModel.CreatedAt, &c.ShallowModel.UpdatedAt, &c.ShallowModel.ContentType, &c.Content)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c Content) Values() []any {
	r := make([]any, 0)
	if c.Model.ID != c.ID {
		GetLogger(4).Flogger("id set mismatch")
	}
	if c.Model.ID != "" {
		r = append(r, c.Model.ID)
	} else if c.ID != "" {
		r = append(r, c.ID)
	} else {
		r = append(r, uuid.NewString())
	}
	r = append(r, c.Model.CreatedAt.Format(time.RFC3339))
	r = append(r, c.Model.UpdatedAt.Format(time.RFC3339))
	r = append(r, c.Model.ContentType)
	r = append(r, c.Content)
	return r
}

func (c ShallowContent) Values() []any {
	r := make([]any, 0)
	if c.ShallowModel.ID != "" {
		r = append(r, c.ShallowModel.ID)
	} else if c.ID != "" {
		r = append(r, c.ID)
	} else {
		r = append(r, uuid.NewString())
	}
	r = append(r, c.ShallowModel.CreatedAt.Format(time.RFC3339))
	r = append(r, c.ShallowModel.UpdatedAt.Format(time.RFC3339))
	r = append(r, c.ShallowModel.ContentType)
	r = append(r, c.Content)
	return r
}

func (c Content) New(ct string) Content {
	c.Model.ID = uuid.NewString()
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
	c.Model.ContentType = ct
	c.Model.Columns = "id, created_at, updated_at, content_type, content"
	c.Model.Values = "$1::uuid, $2, $3, $4, $5::jsonb"
	c.Model.Conflict = "DO UPDATE SET updated_at = $3, content = $5::jsonb"
	return c
} 
