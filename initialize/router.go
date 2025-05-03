package initialize

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/controller"
	"go-netdisk/middleware"
)

func initRouter(r *gin.Engine) {
	// 无需鉴权的公共接口
	api := r.Group("/api")
	{
		api.POST("/login", controller.Login)
		api.GET("/ping", controller.Ping)
		api.POST("/register", controller.Register)
	}

	// 🔥 需要鉴权的接口组
	authApi := api.Group("")
	authApi.Use(middleware.AuthMiddle)
	{
		authApi.GET("/auth/ping", controller.AuthPing)
		files := authApi.Group("/files")
		{
			files.GET("", controller.Files)
			files.DELETE("/delete", controller.Delete)
			files.GET("/download", controller.Download)
			//http://localhost:8080/api/files/download?file_id=3
		}
		//authApi.GET("/files", controller.Files)
		authApi.POST("/uploadfile", controller.UploadFile)
	}
}
