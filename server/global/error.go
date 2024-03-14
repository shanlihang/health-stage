package global

// 统一定义错误码
const (
	SUCCESS        = 200
	PAGE_NOT_FOUND = 404
)

// 错误提示信息
var errMsg = map[int]string{
	SUCCESS:        "请求成功",
	PAGE_NOT_FOUND: "页面不存在",
}

//根据错误码获取错误提示信息
func GetMsg(code int) string {
	msg, isOK := errMsg[code]
	if isOK {
		return msg
	}
	return "错误码未定义"
}
