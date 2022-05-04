package router

import (
	_ "bilibili_demo/docs"
	"bilibili_demo/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/problem", service.GetProblemList)
	r.GET("problem-detail", service.GetProblemListDetail)

	// 用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/email", service.SendEmail)
	r.POST("/register", service.Register)
	r.GET("/rank-list", service.GetRankList)

	// 提交列表
	r.GET("/submit-list", service.GetSubmitList)

	return r
}
