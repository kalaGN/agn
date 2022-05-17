// Package middlewares Gin 中间件
package middlewares

import (
	"agn/pkg/response"
	"crypto/sha1"
	"encoding/hex"
	"sort"

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
			response.Unauthorized(c, "auth param error")
			return
		}

		var tempArray = []string{partner_key, timestr}
		sort.Strings(tempArray)
		//将三个参数字符串拼接成一个字符串进行sha1加密
		var sha1String string = ""
		for _, v := range tempArray {
			sha1String += v
		}

		h := sha1.New()
		h.Write([]byte(sha1String))

		sha1String = hex.EncodeToString(h.Sum([]byte("")))
		//获得加密后的字符串可与signature对比
		if sha1String != interface_sign {
			response.Unauthorized(c, "sign error")
			return
		}

		c.Next()
	}
}
