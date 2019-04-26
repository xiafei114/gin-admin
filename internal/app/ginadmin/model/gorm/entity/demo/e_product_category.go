package demo

import (
	"context"

	schema "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/gormplus"

	"gin-admin/internal/app/ginadmin/model/gorm/entity"
)

// GetProductCategoryDB 获取ProductCategory存储
func GetProductCategoryDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return entity.GetDBWithModel(ctx, defDB, ProductCategory{})
}

// SchemaProductCategory ProductCategory对象
type SchemaProductCategory schema.ProductCategory

// ToProductCategory 转换为ProductCategory实体
func (a SchemaProductCategory) ToProductCategory() *ProductCategory {
	item := &ProductCategory{
		RecordID: a.RecordID,
		NumCode:  a.NumCode,
		ChnName:  a.ChnName,
	}
	return item
}

// ProductCategory ProductCategory实体
type ProductCategory struct {
	entity.Model
	RecordID string `gorm:"column:record_id;size:36;index;"` // 记录内码
	NumCode  string `gorm:"column:num_Code;size:50;"`
	ChnName  string `gorm:"column:chn_Name;size:200;"`
}

func (a ProductCategory) String() string {
	return entity.ToString(a)
}

// TableName 表名
func (a ProductCategory) TableName() string {
	return a.Model.TableName("Product_Category")
}

// ToSchemaProductCategory 转换为ProductCategory对象
func (a ProductCategory) ToSchemaProductCategory() *schema.ProductCategory {
	item := &schema.ProductCategory{
		RecordID: a.RecordID,
		NumCode:  a.NumCode,
		ChnName:  a.ChnName,
	}
	return item
}

// ProductCategorys ProductCategory列表
type ProductCategorys []*ProductCategory

// ToSchemaProductCategorys 转换为ProductCategory对象列表
func (a ProductCategorys) ToSchemaProductCategorys() []*schema.ProductCategory {
	list := make([]*schema.ProductCategory, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaProductCategory()
	}
	return list
}
