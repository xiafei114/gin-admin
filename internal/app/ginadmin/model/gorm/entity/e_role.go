package entity

import (
	"context"
	"strings"

	"gin-admin/internal/app/ginadmin/schema"
	"gin-admin/pkg/gormplus"
)

// GetRoleDB 获取角色存储
func GetRoleDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return GetDBWithModel(ctx, defDB, Role{})
}

// GetRolePermissionDB 获取角色权力关联存储
func GetRolePermissionDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return GetDBWithModel(ctx, defDB, RolePermission{})
}

// SchemaRole 角色对象
type SchemaRole schema.Role

// ToRole 转换为角色实体
func (a SchemaRole) ToRole() *Role {
	item := &Role{
		RecordID:  a.RecordID,
		IndexCode: a.IndexCode,
		Name:      a.Name,
		Sequence:  a.Sequence,
		Memo:      a.Memo,
		Status:    &a.Status,
		Creator:   a.Creator,
	}
	return item
}

// ToRolePermissions 转换为角色权力实体列表
func (a SchemaRole) ToRolePermissions() []*RolePermission {
	list := make([]*RolePermission, len(a.Permissions))
	for i, item := range a.Permissions {
		list[i] = &RolePermission{
			RoleID:       a.RecordID,
			PermissionID: item.PermissionID,
			Action:       strings.Join(item.Actions, ","),
			Resource:     strings.Join(item.Resources, ","),
		}
	}
	return list
}

// Role 角色实体
type Role struct {
	Model
	RecordID  string `gorm:"column:record_id;size:36;index;"`  // 记录内码
	IndexCode string `gorm:"column:index_code;size:50;index;"` // 唯一编码
	Name      string `gorm:"column:name;size:100;index;"`      // 角色名称
	Sequence  int    `gorm:"column:sequence;index;"`           // 排序值
	Memo      string `gorm:"column:memo;size:200;"`            // 备注
	Status    *int   `gorm:"column:status;index;"`             // 状态(0:不隐藏 1:隐藏)
	Creator   string `gorm:"column:creator;size:36;"`          // 创建者
}

func (a Role) String() string {
	return ToString(a)
}

// TableName 表名
func (a Role) TableName() string {
	return a.Model.TableName("role")
}

// ToSchemaRole 转换为角色对象
func (a Role) ToSchemaRole() *schema.Role {
	item := &schema.Role{
		RecordID:  a.RecordID,
		IndexCode: a.IndexCode,
		Name:      a.Name,
		Sequence:  a.Sequence,
		Memo:      a.Memo,
		Status:    *a.Status,
		Creator:   a.Creator,
		CreatedAt: &a.CreatedAt,
		UpdatedAt: &a.UpdatedAt,
	}
	return item
}

// Roles 角色实体列表
type Roles []*Role

// ToSchemaRoles 转换为角色对象列表
func (a Roles) ToSchemaRoles() []*schema.Role {
	list := make([]*schema.Role, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaRole()
	}
	return list
}

// SchemaRolePermission 角色权力对象
type SchemaRolePermission schema.RolePermission

// ToRolePermission 转换为角色权力实体
func (a SchemaRolePermission) ToRolePermission(roleID string) *RolePermission {
	return &RolePermission{
		RoleID:       roleID,
		PermissionID: a.PermissionID,
		Action:       strings.Join(a.Actions, ","),
		Resource:     strings.Join(a.Resources, ","),
	}
}

// RolePermission 角色权力关联实体
type RolePermission struct {
	Model
	RoleID       string `gorm:"column:role_id;size:36;index;"`       // 角色内码
	PermissionID string `gorm:"column:permission_id;size:36;index;"` // 权力内码
	Action       string `gorm:"column:action;size:2048;"`            // 动作权限(多个以英文逗号分隔)
	Resource     string `gorm:"column:resource;size:2048;"`          // 资源权限(多个以英文逗号分隔)
}

// TableName 表名
func (a RolePermission) TableName() string {
	return a.Model.TableName("role_Permission")
}

// ToSchemaRolePermission 转换为角色权力对象
func (a RolePermission) ToSchemaRolePermission() *schema.RolePermission {
	return &schema.RolePermission{
		PermissionID: a.PermissionID,
		Actions:      strings.Split(a.Action, ","),
		Resources:    strings.Split(a.Resource, ","),
	}
}

// RolePermissions 角色权力关联实体列表
type RolePermissions []*RolePermission

// GetByRoleID 根据角色ID获取角色权力对象列表
func (a RolePermissions) GetByRoleID(roleID string) []*schema.RolePermission {
	var list []*schema.RolePermission
	for _, item := range a {
		if item.RoleID == roleID {
			list = append(list, item.ToSchemaRolePermission())
		}
	}
	return list
}

// ToSchemaRolePermissions 转换为角色权力对象列表
func (a RolePermissions) ToSchemaRolePermissions() []*schema.RolePermission {
	list := make([]*schema.RolePermission, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaRolePermission()
	}
	return list
}

// ToMap 转换为键值映射
func (a RolePermissions) ToMap() map[string]*RolePermission {
	m := make(map[string]*RolePermission)
	for _, item := range a {
		m[item.PermissionID] = item
	}
	return m
}
