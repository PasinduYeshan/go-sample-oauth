package response

import (
	"net/http"

	"github.com/PasinduYeshan/go-sample-oauth/internal/common/apperror" // Added import
	"github.com/labstack/echo/v4"
)

// APIErrorDetail holds structured information about a specific error.
type APIErrorDetail struct {
	Code    string `json:"code"`            // Application-specific error code
	Message string `json:"message"`         // Detailed error message
	Field   string `json:"field,omitempty"` // Optional: for validation errors, indicates the problematic field
}

// APIResponse defines the standard structure for API responses.
type APIResponse struct {
	Status   string           `json:"status"`             // "success" or "error"
	Message  string           `json:"message,omitempty"`  // General message, typically for success or a summary for errors
	Data     interface{}      `json:"data,omitempty"`     // Payload for success responses
	Metadata interface{}      `json:"metadata,omitempty"` // Additional metadata
	Errors   []APIErrorDetail `json:"errors,omitempty"`   // List of specific errors for error responses
}

// SuccessResponse constructs a standard success response.
func SuccessResponse(c echo.Context, message string, data interface{}, metadata interface{}) error {
	return c.JSON(http.StatusOK, &APIResponse{
		Status:   "success",
		Message:  message,
		Data:     data,
		Metadata: metadata,
	})
}

// ErrorResponse constructs a standard error response from an AppError.
// The main 'Message' field in APIResponse can be a general summary, while the 'Errors' slice contains specifics.
func ErrorResponse(c echo.Context, appErr *apperror.AppError) error {
	// If a custom message is set in appErr (e.g., via WithMessage), use it. Otherwise, use DefaultMessage.
	// The appErr.Error() method already handles appending OriginalError if present.
	errorMessage := appErr.Error()

	return c.JSON(appErr.HTTPStatus, APIResponse{
		Status:  "error",
		Message: errorMessage, // Use the error's message, which might include original error details
		Errors: []APIErrorDetail{
			{Code: appErr.Code, Message: appErr.DefaultMessage}, // DefaultMessage for the structured error part
		},
	})
}

// ValidationErrorResponse constructs a response for validation errors.
// It uses a predefined error code for validation failures.
func ValidationErrorResponse(c echo.Context, errors []APIErrorDetail) error {
	// Ensure each error detail has a code, if not, assign the generic validation failed code.
	// This part might need its own AppError definition if we want to standardize fully.
	// For now, keeping it as is but noting that ErrCodeValidationFailed would need to be an AppError.
	// Let's assume a generic apperror.ErrValidation for now if we were to refactor this fully.
	// For now, this function's direct AppError integration is deferred as it handles a slice of APIErrorDetail.
	// A possible refactor: create an AppError for validation and pass it, then populate APIErrorDetail from input.

	// For the purpose of this step, we focus on functions that take a single error.
	// We can create a specific AppError for validation if needed.
	// var validationAppError *apperror.AppError // = apperror.ErrValidation (if defined)

	// This function's signature and internal logic for constructing APIResponse
	// from a []APIErrorDetail is a bit different. We'll keep it for now,
	// but ideally, it should also be based on a specific AppError type for "Validation Error".
	// For now, we'll assume ErrCodeValidationFailed is still a string constant.
	// If it's removed, this function will need a dedicated AppError.
	// Let's assume ErrCodeValidationFailed is defined elsewhere or we use a generic one.
	var errCodeValidation = "VALIDATION_FAILED" // Placeholder if ErrCodeValidationFailed is removed

	for i := range errors {
		if errors[i].Code == "" {
			errors[i].Code = errCodeValidation
		}
	}
	return c.JSON(http.StatusBadRequest, APIResponse{
		Status:  "error",
		Message: "Validation failed", // General message for validation errors
		Errors:  errors,
	})
}

// NotFoundResponse constructs a 404 Not Found error response using apperror.ErrNotFound.
func NotFoundResponse(c echo.Context) error {
	return ErrorResponse(c, apperror.ErrNotFound)
}

// UnauthorizedResponse constructs a 401 Unauthorized error response using apperror.ErrUnauthorized.
// It can be customized if a more specific message or error is needed.
func UnauthorizedResponse(c echo.Context, specificError *apperror.AppError) error {
	if specificError != nil {
		return ErrorResponse(c, specificError)
	}
	return ErrorResponse(c, apperror.ErrUnauthorized)
}

// ForbiddenResponse constructs a 403 Forbidden error response using apperror.ErrForbidden.
func ForbiddenResponse(c echo.Context, specificError *apperror.AppError) error {
	if specificError != nil {
		return ErrorResponse(c, specificError)
	}
	return ErrorResponse(c, apperror.ErrForbidden)
}

// InternalServerErrorResponse constructs a 500 Internal Server Error response using apperror.ErrInternalServer.
func InternalServerErrorResponse(c echo.Context, specificError *apperror.AppError) error {
	if specificError != nil {
		// If a specific error is provided, and it has an OriginalError, log it here.
		// log.Printf("Internal Server Error: %v, Original: %v", specificError.DefaultMessage, specificError.OriginalError)
		return ErrorResponse(c, specificError)
	}
	// log.Printf("Internal Server Error: %v", apperror.ErrInternalServer.DefaultMessage)
	return ErrorResponse(c, apperror.ErrInternalServer)
}

// BadRequestResponse constructs a 400 Bad Request error response using apperror.ErrBadRequest.
func BadRequestResponse(c echo.Context, specificError *apperror.AppError) error {
	if specificError != nil {
		return ErrorResponse(c, specificError)
	}
	return ErrorResponse(c, apperror.ErrBadRequest)
}
