package ctl

import (
	"gin-admin/internal/app/ginadmin/bll"
	"gin-admin/internal/app/ginadmin/ginplus"
	"gin-admin/internal/app/ginadmin/schema"
	"gin-admin/pkg/errors"

	"github.com/gin-gonic/gin"
)

// NewRole 创建角色管理控制器
func NewRole(b *bll.Common) *Role {
	return &Role{
		RoleBll:       b.Role,
		PermissionBll: b.Permission,
	}
}

// Role 角色管理
// @Name Role
// @Description 角色管理接口
type Role struct {
	RoleBll       *bll.Role
	PermissionBll *bll.Permission
}

// Query 查询数据
func (a *Role) Query(c *gin.Context) {
	switch c.Query("q") {
	case "page":
		a.QueryPage(c)
	case "list":
		a.QueryList(c)
	case "select":
		a.QuerySelect(c)
	default:
		ginplus.ResError(c, errors.NewBadRequestError("未知的查询类型"))
	}
}

// QueryPage 查询分页数据
// @Summary 查询分页数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" 1
// @Param pageSize query int true "分页大小" 10
// @Param name query string false "角色名称(模糊查询)"
// @Param status query int false "状态(1:启用 2:停用)"
// @Success 200 []schema.Role "分页查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/roles?q=page
func (a *Role) QueryPage(c *gin.Context) {
	var params schema.RoleQueryParam
	params.LikeName = c.Query("name")

	items, pr, err := a.RoleBll.QueryPage(ginplus.NewContext(c), params, ginplus.GetPaginationParam(c))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	// permissions, err := a.PermissionBll.QueryList(ginplus.NewContext(c))
	// if err != nil {
	// 	ginplus.ResError(c, err)
	// 	return
	// }

	ginplus.ResPage(c, items, pr)

	// timeUnix := time.Now().Unix()
	// pageSize := ginplus.GetPageSize(c)
	// response := schema.HTTPResponse{
	// 	Message: "",
	// 	Result: &schema.HTTPRoleResponse{
	// 		HTTPPage: schema.HTTPPage{
	// 			Data:       items,
	// 			PageNo:     ginplus.GetPageIndex(c),
	// 			PageSize:   ginplus.GetPageSize(c),
	// 			TotalPage:  pr.Total / pageSize,
	// 			TotalCount: pr.Total,
	// 		},
	// 		Rules: permissions,
	// 	},
	// 	Status:    200,
	// 	Timestamp: timeUnix,
	// }

	// ginplus.ResSuccess(c, response)
}

// QueryList 查询数据
// @Summary 查询数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param name query string false "角色名称(模糊查询)"
// @Param status query int false "状态(1:启用 2:停用)"
// @Success 200 []schema.Role "分页查询结果：{data:列表数据}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/roles?q=list
func (a *Role) QueryList(c *gin.Context) {
	var params schema.RoleQueryParam
	params.LikeName = c.Query("name")

	items, err := a.RoleBll.QueryList(ginplus.NewContext(c))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResData(c, items)
}

// QuerySelect 查询选择数据
// @Summary 查询选择数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 []schema.Role "{list:角色列表}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/roles?q=select
func (a *Role) QuerySelect(c *gin.Context) {
	items, err := a.RoleBll.QuerySelect(ginplus.NewContext(c))
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
// @Success 200 schema.Role
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 schema.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/roles/{id}
func (a *Role) Get(c *gin.Context) {
	item, err := a.RoleBll.Get(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResData(c, item)
}

// Create 创建数据
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.Role true
// @Success 200 schema.Role
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router POST /api/v1/roles
func (a *Role) Create(c *gin.Context) {
	var item schema.Role
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.RoleBll.Create(ginplus.NewContext(c), item)
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
// @Param body body schema.Role true
// @Success 200 schema.Role
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PUT /api/v1/roles/{id}
func (a *Role) Update(c *gin.Context) {
	var item schema.Role
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.RoleBll.Update(ginplus.NewContext(c), c.Param("id"), item)
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
// @Router DELETE /api/v1/roles/{id}
func (a *Role) Delete(c *gin.Context) {
	err := a.RoleBll.Delete(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
