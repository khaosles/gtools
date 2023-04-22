package gresponse

/*
   @File: ErrorCode.go
   @Author: khaosles
   @Time: 2023/3/7 21:54
   @Desc:
*/

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s Status) SetMsg(msg string) Status {
	s.Msg = msg
	return s
}

var (
	SUCCESS              = Status{20000, "ok"}      // 运行成功
	PARAMS_ERROR         = Status{40000, "请求参数错误"}  // 参数错误
	NOT_LOGIN_ERROR      = Status{40100, "未登录"}     // 账号未登录
	NO_AUTH_ERROR        = Status{40101, "账号无权限"}   // 账号无权限
	TOKEN_EXPIRE         = Status{40102, "登录信息过期"}  // 登录过期
	TOKEN_REMOTING_LOGIN = Status{40103, "账号异地登录"}  // 异地登录
	NOT_FOUND_ERROR      = Status{40400, "请求数据不存在"} // 数据不存在
	FORBIDDEN_ERROR      = Status{40300, "禁止访问"}    // 数据禁止访问
	SYSTEM_ERROR         = Status{50000, "系统内部异常"}  // 系统内部出错
	OPERATION_ERROR      = Status{50001, "操作失败"}    // 操作错误
	CUSTOM_ERROR         = Status{50100, ""}        // 自定义
)
