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

// NewRole 创建角色存储实例
func NewRole(db *gormplus.DB) *Role {
	return &Role{db}
}

// Role 角色存储
type Role struct {
	db *gormplus.DB
}

func (a *Role) getFuncName(name string) string {
	return fmt.Sprintf("gorm.model.Role.%s", name)
}

func (a *Role) getQueryOption(opts ...schema.RoleQueryOptions) schema.RoleQueryOptions {
	var opt schema.RoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Role) Query(ctx context.Context, params schema.RoleQueryParam, opts ...schema.RoleQueryOptions) (*schema.RoleQueryResult, error) {
	span := logger.StartSpan(ctx, "查询数据", a.getFuncName("Query"))
	defer span.Finish()

	db := entity.GetRoleDB(ctx, a.db).DB
	if v := params.RecordIDs; len(v) > 0 {
		db = db.Where("record_id IN(?)", v)
	}
	if v := params.Name; v != "" {
		db = db.Where("name=?", v)
	}
	if v := params.LikeName; v != "" {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.UserID; v != "" {
		subQuery := entity.GetUserRoleDB(ctx, a.db).Where("user_id=?", v).Select("role_id").SubQuery()
		db = db.Where("record_id IN(?)", subQuery)
	}
	db = db.Order("sequence DESC,id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.Roles
	pr, err := WrapPageQuery(db, opt.PageParam, &list)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询数据发生错误")
	}
	qr := &schema.RoleQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaRoles(),
	}

	err = a.fillSchameRoles(ctx, qr.Data, opts...)
	if err != nil {
		return nil, err
	}

	return qr, nil
}

// 填充角色对象
func (a *Role) fillSchameRoles(ctx context.Context, items []*schema.Role, opts ...schema.RoleQueryOptions) error {
	opt := a.getQueryOption(opts...)

	if opt.IncludePermissions {

		roleIDs := make([]string, len(items))
		for i, item := range items {
			roleIDs[i] = item.RecordID
		}

		var PermissionList entity.RolePermissions
		if opt.IncludePermissions {
			items, err := a.queryPermissions(ctx, roleIDs...)
			if err != nil {
				return err
			}
			PermissionList = items
		}

		for i, item := range items {
			if len(PermissionList) > 0 {
				items[i].Permissions = PermissionList.GetByRoleID(item.RecordID)
			}
		}
	}
	return nil
}

// Get 查询指定数据
func (a *Role) Get(ctx context.Context, recordID string, opts ...schema.RoleQueryOptions) (*schema.Role, error) {
	span := logger.StartSpan(ctx, "查询指定数据", a.getFuncName("Get"))
	defer span.Finish()

	var role entity.Role
	ok, err := a.db.FindOne(entity.GetRoleDB(ctx, a.db).Where("record_id=?", recordID), &role)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询指定数据发生错误")
	} else if !ok {
		return nil, nil
	}

	sitem := role.ToSchemaRole()
	err = a.fillSchameRoles(ctx, []*schema.Role{sitem}, opts...)
	if err != nil {
		return nil, err
	}

	return sitem, nil
}

// Create 创建数据
func (a *Role) Create(ctx context.Context, item schema.Role) error {
	span := logger.StartSpan(ctx, "创建数据", a.getFuncName("Create"))
	defer span.Finish()

	return ExecTrans(ctx, a.db, func(ctx context.Context) error {
		sitem := entity.SchemaRole(item)
		result := entity.GetRoleDB(ctx, a.db).Create(sitem.ToRole())
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("创建角色数据发生错误")
		}

		for _, item := range sitem.ToRolePermissions() {
			result := entity.GetRolePermissionDB(ctx, a.db).Create(item)
			if err := result.Error; err != nil {
				span.Errorf(err.Error())
				return errors.New("创建角色权力数据发生错误")
			}
		}

		return nil
	})
}

// 对比并获取需要新增，修改，删除的权力项
func (a *Role) compareUpdatePermission(oldList, newList entity.RolePermissions) (clist, dlist, ulist entity.RolePermissions) {
	oldMap, newMap := oldList.ToMap(), newList.ToMap()

	for _, nitem := range newList {
		if _, ok := oldMap[nitem.PermissionID]; ok {
			ulist = append(ulist, nitem)
			continue
		}
		clist = append(clist, nitem)
	}

	for _, oitem := range oldList {
		if _, ok := newMap[oitem.PermissionID]; !ok {
			dlist = append(dlist, oitem)
		}
	}
	return clist, dlist, ulist
}

// 更新权力数据
func (a *Role) updatePermissions(ctx context.Context, span *logger.Entry, roleID string, items entity.RolePermissions) error {
	list, err := a.queryPermissions(ctx, roleID)
	if err != nil {
		return err
	}

	clist, dlist, ulist := a.compareUpdatePermission(list, items)
	for _, item := range clist {
		result := entity.GetRolePermissionDB(ctx, a.db).Create(item)
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("创建角色权力数据发生错误")
		}
	}

	for _, item := range dlist {
		result := entity.GetRolePermissionDB(ctx, a.db).Where("role_id=? AND permission_id=?", roleID, item.PermissionID).Delete(entity.RolePermission{})
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("删除角色权力数据发生错误")
		}
	}

	for _, item := range ulist {
		result := entity.GetRolePermissionDB(ctx, a.db).Where("role_id=? AND permission_id=?", roleID, item.PermissionID).Omit("role_id", "permission_id").Updates(item)
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("更新角色权力数据发生错误")
		}
	}
	return nil
}

// Update 更新数据
func (a *Role) Update(ctx context.Context, recordID string, item schema.Role) error {
	span := logger.StartSpan(ctx, "更新数据", a.getFuncName("Update"))
	defer span.Finish()

	return ExecTrans(ctx, a.db, func(ctx context.Context) error {
		sitem := entity.SchemaRole(item)
		result := entity.GetRoleDB(ctx, a.db).Where("record_id=?", recordID).Omit("record_id", "creator").Updates(sitem.ToRole())
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("更新角色数据发生错误")
		}

		err := a.updatePermissions(ctx, span, recordID, sitem.ToRolePermissions())
		if err != nil {
			return err
		}

		return nil
	})
}

// Delete 删除数据
func (a *Role) Delete(ctx context.Context, recordID string) error {
	span := logger.StartSpan(ctx, "删除数据", a.getFuncName("Delete"))
	defer span.Finish()

	return ExecTrans(ctx, a.db, func(ctx context.Context) error {
		result := entity.GetRoleDB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.Role{})
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("删除角色数据发生错误")
		}

		result = entity.GetRolePermissionDB(ctx, a.db).Where("role_id=?", recordID).Delete(entity.RolePermission{})
		if err := result.Error; err != nil {
			span.Errorf(err.Error())
			return errors.New("删除角色权力数据发生错误")
		}

		return nil
	})
}

func (a *Role) queryPermissions(ctx context.Context, roleIDs ...string) (entity.RolePermissions, error) {
	span := logger.StartSpan(ctx, "查询角色权力数据", a.getFuncName("queryPermissions"))
	defer span.Finish()

	var list entity.RolePermissions
	result := entity.GetRolePermissionDB(ctx, a.db).Where("role_id IN(?)", roleIDs).Find(&list)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询角色权力数据发生错误")
	}

	return list, nil
}
