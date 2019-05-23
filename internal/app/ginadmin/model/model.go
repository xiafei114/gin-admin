package model

// Common 提供统一的存储接口
type Common struct {
	Trans      ITrans
	Demo       IDemo
	Permission IPermission
	Role       IRole
	User       IUser
	Product    IProduct
	Media      IMedia
	Common     ICommon
}
