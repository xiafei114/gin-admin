package demo

import (
	"context"

	"gin-admin/internal/app/ginadmin/model"
	"gin-admin/internal/app/ginadmin/schema"
	schemaProject "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/util"
)

// NewCommon 创建Common
func NewCommon(m *model.Common) *Common {
	return &Common{
		CommonModel: m.Common,
	}
}

// Common 示例程序
type Common struct {
	CommonModel model.ICommon
}

// QueryPage 查询分页数据
func (a *Common) QueryPage(ctx context.Context, params schema.CommonQueryParam, pp *schema.PaginationParam) (interface{}, *schema.PaginationResult, error) {
	result, err := a.CommonModel.Query(ctx, params, schema.CommonQueryOptions{PageParam: pp})
	if err != nil {
		return nil, nil, err
	}
	return result.Data, result.PageResult, nil
}

// Get 查询指定数据
func (a *Common) Get(ctx context.Context, recordID string) (*schemaProject.CommonFile, error) {
	item, err := a.CommonModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *Common) checkCode(ctx context.Context, code string) error {
	result, err := a.CommonModel.Query(ctx, schema.CommonQueryParam{
		Code: code,
	}, schema.CommonQueryOptions{
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
func (a *Common) Create(ctx context.Context, item schemaProject.CommonFile) (*schemaProject.CommonFile, error) {
	item.RecordID = util.MustUUID()
	// item.Creator = bll.GetUserID(ctx)
	err := a.CommonModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, item.RecordID)
}

// Update 更新数据
func (a *Common) Update(ctx context.Context, recordID string, item schemaProject.CommonFile) (*schemaProject.CommonFile, error) {
	oldItem, err := a.CommonModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	}

	err = a.CommonModel.Update(ctx, recordID, item)
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, recordID)
}

// Delete 删除数据
func (a *Common) Delete(ctx context.Context, recordID string) error {
	err := a.CommonModel.Delete(ctx, recordID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatus 更新状态
func (a *Common) UpdateStatus(ctx context.Context, recordID string, status int) error {
	return a.CommonModel.UpdateStatus(ctx, recordID, status)
}
