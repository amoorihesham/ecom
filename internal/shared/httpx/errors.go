package httpx

type ErrorCode string

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Err     error     `json:"-"`
}

var (
	ErrBadRequest   ErrorCode = "BAD_REQUEST"
	ErrUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrForbidden    ErrorCode = "FORBIDDEN"
	ErrNotFound     ErrorCode = "NOT_FOUND"
	ErrConflict     ErrorCode = "CONFLICT"
	ErrInternal     ErrorCode = "INTERNAL_ERROR"
)

func (ap *AppError) Error() string {
	return ap.Err.Error()
}

func NewError(code ErrorCode, message string, err error) error {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
