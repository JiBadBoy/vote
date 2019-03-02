package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",
	MobileError:   "手机号码格式不对",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
