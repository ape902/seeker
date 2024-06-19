package codex

const (
	Success           = 200000
	Failure           = 200001
	InvalidParameter  = 210001 //无效参数
	ExecutionFailed   = 210002 //执行失败
	AlreadyExists     = 210003 //数据存在
	UserOrPassError   = 210004 //用户名或密码错误
	TokenCreateFailed = 210005 //Token生成失败
)

func CodeText(code int) string {
	switch code {
	case Success:
		return "成功"
	case Failure:
		return "失败"
	case InvalidParameter:
		return "无效参数"
	case ExecutionFailed:
		return "执行失败"
	case AlreadyExists:
		return "数据存在"
	case UserOrPassError:
		return "用户名或密码错误"
	case TokenCreateFailed:
		return "Token生成失败"
	default:
		return ""
	}
}
