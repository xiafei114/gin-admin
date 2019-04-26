package demo

import (
	"context"

	"gin-admin/internal/app/ginadmin/model"
	"gin-admin/internal/app/ginadmin/schema"
	sDemo "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/util"
)

// NewProduct 创建Product
func NewProduct(m *model.Common) *Product {
	return &Product{
		ProductModel: m.Product,
	}
}

// Product 示例程序
type Product struct {
	ProductModel model.IProduct
}

// QueryPage 查询分页数据
func (a *Product) QueryPage(ctx context.Context, params sDemo.ProductQueryParam, pp *schema.PaginationParam) ([]*sDemo.Product, *schema.PaginationResult, error) {
	result, err := a.ProductModel.Query(ctx, params, sDemo.ProductQueryOptions{PageParam: pp})
	if err != nil {
		return nil, nil, err
	}
	return result.Data, result.PageResult, nil
}

// Get 查询指定数据
func (a *Product) Get(ctx context.Context, recordID string) (*sDemo.Product, error) {
	item, err := a.ProductModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *Product) checkCode(ctx context.Context, code string) error {
	result, err := a.ProductModel.Query(ctx, sDemo.ProductQueryParam{
		Code: code,
	}, sDemo.ProductQueryOptions{
		PageParam: &schema.PaginationParam{PageSize: -1},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.NewBadRequestError("编号已经存在")
	}
	return nil
}

// Create 创建数据
func (a *Product) Create(ctx context.Context, item sDemo.Product) (*sDemo.Product, error) {
	err := a.checkCode(ctx, item.ProductCode)
	if err != nil {
		return nil, err
	}

	item.RecordID = util.MustUUID()
	// item.Creator = bll.GetUserID(ctx)
	err = a.ProductModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, item.RecordID)
}

// Update 更新数据
func (a *Product) Update(ctx context.Context, recordID string, item sDemo.Product) (*sDemo.Product, error) {
	oldItem, err := a.ProductModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	} else if oldItem.ProductCode != item.ProductCode {
		err := a.checkCode(ctx, item.ProductCode)
		if err != nil {
			return nil, err
		}
	}

	err = a.ProductModel.Update(ctx, recordID, item)
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, recordID)
}

// Delete 删除数据
func (a *Product) Delete(ctx context.Context, recordID string) error {
	err := a.ProductModel.Delete(ctx, recordID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatus 更新状态
func (a *Product) UpdateStatus(ctx context.Context, recordID string, status int) error {
	return a.ProductModel.UpdateStatus(ctx, recordID, status)
}
