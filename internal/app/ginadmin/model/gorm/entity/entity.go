package entity

import (
	"context"
	"fmt"
	"time"

	icontext "gin-admin/internal/app/ginadmin/context"
	"gin-admin/pkg/gormplus"
	"gin-admin/pkg/util"
	// uuid "github.com/satori/go.uuid"
)

// 表名前缀
var tablePrefix string

// SetTablePrefix 设定表名前缀
func SetTablePrefix(prefix string) {
	tablePrefix = prefix
}

// GetTablePrefix 获取表名前缀
func GetTablePrefix() string {
	return tablePrefix
}

// Model base model
type Model struct {
	ID uint `gorm:"column:id;primary_key;auto_increment;"`
	// ID        uuid.UUID  `gorm:"column:id;type:char(36); primary_key"`
	CreatedAt time.Time  `gorm:"column:created_at;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;"`
}

// TableName table name
func (Model) TableName(name string) string {
	return fmt.Sprintf("%s%s", GetTablePrefix(), name)
}

func ToString(v interface{}) string {
	return util.JSONMarshalToString(v)
}

func GetDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	trans, ok := icontext.FromTrans(ctx)
	if ok {
		db, ok := trans.(*gormplus.DB)
		if ok {
			return db
		}
	}
	return defDB
}

func GetDBWithModel(ctx context.Context, defDB *gormplus.DB, m interface{}) *gormplus.DB {
	return gormplus.Wrap(GetDB(ctx, defDB).Model(m))
}
