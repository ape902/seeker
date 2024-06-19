package initialize

import (
	"github.com/ape902/corex/logx"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func InitLogger() {
	mode := os.Getenv("mode")
	if mode == "" {
		mode = "file"
	}

	opt := make([]logx.OptionFunc, 0)
	switch gin.Mode() {
	case "release":

		commpress := os.Getenv("commpress")
		commPressBool := false
		if commpress == "true" {
			commPressBool = true
		}

		opt = append(opt,
			logx.WithMode(os.Getenv("mode")),
			logx.WithLogPath(os.Getenv("log_path")),
			logx.WithServerName("seeker"),
			logx.WithLevel(os.Getenv("level")),
			logx.WithMaxSize(envStringToInt("max_size")),
			logx.WithMaxDay(envStringToInt("max_day")),
			logx.WithBackups(envStringToInt("backups")),
			logx.WithCompress(commPressBool),
		)

	}

	logx.NewLoggerOption(opt...)
}

func envStringToInt(env string) int {
	envValue := os.Getenv(env)

	var value int
	switch env {
	case "max_size":
		value = strconvToInt(envValue, 20)
	case "max_day":
		value = strconvToInt(envValue, 7)
	case "backups":
		value = strconvToInt(envValue, 3)
	}

	return value
}

func strconvToInt(envValue string, value int) int {
	if envValue != "" {
		value, _ = strconv.Atoi(envValue)
	}

	return value
}
