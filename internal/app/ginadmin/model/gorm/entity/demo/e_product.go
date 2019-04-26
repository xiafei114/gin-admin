package demo

import (
	"context"
	"time"

	schema "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/gormplus"

	"gin-admin/internal/app/ginadmin/model/gorm/entity"
)

// GetProductDB 获取Product存储
func GetProductDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return entity.GetDBWithModel(ctx, defDB, Product{})
}

// SchemaProduct Product对象
type SchemaProduct schema.Product

// ToProduct 转换为Product实体
func (a SchemaProduct) ToProduct() *Product {
	item := &Product{
		RecordID:      a.RecordID,
		ProductCode:   a.ProductCode,
		ProductName:   a.ProductName,
		ProductCateID: a.ProductCateID,
		Price:         a.Price,
		Number:        a.Number,
		StartDate:     a.StartDate,
		EndDate:       a.EndDate,
		ProductType:   a.ProductType,
		ProductColor:  a.ProductColor,
		IsValid:       a.IsValid,
		Content:       a.Content,
	}
	return item
}

// Product Product实体
type Product struct {
	entity.Model
	RecordID      string    `gorm:"column:record_id;size:36;index;"` // 记录内码
	ProductCode   string    `gorm:"column:product_Code;size:50;"`
	ProductName   string    `gorm:"column:product_Name;size:200;"`
	ProductCateID string    `gorm:"column:product_Cate_Id;size:36;index"` // 类型（使用弹出选择）
	Price         float64   `gorm:"column:price;"`
	Number        int       `gorm:"column:number;"`
	StartDate     time.Time `gorm:"column:start_Date;"`    // 开始有效期
	EndDate       time.Time `gorm:"column:end_Date;"`      // 结束有效期
	ProductType   string    `gorm:"column:product_Type;"`  // 类别 select
	ProductColor  string    `gorm:"column:product_Color;"` // 颜色 radio
	IsValid       bool      `gorm:"column:is_Valid;"`      // 是否有效
	Content       string    `gorm:"column:content;"`       // 内容（使用富文本保存）
}

func (a Product) String() string {
	return entity.ToString(a)
}

// TableName 表名
func (a Product) TableName() string {
	return a.Model.TableName("Product")
}

// ToSchemaProduct 转换为Product对象
func (a Product) ToSchemaProduct() *schema.Product {
	item := &schema.Product{
		RecordID:      a.RecordID,
		ProductCode:   a.ProductCode,
		ProductName:   a.ProductName,
		ProductCateID: a.ProductCateID,
		Price:         a.Price,
		Number:        a.Number,
		StartDate:     a.StartDate,
		EndDate:       a.EndDate,
		ProductType:   a.ProductType,
		ProductColor:  a.ProductColor,
		IsValid:       a.IsValid,
		Content:       a.Content,
	}
	return item
}

// Products Product列表
type Products []*Product

// ToSchemaProducts 转换为Product对象列表
func (a Products) ToSchemaProducts() []*schema.Product {
	list := make([]*schema.Product, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaProduct()
	}
	return list
}
