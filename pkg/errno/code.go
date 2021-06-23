package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "生成token失败"}

	// user errors
	ErrUserNotFound      = &Errno{Code: 20102, Message: "没有找到该用户"}
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "token 不正确"}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "密码不正确"}
	ErrUserNameExist     = &Errno{Code: 20105, Message: "用户名已存在"}
	ErrCreateUser        = &Errno{Code: 20106, Message: "创建用户失败"}
)
