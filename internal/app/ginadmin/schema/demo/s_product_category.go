package demo

import (
	schema "gin-admin/internal/app/ginadmin/schema"
)

// ProductCategory ProductCategory对象
type ProductCategory struct {
	RecordID string `json:"record_id" swaggo:"false,记录ID"`
	NumCode  string `json:"num_code" binding:"required" swaggo:"true,编号"`
	ChnName  string `json:"chn_name" binding:"required" swaggo:"true,名称"`
}

// ProductCategoryQueryParam 查询条件
type ProductCategoryQueryParam struct {
	Code     string // 编号
	Status   int    // 状态(1:启用 2:停用)
	LikeCode string // 编号(模糊查询)
	LikeName string // 名称(模糊查询)
}

// ProductCategoryQueryOptions ProductCategory对象查询可选参数项
type ProductCategoryQueryOptions struct {
	PageParam *schema.PaginationParam // 分页参数
}

// ProductCategoryQueryResult ProductCategory对象查询结果
type ProductCategoryQueryResult struct {
	Data       []*ProductCategory
	PageResult *schema.PaginationResult
}
