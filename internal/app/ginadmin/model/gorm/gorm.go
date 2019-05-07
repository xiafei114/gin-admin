package gorm

import (
	"gin-admin/internal/app/ginadmin/model"
	"gin-admin/internal/app/ginadmin/model/gorm/entity"
	demoEntity "gin-admin/internal/app/ginadmin/model/gorm/entity/demo"
	gmodel "gin-admin/internal/app/ginadmin/model/gorm/model"
	"gin-admin/pkg/gormplus"
)

// SetTablePrefix 设定表名前缀
func SetTablePrefix(prefix string) {
	entity.SetTablePrefix(prefix)
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gormplus.DB) error {
	return db.AutoMigrate(
		new(entity.User),
		new(entity.UserRole),
		new(entity.Role),
		new(entity.RolePermission),
		new(entity.Permission),
		new(entity.PermissionAction),
		new(entity.PermissionResource),
		new(entity.Demo),
		new(demoEntity.Product),
		new(demoEntity.ProductCategory),
	).Error
}

// NewModel 创建gorm存储，实现统一的存储接口
func NewModel(db *gormplus.DB) *model.Common {
	return &model.Common{
		Trans: gmodel.NewTrans(db),
		Demo:  gmodel.NewDemo(db),
		Permission:  gmodel.NewPermission(db),
		Role:  gmodel.NewRole(db),
		User:  gmodel.NewUser(db),
	}
}
