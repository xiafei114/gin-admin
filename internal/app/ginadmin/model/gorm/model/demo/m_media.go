package model

import (
	"context"
	"fmt"

	entity "gin-admin/internal/app/ginadmin/model/gorm/entity/demo"
	model "gin-admin/internal/app/ginadmin/model/gorm/model"
	schema "gin-admin/internal/app/ginadmin/schema"
	schemaProject "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/gormplus"
	"gin-admin/pkg/logger"
)

// NewMedia 创建Media存储实例
func NewMedia(db *gormplus.DB) *Media {
	return &Media{db}
}

// Media Media存储
type Media struct {
	db *gormplus.DB
}

func (a *Media) getFuncName(name string) string {
	return fmt.Sprintf("gorm.model.Media.%s", name)
}

func (a *Media) getQueryOption(opts ...schema.CommonQueryOptions) schema.CommonQueryOptions {
	var opt schema.CommonQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Media) Query(ctx context.Context, params schema.CommonQueryParam, opts ...schema.CommonQueryOptions) (*schema.CommonQueryResult, error) {
	span := logger.StartSpan(ctx, "查询数据", a.getFuncName("Query"))
	defer span.Finish()

	db := entity.GetMediaDB(ctx, a.db).DB
	if v := params.Code; v != "" {
		db = db.Where("code=?", v)
	}
	if v := params.LikeCode; v != "" {
		db = db.Where("code LIKE ?", "%"+v+"%")
	}
	if v := params.LikeName; v != "" {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}
	db = db.Order("id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.Medias
	pr, err := model.WrapPageQuery(db, opt.PageParam, &list)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询数据发生错误")
	}
	qr := &schema.CommonQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaMedias(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Media) Get(ctx context.Context, recordID string, opts ...schema.CommonQueryOptions) (*schemaProject.Media, error) {
	span := logger.StartSpan(ctx, "查询指定数据", a.getFuncName("Get"))
	defer span.Finish()

	db := entity.GetMediaDB(ctx, a.db).Where("record_id=?", recordID)
	var item entity.Media
	ok, err := a.db.FindOne(db, &item)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询指定数据发生错误")
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaMedia(), nil
}

// Create 创建数据
func (a *Media) Create(ctx context.Context, item schemaProject.Media) error {
	span := logger.StartSpan(ctx, "创建数据", a.getFuncName("Create"))
	defer span.Finish()

	Media := entity.SchemaMedia(item).ToMedia()
	result := entity.GetMediaDB(ctx, a.db).Create(Media)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("创建数据发生错误")
	}
	return nil
}

// Update 更新数据
func (a *Media) Update(ctx context.Context, recordID string, item schemaProject.Media) error {
	span := logger.StartSpan(ctx, "更新数据", a.getFuncName("Update"))
	defer span.Finish()

	Media := entity.SchemaMedia(item).ToMedia()
	result := entity.GetMediaDB(ctx, a.db).Where("record_id=?", recordID).Omit("record_id", "creator").Updates(Media)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("更新数据发生错误")
	}
	return nil
}

// Delete 删除数据
func (a *Media) Delete(ctx context.Context, recordID string) error {
	span := logger.StartSpan(ctx, "删除数据", a.getFuncName("Delete"))
	defer span.Finish()

	result := entity.GetMediaDB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.Media{})
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("删除数据发生错误")
	}
	return nil
}

// UpdateStatus 更新状态
func (a *Media) UpdateStatus(ctx context.Context, recordID string, status int) error {
	span := logger.StartSpan(ctx, "更新状态", a.getFuncName("UpdateStatus"))
	defer span.Finish()

	result := entity.GetMediaDB(ctx, a.db).Where("record_id=?", recordID).Update("status", status)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("更新状态发生错误")
	}
	return nil
}
