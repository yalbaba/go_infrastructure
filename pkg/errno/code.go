package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 1, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 2, Message: "Error occurred while binding the request body to the struct"}
	ErrParam            = &Errno{Code: 3, Message: "Params error"}
)
