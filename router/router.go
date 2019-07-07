package router

import (
	"github.com/gin-gonic/gin"

	"github.com/TechLoCo/Zousan-api/controllers"
)

// Init actions initialize router
func Init() {
	r := route()
	r.Run()
}

func route() *gin.Engine {
	r := gin.Default()

	u := r.Group("users")
	{
		// HACK: controller controllerしててなんかヤダ
		ctrl := controller.UserController{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
	}

	return r
}
