package router

import (
	"github.com/ape902/seeker/pkg/api/seeker"
	"github.com/ape902/seeker/pkg/tools/ginx/middleware"
	"github.com/gin-gonic/gin"
)

func InitSeekerRouter(e *gin.RouterGroup) {
	e.POST("/login", seeker.Login)
	e.GET("/prom/discover", seeker.HttpSDConfig)
	system := e.Group("/user", middleware.JWTAuth())
	{
		system.POST("/add", seeker.UserCenterCreate)
		system.GET("/list", seeker.UserCenterFindPage)
		system.GET("/info", seeker.UserCenterFindById)
		system.POST("/modify", seeker.UserCenterUpdate)
		system.POST("/del", seeker.UserCenterDeleteById)
	}
	cmdb := e.Group("cmdb", middleware.JWTAuth())
	{
		cmdb.POST("/put/sftp", seeker.SftpPut)
		cmdb.POST("/command", seeker.RunCommand)
		cmdb.GET("/discovery", seeker.Discovery)
		cmdb.POST("/create", seeker.HostInfoCreate)
		cmdb.GET("/list", seeker.HostInfoFindPage)
		cmdb.POST("/delete", seeker.HostInfoDelete)
		cmdb.POST("/update", seeker.HostInfoUpdateHost)
		cmdb.POST("/modify/auth", seeker.HostInfoUpdateAuthentication)
		cmdb.GET("/info", seeker.HostInfoFindById)
	}

	storage := e.Group("/storage")
	{
		storage.GET("/bucket/list", seeker.MinioBucketList)
	}

}
