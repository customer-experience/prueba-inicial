package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine) {
	//c := controller.NewController()
	//
	//v1 := r.Group("/api/v1")
	//{
	//	envios := v1.Group("/envios")
	//	{
	//		envios.GET(":id", c.ShowShipment)
	//	}
	//}
	r.GET("/api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // ../api/doc/index.html
}
