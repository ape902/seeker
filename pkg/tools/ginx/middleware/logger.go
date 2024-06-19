package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type loggerStruct struct {
	Runtime       string        `json:"runtime"`
	Status        int           `json:"status"`
	LatencyTime   time.Duration `json:"latency_time"`
	ClientIP      string        `json:"client_ip"`
	RequestMethod string        `json:"request_method"`
	RequestURI    string        `json:"request_uri"`
}

func (l *loggerStruct) String() string {
	return fmt.Sprintf("runtime: %s status: %d latency_time: %v client_ip: %s method: %s uri: %s",
		l.Runtime, l.Status, l.LatencyTime, l.ClientIP, l.RequestMethod, l.RequestURI)
}

// logger 初始化Gin日志输出
func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ls loggerStruct
		// 开始时间
		startTime := time.Now()
		ls.Runtime = startTime.Format("2006-01-02 15:04:05.9999")

		// 结束时间
		endTime := time.Now()

		// 执行时间
		ls.LatencyTime = endTime.Sub(startTime)

		// 请求方式
		ls.RequestMethod = c.Request.Method

		// 请求路由
		ls.RequestURI = c.Request.RequestURI

		// 状态码
		ls.Status = c.Writer.Status()

		// 请求IP
		ls.ClientIP = c.ClientIP()

		fmt.Println(ls.String())

		// 处理请求
		c.Next()
	}
}
