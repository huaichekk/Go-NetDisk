package initialize

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/middleware"
)

func InitEngine(r *gin.Engine) {
	middleware.Cors(r)

	initRouter(r)
}
