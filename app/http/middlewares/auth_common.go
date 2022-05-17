// Package middlewares Gin 中间件
package middlewares

import (
	"agn/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthCommon() gin.HandlerFunc {
	return func(c *gin.Context) {

		//
		partner_key := c.Query("partner_key")
		timestr := c.Query("timestr")
		interface_sign := c.Query("interface_sign")
		// JWT 解析失败，有错误发生
		if len(partner_key) == 0 || len(timestr) == 0 || len(interface_sign) == 0 {
			response.Unauthorized(c, "param error")
			return
		}
		//TODO check sign
		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		// c.Set("current_user_id", userModel.GetStringID())
		// c.Set("current_user_name", userModel.Name)
		// c.Set("current_user", userModel)

		c.Next()
	}
}
