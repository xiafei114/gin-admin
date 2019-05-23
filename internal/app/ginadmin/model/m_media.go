package model

import (
	"context"

	schema "gin-admin/internal/app/ginadmin/schema"
	schemaProject "gin-admin/internal/app/ginadmin/schema/demo"
)

// IMedia Media存储接口
type IMedia interface {
	// 查询数据
	Query(ctx context.Context, params schema.CommonQueryParam, opts ...schema.CommonQueryOptions) (*schema.CommonQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.CommonQueryOptions) (*schemaProject.Media, error)
	// 创建数据
	Create(ctx context.Context, item schemaProject.Media) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schemaProject.Media) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
	// 更新状态
	UpdateStatus(ctx context.Context, recordID string, status int) error
}