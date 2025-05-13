package apperror

import "net/http"

// AppError defines a standard application error.
type AppError struct {
	Code           string
	DefaultMessage string
	HTTPStatus     int
	OriginalError  error // To wrap an underlying error, optional
}

// New creates a new AppError pointer.
func New(code, defaultMessage string, httpStatus int) *AppError {
	return &AppError{
		Code:           code,
		DefaultMessage: defaultMessage,
		HTTPStatus:     httpStatus,
	}
}

// WithMessage returns a new AppError instance with a specific message,
// allowing to override the DefaultMessage for a particular occurrence.
func (e *AppError) WithMessage(message string) *AppError {
	// Return a new instance to avoid modifying the global error definition
	return &AppError{
		Code:           e.Code,
		DefaultMessage: message, // This becomes the message for this specific instance
		HTTPStatus:     e.HTTPStatus,
		OriginalError:  e.OriginalError,
	}
}

// WithError returns a new AppError instance that wraps an original error.
// This is useful for preserving context from underlying library errors.
func (e *AppError) WithError(err error) *AppError {
	// Return a new instance
	return &AppError{
		Code:           e.Code,
		DefaultMessage: e.DefaultMessage,
		HTTPStatus:     e.HTTPStatus,
		OriginalError:  err,
	}
}

// Error makes AppError satisfy the error interface.
// It returns the DefaultMessage. If an OriginalError is wrapped, its message is appended.
func (e *AppError) Error() string {
	msg := e.DefaultMessage
	if e.OriginalError != nil {
		msg += ": " + e.OriginalError.Error()
	}
	return msg
}

// --- Generic Application Errors ---
var (
	ErrBadRequest     = New("APP_BAD_REQUEST", "The request was invalid or cannot be otherwise served.", http.StatusBadRequest)
	ErrUnauthorized   = New("APP_UNAUTHORIZED", "Authentication is required and has failed or has not yet been provided.", http.StatusUnauthorized)
	ErrForbidden      = New("APP_FORBIDDEN", "You do not have permission to perform this action.", http.StatusForbidden)
	ErrNotFound       = New("APP_NOT_FOUND", "The requested resource could not be found.", http.StatusNotFound)
	ErrInternalServer = New("APP_INTERNAL_ERROR", "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
	// Add more generic errors as needed, e.g., ErrConflict, ErrValidation, etc.
)

// --- Middleware Specific Errors ---
// These could also live in a middleware-specific errors file if preferred,
// but for now, centralizing them here is fine.
var (
	ErrScopeConfigNotLoaded = New("MDW_CFG_001", "Critical: API scope configuration is not loaded.", http.StatusInternalServerError)
	ErrInsufficientScope    = New("MDW_ATH_001", "Access Denied: You do not have sufficient scope to access this resource.", http.StatusForbidden)
	ErrMissingAuthHeader    = New("MDW_ATH_002", "Authorization header is missing.", http.StatusUnauthorized)
	ErrInvalidAuthHeader    = New("MDW_ATH_003", "Authorization header is malformed or token type is not Bearer.", http.StatusUnauthorized)
	ErrInvalidToken         = New("MDW_ATH_004", "The provided token is invalid, expired, or malformed.", http.StatusUnauthorized)
	// ErrRateLimited          = New("MDW_RAT_001", "You have exceeded the request rate limit.", http.StatusTooManyRequests)
)

// --- Service/Handler Specific Errors (Examples) ---
// It's often good practice to define more specific errors closer to where they occur,
// or define them here if they are shared across services.

// Example: Merchant Service Errors
// var (
// 	ErrMerchantNotFound = New("SVC_MER_001", "Merchant not found.", http.StatusNotFound)
// 	ErrInvalidMerchantData = New("SVC_MER_002", "Invalid merchant data provided.", http.StatusBadRequest)
// )

// Example: Ads Service Errors
// var (
// 	ErrAdNotFound = New("SVC_ADS_001", "Advertisement not found.", http.StatusNotFound)
// )
