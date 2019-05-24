package demo

import (
	"context"

	"gin-admin/internal/app/ginadmin/model"
	"gin-admin/internal/app/ginadmin/schema"
	schemaProject "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/util"
)

// NewMedia 创建Media
func NewMedia(m *model.Common) *Media {
	return &Media{
		MediaModel: m.Media,
	}
}

// Media 示例程序
type Media struct {
	MediaModel model.IMedia
}

// QueryPage 查询分页数据
func (a *Media) QueryPage(ctx context.Context, params schema.CommonQueryParam, pp *schema.PaginationParam) (interface{}, *schema.PaginationResult, error) {
	result, err := a.MediaModel.Query(ctx, params, schema.CommonQueryOptions{PageParam: pp})
	if err != nil {
		return nil, nil, err
	}
	return result.Data, result.PageResult, nil
}

// Get 查询指定数据
func (a *Media) Get(ctx context.Context, recordID string, hostName string) (*schemaProject.Media, error) {
	item, err := a.MediaModel.Get(ctx, recordID, hostName)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *Media) checkCode(ctx context.Context, code string) error {
	result, err := a.MediaModel.Query(ctx, schema.CommonQueryParam{
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

// Upload 创建数据
func (a *Media) Upload(ctx context.Context, item schemaProject.CommonFile, hostName string) (*schemaProject.Media, error) {
	var media schemaProject.Media
	media.RecordID = util.MustUUID()
	// item.Creator = bll.GetUserID(ctx)
	err := a.MediaModel.Upload(ctx, media, item)
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, media.RecordID, hostName)
}

// Update 更新数据
func (a *Media) Update(ctx context.Context, recordID string, item schemaProject.Media, hostName string) (*schemaProject.Media, error) {
	oldItem, err := a.MediaModel.Get(ctx, recordID, hostName)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	} else if oldItem.InfoNo != item.InfoNo {
		err := a.checkCode(ctx, item.InfoNo)
		if err != nil {
			return nil, err
		}
	}

	err = a.MediaModel.Update(ctx, recordID, item)
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, recordID, hostName)
}

// Delete 删除数据
func (a *Media) Delete(ctx context.Context, recordID string) error {
	err := a.MediaModel.Delete(ctx, recordID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatus 更新状态
func (a *Media) UpdateStatus(ctx context.Context, recordID string, status int) error {
	return a.MediaModel.UpdateStatus(ctx, recordID, status)
}
