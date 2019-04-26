package model

import (
	"context"

	schema "gin-admin/internal/app/ginadmin/schema/demo"
)

// IProduct Product存储接口
type IProduct interface {
	// 查询数据
	Query(ctx context.Context, params schema.ProductQueryParam, opts ...schema.ProductQueryOptions) (*schema.ProductQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.ProductQueryOptions) (*schema.Product, error)
	// 创建数据
	Create(ctx context.Context, item schema.Product) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.Product) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
	// 更新状态
	UpdateStatus(ctx context.Context, recordID string, status int) error
}
