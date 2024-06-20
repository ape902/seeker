package ginx

import (
	"net/http"
	"time"

	"github.com/ape902/seeker/pkg/tools/codex"
	"github.com/gin-gonic/gin"
)

type (
	response struct {
		Code      int         `json:"code"`
		Data      interface{} `json:"data"`
		Message   string      `json:"message"`
		Timestamp int64       `json:"timestamp"`
	}
)

func RESP(c *gin.Context, code int, data interface{}) {
	var resp response
	resp.Code = code
	resp.Data = data
	resp.Message = codex.CodeText(code)
	resp.Timestamp = time.Now().Unix()

	c.JSON(http.StatusOK, resp)
}

func Page(total int64, records interface{}) map[string]interface{} {
	data := make(map[string]interface{})

	data = map[string]interface{}{
		"total":   total,
		"records": records,
	}

	return data
}
