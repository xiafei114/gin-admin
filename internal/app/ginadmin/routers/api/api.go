package api

import (
	"gin-admin/internal/app/ginadmin/bll"
	"gin-admin/internal/app/ginadmin/middleware"
	"gin-admin/internal/app/ginadmin/routers/api/ctl"

	"gin-admin/pkg/auth"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册/api路由
func RegisterRouter(app *gin.Engine, b *bll.Common, a auth.Auther, enforcer *casbin.Enforcer) {
	g := app.Group("/api")

	// 用户身份授权
	g.Use(middleware.UserAuthMiddleware(
		a,
		middleware.AllowMethodAndPathPrefixSkipper(
			middleware.JoinRouter("GET", "/api/v1/login"),
			middleware.JoinRouter("POST", "/api/v1/login"),
		),
	))

	// casbin权限校验中间件
	g.Use(middleware.CasbinMiddleware(enforcer,
		middleware.AllowMethodAndPathPrefixSkipper(
			middleware.JoinRouter("GET", "/api/v1/login"),
			middleware.JoinRouter("POST", "/api/v1/login"),
			middleware.JoinRouter("POST", "/api/v1/refresh_token"),
			middleware.JoinRouter("PUT", "/api/v1/current/password"),
			middleware.JoinRouter("GET", "/api/v1/current/user"),
			// middleware.JoinRouter("GET", "/api/v1/current/permissiontree"),
		),
	))

	// 请求频率限制中间件
	g.Use(middleware.RateLimiterMiddleware())

	demoCCtl := ctl.NewDemo(b)
	loginCtl := ctl.NewLogin(b)
	PermissionCtl := ctl.NewPermission(b)
	roleCtl := ctl.NewRole(b)
	userCtl := ctl.NewUser(b)

	productCtl := ctl.NewProduct(b)

	v1 := g.Group("/v1")
	{
		// 注册/api/v1/login
		v1.GET("/login/captchaid", loginCtl.GetCaptchaID)
		v1.GET("/login/captcha", loginCtl.GetCaptcha)
		v1.POST("/login", loginCtl.Login)
		v1.POST("/login/exit", loginCtl.Logout)

		// 注册/api/v1/refresh_token
		v1.POST("/refresh_token", loginCtl.RefreshToken)

		// 注册/api/v1/current
		v1.PUT("/current/password", loginCtl.UpdatePassword)
		v1.GET("/current/user", loginCtl.GetUserInfo)
		// v1.GET("/current/Permissiontree", loginCtl.QueryUserPermissionTree)

		// 注册/api/v1/demos
		v1.GET("/demos", demoCCtl.Query)
		v1.GET("/demos/:id", demoCCtl.Get)
		v1.POST("/demos", demoCCtl.Create)
		v1.PUT("/demos/:id", demoCCtl.Update)
		v1.DELETE("/demos/:id", demoCCtl.Delete)
		v1.PATCH("/demos/:id/enable", demoCCtl.Enable)
		v1.PATCH("/demos/:id/disable", demoCCtl.Disable)

		// 注册/api/v1/permission
		v1.GET("/permission", PermissionCtl.Query)
		v1.GET("/permission/:id", PermissionCtl.Get)
		v1.POST("/permission", PermissionCtl.Create)
		v1.PUT("/permission/:id", PermissionCtl.Update)
		v1.DELETE("/permission/:id", PermissionCtl.Delete)

		// 注册/api/v1/roles
		v1.GET("/roles", roleCtl.Query)
		v1.GET("/roles/:id", roleCtl.Get)
		v1.POST("/roles", roleCtl.Create)
		v1.PUT("/roles/:id", roleCtl.Update)
		v1.DELETE("/roles/:id", roleCtl.Delete)

		// 注册/api/v1/users
		v1.GET("/users", userCtl.Query)
		v1.GET("/users/:id", userCtl.Get)
		v1.POST("/users", userCtl.Create)
		v1.PUT("/users/:id", userCtl.Update)
		v1.DELETE("/users/:id", userCtl.Delete)
		v1.PATCH("/users/:id/enable", userCtl.Enable)
		v1.PATCH("/users/:id/disable", userCtl.Disable)

		// 注册/api/v1/product
		v1.GET("/product", productCtl.Query)
		v1.GET("/product/:id", productCtl.Get)
		v1.POST("/products", productCtl.Create)
		v1.PUT("/products/:id", productCtl.Update)
		v1.DELETE("/products/:id", productCtl.Delete)
		v1.PATCH("/products/:id/enable", productCtl.Enable)
		v1.PATCH("/products/:id/disable", productCtl.Disable)
	}
}
