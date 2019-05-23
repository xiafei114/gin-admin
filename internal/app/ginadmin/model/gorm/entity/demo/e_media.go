package demo

import (
	"context"

	schema "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/gormplus"

	"gin-admin/internal/app/ginadmin/model/gorm/entity"
)

// GetMediaDB 获取Media存储
func GetMediaDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return entity.GetDBWithModel(ctx, defDB, Media{})
}

// SchemaMedia Media对象
type SchemaMedia schema.Media

// ToMedia 转换为Media实体
func (a SchemaMedia) ToMedia() *Media {
	item := &Media{
		RecordID:     a.RecordID,
		CommonFileID: a.CommonFileID,
		InfoNo:       a.InfoNo,
		InfoDesc:     a.InfoDesc,
	}
	return item
}

// Media 通用文件
type Media struct {
	entity.Model
	RecordID     string `gorm:"column:record_id;size:36;index;"`     // 记录内码
	CommonFileID string `gorm:"column:common_file_id;size:36;index"` // 类型（使用弹出选择）
	InfoNo       string `gorm:"column:info_no;size:200;"`            // 文件名
	InfoDesc     string `gorm:"column:info_desc;size:1000;index"`    // 存放位置
}

func (a Media) String() string {
	return entity.ToString(a)
}

// TableName 表名
func (a Media) TableName() string {
	return a.Model.TableName("Media")
}

// ToSchemaMedia 转换为Media对象
func (a Media) ToSchemaMedia() *schema.Media {
	item := &schema.Media{
		RecordID:     a.RecordID,
		CommonFileID: a.CommonFileID,
		InfoNo:       a.InfoNo,
		InfoDesc:     a.InfoDesc,
	}
	return item
}

// Medias Media列表
type Medias []*Media

// ToSchemaMedias 转换为Media对象列表
func (a Medias) ToSchemaMedias() []*schema.Media {
	list := make([]*schema.Media, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaMedia()
	}
	return list
}
