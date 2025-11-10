package constants

// Error codes following KBTG workshop standards
const (
	// General error codes
	ErrCodeBadRequest     = "BAD_REQUEST"
	ErrCodeUnauthorized   = "UNAUTHORIZED"
	ErrCodeForbidden      = "FORBIDDEN"
	ErrCodeNotFound       = "NOT_FOUND"
	ErrCodeConflict       = "CONFLICT"
	ErrCodeInternalServer = "INTERNAL_SERVER_ERROR"
	ErrCodeValidation     = "VALIDATION_ERROR"

	// User specific error codes
	ErrCodeUserNotFound    = "USER_NOT_FOUND"
	ErrCodeUserEmailExists = "USER_EMAIL_EXISTS"
	ErrCodeUserInvalidData = "USER_INVALID_DATA"

	// Transfer specific error codes
	ErrCodeTransferNotFound       = "TRANSFER_NOT_FOUND"
	ErrCodeInsufficientBalance    = "INSUFFICIENT_BALANCE"
	ErrCodeSelfTransferNotAllowed = "SELF_TRANSFER_NOT_ALLOWED"
	ErrCodeTransferFailed         = "TRANSFER_FAILED"
	ErrCodeInvalidAmount          = "INVALID_AMOUNT"
)

// HTTP Status Messages
const (
	MsgSuccess         = "Success"
	MsgCreated         = "Created successfully"
	MsgUpdated         = "Updated successfully"
	MsgDeleted         = "Deleted successfully"
	MsgBadRequest      = "Bad request"
	MsgUnauthorized    = "Unauthorized"
	MsgForbidden       = "Forbidden"
	MsgNotFound        = "Not found"
	MsgInternalError   = "Internal server error"
	MsgValidationError = "Validation error"
)

// Default pagination values
const (
	DefaultLimit  = 10
	MaxLimit      = 100
	DefaultOffset = 0
)
