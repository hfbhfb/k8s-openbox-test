package innhttpcode

const (
	STATUS_OK              = 10000
	ERR_INVAID_PARAM       = 10001
	ERR_USERNAME_OR_PASSWD = 10002
	ERR_TOKEN              = 10003
	ERR_CREATE_MENU        = 10004
	ERR_NOT_ENOUGH_COIN    = 10005
	ERR_COMMON             = 10006
)

var errorCode = map[int]string{
	10000: "请求成功",
	10001: "无效的参数",
	10002: "用户名或者密码错误",
	10003: "缺少Token",
	10004: "创建Menu项失败",
	10005: "余额不足",
	10006: "操作出错",
}

func ErrorText(code int) string {
	if s, ok := errorCode[code]; ok {
		return s
	}
	return "未知的内部错误"
}
