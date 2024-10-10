package codex

const (
	Success                 = 200000
	Failure                 = 200001
	InvalidParameter        = 210001 //无效参数
	ExecutionFailed         = 210002 //执行失败
	AlreadyExists           = 210003 //已存在
	UserOrPassError         = 210004 //用户名或密码错误
	TokenCreateFailed       = 210005 //Token生成失败
	EncryptCreateFailed     = 210006 //密钥创建失败
	DatabaseExecutionFailed = 210007 //数据库执行失败
	GRPCConnectionFailed    = 210008 //GRPC连接失败
	SSHConnectionFailed     = 210009 //SSH连接失败
	NotExists               = 210010 //不存在
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
	case EncryptCreateFailed:
		return "密钥创建失败"
	case DatabaseExecutionFailed:
		return "数据库执行失败"
	case GRPCConnectionFailed:
		return "GRPC连接失败"
	case NotExists:
		return "不存在"
	default:
		return ""
	}
}
