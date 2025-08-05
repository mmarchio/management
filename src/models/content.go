package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mmarchio/management/database"
	merrors "github.com/mmarchio/management/errors"
)

type Content struct {
	Model
	ID string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ContentType string `json:"content_type"`
	Content string `json:"content"`
}

type ShallowContent struct {
	ShallowModel
	ID string `json:"id"`
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

func (c Content) ShallowGetIn(ctx context.Context) ([]ShallowContent, error) {
	c.Init()
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	q := fmt.Sprintf("SELECT %s FROM content WHERE id IN ('%s'::uuid)", c.Columns, strings.Join(c.Model.Manifest, "'::uuid, '"))
	rows, err := db.Query(q)
	if err != nil {
		return nil, merrors.ContentGetError{Info: q, Package: "models", Struct: "Content", Function: "ShallowGetIn"}.Wrap(err)
	}
	ta := make([]ShallowContent, 0)
	for rows.Next() {
		t, err := c.ShallowScan(ctx, rows)
		if err != nil {
			return nil, err
		}
		ta = append(ta, t)
	}
	return ta, nil
}

func (c Content) GetIn(ctx context.Context) ([]Content, error) {
	c.Init()
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	q := fmt.Sprintf("SELECT %s FROM content WHERE id IN ('%s'::uuid)", c.Columns, strings.Join(c.Model.Manifest, "'::uuid, '"))
	rows, err := db.Query(q)
	if err != nil {
		return nil, merrors.ContentGetError{Info: q, Package: "models", Struct: "Content", Function: "GetIn"}.Wrap(err)
	}
	ta := make([]Content, 0)
	for rows.Next() {
		t, err := c.Scan(ctx, rows)
		if err != nil {
			return nil, err
		}
		ta = append(ta, t)
	}
	return ta, nil
}

func (c *Content) Get(ctx context.Context) error {
	c.Init()
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	var id string
	if c.Model.ID != "" {
		id = c.Model.ID
	}
	if c.ID != "" {
		id = c.ID
	}
	defer db.Close()
	q := fmt.Sprintf("SELECT %s FROM content WHERE id = '%s'::uuid OR content @> '{\"id\":\"%s\"}'", c.Columns, id, id)
	rows, err := db.Query(q)
	if err != nil {
		return merrors.ContentGetError{Info: fmt.Sprintf("id: %s, q: %s", id, q)}.Wrap(err)
	}
	var t Content
	for rows.Next() {
		t, err = c.Scan(ctx, rows)
		if err != nil {
			return merrors.DBContentScanError{Info: fmt.Sprintf("id: %s, q: %s", id, q)}.Wrap(err)
		}
		*c = t
	}
	return nil
}

func (c *ShallowContent) Get(ctx context.Context) error {
	c.Init()
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	var id string
	if c.ShallowModel.ID != "" {
		id = c.ShallowModel.ID
	}
	if c.ID != "" {
		id = c.ID
	}
	defer db.Close()
	q := fmt.Sprintf("SELECT %s FROM content WHERE id = '%s'::uuid OR content @> '{\"id\":\"%s\"}'", c.Columns, id, id)
	rows, err := db.Query(q)
	if err != nil {
		return merrors.ContentGetError{Info: fmt.Sprintf("id: %s, q: %s", id, q)}.Wrap(err)
	}
	var t ShallowContent
	for rows.Next() {
		t, err = c.Scan(ctx, rows)
		if err != nil {
			return merrors.DBContentScanError{Info: fmt.Sprintf("id: %s, q: %s", id, q)}.Wrap(err)
		}
		*c = t
	}
	return nil
}

func (c *Content) FindBy(ctx context.Context, key, value string) error {
	c.Init()
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	q := fmt.Sprintf("SELECT %s FROM content WHERE content @> '{\"%s\":\"%s\"}'", c.Columns, key, value)
	rows, err := db.Query(q)
	if err != nil {
		return merrors.ContentFindByError{Info: fmt.Sprintf("id: %s, q: %s", c.Model.ID, q)}.Wrap(err)
	}
	var t Content
	ctr := 0
	for rows.Next() {
		ctr++
		t, err = c.Scan(ctx, rows)
		if err != nil {
			return merrors.DBContentScanError{Info: fmt.Sprintf("id: %s, q: %s", c.Model.ID, q)}.Wrap(err)
		}
	}
	if ctr == 0 {
		return merrors.NilContentError{Info: q, Package: "models", Struct: "Content", Function: "FindBy", Code: 404}
	}
	if t.Content == "" {
		return merrors.NilContentError{Package: "models", Struct: "Content", Function: "FindBy", Code: 404}.Wrap(err).BubbleCode()
	}
	*c = t
	return nil
}

func (c Content) Set(ctx context.Context) error {
	c.Init()
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	tx := database.GetPQTx(database.GetPQContext(ctx))
	q := fmt.Sprintf("INSERT INTO content (%s) VALUES (%s) ON CONFLICT(id) %s RETURNING id", c.Model.Columns, c.Model.Values, c.Model.Conflict)
	_, err := tx.Exec(q, c.Values()...)
	if err != nil {
		tx.Rollback()
		return merrors.SQLQueryError{Info: fmt.Sprintf("q: %s, values: %#v", q, c.Values()[0])}.Wrap(err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return merrors.TransactionCommitError{Info: "content set"}.Wrap(err)
	}
	return nil
}

func (c ShallowContent) Set(ctx context.Context) error {
	c.Init()
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	tx := database.GetPQTx(ctx)
	q := fmt.Sprintf("INSERT INTO content (%s) VALUES (%s) ON CONFLICT(id) %s RETURNING id", c.ShallowModel.Columns, c.ShallowModel.Values, c.ShallowModel.Conflict)
	_, err := tx.Exec(q, c.Values()...)
	if err != nil {
		tx.Rollback()
		return merrors.SQLQueryError{Info: fmt.Sprintf("q: %s, values: %#v", q, c.Values()[0])}.Wrap(err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return merrors.TransactionCommitError{Info: "content set"}.Wrap(err)
	}
	return nil
}

func (c Content) List(ctx context.Context) ([]Content, error) {
	c.Init()
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	if db == nil {
		return nil, merrors.DBConnectionError{}.Wrap(fmt.Errorf("db is nil"))
	}

	textOut := strings.Replace(c.Model.Columns, "e, content", "e, content::text", 1)
	q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = $1", textOut)
	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, merrors.DBPrepareStatementError{Info: q}.Wrap(err)
	}
	rows, err := stmt.Query(c.Model.ContentType)
	if err != nil {
		return nil, merrors.DBStatementQueryQueryError{Info: q}.Wrap(err)
	}
	r := make([]Content, 0)
	for rows.Next() {
		content, err := c.Scan(ctx, rows)
		r = append(r, content)
		if err != nil {
			return nil, merrors.DBContentScanError{}.Wrap(err)
		}
		if rows.Err() != nil {
			return nil, merrors.DBConnectionError{}.Wrap(err)
		}
	}
	if rows.Err() != nil {
		return nil, merrors.DBContentScanError{}.Wrap(err)
	}
	rows.Close()
	return r, nil
}

func (c Content) ListBy(ctx context.Context, key, value interface{}) ([]Content, error) {
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	c.Init()
	q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = $1 AND content @> '{\"%s\":\"%v\"}'", c.Model.Columns, key, value)
	rows, err := db.Query(q, c.Model.ContentType)
	if err != nil {
		return nil, merrors.DBQueryError{Info: q, Package: "models", Struct: "Content", Function: "ListBy"}.Wrap(err)
	}
	r := make([]Content, 0)
	for rows.Next() {
		content, err := c.Scan(ctx, rows)
		if err != nil {
			return nil, merrors.DBContentScanError{Info: q, Package: "models", Struct: "Content", Function: "ListBy"}.Wrap(err)
		}
		r = append(r, content)
	}
	return r, nil
}

func (c ShallowContent) ListBy(ctx context.Context, key, value interface{}) ([]ShallowContent, error) {
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	c.Init()
	q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = $1 AND content @> '{\"%s\":\"%v\"}'", c.ShallowModel.Columns, key, value)
	rows, err := db.Query(q, c.ShallowModel.ContentType)
	if err != nil {
		return nil, merrors.DBQueryError{Info: q, Package: "models", Struct: "Content", Function: "ListBy"}.Wrap(err)
	}
	r := make([]ShallowContent, 0)
	for rows.Next() {
		content, err := c.Scan(ctx, rows)
		if err != nil {
			return nil, merrors.DBContentScanError{Info: q, Package: "models", Struct: "Content", Function: "ListBy"}.Wrap(err)
		}
		r = append(r, content)
	}
	return r, nil
}

func (c Content) Delete(ctx context.Context) error {
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	tx := database.GetPQTx(ctx)
	q := fmt.Sprintf("DELETE FROM content WHERE id = $1 OR content @> '{\"id\":\"%s\"}'", c.Model.ID)
	_, err := tx.Exec(q, c.Model.ID)
	if err != nil {
		tx.Rollback()
		return merrors.SQLDeleteErorr{Info: q, Package: "models", Struct: "Content", Function: "Delete"}.Wrap(err)
	}
	if err = tx.Commit(); err != nil {
		fmt.Printf("rollback: models:content:delete\n")
		tx.Rollback()
		return merrors.TransactionCommitError{Package: "models", Struct: "Content", Function: "Delete"}.Wrap(err)
	}
	
	return nil
}

func (c Content) CustomQuery(ctx context.Context, write bool, q string, vars ...any) ([]Content, error) {
	c.Init()
	ctx = database.GetPQContext(ctx)
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	if write {
		tx := database.GetPQTx(ctx)
		vals := make([]any, 0)
		vals = append(vals, c.Model.Columns)
		vals = append(vals, c.Model.Values)
		vals = append(vals, c.Model.Conflict)
		vals = append(vals, vars)
		_, err := tx.Exec(q, vals...)
		if err != nil {
			if err = tx.Rollback(); err != nil {
				panic(fmt.Errorf("rollback error: %w", err))
			}
			if err = tx.Commit(); err != nil {
				panic(fmt.Errorf("commit error: %w", err))
			}
			return nil, merrors.SQLQueryError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery"}.Wrap(err)
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
		return nil, merrors.DBQueryError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery"}.Wrap(err)			
	}
	ctr := 0
	r := make([]Content, 0)
	for rows.Next() {
		ctr++
		content, err := c.Scan(ctx, rows)
		if err != nil {
			return nil, merrors.DBContentScanError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery"}.Wrap(err)
		}
		r = append(r, content)
	}
	if rows.Err() != nil {
			return nil, merrors.DBContentScanError{Info: q, Package: "models", Struct: "Content", Function: "CustomQuery"}.Wrap(rows.Err())
	}
	if ctr == 0 {
		return nil, merrors.NilContentError{Info: q, Package: "models", Struct: "Content", Function: "FindBy", Code: 404}.Wrap(err).BubbleCode()
	}
	for _, t := range r {
		if t.Content == "" {
			return nil, merrors.NilContentError{Package: "models", Struct: "Content", Function: "FindBy", Code: 500}.Wrap(err)
		}
	}
	return r, nil
}

func (c Content) Scan(ctx context.Context, rows Scannable) (Content, error) {
	err := rows.Scan(&c.Model.ID, &c.Model.CreatedAt, &c.Model.UpdatedAt, &c.Model.ContentType, &c.Content)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c ShallowContent) Scan(ctx context.Context, rows Scannable) (ShallowContent, error) {
	err := rows.Scan(&c.ShallowModel.ID, &c.ShallowModel.CreatedAt, &c.ShallowModel.UpdatedAt, &c.ShallowModel.ContentType, &c.Content)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c Content) ShallowScan(ctx context.Context, rows Scannable) (ShallowContent, error) {
	sc := ShallowContent{}
	err := rows.Scan(&sc.ShallowModel.ID, &sc.ShallowModel.CreatedAt, &sc.ShallowModel.UpdatedAt, &sc.ShallowModel.ContentType, &sc.Content)
	if err != nil {
		return sc, err
	}
	return sc, nil
}

func (c Content) Values() []any {
	r := make([]any, 0)
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
	c.Model.ID = uuid.New().String()
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
	c.Model.ContentType = ct
	c.Model.Columns = "id, created_at, updated_at, content_type, content"
	c.Model.Values = "$1::uuid, $2, $3, $4, $5::jsonb"
	c.Model.Conflict = "DO UPDATE SET updated_at = $3, content = $5::jsonb"
	return c
} 
