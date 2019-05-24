package ctl

import (
	"fmt"
	"gin-admin/internal/app/ginadmin/bll"
	demoBll "gin-admin/internal/app/ginadmin/bll/demo"
	"gin-admin/internal/app/ginadmin/ginplus"
	"gin-admin/internal/app/ginadmin/schema"
	schemaProject "gin-admin/internal/app/ginadmin/schema/demo"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/util"
	"gin-admin/pkg/util/upload"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// NewMedia 创建Media控制器
func NewMedia(b *bll.Common) *Media {
	return &Media{
		MediaBll: b.Media,
	}
}

// Media Media
// @Name Media
// @Description Media
type Media struct {
	MediaBll *demoBll.Media
}

// Query 查询数据
func (a *Media) Query(c *gin.Context) {
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
// @Success 200 []schemaProject.Media "查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/medias?q=page
func (a *Media) QueryPage(c *gin.Context) {
	var params schema.CommonQueryParam
	params.LikeCode = c.Query("code")
	params.LikeName = c.Query("name")
	params.Status = util.S(c.Query("status")).Int()

	items, pr, err := a.MediaBll.QueryPage(ginplus.NewContext(c), params, ginplus.GetPaginationParam(c))
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
// @Success 200 schemaProject.Media
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 schema.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/medias/{id}
func (a *Media) Get(c *gin.Context) {
	hostName := GetHostName(c.Request.Referer())
	item, err := a.MediaBll.Get(ginplus.NewContext(c), c.Param("id"), hostName)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// Upload 上传文件
// @Summary 创建数据
// @Produce  json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param fileData form file true "File"
// @Success 200 schemaProject.Media
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router POST /api/v1/medias/upload
func (a *Media) Upload(c *gin.Context) {

	hostName := GetHostName(c.Request.Referer())

	file, header, err := c.Request.FormFile("fileData")
	if err != nil {
		ginplus.ResError(c, err, http.StatusInternalServerError)
		return
	}

	if header == nil {
		ginplus.ResError(c, err, http.StatusBadRequest)
		return
	}

	if !upload.CheckFileSize(file) {
		ginplus.ResError(c, err, http.StatusBadRequest)
		return
	}

	fileName, ext := upload.GetFileName(header.Filename)
	fullPath, relativePath := upload.GetFilePath("media/")
	saveFilePath := fullPath + fileName + ext
	relativeFilePath := relativePath + fileName + ext
	err = upload.CheckImage(fullPath)
	if err != nil {
		ginplus.ResError(c, err, http.StatusInternalServerError)
		return
	}

	if err := c.SaveUploadedFile(header, saveFilePath); err != nil {
		ginplus.ResError(c, err, http.StatusInternalServerError)
		return
	}

	var item schemaProject.CommonFile
	item.RecordID = fileName
	item.FileName = header.Filename
	item.FilePath = relativeFilePath

	nitem, err := a.MediaBll.Upload(ginplus.NewContext(c), item, hostName)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	// ginplus.ResSuccess(c, map[string]string{
	// 	"image_url": upload.GetFileFullUrl(hostName, relativeFilePath),
	// 	"hostName":  hostName,
	// })
	ginplus.ResSuccess(c, nitem)
}

func trim(url string) string {
	url = strings.Replace(url, " ", "", -1)
	return url
}

// GetHostName 获得HostName
func GetHostName(s string) string {
	s = trim(s)
	u, err := url.Parse(s)
	if err != nil {
		log.Panicln(err)
	}

	host := fmt.Sprintf("%s://%s", u.Scheme, u.Host)

	return host
}

// Update 更新数据
// @Summary 更新数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Param body body schemaProject.Media true
// @Success 200 schemaProject.Media
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PUT /api/v1/medias/{id}
func (a *Media) Update(c *gin.Context) {
	hostName := GetHostName(c.Request.Referer())
	var item schemaProject.Media
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.MediaBll.Update(ginplus.NewContext(c), c.Param("id"), item, hostName)
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
// @Router DELETE /api/v1/medias/{id}
func (a *Media) Delete(c *gin.Context) {
	err := a.MediaBll.Delete(ginplus.NewContext(c), c.Param("id"))
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
// @Router PATCH /api/v1/medias/{id}/enable
func (a *Media) Enable(c *gin.Context) {
	err := a.MediaBll.UpdateStatus(ginplus.NewContext(c), c.Param("id"), 1)
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
// @Router PATCH /api/v1/medias/{id}/disable
func (a *Media) Disable(c *gin.Context) {
	err := a.MediaBll.UpdateStatus(ginplus.NewContext(c), c.Param("id"), 2)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
