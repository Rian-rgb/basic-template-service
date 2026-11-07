package errors

import "net/http"

func (c Code) ToHTTPStatus() int {
	switch c {
	case ErrCodeUnknown:
		return http.StatusInternalServerError

	case ErrCodeBadRequest:
		return http.StatusBadRequest

	case ErrCodeUnauthorized, ErrCodeInvalidPassword:
		return http.StatusUnauthorized

	case ErrCodeForbidden:
		return http.StatusForbidden

	case ErrCodeNotFound,
		ErrCodeUserNotFound,
		ErrCodeSessionNotFound:

		return http.StatusNotFound

	case ErrCodeDuplicateReference,
		ErrCodeFinancialReportAlreadyUploaded:
		return http.StatusConflict // HTTP 409 sutable for conflict data in database

	case ErrCodeTooManyRequests:
		return http.StatusTooManyRequests

	default:
		return http.StatusInternalServerError
	}
}
