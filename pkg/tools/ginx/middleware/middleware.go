package middleware

import "github.com/gin-gonic/gin"

// 自定义gin中间件
func Middleware(e *gin.Engine) {
	// 使用自定义日志
	// noCache 防止客户端缓存http响应
	// option 跨域请求头
	// secure 附加安全和资源访问请求头
	// logger 自定义日志输出
	e.Use(noCache(), options(), secure(), logger())
}
