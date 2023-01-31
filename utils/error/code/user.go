package code

const (
	/**
	Describe 用户模块错误码
	1000 - 1999
	*/

	UserNotExist      = 1000 // 用户不存在
	UserPasswordWrong = 1001 // 用户密码错误
	UserAlreadyExist  = 1002 // 用户已存在
	UserIsDisabled    = 1003 // 用户已被禁用
	UserIsDeleted     = 1004 // 用户已被删除
	UserParam         = 1005 // 用户已被删除
)

var userCode = map[int]string{
	UserNotExist:      "用户不存在",
	UserPasswordWrong: "用户密码错误",
	UserAlreadyExist:  "用户已存在",
	UserIsDisabled:    "用户已被禁用",
	UserIsDeleted:     "用户已被删除",
	UserParam:         "参数错误",
}

func GetUserMsg(code int) string {
	msg, ok := userCode[code]
	if ok {
		return msg
	}

	return userCode[UserNotExist]
}
