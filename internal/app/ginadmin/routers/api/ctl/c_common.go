package ctl

import (
	"gin-admin/internal/app/ginadmin/bll"
	demoBll "gin-admin/internal/app/ginadmin/bll/demo"
	"gin-admin/internal/app/ginadmin/ginplus"
	"gin-admin/internal/app/ginadmin/schema"
	schemaProject "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/util"

	"github.com/gin-gonic/gin"
)

// NewCommon 创建Common控制器
func NewCommon(b *bll.Common) *Common {
	return &Common{
		CommonBll: b.Common,
	}
}

// Common Common
// @Name Common
// @Description Common
type Common struct {
	CommonBll *demoBll.Common
}

// Query 查询数据
func (a *Common) Query(c *gin.Context) {
	switch c.Query("q") {
	case "page":
		a.QueryPage(c)
	default:
		ginplus.ResError(c, errors.NewBadRequestError("未知的查询类型"))
	}
}

// QueryPage 查询分页数据
// @Summary 查询分页数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" 1
// @Param pageSize query int true "分页大小" 10
// @Param code query string false "编号"
// @Param name query string false "名称"
// @Param status query int false "状态(1:启用 2:停用)"
// @Success 200 []schemaProject.CommonFile "查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/commons?q=page
func (a *Common) QueryPage(c *gin.Context) {
	var params schema.CommonQueryParam
	params.LikeCode = c.Query("code")
	params.LikeName = c.Query("name")
	params.Status = util.S(c.Query("status")).Int()

	items, pr, err := a.CommonBll.QueryPage(ginplus.NewContext(c), params, ginplus.GetPaginationParam(c))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResPage(c, items, pr)
}

// Get 查询指定数据
// @Summary 查询指定数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schemaProject.CommonFile
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 schema.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/commons/{id}
func (a *Common) Get(c *gin.Context) {
	item, err := a.CommonBll.Get(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// Create 创建数据
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schemaProject.CommonFile true
// @Success 200 schemaProject.CommonFile
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router POST /api/v1/commons/upload
func (a *Common) Upload(c *gin.Context) {
	var item schemaProject.CommonFile
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.CommonBll.Create(ginplus.NewContext(c), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
}

// Create 创建数据
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schemaProject.CommonFile true
// @Success 200 schemaProject.CommonFile
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router POST /api/v1/commons
func (a *Common) Create(c *gin.Context) {
	var item schemaProject.CommonFile
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.CommonBll.Create(ginplus.NewContext(c), item)
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
// @Param body body schemaProject.CommonFile true
// @Success 200 schemaProject.CommonFile
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PUT /api/v1/commons/{id}
func (a *Common) Update(c *gin.Context) {
	var item schemaProject.CommonFile
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.CommonBll.Update(ginplus.NewContext(c), c.Param("id"), item)
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
// @Router DELETE /api/v1/commons/{id}
func (a *Common) Delete(c *gin.Context) {
	err := a.CommonBll.Delete(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Enable 启用数据
// @Summary 启用数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PATCH /api/v1/commons/{id}/enable
func (a *Common) Enable(c *gin.Context) {
	err := a.CommonBll.UpdateStatus(ginplus.NewContext(c), c.Param("id"), 1)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Disable 禁用数据
// @Summary 禁用数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PATCH /api/v1/commons/{id}/disable
func (a *Common) Disable(c *gin.Context) {
	err := a.CommonBll.UpdateStatus(ginplus.NewContext(c), c.Param("id"), 2)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
