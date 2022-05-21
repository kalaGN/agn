// Package auth 处理用户身份认证相关逻辑
package interfacett

import (
	"agn/app/http/controllers/api/external"
	"agn/app/models/dial_record"
	"agn/app/requests"
	"agn/pkg/response"

	"github.com/gin-gonic/gin"
)

// BaseAPIController 基础控制器
type InterfacettController struct {
	external.BaseExternalController
}

// 记录列表
func (it *InterfacettController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := dial_record.Paginate(c, 10)
	response.JSON(c, gin.H{
		"code":  0,
		"msg":   "",
		"data":  data,
		"pager": pager,
	})
}
