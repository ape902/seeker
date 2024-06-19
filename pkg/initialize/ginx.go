package initialize

import (
	"github.com/ape902/seeker/pkg/router"
	"github.com/ape902/seeker/pkg/tools/ginx/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Engine(mode string) *gin.Engine {
	e := gin.New()

	// 自定义中间件
	middleware.Middleware(e)

	// 除了release模式其它模式都加载pprof功能
	if mode != "release" {
		// pprof注册到gin中
		pprof.Register(e)
	}

	v1 := e.Group("/api/v1")
	router.InitSeekerRouter(v1)

	return e
}
