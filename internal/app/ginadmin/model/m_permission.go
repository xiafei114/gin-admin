package model

import (
	"context"

	"gin-admin/internal/app/ginadmin/schema"
)

// IPermission 权限管理存储接口
type IPermission interface {
	// 查询数据
	Query(ctx context.Context, params schema.PermissionQueryParam, opts ...schema.PermissionQueryOptions) (*schema.PermissionQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.PermissionQueryOptions) (*schema.Permission, error)
	// 创建数据
	Create(ctx context.Context, item schema.Permission) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.Permission) error
	// 更新父级路径
	UpdateParentPath(ctx context.Context, recordID, parentPath string) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}
