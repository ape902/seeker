// 状态码常量定义
export const CodeStatus = {
  Success: 200000,
  Failure: 200001,
  InvalidParameter: 210001,
  ExecutionFailed: 210002,
  AlreadyExists: 210003,
  UserOrPassError: 210004,
  TokenCreateFailed: 210005,
  EncryptCreateFailed: 210006,
  DatabaseExecutionFailed: 210007,
  GRPCConnectionFailed: 210008,
  SSHConnectionFailed: 210009,
  NotExists: 210010,
} as const;

// 状态码文本映射
export const getCodeText = (code: number): string => {
  switch (code) {
    case CodeStatus.Success:
      return '成功';
    case CodeStatus.Failure:
      return '失败';
    case CodeStatus.InvalidParameter:
      return '无效参数';
    case CodeStatus.ExecutionFailed:
      return '执行失败';
    case CodeStatus.AlreadyExists:
      return '数据存在';
    case CodeStatus.UserOrPassError:
      return '用户名或密码错误';
    case CodeStatus.TokenCreateFailed:
      return 'Token生成失败';
    case CodeStatus.EncryptCreateFailed:
      return '密钥创建失败';
    case CodeStatus.DatabaseExecutionFailed:
      return '数据库执行失败';
    case CodeStatus.GRPCConnectionFailed:
      return 'GRPC连接失败';
    case CodeStatus.NotExists:
      return '不存在';
    default:
      return '';
  }
};