package code

type ErrorCode string

var (
	CodeOK ErrorCode = "ok"

	CodeBadRequest   ErrorCode = "bad request"
	CodeUnauthorized ErrorCode = "unauthorized"
	CodeValidParam   ErrorCode = "valid parameter"
	CodeNotFound     ErrorCode = "not found"
	CodeOutOfTerm    ErrorCode = "out of term"

	CodeDatabase ErrorCode = "database error"
	CodeInternal ErrorCode = "internal error"

	CodeUnknown ErrorCode = "unknown error"
)

func (e ErrorCode) Error() string {
	return string(e)
}

var _ error = (*ErrorCode)(nil)
