package schema

// LoginParam 登录参数
type LoginParam struct {
	UserName    string `json:"user_name" binding:"required" swaggo:"true,用户名"`
	Password    string `json:"password" binding:"required" swaggo:"true,密码(md5加密)"`
	CaptchaID   string `json:"captcha_id" swaggo:"true,验证码ID"`
	CaptchaCode string `json:"captcha_code" swaggo:"true,验证码"`
}

// UserLoginInfo 用户登录信息
type UserLoginInfo struct {
	ID        string          `json:"id" swaggo:"true,id"`
	Name      string          `json:"name" swaggo:"true,姓名"`
	UserName  string          `json:"user_name" swaggo:"true,用户名"`
	RealName  string          `json:"real_name" swaggo:"true,真实姓名"`
	RoleNames []string        `json:"role_names" swaggo:"true,角色名列表"`
	Role      []UserLoginRole `json:"role" swaggo:"true,用户角色"`
}

// UserLoginRole 获得登录用户信息
type UserLoginRole struct {
	ID          string                `json:"id" swaggo:"true,id"`
	Name        string                `json:"name" swaggo:"true,名称"`
	Describe    string                `json:"describe" swaggo:"true,描述"`
	Permissions []UserLoginPermission `json:"permissions,omitempty" swaggo:"true,许可"`
}

// UserLoginPermission 获得登录用户信息
type UserLoginPermission struct {
	PermissionID    string      `json:"permissionID" swaggo:"true,id"`
	PermissionName  string      `json:"permissionName" swaggo:"true,名称"`
	ActionEntitySet PermissionActions `json:"actionEntitySets" swaggo:"false,动作列表"`
}

// UserLoginedInfo 获得登录用户信息
type UserLoginedInfo struct {
	Message string        `json:"message" swaggo:"true,消息"`
	Result  UserLoginInfo `json:"result" swaggo:"true,返回结果"`
}

// UpdatePasswordParam 更新密码请求参数
type UpdatePasswordParam struct {
	OldPassword string `json:"old_password" binding:"required" swaggo:"true,旧密码(md5加密)"`
	NewPassword string `json:"new_password" binding:"required" swaggo:"true,新密码(md5加密)"`
}

// LoginCaptcha 登录验证码
type LoginCaptcha struct {
	CaptchaID string `json:"captcha_id" swaggo:"true,验证码ID"`
}

// LoginTokenInfo 登录令牌信息
type LoginTokenInfo struct {
	AccessToken string `json:"access_token" swaggo:"true,访问令牌"`
	TokenType   string `json:"token_type" swaggo:"true,令牌类型"`
	ExpiresAt   int64  `json:"expires_at" swaggo:"true,令牌到期时间"`
}
