package entity

import (
	"context"
	"strconv"

	"gin-admin/internal/app/ginadmin/schema"
	"gin-admin/pkg/gormplus"
)

// GetPermissionDB 获取权力存储
func GetPermissionDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return GetDBWithModel(ctx, defDB, Permission{})
}

// GetPermissionActionDB 获取权力动作存储
func GetPermissionActionDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return GetDBWithModel(ctx, defDB, PermissionAction{})
}

// GetPermissionResourceDB 获取权力资源存储
func GetPermissionResourceDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return GetDBWithModel(ctx, defDB, PermissionResource{})
}

// SchemaPermission 权力对象
type SchemaPermission schema.Permission

// ToPermission 转换为权力实体
func (a SchemaPermission) ToPermission() *Permission {
	item := &Permission{
		RecordID:   a.RecordID,
		Name:       a.Name,
		Sequence:   a.Sequence,
		Icon:       a.Icon,
		Router:     a.Router,
		Hidden:     &a.Hidden,
		ParentID:   a.ParentID,
		ParentPath: a.ParentPath,
		Creator:    a.Creator,
	}
	return item
}

// ToPermissionActions 转换为权力动作列表
func (a SchemaPermission) ToPermissionActions() []*PermissionAction {
	list := make([]*PermissionAction, len(a.Actions))
	for i, item := range a.Actions {
		list[i] = SchemaPermissionAction(*item).ToPermissionAction(a.RecordID)
	}
	return list
}

// ToPermissionResources 转换为权力资源列表
func (a SchemaPermission) ToPermissionResources() []*PermissionResource {
	list := make([]*PermissionResource, len(a.Resources))
	for i, item := range a.Resources {
		list[i] = SchemaPermissionResource(*item).ToPermissionResource(a.RecordID)
	}
	return list
}

// Permission 权力实体
type Permission struct {
	Model
	RecordID   string `gorm:"column:record_id;size:36;index;"`    // 记录内码
	Name       string `gorm:"column:name;size:50;index;"`         // 权力名称
	Sequence   int    `gorm:"column:sequence;index;"`             // 排序值
	Icon       string `gorm:"column:icon;size:255;"`              // 权力图标
	Router     string `gorm:"column:router;size:255;"`            // 访问路由
	Hidden     *int   `gorm:"column:hidden;index;"`               // 隐藏权力(0:不隐藏 1:隐藏)
	ParentID   string `gorm:"column:parent_id;size:36;index;"`    // 父级内码
	ParentPath string `gorm:"column:parent_path;size:518;index;"` // 父级路径
	Creator    string `gorm:"column:creator;size:36;"`            // 创建人
}

func (a Permission) String() string {
	return ToString(a)
}

// TableName 表名
func (a Permission) TableName() string {
	return a.Model.TableName("permission")
}

// ToSchemaPermission 转换为权力对象
func (a Permission) ToSchemaPermission() *schema.Permission {
	item := &schema.Permission{
		RecordID:   a.RecordID,
		Name:       a.Name,
		Sequence:   a.Sequence,
		Icon:       a.Icon,
		Router:     a.Router,
		Hidden:     *a.Hidden,
		ParentID:   a.ParentID,
		ParentPath: a.ParentPath,
		Creator:    a.Creator,
		CreatedAt:  &a.CreatedAt,
		UpdatedAt:  &a.UpdatedAt,
	}
	return item
}

// Permissions 权力实体列表
type Permissions []*Permission

// ToSchemaPermissions 转换为权力对象列表
func (a Permissions) ToSchemaPermissions() []*schema.Permission {
	list := make([]*schema.Permission, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaPermission()
	}
	return list
}

// SchemaPermissionAction 权力动作对象
type SchemaPermissionAction schema.PermissionAction

// ToPermissionAction 转换为权力动作实体
func (a SchemaPermissionAction) ToPermissionAction(PermissionID string) *PermissionAction {
	return &PermissionAction{
		PermissionID: PermissionID,
		Code:         a.Code,
		Name:         a.Name,
	}
}

// PermissionAction 权力动作关联实体
type PermissionAction struct {
	Model
	PermissionID string `gorm:"column:permission_id;size:36;index;"` // 权力ID
	Code         string `gorm:"column:code;size:50;index;"`          // 动作编号
	Name         string `gorm:"column:name;size:50;"`                // 动作名称
}

// TableName 表名
func (a PermissionAction) TableName() string {
	return a.Model.TableName("permission_action")
}

// ToSchemaPermissionAction 转换为权力动作对象
func (a PermissionAction) ToSchemaPermissionAction() *schema.PermissionAction {
	return &schema.PermissionAction{
		ID:   strconv.Itoa(int(a.ID)),
		Code: a.Code,
		Name: a.Name,
	}
}

// PermissionActions 权力动作关联实体列表
type PermissionActions []*PermissionAction

// GetByPermissionID 根据权力ID获取权力动作列表
func (a PermissionActions) GetByPermissionID(PermissionID string) []*schema.PermissionAction {
	var list []*schema.PermissionAction
	for _, item := range a {
		if item.PermissionID == PermissionID {
			list = append(list, item.ToSchemaPermissionAction())
		}
	}
	return list
}

// ToSchemaPermissionActions 转换为权力动作列表
func (a PermissionActions) ToSchemaPermissionActions() []*schema.PermissionAction {
	list := make([]*schema.PermissionAction, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaPermissionAction()
	}
	return list
}

// ToMap 转换为键值映射
func (a PermissionActions) ToMap() map[string]*PermissionAction {
	m := make(map[string]*PermissionAction)
	for _, item := range a {
		m[item.Code] = item
	}
	return m
}

// SchemaPermissionResource 权力资源对象
type SchemaPermissionResource schema.PermissionResource

// ToPermissionResource 转换为权力资源实体
func (a SchemaPermissionResource) ToPermissionResource(PermissionID string) *PermissionResource {
	return &PermissionResource{
		PermissionID: PermissionID,
		Code:         a.Code,
		Name:         a.Name,
		Method:       a.Method,
		Path:         a.Path,
	}
}

// PermissionResource 权力资源关联实体
type PermissionResource struct {
	Model
	PermissionID string `gorm:"column:permission_id;size:36;index;"` // 权力ID
	Code         string `gorm:"column:code;size:50;index;"`          // 资源编号
	Name         string `gorm:"column:name;size:50;"`                // 资源名称
	Method       string `gorm:"column:method;size:50;"`              // 请求方式
	Path         string `gorm:"column:path;size:255;"`               // 请求路径
}

// TableName 表名
func (a PermissionResource) TableName() string {
	return a.Model.TableName("permission_resource")
}

// ToSchemaPermissionResource 转换为权力资源对象
func (a PermissionResource) ToSchemaPermissionResource() *schema.PermissionResource {
	return &schema.PermissionResource{
		Code:   a.Code,
		Name:   a.Name,
		Method: a.Method,
		Path:   a.Path,
	}
}

// PermissionResources 权力资源关联实体列表
type PermissionResources []*PermissionResource

// GetByPermissionID 根据权力ID获取权力资源列表
func (a PermissionResources) GetByPermissionID(PermissionID string) []*schema.PermissionResource {
	var list []*schema.PermissionResource
	for _, item := range a {
		if item.PermissionID == PermissionID {
			list = append(list, item.ToSchemaPermissionResource())
		}
	}
	return list
}

// ToSchemaPermissionResources 转换为权力资源列表
func (a PermissionResources) ToSchemaPermissionResources() []*schema.PermissionResource {
	list := make([]*schema.PermissionResource, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaPermissionResource()
	}
	return list
}

// ToMap 转换为键值映射
func (a PermissionResources) ToMap() map[string]*PermissionResource {
	m := make(map[string]*PermissionResource)
	for _, item := range a {
		m[item.Code] = item
	}
	return m
}
