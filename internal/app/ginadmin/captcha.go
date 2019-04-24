package ginadmin

import (
	"gin-admin/internal/app/ginadmin/config"
	"gin-admin/pkg/logger"

	"github.com/LyricTian/captcha"
	"github.com/LyricTian/captcha/store"
)

// InitCaptcha 初始化图形验证码
func InitCaptcha() {
	cfg := config.GetGlobalConfig().Captcha
	if cfg.Store == "redis" {
		rc := config.GetGlobalConfig().Redis
		captcha.SetCustomStore(store.NewRedisStore(&store.RedisOptions{
			Addr:     rc.Addr,
			Password: rc.Password,
			DB:       cfg.RedisDB,
		}, captcha.Expiration, logger.StandardLogger(), cfg.RedisPrefix))
	}
}
