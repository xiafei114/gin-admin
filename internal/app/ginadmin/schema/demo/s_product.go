package demo

import (
	schema "gin-admin/internal/app/ginadmin/schema"
	"time"
)

// Product Product对象
type Product struct {
	RecordID        string    `json:"record_id" swaggo:"false,记录ID"`
	ProductCode     string    `json:"product_code" binding:"required" swaggo:"true,编号"`
	ProductName     string    `json:"product_name" binding:"required" swaggo:"true,名称"`
	ProductCateName string    `json:"product_cate_name"  swaggo:"false,类别名称"`                // 类型（使用弹出选择）
	ProductCateID   string    `json:"product_cate_id" binding:"required" swaggo:"true,类别主键"` // 类型（使用弹出选择）
	Price           float64   `json:"price"  swaggo:"false,价格"`
	Number          int       `json:"number"  swaggo:"false,数量"`
	StartDate       time.Time `json:"startDate"  swaggo:"false,开始有效期"`                 // 开始有效期
	EndDate         time.Time `json:"endDate"  swaggo:"false,结束有效期"`                   // 结束有效期
	ProductType     string    `json:"product_type"  swaggo:"false,类别"`                 // 类别 select
	ProductColor    string    `json:"product_color"  swaggo:"false,颜色"`                // 颜色 radio
	IsValid         bool      `json:"is_Valid" binding:"required" swaggo:"false,是否有效"` // 是否有效
	Content         string    `json:"content" swaggo:"false,内容"`                       // 内容（使用富文本保存）
}

// ProductQueryParam 查询条件
type ProductQueryParam struct {
	Code     string // 编号
	Status   int    // 状态(1:启用 2:停用)
	LikeCode string // 编号(模糊查询)
	LikeName string // 名称(模糊查询)
}

// ProductQueryOptions Product对象查询可选参数项
type ProductQueryOptions struct {
	PageParam *schema.PaginationParam // 分页参数
}

// ProductQueryResult Product对象查询结果
type ProductQueryResult struct {
	Data       []*Product
	PageResult *schema.PaginationResult
}
