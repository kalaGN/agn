package config

import (
	"agn/pkg/config"
)

func init() {

	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{

			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"password": config.Env("REDIS_PASSWORD", ""),

			// 业务类存储使用 1 (图片验证码、短信验证码、会话)
			"database":       config.Env("REDIS_MAIN_DB", 1),
			"database_cache": config.Env("REDIS_CACHE_DB", 1),
		}
	})

	config.Add("limitredis", func() map[string]interface{} {
		return map[string]interface{}{

			"host":     config.Env("LIMITREDIS_HOST", "127.0.0.1"),
			"port":     config.Env("LIMITREDIS_PORT", "6379"),
			"password": config.Env("LIMITREDIS_PASSWORD", ""),

			// 业务类存储使用 1 (图片验证码、短信验证码、会话)
			"database": config.Env("LIMITREDIS_MAIN_DB", 1),
		}
	})
}
