package errno

var (
	// Common errors
	OK                    = &Errno{Code: 0, Message: "OK"}
	InternalServerError   = &Errno{Code: 1, Message: "Internal server error"}
	ErrBind               = &Errno{Code: 2, Message: "Error occurred while binding the request body to the struct"}
	ErrParam              = &Errno{Code: 3, Message: "Params error"}
	AuthorizationRequired = &Errno{Code: 4, Message: "Authorization required"}
	VerifyCodeError       = &Errno{Code: 5, Message: "验证码错误"}
	AuthorizationError    = &Errno{Code: 6, Message: "权限验证失败"}
	AuthorizationLess     = &Errno{Code: 7, Message: "Authorization less"}
	AccountLogoff         = &Errno{Code: 8, Message: "账户已注销"}

	// user errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found"}
	ErrOpNotFound   = &Errno{Code: 20103, Message: "The operate was not found"}

	// admin        10000~19999
	// user         20000~29999
	// content      30000~39999
	OperationRepeat = &Errno{Code: 30001, Message: "操作重复"}
	// im           40000~49999
	// sale         50000~59999
	// search       60000~69999
	// message      70000~79999
	// guide        80000~89999
	ActivityFinished = &Errno{Code: 80001, Message: "活动已结束"}
	PhoneExist       = &Errno{Code: 80002, Message: "该手机号已经报名"}
)
