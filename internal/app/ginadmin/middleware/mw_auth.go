package middleware

import (
	"gin-admin/internal/app/ginadmin/config"
	"gin-admin/internal/app/ginadmin/ginplus"
	"gin-admin/pkg/auth"
	"gin-admin/pkg/errors"
	"gin-admin/pkg/logger"

	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auther, skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		if t := ginplus.GetToken(c); t != "" {
			id, err := a.ParseUserID(t)
			if err != nil {
				if err == auth.ErrInvalidToken {
					ginplus.ResError(c, errors.NewUnauthorizedError())
					return
				}
				logger.StartSpan(ginplus.NewContext(c), "用户授权中间件", "UserAuthMiddleware").Errorf(err.Error())
				ginplus.ResError(c, errors.NewInternalServerError())
				return
			}
			userID = id
		}

		if userID != "" {
			c.Set(ginplus.UserIDKey, userID)
		}

		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		if userID == "" {
			if config.GetGlobalConfig().RunMode == "debug" {
				c.Set(ginplus.UserIDKey, config.GetGlobalConfig().Root.UserName)
				c.Next()
				return
			}
			ginplus.ResError(c, errors.NewUnauthorizedError("用户未登录"))
		}
	}
}
