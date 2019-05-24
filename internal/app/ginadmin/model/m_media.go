package model

import (
	"context"

	schema "gin-admin/internal/app/ginadmin/schema"
	schemaProject "gin-admin/internal/app/ginadmin/schema/demo"
)

// IMedia Media存储接口
type IMedia interface {
	// 查询数据
	Query(ctx context.Context, params schema.CommonQueryParam, hostName string, opts ...schema.CommonQueryOptions) (*schema.CommonQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, hostName string, opts ...schema.CommonQueryOptions) (*schemaProject.Media, error)
	// 上传文件
	Upload(ctx context.Context, smedia schemaProject.Media, item schemaProject.CommonFile) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schemaProject.Media) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
	// 更新状态
	UpdateStatus(ctx context.Context, recordID string, status int) error
}
