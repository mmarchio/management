package merrors

import (
	"database/sql"
	"fmt"

	"github.com/mmarchio/management/logger"
)

type Merror struct {
	Err      error
	Info     string
	Package  string
	Struct   string
	Function string
	Wrapped  error
	Code     ErrorCode
	CalledBy string
	DB 		*sql.DB
}

func (c Merror) DBCloseAll(db *sql.DB) {
	if c.DB != nil {
        c.DB.Close()
    }
}

type ErrorCode int16

//Echo Errors
type EchoBindError Merror

func GetLogger(severity int) logger.LoggingContext {
	r := logger.LoggingContext{Depth: 3, Severity: severity}
	r.Init()
	return r
}

func (c EchoBindError) New(s string, vars ...any) EchoBindError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c EchoBindError) Wrap(err error) EchoBindError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("EchoBindError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c EchoBindError) Log() EchoBindError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c DBConnectionError) New(s string, vars ...any) DBConnectionError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c DBConnectionError) Wrap(err error) DBConnectionError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBConnectionError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c DBConnectionError) Log() DBConnectionError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c DBConnectionError) Error() string {
	return c.Err.Error()
}

func (c DBConnectionError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DBQueryError Merror

func (c DBQueryError) New(s string, vars ...any) DBQueryError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c DBQueryError) Wrap(err error) DBQueryError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBConnectionError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c DBQueryError) Log() DBQueryError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c DBQueryError) Error() string {
	return c.Err.Error()
}

func (c DBQueryError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DBContentScanError Merror

func (c DBContentScanError) New(s string, vars ...any) DBContentScanError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c DBContentScanError) Wrap(err error) DBContentScanError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBContentScanError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c DBContentScanError) Log() DBContentScanError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c DBContentScanError) Error() string {
	return c.Err.Error()
}

func (c DBContentScanError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DBTransactionCommitError Merror

func (c DBTransactionCommitError) New(s string, vars ...any) DBTransactionCommitError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c DBTransactionCommitError) Wrap(err error) DBTransactionCommitError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBTransactionCommitError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c DBTransactionCommitError) Log() DBTransactionCommitError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c DBTransactionCommitError) Error() string {
	return c.Err.Error()
}

func (c DBTransactionCommitError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type SQLDeleteErorr Merror

func (c SQLDeleteErorr) New(s string, vars ...any) SQLDeleteErorr {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c SQLDeleteErorr) Wrap(err error) SQLDeleteErorr {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SQLDeleteErorr", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c SQLDeleteErorr) Log() SQLDeleteErorr {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c SQLDeleteErorr) Error() string {
	return c.Err.Error()
}

func (c SQLDeleteErorr) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type SQLQueryError Merror

func (c SQLQueryError) New(s string, vars ...any) SQLQueryError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c SQLQueryError) Wrap(err error) SQLQueryError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SQLQueryError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c SQLQueryError) Log() SQLQueryError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c SQLQueryError) Error() string {
	return c.Err.Error()
}

func (c SQLQueryError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type TransactionCommitError Merror

func (c TransactionCommitError) New(s string, vars ...any) TransactionCommitError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c TransactionCommitError) Wrap(err error) TransactionCommitError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("TransactionCommitError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c TransactionCommitError) Log() TransactionCommitError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c TransactionCommitError) Error() string {
	return c.Err.Error()
}

func (c TransactionCommitError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DBPrepareStatementError Merror

func (c DBPrepareStatementError) New(s string, vars ...any) DBPrepareStatementError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c DBPrepareStatementError) Wrap(err error) DBPrepareStatementError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBPrepareStatementError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c DBPrepareStatementError) Log() DBPrepareStatementError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c DBPrepareStatementError) Error() string {
	return c.Err.Error()
}

func (c DBPrepareStatementError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DBStatementQueryQueryError Merror

func (c DBStatementQueryQueryError) New(s string, vars ...any) DBStatementQueryQueryError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c DBStatementQueryQueryError) Wrap(err error) DBStatementQueryQueryError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DBStatementQueryQueryError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c DBStatementQueryQueryError) Log() DBStatementQueryQueryError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c JSONUnmarshallingError) New(s string, vars ...any) JSONUnmarshallingError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c JSONUnmarshallingError) Wrap(err error) JSONUnmarshallingError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JSONUnmarshallingError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c JSONUnmarshallingError) Log() JSONUnmarshallingError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c JSONUnmarshallingError) Error() string {
	return c.Err.Error()
}

func (c JSONUnmarshallingError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JSONMarshallingError Merror

func (c JSONMarshallingError) New(s string, vars ...any) JSONMarshallingError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c JSONMarshallingError) Wrap(err error) JSONMarshallingError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JSONMarshallingError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c JSONMarshallingError) Log() JSONMarshallingError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c IDSetError) New(s string, vars ...any) IDSetError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c IDSetError) Wrap(err error) IDSetError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("IDSetError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c IDSetError) Log() IDSetError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c ContextSetError) New(s string, vars ...any) ContextSetError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContextSetError) Wrap(err error) ContextSetError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContextSetError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContextSetError) Log() ContextSetError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c ContextSetError) Error() string {
	return c.Err.Error()
}

func (c ContextSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContextGetError Merror

func (c ContextGetError) New(s string, vars ...any) ContextGetError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContextGetError) Wrap(err error) ContextGetError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContextGetError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContextGetError) Log() ContextGetError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c ContentGetError) New(s string, vars ...any) ContentGetError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentGetError) Wrap(err error) ContentGetError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentGetError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentGetError) Log() ContentGetError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c ContentGetError) Error() string {
	return c.Err.Error()
}

func (c ContentGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentSetError Merror

func (c ContentSetError) New(s string, vars ...any) ContentSetError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentSetError) Wrap(err error) ContentSetError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentSetError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentSetError) Log() ContentSetError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c ContentSetError) Error() string {
	return c.Err.Error()
}

func (c ContentSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentListError Merror

func (c ContentListError) New(s string, vars ...any) ContentListError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentListError) Wrap(err error) ContentListError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentListError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentListError) Log() ContentListError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c ContentListError) Error() string {
	return c.Err.Error()
}

func (c ContentListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentListByError Merror

func (c ContentListByError) New(s string, vars ...any) ContentListByError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentListByError) Wrap(err error) ContentListByError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentListByError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentListByError) Log() ContentListByError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c ContentListByError) Error() string {
	return c.Err.Error()
}

func (c ContentListByError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentModelDeleteError Merror

func (c ContentModelDeleteError) New(s string, vars ...any) ContentModelDeleteError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentModelDeleteError) Wrap(err error) ContentModelDeleteError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentModelDeleteError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentModelDeleteError) Log() ContentModelDeleteError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c ContentModelDeleteError) Error() string {
	return c.Err.Error()
}

func (c ContentModelDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentFindByError Merror

func (c ContentFindByError) New(s string, vars ...any) ContentFindByError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentFindByError) Wrap(err error) ContentFindByError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentFindByError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentFindByError) Log() ContentFindByError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c ContentFindByError) Error() string {
	return c.Err.Error()
}

func (c ContentFindByError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentValidationError Merror 

func (c ContentValidationError) New(s string, vars ...any) ContentValidationError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentValidationError) Wrap(err error) ContentValidationError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentValidationError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentValidationError) Log() ContentValidationError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c ContentValidationError) Error() string {
	return c.Err.Error()
}

func (c ContentValidationError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type NilContentError Merror

func (c NilContentError) New(s string, vars ...any) NilContentError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c NilContentError) Wrap(err error) NilContentError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("NilContentError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c NilContentError) Log() NilContentError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c ContentCustomQueryError) New(s string, vars ...any) ContentCustomQueryError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentCustomQueryError) Wrap(err error) ContentCustomQueryError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentCustomQueryError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentCustomQueryError) Log() ContentCustomQueryError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c ContentDeleteError) New(s string, vars ...any) ContentDeleteError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentDeleteError) Wrap(err error) ContentDeleteError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentDeleteError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentDeleteError) Log() ContentDeleteError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c MSIConversionError) New(s string, vars ...any) MSIConversionError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c MSIConversionError) Wrap(err error) MSIConversionError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("MSIConversionError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c MSIConversionError) Log() MSIConversionError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c HTTPRequestError) New(s string, vars ...any) HTTPRequestError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c HTTPRequestError) Wrap(err error) HTTPRequestError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("HTTPRequestError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c HTTPRequestError) Log() HTTPRequestError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c SetContextError) New(s string, vars ...any) SetContextError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c SetContextError) Wrap(err error) SetContextError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SetContextError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c SetContextError) Log() SetContextError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c ContentToTypeError) New(s string, vars ...any) ContentToTypeError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeError) Wrap(err error) ContentToTypeError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeError) Log() ContentToTypeError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c ContentToTypeGetError) New(s string, vars ...any) ContentToTypeGetError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeGetError) Wrap(err error) ContentToTypeGetError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeGetError) Log() ContentToTypeGetError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c ContentToTypeSetError) New(s string, vars ...any) ContentToTypeSetError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeSetError) Wrap(err error) ContentToTypeSetError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeSetError) Log() ContentToTypeSetError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c ContentToTypeListError) New(s string, vars ...any) ContentToTypeListError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeListError) Wrap(err error) ContentToTypeListError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeListError) Log() ContentToTypeListError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c ContentToTypeListByError) New(s string, vars ...any) ContentToTypeListByError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeListByError) Wrap(err error) ContentToTypeListByError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeListByError) Log() ContentToTypeListByError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c ContentToTypeFindByError) New(s string, vars ...any) ContentToTypeFindByError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentToTypeFindByError) Wrap(err error) ContentToTypeFindByError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentToTypeError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentToTypeFindByError) Log() ContentToTypeFindByError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

func (c JPATHError) New(s string, vars ...any) JPATHError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c JPATHError) Wrap(err error) JPATHError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JPATHError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c JPATHError) Log() JPATHError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
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

type ParseContent Merror

func (c ParseContent) New(s string, vars ...any) ParseContent {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ParseContent) Wrap(err error) ParseContent {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ParseContent", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ParseContent) Log() ParseContent {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c ParseContent) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c ParseContent) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c ParseContent) GetCode() ErrorCode {
	return c.Code
}

func (c ParseContent) BubbleCode() ParseContent {
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

type WebsocketDialError Merror

func (c WebsocketDialError) New(s string, vars ...any) WebsocketDialError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c WebsocketDialError) Wrap(err error) WebsocketDialError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("WebsocketDialError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c WebsocketDialError) Log() WebsocketDialError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c WebsocketDialError) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c WebsocketDialError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c WebsocketDialError) GetCode() ErrorCode {
	return c.Code
}

func (c WebsocketDialError) BubbleCode() WebsocketDialError {
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

type WebsocketWriteError Merror

func (c WebsocketWriteError) New(s string, vars ...any) WebsocketWriteError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c WebsocketWriteError) Wrap(err error) WebsocketWriteError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("WebsocketWriteError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c WebsocketWriteError) Log() WebsocketWriteError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c WebsocketWriteError) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c WebsocketWriteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c WebsocketWriteError) GetCode() ErrorCode {
	return c.Code
}

func (c WebsocketWriteError) BubbleCode() WebsocketWriteError {
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

type WebsocketReadError Merror

func (c WebsocketReadError) New(s string, vars ...any) WebsocketReadError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c WebsocketReadError) Wrap(err error) WebsocketReadError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("WebsocketReadError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c WebsocketReadError) Log() WebsocketReadError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c WebsocketReadError) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c WebsocketReadError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c WebsocketReadError) GetCode() ErrorCode {
	return c.Code
}

func (c WebsocketReadError) BubbleCode() WebsocketReadError {
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

type ContentCheckError Merror

func (c ContentCheckError) New(s string, vars ...any) ContentCheckError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c ContentCheckError) Wrap(err error) ContentCheckError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ContentCheckError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c ContentCheckError) Log() ContentCheckError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c ContentCheckError) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c ContentCheckError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c ContentCheckError) GetCode() ErrorCode {
	return c.Code
}

func (c ContentCheckError) BubbleCode() ContentCheckError {
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

type NodeExecError Merror

func (c NodeExecError) New(s string, vars ...any) NodeExecError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c NodeExecError) Wrap(err error) NodeExecError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("NodeExecError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c NodeExecError) Log() NodeExecError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c NodeExecError) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c NodeExecError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c NodeExecError) GetCode() ErrorCode {
	return c.Code
}

func (c NodeExecError) BubbleCode() NodeExecError {
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

type CurlError Merror

func (c CurlError) New(s string, vars ...any) CurlError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c CurlError) Wrap(err error) CurlError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("CurlError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c CurlError) Log() CurlError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c CurlError) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c CurlError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c CurlError) GetCode() ErrorCode {
	return c.Code
}

func (c CurlError) BubbleCode() CurlError {
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

type SSHError Merror

func (c SSHError) New(s string, vars ...any) SSHError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c SSHError) Wrap(err error) SSHError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SSHError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c SSHError) Log() SSHError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c SSHError) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c SSHError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c SSHError) GetCode() ErrorCode {
	return c.Code
}

func (c SSHError) BubbleCode() SSHError {
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

type StepValidationError Merror

func (c StepValidationError) New(s string, vars ...any) StepValidationError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c StepValidationError) Wrap(err error) StepValidationError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("StepValidationError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c StepValidationError) Log() StepValidationError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c StepValidationError) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c StepValidationError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c StepValidationError) GetCode() ErrorCode {
	return c.Code
}

func (c StepValidationError) BubbleCode() StepValidationError {
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

type GeneralError Merror

func (c GeneralError) New(s string, vars ...any) GeneralError {
	c = c.Wrap(fmt.Errorf(s, vars...))

	return c
}

func (c GeneralError) Wrap(err error) GeneralError {
	if c.DB != nil {
        c.DB.Close()
    }
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("StepValidationError", c.Info, c.Package, c.Struct, c.Function, c.CalledBy, c.Err), c.Wrapped)
	return c
}

func (c GeneralError) Log() GeneralError {
	GetLogger(1).Flogger("err: %s", c.Err.Error())
	return c
}

func (c GeneralError) Error() string {
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}

func (c GeneralError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c GeneralError) GetCode() ErrorCode {
	return c.Code
}

func (c GeneralError) BubbleCode() GeneralError {
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

func ErrString(Type, Info, Package, Struct, Function, CalledBy string, Err error) string {
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
	if CalledBy != "" {
		errString += fmt.Sprintf("CalledBy: %s\n", CalledBy)
	}
	return errString
}
