package demo

import (
	"context"

	schema "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/gormplus"

	"gin-admin/internal/app/ginadmin/model/gorm/entity"
)

// GetCommonFileDB 获取CommonFile存储
func GetCommonFileDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return entity.GetDBWithModel(ctx, defDB, CommonFile{})
}

// SchemaCommonFile CommonFile对象
type SchemaCommonFile schema.CommonFile

// ToCommonFile 转换为CommonFile实体
func (a SchemaCommonFile) ToCommonFile() *CommonFile {
	item := &CommonFile{
		RecordID:    a.RecordID,
		ContentType: a.ContentType,
		FileName:    a.FileName,
		FilePath:    a.FilePath,
		FileType:    a.FileType,
	}
	return item
}

// CommonFile 通用文件
type CommonFile struct {
	entity.Model
	RecordID    string `gorm:"column:record_id;size:36;index;"`  // 记录内码
	ContentType string `gorm:"column:content_Type;size:500;"`    // 文件类型
	FileName    string `gorm:"column:file_Name;size:200;"`       // 文件名
	FilePath    string `gorm:"column:file_Path;size:1000;index"` // 存放位置
	FileType    int    `gorm:"column:file_Type;"`                // 文件类型
}

func (a CommonFile) String() string {
	return entity.ToString(a)
}

// TableName 表名
func (a CommonFile) TableName() string {
	return a.Model.TableName("Common_File")
}

// ToSchemaCommonFile 转换为CommonFile对象
func (a CommonFile) ToSchemaCommonFile() *schema.CommonFile {
	item := &schema.CommonFile{
		RecordID:    a.RecordID,
		ContentType: a.ContentType,
		FileName:    a.FileName,
		FilePath:    a.FilePath,
		FileType:    a.FileType,
	}
	return item
}

// CommonFiles CommonFile列表
type CommonFiles []*CommonFile

// ToSchemaCommonFiles 转换为CommonFile对象列表
func (a CommonFiles) ToSchemaCommonFiles() []*schema.CommonFile {
	list := make([]*schema.CommonFile, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaCommonFile()
	}
	return list
}
