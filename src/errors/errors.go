package merrors

import (
	"database/sql"
	"fmt"
)

type Merror struct {
	Err      error
	Info     string
	Package  string
	Struct   string
	Function string
	Wrapped  error
	Code     ErrorCode
}

func (c Merror) DBCloseAll(db *sql.DB) {
	if db != nil {
        db.Close()
    }
}

type ErrorCode int16

//Echo Errors
type EchoBindError Merror

func (c EchoBindError) New(db *sql.DB, s string, vars ...any) EchoBindError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c EchoBindError) Wrap(db *sql.DB, err error) EchoBindError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("EchoBindError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c EchoBindError) Error() string {
	return c.Err.Error()
}

func (c EchoBindError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

//DB Errors
type DBConnectionError Merror

func (c DBConnectionError) New(db *sql.DB, s string, vars ...any) DBConnectionError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c DBConnectionError) Wrap(db *sql.DB, err error) DBConnectionError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBConnectionError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DBConnectionError) Error() string {
	return c.Err.Error()
}

func (c DBConnectionError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DBQueryError Merror

func (c DBQueryError) New(db *sql.DB, s string, vars ...any) DBQueryError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c DBQueryError) Wrap(db *sql.DB, err error) DBQueryError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBConnectionError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DBQueryError) Error() string {
	return c.Err.Error()
}

func (c DBQueryError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DBContentScanError Merror

func (c DBContentScanError) New(db *sql.DB, s string, vars ...any) DBContentScanError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c DBContentScanError) Wrap(db *sql.DB, err error) DBContentScanError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBContentScanError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DBContentScanError) Error() string {
	return c.Err.Error()
}

func (c DBContentScanError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DBTransactionCommitError Merror

func (c DBTransactionCommitError) New(db *sql.DB, s string, vars ...any) DBTransactionCommitError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c DBTransactionCommitError) Wrap(db *sql.DB, err error) DBTransactionCommitError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBTransactionCommitError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DBTransactionCommitError) Error() string {
	return c.Err.Error()
}

func (c DBTransactionCommitError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type SQLDeleteErorr Merror

func (c SQLDeleteErorr) New(db *sql.DB, s string, vars ...any) SQLDeleteErorr {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c SQLDeleteErorr) Wrap(db *sql.DB, err error) SQLDeleteErorr {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SQLDeleteErorr", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c SQLDeleteErorr) Error() string {
	return c.Err.Error()
}

func (c SQLDeleteErorr) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type SQLQueryError Merror

func (c SQLQueryError) New(db *sql.DB, s string, vars ...any) SQLQueryError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c SQLQueryError) Wrap(db *sql.DB, err error) SQLQueryError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SQLQueryError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c SQLQueryError) Error() string {
	return c.Err.Error()
}

func (c SQLQueryError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type TransactionCommitError Merror

func (c TransactionCommitError) New(db *sql.DB, s string, vars ...any) TransactionCommitError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c TransactionCommitError) Wrap(db *sql.DB, err error) TransactionCommitError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("TransactionCommitError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c TransactionCommitError) Error() string {
	return c.Err.Error()
}

func (c TransactionCommitError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DBPrepareStatementError Merror

func (c DBPrepareStatementError) New(db *sql.DB, s string, vars ...any) DBPrepareStatementError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c DBPrepareStatementError) Wrap(db *sql.DB, err error) DBPrepareStatementError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBPrepareStatementError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DBPrepareStatementError) Error() string {
	return c.Err.Error()
}

func (c DBPrepareStatementError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DBStatementQueryQueryError Merror

func (c DBStatementQueryQueryError) New(db *sql.DB, s string, vars ...any) DBStatementQueryQueryError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c DBStatementQueryQueryError) Wrap(db *sql.DB, err error) DBStatementQueryQueryError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBStatementQueryQueryError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DBStatementQueryQueryError) Error() string {
	return c.Err.Error()
}

func (c DBStatementQueryQueryError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

//JSON Errors
type JSONUnmarshallingError Merror

func (c JSONUnmarshallingError) New(db *sql.DB, s string, vars ...any) JSONUnmarshallingError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c JSONUnmarshallingError) Wrap(db *sql.DB, err error) JSONUnmarshallingError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JSONUnmarshallingError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JSONUnmarshallingError) Error() string {
	return c.Err.Error()
}

func (c JSONUnmarshallingError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JSONMarshallingError Merror

func (c JSONMarshallingError) New(db *sql.DB, s string, vars ...any) JSONMarshallingError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c JSONMarshallingError) Wrap(db *sql.DB, err error) JSONMarshallingError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JSONMarshallingError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JSONMarshallingError) Error() string {
	return c.Err.Error()
}

func (c JSONMarshallingError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

//ID errors
type IDSetError Merror

func (c IDSetError) New(db *sql.DB, s string, vars ...any) IDSetError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c IDSetError) Wrap(db *sql.DB, err error) IDSetError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("IDSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c IDSetError) Error() string {
	return c.Err.Error()
}

func (c IDSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

//Context Error
type ContextSetError Merror

func (c ContextSetError) New(db *sql.DB, s string, vars ...any) ContextSetError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContextSetError) Wrap(db *sql.DB, err error) ContextSetError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContextSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContextSetError) Error() string {
	return c.Err.Error()
}

func (c ContextSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContextGetError Merror

func (c ContextGetError) New(db *sql.DB, s string, vars ...any) ContextGetError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContextGetError) Wrap(db *sql.DB, err error) ContextGetError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContextGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContextGetError) Error() string {
	return c.Err.Error()
}

func (c ContextGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}


//Content Errors
type ContentGetError Merror

func (c ContentGetError) New(db *sql.DB, s string, vars ...any) ContentGetError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentGetError) Wrap(db *sql.DB, err error) ContentGetError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentGetError) Error() string {
	return c.Err.Error()
}

func (c ContentGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentSetError Merror

func (c ContentSetError) New(db *sql.DB, s string, vars ...any) ContentSetError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentSetError) Wrap(db *sql.DB, err error) ContentSetError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentSetError) Error() string {
	return c.Err.Error()
}

func (c ContentSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentListError Merror

func (c ContentListError) New(db *sql.DB, s string, vars ...any) ContentListError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentListError) Wrap(db *sql.DB, err error) ContentListError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentListError) Error() string {
	return c.Err.Error()
}

func (c ContentListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentListByError Merror

func (c ContentListByError) New(db *sql.DB, s string, vars ...any) ContentListByError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentListByError) Wrap(db *sql.DB, err error) ContentListByError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentListByError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentListByError) Error() string {
	return c.Err.Error()
}

func (c ContentListByError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentModelDeleteError Merror

func (c ContentModelDeleteError) New(db *sql.DB, s string, vars ...any) ContentModelDeleteError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentModelDeleteError) Wrap(db *sql.DB, err error) ContentModelDeleteError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentModelDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentModelDeleteError) Error() string {
	return c.Err.Error()
}

func (c ContentModelDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentFindByError Merror

func (c ContentFindByError) New(db *sql.DB, s string, vars ...any) ContentFindByError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentFindByError) Wrap(db *sql.DB, err error) ContentFindByError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentFindByError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentFindByError) Error() string {
	return c.Err.Error()
}

func (c ContentFindByError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentValidationError Merror 

func (c ContentValidationError) New(db *sql.DB, s string, vars ...any) ContentValidationError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentValidationError) Wrap(db *sql.DB, err error) ContentValidationError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentValidationError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentValidationError) Error() string {
	return c.Err.Error()
}

func (c ContentValidationError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type NilContentError Merror

func (c NilContentError) New(db *sql.DB, s string, vars ...any) NilContentError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c NilContentError) Wrap(db *sql.DB, err error) NilContentError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("NilContentError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c NilContentError) Error() string {
	return c.Err.Error()
}

func (c NilContentError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c NilContentError) GetCode() ErrorCode {
	return c.Code
}

func (c NilContentError) BubbleCode() NilContentError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type ContentCustomQueryError Merror

func (c ContentCustomQueryError) New(db *sql.DB, s string, vars ...any) ContentCustomQueryError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentCustomQueryError) Wrap(db *sql.DB, err error) ContentCustomQueryError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentCustomQueryError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentCustomQueryError) Error() string {
	return c.Err.Error()
}

func (c ContentCustomQueryError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c ContentCustomQueryError) GetCode() ErrorCode {
	return c.Code
}

func (c ContentCustomQueryError) BubbleCode() ContentCustomQueryError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type ContentDeleteError Merror

func (c ContentDeleteError) New(db *sql.DB, s string, vars ...any) ContentDeleteError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentDeleteError) Wrap(db *sql.DB, err error) ContentDeleteError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentDeleteError) Error() string {
	return c.Err.Error()
}

func (c ContentDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c ContentDeleteError) GetCode() ErrorCode {
	return c.Code
}

func (c ContentDeleteError) BubbleCode() ContentDeleteError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type MSIConversionError Merror

func (c MSIConversionError) New(db *sql.DB, s string, vars ...any) MSIConversionError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c MSIConversionError) Wrap(db *sql.DB, err error) MSIConversionError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("MSIConversionError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c MSIConversionError) Error() string {
	return c.Err.Error()
}

func (c MSIConversionError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c MSIConversionError) GetCode() ErrorCode {
	return c.Code
}

func (c MSIConversionError) BubbleCode() MSIConversionError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type HTTPRequestError Merror

func (c HTTPRequestError) New(db *sql.DB, s string, vars ...any) HTTPRequestError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c HTTPRequestError) Wrap(db *sql.DB, err error) HTTPRequestError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("HTTPRequestError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c HTTPRequestError) Error() string {
	return c.Err.Error()
}

func (c HTTPRequestError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c HTTPRequestError) GetCode() ErrorCode {
	return c.Code
}

func (c HTTPRequestError) BubbleCode() HTTPRequestError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type SetContextError Merror

func (c SetContextError) New(db *sql.DB, s string, vars ...any) SetContextError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c SetContextError) Wrap(db *sql.DB, err error) SetContextError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SetContextError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c SetContextError) Error() string {
	return c.Err.Error()
}

func (c SetContextError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c SetContextError) GetCode() ErrorCode {
	return c.Code
}

func (c SetContextError) BubbleCode() SetContextError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type ContentToTypeError Merror 

func (c ContentToTypeError) New(db *sql.DB, s string, vars ...any) ContentToTypeError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeError) Wrap(db *sql.DB, err error) ContentToTypeError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeError) Error() string {
	return c.Err.Error()
}

func (c ContentToTypeError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c ContentToTypeError) GetCode() ErrorCode {
	return c.Code
}

func (c ContentToTypeError) BubbleCode() ContentToTypeError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type ContentToTypeGetError Merror

func (c ContentToTypeGetError) New(db *sql.DB, s string, vars ...any) ContentToTypeGetError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeGetError) Wrap(db *sql.DB, err error) ContentToTypeGetError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeGetError) Error() string {
	return c.Err.Error()
}

func (c ContentToTypeGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c ContentToTypeGetError) GetCode() ErrorCode {
	return c.Code
}

func (c ContentToTypeGetError) BubbleCode() ContentToTypeGetError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type ContentToTypeSetError Merror 

func (c ContentToTypeSetError) New(db *sql.DB, s string, vars ...any) ContentToTypeSetError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeSetError) Wrap(db *sql.DB, err error) ContentToTypeSetError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeSetError) Error() string {
	return c.Err.Error()
}

func (c ContentToTypeSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c ContentToTypeSetError) GetCode() ErrorCode {
	return c.Code
}

func (c ContentToTypeSetError) BubbleCode() ContentToTypeSetError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type ContentToTypeListError Merror

func (c ContentToTypeListError) New(db *sql.DB, s string, vars ...any) ContentToTypeListError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeListError) Wrap(db *sql.DB, err error) ContentToTypeListError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeListError) Error() string {
	return c.Err.Error()
}

func (c ContentToTypeListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c ContentToTypeListError) GetCode() ErrorCode {
	return c.Code
}

func (c ContentToTypeListError) BubbleCode() ContentToTypeListError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type ContentToTypeListByError Merror

func (c ContentToTypeListByError) New(db *sql.DB, s string, vars ...any) ContentToTypeListByError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeListByError) Wrap(db *sql.DB, err error) ContentToTypeListByError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeListByError) Error() string {
	return c.Err.Error()
}

func (c ContentToTypeListByError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c ContentToTypeListByError) GetCode() ErrorCode {
	return c.Code
}

func (c ContentToTypeListByError) BubbleCode() ContentToTypeListByError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type ContentToTypeFindByError Merror

func (c ContentToTypeFindByError) New(db *sql.DB, s string, vars ...any) ContentToTypeFindByError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeFindByError) Wrap(db *sql.DB, err error) ContentToTypeFindByError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeFindByError) Error() string {
	return c.Err.Error()
}

func (c ContentToTypeFindByError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c ContentToTypeFindByError) GetCode() ErrorCode {
	return c.Code
}

func (c ContentToTypeFindByError) BubbleCode() ContentToTypeFindByError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type JPATHError Merror

func (c JPATHError) New(db *sql.DB, s string, vars ...any) JPATHError {
	c = c.Wrap(db, fmt.Errorf(s, vars...))

	return c
}

func (c JPATHError) Wrap(db *sql.DB, err error) JPATHError {
	if db != nil {
        db.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JPATHError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JPATHError) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c JPATHError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c JPATHError) GetCode() ErrorCode {
	return c.Code
}

func (c JPATHError) BubbleCode() JPATHError {
	if c.Code == 0 {
		c.Code = 500
	}
	if e, ok := c.Err.(WrappedError); ok {
		if e.GetCode() != 500 {
			c.Code = e.GetCode()
		}
	}
	return c
}

type WrappedError interface {
	Wrap(*sql.DB, error)
	ErrorCode(int16)
	GetCode() ErrorCode
	Error() string
}

func ErrString(Type, Info, Package, Struct, Function string, Err error) string {
	errString := fmt.Sprintf("%s\n", Type)
	errString += fmt.Sprintf("Info: %s\n", Info)
	if Package != "" {
		errString += fmt.Sprintf("Package: %s\n", Package)
	}
	if Struct != "" {
		errString += fmt.Sprintf("Struct: %s\n", Struct)
	}
	if Function != "" {
		errString += fmt.Sprintf("Function: %s\n", Function)
	}
	return errString
}
