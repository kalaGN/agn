// Package auth 处理用户身份认证相关逻辑
package interfacett

import (
	"agn/app/http/controllers/api/external"
	dial_record "agn/app/models/bd_dial_record"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BaseAPIController 基础控制器
type InterfacettController struct {
	external.BaseExternalController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *InterfacettController) IsDataExist(c *gin.Context) {

	// 初始化请求对象
	// 请求对象
	type PhoneExistRequest struct {
		Enterprise_id string `json:"enterprise_id"`
	}
	request := PhoneExistRequest{}
	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "",
		"data": dial_record.IsDataExist(request.Enterprise_id),
	})
}