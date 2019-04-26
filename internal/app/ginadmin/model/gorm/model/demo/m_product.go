package model

import (
	"context"
	"fmt"

	entity "gin-admin/internal/app/ginadmin/model/gorm/entity/demo"
	schema "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/gormplus"
	"gin-admin/pkg/logger"
)

// NewProduct 创建Product存储实例
func NewProduct(db *gormplus.DB) *Product {
	return &Product{db}
}

// Product Product存储
type Product struct {
	db *gormplus.DB
}

func (a *Product) getFuncName(name string) string {
	return fmt.Sprintf("gorm.model.Product.%s", name)
}

func (a *Product) getQueryOption(opts ...schema.ProductQueryOptions) schema.ProductQueryOptions {
	var opt schema.ProductQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Product) Query(ctx context.Context, params schema.ProductQueryParam, opts ...schema.ProductQueryOptions) (*schema.ProductQueryResult, error) {
	span := logger.StartSpan(ctx, "查询数据", a.getFuncName("Query"))
	defer span.Finish()

	db := entity.GetProductDB(ctx, a.db).DB
	if v := params.Code; v != "" {
		db = db.Where("code=?", v)
	}
	if v := params.LikeCode; v != "" {
		db = db.Where("code LIKE ?", "%"+v+"%")
	}
	if v := params.LikeName; v != "" {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}
	db = db.Order("id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.Products
	pr, err := WrapPageQuery(db, opt.PageParam, &list)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询数据发生错误")
	}
	qr := &schema.ProductQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaProducts(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Product) Get(ctx context.Context, recordID string, opts ...schema.ProductQueryOptions) (*schema.Product, error) {
	span := logger.StartSpan(ctx, "查询指定数据", a.getFuncName("Get"))
	defer span.Finish()

	db := entity.GetProductDB(ctx, a.db).Where("record_id=?", recordID)
	var item entity.Product
	ok, err := a.db.FindOne(db, &item)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询指定数据发生错误")
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaProduct(), nil
}

// Create 创建数据
func (a *Product) Create(ctx context.Context, item schema.Product) error {
	span := logger.StartSpan(ctx, "创建数据", a.getFuncName("Create"))
	defer span.Finish()

	Product := entity.SchemaProduct(item).ToProduct()
	result := entity.GetProductDB(ctx, a.db).Create(Product)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("创建数据发生错误")
	}
	return nil
}

// Update 更新数据
func (a *Product) Update(ctx context.Context, recordID string, item schema.Product) error {
	span := logger.StartSpan(ctx, "更新数据", a.getFuncName("Update"))
	defer span.Finish()

	Product := entity.SchemaProduct(item).ToProduct()
	result := entity.GetProductDB(ctx, a.db).Where("record_id=?", recordID).Omit("record_id", "creator").Updates(Product)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("更新数据发生错误")
	}
	return nil
}

// Delete 删除数据
func (a *Product) Delete(ctx context.Context, recordID string) error {
	span := logger.StartSpan(ctx, "删除数据", a.getFuncName("Delete"))
	defer span.Finish()

	result := entity.GetProductDB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.Product{})
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("删除数据发生错误")
	}
	return nil
}

// UpdateStatus 更新状态
func (a *Product) UpdateStatus(ctx context.Context, recordID string, status int) error {
	span := logger.StartSpan(ctx, "更新状态", a.getFuncName("UpdateStatus"))
	defer span.Finish()

	result := entity.GetProductDB(ctx, a.db).Where("record_id=?", recordID).Update("status", status)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("更新状态发生错误")
	}
	return nil
}
