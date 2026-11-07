package errors

type Code string

const (

	// ---- GENERAL ERRORS ----
	ErrCodeUnknown                        Code = "UNKNOWN_ERROR"
	ErrCodeInternalServerError            Code = "INTERNAL_SERVER_ERROR"
	ErrCodeBadRequest                     Code = "BAD_REQUEST"
	ErrCodeUnauthorized                   Code = "UNAUTHORIZED"
	ErrCodeDuplicateReference             Code = "DUPLICATE_REFERENCE"
	ErrCodeForbidden                      Code = "FORBIDDEN"
	ErrCodeNotFound                       Code = "NOT_FOUND"
	ErrCodeFinancialReportAlreadyUploaded Code = "FINANCIAL_REPORT_ALREADY_UPLOADED"
	ErrCodeTooManyRequests                Code = "TOO_MANY_REQUESTS"

	// ---- AUTH & USER ERRORS ----
	ErrCodeUserNotFound    Code = "USER_NOT_FOUND"
	ErrCodeSessionNotFound Code = "SESSION_NOT_FOUND"
	ErrCodeInvalidPassword Code = "INVALID_PASSWORD"
)
