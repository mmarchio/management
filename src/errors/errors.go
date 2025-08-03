package merrors

import "fmt"

type Merror struct{
	Err error
	Info string
	Package string
	Struct string
	Function string
	Wrapped error
	Code ErrorCode
}

type ErrorCode int16

type DBConnectionError Merror

func (c DBConnectionError) Wrap(err error) DBConnectionError {
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

func (c DBQueryError) Wrap(err error) DBQueryError {
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

func (c DBContentScanError) Wrap(err error) DBContentScanError {
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

func (c DBTransactionCommitError) Wrap(err error) DBTransactionCommitError {
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

type JobRunGetError Merror

func (c JobRunGetError) Wrap(err error) JobRunGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobRunGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobRunGetError) Error() string {
	return c.Err.Error()
}

func (c JobRunGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobRunSetError Merror

func (c JobRunSetError) Wrap(err error) JobRunSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobRunSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobRunSetError) Error() string {
	return c.Err.Error()
}

func (c JobRunSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobStatusSetError Merror

func (c JobStatusSetError) Wrap(err error) JobStatusSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobStatusSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobStatusSetError) Error() string {
	return c.Err.Error()
}

func (c JobStatusSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobStatusGetError Merror

func (c JobStatusGetError) Wrap(err error) JobStatusGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobStatusGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobStatusGetError) Error() string {
	return c.Err.Error()
}

func (c JobStatusGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobStatusListError Merror

func (c JobStatusListError) Wrap(err error) JobStatusListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobStatusListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobStatusListError) Error() string {
	return c.Err.Error()
}

func (c JobStatusListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContextSetError Merror

func (c ContextSetError) Wrap(err error) ContextSetError {
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

func (c ContextGetError) Wrap(err error) ContextGetError {
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

type JSONUnmarshallingError Merror

func (c JSONUnmarshallingError) Wrap(err error) JSONUnmarshallingError {
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

func (c JSONMarshallingError) Wrap(err error) JSONMarshallingError {
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

type ComfyUITemplateDeleteError Merror

func (c ComfyUITemplateDeleteError) Wrap(err error) ComfyUITemplateDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ComfyUITemplateDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ComfyUITemplateDeleteError) Error() string {
	return c.Err.Error()
}

func (c ComfyUITemplateDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type PromptSetError Merror

func (c PromptSetError) Wrap(err error) PromptSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptSetError) Error() string {
	return c.Err.Error()
}

func (c PromptSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type PromptGetError Merror

func (c PromptGetError) Wrap(err error) PromptGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptGetError) Error() string {
	return c.Err.Error()
}

func (c PromptGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type PromptListError Merror

func (c PromptListError) Wrap(err error) PromptListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptListError) Error() string {
	return c.Err.Error()
}

func (c PromptListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DispositionGetError Merror

func (c DispositionGetError) Wrap(err error) DispositionGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DispositionGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DispositionGetError) Error() string {
	return c.Err.Error()
}

func (c DispositionGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DispositionSetError Merror

func (c DispositionSetError) Wrap(err error) DispositionSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DispositionSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DispositionSetError) Error() string {
	return c.Err.Error()
}

func (c DispositionSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DispositionListError Merror

func (c DispositionListError) Wrap(err error) DispositionListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DispositionListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DispositionListError) Error() string {
	return c.Err.Error()
}

func (c DispositionListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ComfyUITemplateGetError Merror

func (c ComfyUITemplateGetError) Wrap(err error) ComfyUITemplateGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ComfyUITemplateGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ComfyUITemplateGetError) Error() string {
	return c.Err.Error()
}

func (c ComfyUITemplateGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ComfyUITemplateSetError Merror

func (c ComfyUITemplateSetError) Wrap(err error) ComfyUITemplateSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ComfyUITemplateSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ComfyUITemplateSetError) Error() string {
	return c.Err.Error()
}

func (c ComfyUITemplateSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ComfyUITemplateListError Merror

func (c ComfyUITemplateListError) Wrap(err error) ComfyUITemplateListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("ComfyUITemplateListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c ComfyUITemplateListError) Error() string {
	return c.Err.Error()
}

func (c ComfyUITemplateListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobListError Merror

func (c JobListError) Wrap(err error) JobListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobListError) Error() string {
	return c.Err.Error()
}

func (c JobListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobGetError Merror

func (c JobGetError) Wrap(err error) JobGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobGetError) Error() string {
	return c.Err.Error()
}

func (c JobGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobSetError Merror

func (c JobSetError) Wrap(err error) JobSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobSetError) Error() string {
	return c.Err.Error()
}

func (c JobSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobRunListError Merror

func (c JobRunListError) Wrap(err error) JobRunListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobRunListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobRunListError) Error() string {
	return c.Err.Error()
}

func (c JobRunListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type EchoBindError Merror

func (c EchoBindError) Wrap(err error) EchoBindError {
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

type PromptBindError Merror

func (c PromptBindError) Wrap(err error) PromptBindError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptBindError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptBindError) Error() string {
	return c.Err.Error()
}

func (c PromptBindError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type SQLDeleteErorr Merror

func (c SQLDeleteErorr) Wrap(err error) SQLDeleteErorr {
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

type HandlePromptSaveError Merror

func (c HandlePromptSaveError) Wrap(err error) HandlePromptSaveError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("HandlePromptSaveError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c HandlePromptSaveError) Error() string {
	return c.Err.Error()
}

func (c HandlePromptSaveError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type PromptModelScanError Merror

func (c PromptModelScanError) Wrap(err error) PromptModelScanError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptModelScanError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptModelScanError) Error() string {
	return c.Err.Error()
}

func (c PromptModelScanError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type Base64DecodingError Merror

func (c Base64DecodingError) Wrap(err error) Base64DecodingError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("Base64DecodingError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c Base64DecodingError) Error() string {
	return c.Err.Error()
}

func (c Base64DecodingError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DispositionBase64DecodingError Merror

func (c DispositionBase64DecodingError) Wrap(err error) DispositionBase64DecodingError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DispositionBase64DecodingError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DispositionBase64DecodingError) Error() string {
	return c.Err.Error()
}

func (c DispositionBase64DecodingError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DispositionModelScanError Merror

func (c DispositionModelScanError) Wrap(err error) DispositionModelScanError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DispositionModelScanError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DispositionModelScanError) Error() string {
	return c.Err.Error()
}

func (c DispositionModelScanError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DispositionNewEntitlementsError Merror

func (c DispositionNewEntitlementsError) Wrap(err error) DispositionNewEntitlementsError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DispositionNewEntitlementsError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DispositionNewEntitlementsError) Error() string {
	return c.Err.Error()
}

func (c DispositionNewEntitlementsError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type DispositionNewStepsError Merror

func (c DispositionNewStepsError) Wrap(err error) DispositionNewStepsError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DispositionNewStepsError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DispositionNewStepsError) Error() string {
	return c.Err.Error()
}

func (c DispositionNewStepsError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type SQLQueryError Merror

func (c SQLQueryError) Wrap(err error) SQLQueryError {
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

func (c TransactionCommitError) Wrap(err error) TransactionCommitError {
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

type PromptUnpackError Merror

func (c PromptUnpackError) Wrap(err error) PromptUnpackError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptUnpackError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptUnpackError) Error() string {
	return c.Err.Error()
}

func (c PromptUnpackError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type Base64EncodingError Merror

func (c Base64EncodingError) Wrap(err error) Base64EncodingError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("Base64EncodingError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c Base64EncodingError) Error() string {
	return c.Err.Error()
}

func (c Base64EncodingError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type NestedPackError Merror

func (c NestedPackError) Wrap(err error) NestedPackError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("NestedPackError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c NestedPackError) Error() string {
	return c.Err.Error()
}

func (c NestedPackError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentGetError Merror

func (c ContentGetError) Wrap(err error) ContentGetError {
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

func (c ContentSetError) Wrap(err error) ContentSetError {
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

func (c ContentListError) Wrap(err error) ContentListError {
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

func (c ContentListByError) Wrap(err error) ContentListByError {
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

type JobStatusError Merror

func (c JobStatusError) Wrap(err error) JobStatusError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobStatusError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobStatusError) Error() string {
	return c.Err.Error()
}

func (c JobStatusError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type IDSetError Merror

func (c IDSetError) Wrap(err error) IDSetError {
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

type DBPrepareStatementError Merror

func (c DBPrepareStatementError) Wrap(err error) DBPrepareStatementError {
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

func (c DBStatementQueryQueryError) Wrap(err error) DBStatementQueryQueryError {
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

type SystemPromptSetError Merror

func (c SystemPromptSetError) Wrap(err error) SystemPromptSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SystemPromptSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c SystemPromptSetError) Error() string {
	return c.Err.Error()
}

func (c SystemPromptSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type SystemPromptGetError Merror

func (c SystemPromptGetError) Wrap(err error) SystemPromptGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SystemPromptGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c SystemPromptGetError) Error() string {
	return c.Err.Error()
}

func (c SystemPromptGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type SystemPromptListError Merror

func (c SystemPromptListError) Wrap(err error) SystemPromptListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SystemPromptListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c SystemPromptListError) Error() string {
	return c.Err.Error()
}

func (c SystemPromptListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentModelDeleteError Merror

func (c ContentModelDeleteError) Wrap(err error) ContentModelDeleteError {
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

type DispositionDeleteError Merror

func (c DispositionDeleteError) Wrap(err error) DispositionDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("DispositionDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c DispositionDeleteError) Error() string {
	return c.Err.Error()
}

func (c DispositionDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobDeleteError Merror

func (c JobDeleteError) Wrap(err error) JobDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobDeleteError) Error() string {
	return c.Err.Error()
}

func (c JobDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobRunDeleteError Merror

func (c JobRunDeleteError) Wrap(err error) JobRunDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobRunDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobRunDeleteError) Error() string {
	return c.Err.Error()
}

func (c JobRunDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type PromptDeleteError Merror

func (c PromptDeleteError) Wrap(err error) PromptDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptDeleteError) Error() string {
	return c.Err.Error()
}

func (c PromptDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type SystemPromptDeleteError Merror

func (c SystemPromptDeleteError) Wrap(err error) SystemPromptDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("SystemPromptDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c SystemPromptDeleteError) Error() string {
	return c.Err.Error()
}

func (c SystemPromptDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobStatusDeleteError Merror

func (c JobStatusDeleteError) Wrap(err error) JobStatusDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobStatusDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobStatusDeleteError) Error() string {
	return c.Err.Error()
}

func (c JobStatusDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type JobFindByError Merror


func (c JobFindByError) Wrap(err error) JobFindByError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobFindByError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobFindByError) Error() string {
	return c.Err.Error()
}

func (c JobFindByError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

type ContentFindByError Merror

func (c ContentFindByError) Wrap(err error) ContentFindByError {
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

type NilContentError Merror

func (c NilContentError) Wrap(err error) NilContentError {
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

type JobRunCustomQueryError Merror

func (c JobRunCustomQueryError) Wrap(err error) JobRunCustomQueryError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("JobRunCustomQueryError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c JobRunCustomQueryError) Error() string {
	return c.Err.Error()
}

func (c JobRunCustomQueryError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c JobRunCustomQueryError) GetCode() ErrorCode {
	return c.Code
}

func (c JobRunCustomQueryError) BubbleCode() JobRunCustomQueryError {
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

func (c ContentCustomQueryError) Wrap(err error) ContentCustomQueryError {
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

type AutomationWorkflowListError Merror

func (c AutomationWorkflowListError) Wrap(err error) AutomationWorkflowListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("AutomationWorkflowListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c AutomationWorkflowListError) Error() string {
	return c.Err.Error()
}

func (c AutomationWorkflowListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c AutomationWorkflowListError) GetCode() ErrorCode {
	return c.Code
}

func (c AutomationWorkflowListError) BubbleCode() AutomationWorkflowListError {
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

type AutomationWorkflowGetError Merror 

func (c AutomationWorkflowGetError) Wrap(err error) AutomationWorkflowGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("AutomationWorkflowGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c AutomationWorkflowGetError) Error() string {
	return c.Err.Error()
}

func (c AutomationWorkflowGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c AutomationWorkflowGetError) GetCode() ErrorCode {
	return c.Code
}

func (c AutomationWorkflowGetError) BubbleCode() AutomationWorkflowGetError {
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

type AutomationWorkflowSetError Merror

func (c AutomationWorkflowSetError) Wrap(err error) AutomationWorkflowSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("AutomationWorkflowSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c AutomationWorkflowSetError) Error() string {
	return c.Err.Error()
}

func (c AutomationWorkflowSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c AutomationWorkflowSetError) GetCode() ErrorCode {
	return c.Code
}

type AutomationWorkflowDeleteError Merror

func (c AutomationWorkflowDeleteError) Wrap(err error) AutomationWorkflowDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("AutomationWorkflowDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c AutomationWorkflowDeleteError) Error() string {
	return c.Err.Error()
}

func (c AutomationWorkflowDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c AutomationWorkflowDeleteError) GetCode() ErrorCode {
	return c.Code
}

func (c AutomationWorkflowDeleteError) BubbleCode() AutomationWorkflowDeleteError {
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

type AutomationStepDeleteError Merror 

func (c AutomationStepDeleteError) Wrap(err error) AutomationStepDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("AutomationStepDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c AutomationStepDeleteError) Error() string {
	return c.Err.Error()
}

func (c AutomationStepDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c AutomationStepDeleteError) GetCode() ErrorCode {
	return c.Code
}

func (c AutomationStepDeleteError) BubbleCode() AutomationStepDeleteError {
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

type AutomationStepSetError Merror

func (c AutomationStepSetError) Wrap(err error) AutomationStepSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("AutomationStepSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c AutomationStepSetError) Error() string {
	return c.Err.Error()
}

func (c AutomationStepSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c AutomationStepSetError) GetCode() ErrorCode {
	return c.Code
}

func (c AutomationStepSetError) BubbleCode() AutomationStepSetError {
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

type AutomationStepflowGetError Merror

func (c AutomationStepflowGetError) Wrap(err error) AutomationStepflowGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("AutomationStepflowGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c AutomationStepflowGetError) Error() string {
	return c.Err.Error()
}

func (c AutomationStepflowGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c AutomationStepflowGetError) GetCode() ErrorCode {
	return c.Code
}

func (c AutomationStepflowGetError) BubbleCode() AutomationStepflowGetError {
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

type AutomationStepListError Merror

func (c AutomationStepListError) Wrap(err error) AutomationStepListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("AutomationStepListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c AutomationStepListError) Error() string {
	return c.Err.Error()
}

func (c AutomationStepListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c AutomationStepListError) GetCode() ErrorCode {
	return c.Code
}

func (c AutomationStepListError) BubbleCode() AutomationStepListError {
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

type WorkflowListError Merror

func (c WorkflowListError) Wrap(err error) WorkflowListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("WorkflowListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c WorkflowListError) Error() string {
	return c.Err.Error()
}

func (c WorkflowListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c WorkflowListError) GetCode() ErrorCode {
	return c.Code
}

func (c WorkflowListError) BubbleCode() WorkflowListError {
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

type WorkflowGetError Merror

func (c WorkflowGetError) Wrap(err error) WorkflowGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("WorkflowGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c WorkflowGetError) Error() string {
	return c.Err.Error()
}

func (c WorkflowGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c WorkflowGetError) GetCode() ErrorCode {
	return c.Code
}

func (c WorkflowGetError) BubbleCode() WorkflowGetError {
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

type WorkflowSetError Merror

func (c WorkflowSetError) Wrap(err error) WorkflowSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("WorkflowSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c WorkflowSetError) Error() string {
	return c.Err.Error()
}

func (c WorkflowSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c WorkflowSetError) GetCode() ErrorCode {
	return c.Code
}

func (c WorkflowSetError) BubbleCode() WorkflowSetError {
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

type WorkflowDeleteError Merror

func (c WorkflowDeleteError) Wrap(err error) WorkflowDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("WorkflowDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c WorkflowDeleteError) Error() string {
	return c.Err.Error()
}

func (c WorkflowDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c WorkflowDeleteError) GetCode() ErrorCode {
	return c.Code
}

func (c WorkflowDeleteError) BubbleCode() WorkflowDeleteError {
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

type NodeListError Merror

func (c NodeListError) Wrap(err error) NodeListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("NodeListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c NodeListError) Error() string {
	return c.Err.Error()
}

func (c NodeListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c NodeListError) GetCode() ErrorCode {
	return c.Code
}

func (c NodeListError) BubbleCode() NodeListError {
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

type NodeGetError Merror

func (c NodeGetError) Wrap(err error) NodeGetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("NodeGetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c NodeGetError) Error() string {
	return c.Err.Error()
}

func (c NodeGetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c NodeGetError) GetCode() ErrorCode {
	return c.Code
}

func (c NodeGetError) BubbleCode() NodeGetError {
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

type NodeSetError Merror

func (c NodeSetError) Wrap(err error) NodeSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("NodeSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c NodeSetError) Error() string {
	return c.Err.Error()
}

func (c NodeSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c NodeSetError) GetCode() ErrorCode {
	return c.Code
}

func (c NodeSetError) BubbleCode() NodeSetError {
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

type NodeDeleteError Merror

func (c NodeDeleteError) Wrap(err error) NodeDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("NodeDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c NodeDeleteError) Error() string {
	return c.Err.Error()
}

func (c NodeDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c NodeDeleteError) GetCode() ErrorCode {
	return c.Code
}

func (c NodeDeleteError) BubbleCode() NodeDeleteError {
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

type PromptTemplateDeleteError Merror

func (c PromptTemplateDeleteError) Wrap(err error) PromptTemplateDeleteError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptTemplateDeleteError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptTemplateDeleteError) Error() string {
	return c.Err.Error()
}

func (c PromptTemplateDeleteError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c PromptTemplateDeleteError) GetCode() ErrorCode {
	return c.Code
}

func (c PromptTemplateDeleteError) BubbleCode() PromptTemplateDeleteError {
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

type PromptTemplateSetError Merror

func (c PromptTemplateSetError) Wrap(err error) PromptTemplateSetError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptTemplateSetError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptTemplateSetError) Error() string {
	return c.Err.Error()
}

func (c PromptTemplateSetError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c PromptTemplateSetError) GetCode() ErrorCode {
	return c.Code
}

func (c PromptTemplateSetError) BubbleCode() PromptTemplateSetError {
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

type PromptTemplateError Merror

func (c PromptTemplateError) Wrap(err error) PromptTemplateError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptTemplateError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptTemplateError) Error() string {
	return c.Err.Error()
}

func (c PromptTemplateError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c PromptTemplateError) GetCode() ErrorCode {
	return c.Code
}

func (c PromptTemplateError) BubbleCode() PromptTemplateError {
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

type PromptTemplateListError Merror

func (c PromptTemplateListError) Wrap(err error) PromptTemplateListError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("PromptTemplateListError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c PromptTemplateListError) Error() string {
	return c.Err.Error()
}

func (c PromptTemplateListError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c PromptTemplateListError) GetCode() ErrorCode {
	return c.Code
}

func (c PromptTemplateListError) BubbleCode() PromptTemplateListError {
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

func (c MSIConversionError) Wrap(err error) MSIConversionError {
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

func (c HTTPRequestError) Wrap(err error) HTTPRequestError {
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

type GetPromptTemplateError Merror

func (c GetPromptTemplateError) Wrap(err error) GetPromptTemplateError {
	c.Wrapped = err
	c.Err = fmt.Errorf("%s: %w\n", ErrString("GetPromptTemplateError", c.Info, c.Package, c.Struct, c.Function, c.Err), c.Wrapped)
	return c
}

func (c GetPromptTemplateError) Error() string {
	return c.Err.Error()
}

func (c GetPromptTemplateError) ErrorCode(code int16) {
	c.Code = ErrorCode(code)
}

func (c GetPromptTemplateError) GetCode() ErrorCode {
	return c.Code
}

func (c GetPromptTemplateError) BubbleCode() GetPromptTemplateError {
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


func (c SetContextError) Wrap(err error) SetContextError {
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

type WrappedError interface {
	Wrap(error)
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

