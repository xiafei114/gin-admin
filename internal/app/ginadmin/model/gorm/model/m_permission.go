package model

import (
	"context"
	"fmt"

	"gin-admin/internal/app/ginadmin/model/gorm/entity"
	"gin-admin/internal/app/ginadmin/schema"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/gormplus"
	"gin-admin/pkg/logger"
)

// NewPermission 创建菜单存储实例
func NewPermission(db *gormplus.DB) *Permission {
	return &Permission{db}
}

// Permission 菜单存储
type Permission struct {
	db *gormplus.DB
}

func (a *Permission) getFuncName(name string) string {
	return fmt.Sprintf("gorm.model.Permission.%s", name)
}

func (a *Permission) getQueryOption(opts ...schema.PermissionQueryOptions) schema.PermissionQueryOptions {
	var opt schema.PermissionQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Permission) Query(ctx context.Context, params schema.PermissionQueryParam, opts ...schema.PermissionQueryOptions) (*schema.PermissionQueryResult, error) {
	span := logger.StartSpan(ctx, "查询数据", a.getFuncName("Query"))
	defer span.Finish()

	db := entity.GetPermissionDB(ctx, a.db).DB
	if v := params.RecordIDs; len(v) > 0 {
		db = db.Where("record_id IN(?)", v)
	}
	if v := params.LikeName; v != "" {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.ParentID; v != nil {
		db = db.Where("parent_id=?", *v)
	}
	if v := params.PrefixParentPath; v != "" {
		db = db.Where("parent_path LIKE ?", v+"%")
	}
	if v := params.Hidden; v != nil {
		db = db.Where("hidden=?", *v)
	}
	db = db.Order("sequence DESC,id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.Permissions
	pr, err := WrapPageQuery(db, opt.PageParam, &list)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询数据发生错误")
	}
	qr := &schema.PermissionQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaPermissions(),
	}

	err = a.fillSchemaPermissions(ctx, qr.Data, opts...)
	if err != nil {
		return nil, err
	}

	return qr, nil
}

// 填充菜单对象数据
func (a *Permission) fillSchemaPermissions(ctx context.Context, items []*schema.Permission, opts ...schema.PermissionQueryOptions) error {
	opt := a.getQueryOption(opts...)

	if opt.IncludeActions || opt.IncludeResources {

		PermissionIDs := make([]string, len(items))
		for i, item := range items {
			PermissionIDs[i] = item.RecordID
		}

		var actionList entity.PermissionActions
		var resourceList entity.PermissionResources
		if opt.IncludeActions {
			items, err := a.queryActions(ctx, PermissionIDs...)
			if err != nil {
				return err
			}
			actionList = items
		}

		if opt.IncludeResources {
			items, err := a.queryResources(ctx, PermissionIDs...)
			if err != nil {
				return err
			}
			resourceList = items
		}

		for i, item := range items {
			if len(actionList) > 0 {
				items[i].Actions = actionList.GetByPermissionID(item.RecordID)
			}
			if len(resourceList) > 0 {
				items[i].Resources = resourceList.GetByPermissionID(item.RecordID)
			}
		}
	}

	return nil
}

// Get 查询指定数据
func (a *Permission) Get(ctx context.Context, recordID string, opts ...schema.PermissionQueryOptions) (*schema.Permission, error) {
	span := logger.StartSpan(ctx, "查询指定数据", a.getFuncName("Get"))
	defer span.Finish()

	var item entity.Permission
	ok, err := a.db.FindOne(entity.GetPermissionDB(ctx, a.db).Where("record_id=?", recordID), &item)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询指定数据发生错误")
	} else if !ok {
		return nil, nil
	}

	sitem := item.ToSchemaPermission()
	err = a.fillSchemaPermissions(ctx, []*schema.Permission{sitem}, opts...)
	if err != nil {
		return nil, err
	}

	return sitem, nil
}

// Create 创建数据
func (a *Permission) Create(ctx context.Context, item schema.Permission) error {
	span := logger.StartSpan(ctx, "创建数据", a.getFuncName("Create"))
	defer span.Finish()

	return ExecTrans(ctx, a.db, func(ctx context.Context) error {
		sitem := entity.SchemaPermission(item)
		result := entity.GetPermissionDB(ctx, a.db).Create(sitem.ToPermission())
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("创建菜单数据发生错误")
		}

		for _, item := range sitem.ToPermissionActions() {
			result := entity.GetPermissionActionDB(ctx, a.db).Create(item)
			if err := result.Error; err != nil {
				span.Errorf(err.Error())
				return errors.New("创建菜单动作数据发生错误")
			}
		}

		for _, item := range sitem.ToPermissionResources() {
			result := entity.GetPermissionResourceDB(ctx, a.db).Create(item)
			if err := result.Error; err != nil {
				span.Errorf(err.Error())
				return errors.New("创建菜单资源数据发生错误")
			}
		}

		return nil
	})
}

// 对比并获取需要新增，修改，删除的动作项
func (a *Permission) compareUpdateAction(oldList, newList entity.PermissionActions) (clist, dlist, ulist []*entity.PermissionAction) {
	oldMap, newMap := oldList.ToMap(), newList.ToMap()

	for _, nitem := range newList {
		if _, ok := oldMap[nitem.Code]; ok {
			ulist = append(ulist, nitem)
			continue
		}
		clist = append(clist, nitem)
	}

	for _, oitem := range oldList {
		if _, ok := newMap[oitem.Code]; !ok {
			dlist = append(dlist, oitem)
		}
	}
	return clist, dlist, ulist
}

// 更新动作数据
func (a *Permission) updateActions(ctx context.Context, span *logger.Entry, PermissionID string, items entity.PermissionActions) error {
	list, err := a.queryActions(ctx, PermissionID)
	if err != nil {
		return err
	}

	clist, dlist, ulist := a.compareUpdateAction(list, items)
	for _, item := range clist {
		item.PermissionID = PermissionID
		result := entity.GetPermissionActionDB(ctx, a.db).Create(item)
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("创建菜单动作数据发生错误")
		}
	}

	for _, item := range dlist {
		result := entity.GetPermissionActionDB(ctx, a.db).Where("permission_id=? AND code=?", PermissionID, item.Code).Delete(entity.PermissionAction{})
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("删除菜单动作数据发生错误")
		}
	}

	for _, item := range ulist {
		result := entity.GetPermissionActionDB(ctx, a.db).Where("permission_id=? AND code=?", PermissionID, item.Code).Omit("permission_id", "code").Updates(item)
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("更新菜单动作数据发生错误")
		}
	}
	return nil
}

// 对比并获取需要新增，修改，删除的资源项
func (a *Permission) compareUpdateResource(oldList, newList entity.PermissionResources) (clist, dlist, ulist []*entity.PermissionResource) {
	oldMap, newMap := oldList.ToMap(), newList.ToMap()

	for _, nitem := range newList {
		if _, ok := oldMap[nitem.Code]; ok {
			ulist = append(ulist, nitem)
			continue
		}
		clist = append(clist, nitem)
	}

	for _, oitem := range oldList {
		if _, ok := newMap[oitem.Code]; !ok {
			dlist = append(dlist, oitem)
		}
	}
	return
}

// 更新资源数据
func (a *Permission) updateResources(ctx context.Context, span *logger.Entry, PermissionID string, items entity.PermissionResources) error {
	list, err := a.queryResources(ctx, PermissionID)
	if err != nil {
		return err
	}

	clist, dlist, ulist := a.compareUpdateResource(list, items)
	for _, item := range clist {
		result := entity.GetPermissionResourceDB(ctx, a.db).Create(item)
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("创建菜单资源数据发生错误")
		}
	}

	for _, item := range dlist {
		result := entity.GetPermissionResourceDB(ctx, a.db).Where("permission_id=? AND code=?", PermissionID, item.Code).Delete(entity.PermissionResource{})
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("删除菜单资源数据发生错误")
		}
	}

	for _, item := range ulist {
		result := entity.GetPermissionResourceDB(ctx, a.db).Where("permission_id=? AND code=?", PermissionID, item.Code).Omit("permission_id", "code").Updates(item)
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("更新菜单资源数据发生错误")
		}
	}
	return nil
}

// Update 更新数据
func (a *Permission) Update(ctx context.Context, recordID string, item schema.Permission) error {
	span := logger.StartSpan(ctx, "更新数据", a.getFuncName("Update"))
	defer span.Finish()

	return ExecTrans(ctx, a.db, func(ctx context.Context) error {
		sitem := entity.SchemaPermission(item)
		result := entity.GetPermissionDB(ctx, a.db).Where("record_id=?", recordID).Omit("record_id", "creator").Updates(sitem.ToPermission())
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("更新数据发生错误")
		}

		err := a.updateActions(ctx, span, recordID, sitem.ToPermissionActions())
		if err != nil {
			return err
		}

		err = a.updateResources(ctx, span, recordID, sitem.ToPermissionResources())
		if err != nil {
			return err
		}

		return nil
	})
}

// UpdateParentPath 更新父级路径
func (a *Permission) UpdateParentPath(ctx context.Context, recordID, parentPath string) error {
	span := logger.StartSpan(ctx, "更新父级路径", a.getFuncName("UpdateParentPath"))
	defer span.Finish()

	result := entity.GetPermissionDB(ctx, a.db).Where("record_id=?", recordID).Update("parent_path", parentPath)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("更新父级路径发生错误")
	}
	return nil
}

// Delete 删除数据
func (a *Permission) Delete(ctx context.Context, recordID string) error {
	span := logger.StartSpan(ctx, "删除数据", a.getFuncName("Delete"))
	defer span.Finish()

	return ExecTrans(ctx, a.db, func(ctx context.Context) error {
		result := entity.GetPermissionDB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.Permission{})
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("删除数据发生错误")
		}

		result = entity.GetPermissionActionDB(ctx, a.db).Where("permission_id=?", recordID).Delete(entity.PermissionAction{})
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("删除菜单动作数据发生错误")
		}

		result = entity.GetPermissionResourceDB(ctx, a.db).Where("permission_id=?", recordID).Delete(entity.PermissionResource{})
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("删除菜单资源数据发生错误")
		}
		return nil
	})
}

func (a *Permission) queryActions(ctx context.Context, PermissionIDs ...string) (entity.PermissionActions, error) {
	span := logger.StartSpan(ctx, "查询菜单动作数据", a.getFuncName("queryActions"))
	defer span.Finish()

	var list entity.PermissionActions
	result := entity.GetPermissionActionDB(ctx, a.db).Where("permission_id IN(?)", PermissionIDs).Find(&list)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询菜单动作数据发生错误")
	}

	return list, nil
}

func (a *Permission) queryResources(ctx context.Context, PermissionIDs ...string) (entity.PermissionResources, error) {
	span := logger.StartSpan(ctx, "查询菜单资源数据", a.getFuncName("queryResources"))
	defer span.Finish()

	var list entity.PermissionResources
	result := entity.GetPermissionResourceDB(ctx, a.db).Where("permission_id IN(?)", PermissionIDs).Find(&list)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询菜单资源数据发生错误")
	}

	return list, nil
}
