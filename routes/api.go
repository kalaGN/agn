// Package routes 注册路由
package routes

import (
	interfacett "agn/app/http/controllers/api/external/interface"
	"agn/app/http/controllers/api/v1/auth"
	"agn/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	externalRouter := r.Group("/external")
	{
		suc := new(interfacett.InterfacettController)
		// 判断手机是否已注册
		externalRouter.GET("/interface/getdialrecord", middlewares.AuthCommon(), suc.Index)

	}
}
