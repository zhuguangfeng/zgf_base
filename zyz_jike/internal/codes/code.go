package codes

const (
	Success = 0
)

// user 用户相关错误码
const (

	// UserInvalidInput 统一的用户模块的输入错误
	UserInvalidInput = 401001
	// UserInvalidPhone  用户名错误
	UserInvalidPhone = 401002
	// UserInvalidPassword  密码错误
	UserInvalidPassword = 401003
	// UserDuplicatePhone 用户手机号码冲突冲突
	UserDuplicatePhone = 401004
	// UserInternalServerError 统一的用户模块的系统错误
	UserInternalServerError = 401000
)

const (
	ArticleInvalidInput = 402001
	ArticleInternalServerError
)
