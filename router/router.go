package router

import (
	"github.com/gin-gonic/gin"

	"github.com/TechLoCo/Zousan-api/controller"
)

// Run actions initialize router
func Run() {
	r := route()
	r.Run()
}

func route() *gin.Engine {
	r := gin.Default()

	u := r.Group("users")
	{
		ctrl := controller.User{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
	}

	return r
}
