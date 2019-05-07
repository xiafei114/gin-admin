package bll

import (
	"context"

	"gin-admin/internal/app/ginadmin/model"
	"gin-admin/internal/app/ginadmin/schema"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/util"
)

// NewPermission 创建权力管理实例
func NewPermission(m *model.Common) *Permission {
	return &Permission{
		TransModel: m.Trans,
		PermissionModel:  m.Permission,
	}
}

// Permission 权力管理
type Permission struct {
	PermissionModel  model.IPermission
	TransModel model.ITrans
}

// CheckDataInit 检查数据是否初始化
func (a *Permission) CheckDataInit(ctx context.Context) (bool, error) {
	result, err := a.PermissionModel.Query(ctx, schema.PermissionQueryParam{}, schema.PermissionQueryOptions{
		PageParam: &schema.PaginationParam{PageSize: -1},
	})
	if err != nil {
		return false, err
	}
	return result.PageResult.Total > 0, nil
}

// QueryPage 查询分页数据
func (a *Permission) QueryPage(ctx context.Context, params schema.PermissionQueryParam, pp *schema.PaginationParam) ([]*schema.Permission, *schema.PaginationResult, error) {
	result, err := a.PermissionModel.Query(ctx, params, schema.PermissionQueryOptions{
		PageParam: pp,
	})
	if err != nil {
		return nil, nil, err
	}
	return result.Data, result.PageResult, nil
}

// QueryTree 查询权力树
func (a *Permission) QueryTree(ctx context.Context, includeActions, includeResources bool) ([]*schema.PermissionTree, error) {
	result, err := a.PermissionModel.Query(ctx, schema.PermissionQueryParam{}, schema.PermissionQueryOptions{
		IncludeActions:   includeActions,
		IncludeResources: includeResources,
	})
	if err != nil {
		return nil, err
	}
	return result.Data.ToTrees().ToTree(), nil
}

// Get 查询指定数据
func (a *Permission) Get(ctx context.Context, recordID string) (*schema.Permission, error) {
	item, err := a.PermissionModel.Get(ctx, recordID,
		schema.PermissionQueryOptions{
			IncludeResources: true,
			IncludeActions:   true,
		},
	)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	return item, nil
}

// 获取父级路径
func (a *Permission) getParentPath(ctx context.Context, parentID string) (string, error) {
	if parentID == "" {
		return "", nil
	}

	pitem, err := a.PermissionModel.Get(ctx, parentID)
	if err != nil {
		return "", err
	} else if pitem == nil {
		return "", errors.NewBadRequestError("无效的父级节点")
	}

	var parentPath string
	if v := pitem.ParentPath; v != "" {
		parentPath = v + "/"
	}
	parentPath = parentPath + pitem.RecordID
	return parentPath, nil
}

// Create 创建数据
func (a *Permission) Create(ctx context.Context, item schema.Permission) (*schema.Permission, error) {
	parentPath, err := a.getParentPath(ctx, item.ParentID)
	if err != nil {
		return nil, err
	}

	item.ParentPath = parentPath
	item.RecordID = util.MustUUID()
	item.Creator = GetUserID(ctx)
	err = a.PermissionModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return a.Get(ctx, item.RecordID)
}

// Update 更新数据
func (a *Permission) Update(ctx context.Context, recordID string, item schema.Permission) (*schema.Permission, error) {
	if recordID == item.ParentID {
		return nil, errors.NewBadRequestError("不允许使用节点自身作为父级节点")
	}

	oldItem, err := a.PermissionModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	}
	item.ParentPath = oldItem.ParentPath

	err = ExecTrans(ctx, a.TransModel, func(ctx context.Context) error {
		// 如果父级更新，需要更新当前节点及节点下级的父级路径
		if item.ParentID != oldItem.ParentID {
			parentPath, err := a.getParentPath(ctx, item.ParentID)
			if err != nil {
				return err
			}
			item.ParentPath = parentPath

			opath := oldItem.ParentPath
			if opath != "" {
				opath += "/"
			}
			opath += oldItem.RecordID

			result, err := a.PermissionModel.Query(ctx, schema.PermissionQueryParam{
				PrefixParentPath: opath,
			})
			if err != nil {
				return err
			}

			for _, Permission := range result.Data {
				npath := item.ParentPath + Permission.ParentPath[len(opath):]
				err = a.PermissionModel.UpdateParentPath(ctx, Permission.RecordID, npath)
				if err != nil {
					return err
				}
			}
		}

		return a.PermissionModel.Update(ctx, recordID, item)
	})
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, recordID)
}

// Delete 删除数据
func (a *Permission) Delete(ctx context.Context, recordID string) error {
	result, err := a.PermissionModel.Query(ctx, schema.PermissionQueryParam{
		ParentID: &recordID,
	}, schema.PermissionQueryOptions{PageParam: &schema.PaginationParam{PageSize: -1}})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.NewBadRequestError("含有子级权力，不能删除")
	}

	return a.PermissionModel.Delete(ctx, recordID)
}
