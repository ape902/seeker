package router

import (
	"github.com/ape902/seeker/pkg/api/seeker"
	"github.com/ape902/seeker/pkg/tools/ginx/middleware"
	"github.com/gin-gonic/gin"
)

func InitSeekerRouter(e *gin.RouterGroup) {
	e.POST("/login", seeker.Login)
	system := e.Group("/user", middleware.JWTAuth())
	{
		system.POST("/add", seeker.CreateUser)
	}
	cmdb := e.Group("cmdb", middleware.JWTAuth())
	{
		cmdb.POST("/command", seeker.RunCommand)
	}

}
