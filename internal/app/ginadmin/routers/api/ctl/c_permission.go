package ctl

import (
	"gin-admin/internal/app/ginadmin/bll"
	"gin-admin/internal/app/ginadmin/ginplus"
	"gin-admin/internal/app/ginadmin/schema"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/util"

	"github.com/gin-gonic/gin"
)

// NewPermission 创建权限管理控制器
func NewPermission(b *bll.Common) *Permission {
	return &Permission{
		PermissionBll: b.Permission,
	}
}

// Permission 权力管理
// @Name Permission
// @Description 权力管理接口
type Permission struct {
	PermissionBll *bll.Permission
}

// Query 查询数据
func (a *Permission) Query(c *gin.Context) {
	switch c.Query("q") {
	case "page":
		a.QueryPage(c)
	case "list":
		a.QueryList(c)
	default:
		ginplus.ResError(c, errors.NewBadRequestError("未知的查询类型"))
	}
}

// QueryPage 查询分页数据
// @Summary 查询分页数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" 1
// @Param pageSize query int true "分页大小" 10
// @Param name query string false "名称"
// @Param hidden query int false "隐藏权限(0:不隐藏 1:隐藏)"
// @Param parent_id query string false "父级ID"
// @Success 200 []schema.Permission "分页查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/permission?q=page
func (a *Permission) QueryPage(c *gin.Context) {
	params := schema.PermissionQueryParam{
		LikeName: c.Query("name"),
	}

	if v := c.Query("parent_id"); v != "" {
		params.ParentID = &v
	}

	if v := c.Query("status"); v != "" {
		if status := util.S(v).Int(); status > -1 {
			params.Status = &status
		}
	}

	items, pr, err := a.PermissionBll.QueryPage(ginplus.NewContext(c), params, ginplus.GetPaginationParam(c))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResPage(c, items, pr)
}

// QueryList 查询数据
// @Summary 查询数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 []schema.Permission "分页查询结果：{data:列表数据}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/permission?q=list
func (a *Permission) QueryList(c *gin.Context) {
	items, err := a.PermissionBll.QueryList(ginplus.NewContext(c))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResData(c, items)
}

// Get 查询指定数据
// @Summary 查询指定数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.Permission
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 schema.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/permission/{id}
func (a *Permission) Get(c *gin.Context) {
	item, err := a.PermissionBll.Get(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResData(c, item)
}

// Create 创建数据
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.Permission true
// @Success 200 schema.Permission
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router POST /api/v1/permission
func (a *Permission) Create(c *gin.Context) {
	var item schema.Permission
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.PermissionBll.Create(ginplus.NewContext(c), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
}

// Update 更新数据
// @Summary 更新数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Param body body schema.Permission true
// @Success 200 schema.Permission
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PUT /api/v1/permission/{id}
func (a *Permission) Update(c *gin.Context) {
	var item schema.Permission
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.PermissionBll.Update(ginplus.NewContext(c), c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
}

// Delete 删除数据
// @Summary 删除数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router DELETE /api/v1/permission/{id}
func (a *Permission) Delete(c *gin.Context) {
	err := a.PermissionBll.Delete(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
