package tools

/**
自定义返回接口
*/

type ResultUtil struct {
	//返回0 成功
	//返回-1 缺少参数
	//返回-2 参数不合法
	//返回-3 请求失败
	//返回-4 非法用户
	//返回-5 鉴权失败
	Code int
	Msg  string
	// 定义空接口以支持泛型
	Data interface{}
}

func NewResultUtil(code int, msg string, data interface{}) *ResultUtil {
	return &ResultUtil{Code: code, Msg: msg, Data: data}
}

func NewResultSuccess(data interface{}) *ResultUtil {
	return &ResultUtil{Code: 0, Msg: "ok", Data: data}
}

func NewResult() *ResultUtil {
	return &ResultUtil{Code: 0, Msg: "ok"}
}

func NewResultError(code int, msg string) *ResultUtil {
	return &ResultUtil{Code: code, Msg: msg}
}
