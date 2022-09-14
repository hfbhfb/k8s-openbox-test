package share

const (
	SUCCESS                             = 200
	ERR_PARAM                           = 400
	ERR_USER_UNRECOGNIZED               = 401
	ERR_SERVER_INNER                    = 402
	ERR_ACCESS_DENIED                   = 403
	ERR_DATA_NOT_FOUND                  = 404
	ERR_DATA_ALREADY_EXIST              = 410
	ERR_TOKEN_LACK                      = 501
	ERR_TOKEN_EXPIRE                    = 502
	ERR_TOKEN_CHECK                     = 503
	ERR_TOKEN                           = 504
	ERR_USER_NOT_EXIST                  = 10000
	ERR_PHONE_REGISTERED                = 10001
	ERR_COMPANY_REGISTERED              = 10002
	ERR_COMPANY_NAME_TOO_LONG           = 10003
	ERR_PHONE                           = 10004
	ERR_VERIFY_CODE                     = 10005
	ERR_BLOCK_CHAIN                     = 10006
	ERR_PHONE_NOT_REGISTER              = 10007
	ERR_PASSWORD                        = 10008
	ERR_RESOURCE_NOT_EXIST              = 10009
	ERR_PUSH_DATA_TO_CHAIN              = 10010
	ERR_NODE_NOT_EXIST                  = 10011
	ERR_MODIFY_RESOURCE_ILLEGAL         = 10012
	ERR_DEL_RESOURCE_ILLEGAL            = 10013
	ERR_RESOURCE_TITLE_TOO_LONG         = 10014
	ERR_RESOURCE_DESC_TOO_LONG          = 10015
	ERR_MODIFY_NAME_EMPTY               = 10016
	ERR_NODE_NAME_TOO_LONG              = 10017
	ERR_NAME_MUSTL_HAN                  = 10018
	ERR_LOGIN_AND_MODIFY_PHONE_MISMATCH = 10019
	ERR_CAN_NOT_COOP_MYSELF_RESOURCE    = 10020
	ERR_RESOURCE_ALREADY_COOP           = 10021
	ERR_REGION                          = 10022
	ERR_USER_ABNORMAL                   = 10023
)

var errorCode = map[int]string{
	501:   "缺少Token",
	502:   "token已过期",
	503:   "校验token错误",
	504:   "token错误",
	200:   "请求成功",
	400:   "参数错误",   // 客户端请求的语法错误，服务器无法理解
	401:   "无法识别用户", // 请求要求用户的身份认证
	402:   "未知的内部错误",
	403:   "拒绝访问",  // 服务器理解请求客户端的请求，但是拒绝执行此请求
	404:   "数据不存在", // 服务器找不到请求的网页
	410:   "数据已存在", // 服务器数据已存在
	10000: "用户不存在",
	10001: "手机号已注册",
	10002: "公司名称已注册",
	10003: "公司名称超出限制",
	10004: "手机号错误",
	10005: "验证码错误",
	10006: "区块服务错误",
	10007: "手机号未注册",
	10008: "用户密码错误",
	10009: "资源不存在",
	10010: "数据上链出错",
	10011: "用户节点信息异常",
	10012: "非法修改资源",
	10013: "非法删除资源",
	10014: "资源标题超出限制",
	10015: "资源描述超出限制",
	10016: "修改名称不能为空",
	10017: "节点名称超出限制",
	10018: "公司名称必须是汉字",
	10019: "登陆的手机号和修改的手机号不一致",
	10020: "自己的资源，无须合作",
	10021: "资源已申请合作，请勿重复申请",
	10022: "地区信息错误",
	10023: "账户异常",
}

func ErrorText(code int) string {
	if s, ok := errorCode[code]; ok {
		return s
	}
	return "未知的内部错误"
}
