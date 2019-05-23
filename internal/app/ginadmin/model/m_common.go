package model

import (
	"context"

	schema "gin-admin/internal/app/ginadmin/schema"
	schemaProject "gin-admin/internal/app/ginadmin/schema/demo"
)

// ICommon Common存储接口
type ICommon interface {
	// 查询数据
	Query(ctx context.Context, params schema.CommonQueryParam, opts ...schema.CommonQueryOptions) (*schema.CommonQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.CommonQueryOptions) (*schemaProject.CommonFile, error)
	// 创建数据
	Create(ctx context.Context, item schemaProject.CommonFile) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schemaProject.CommonFile) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
	// 更新状态
	UpdateStatus(ctx context.Context, recordID string, status int) error
}
