package constant

type MsgCode struct {
	Code int
	Msg  string
}

var (
	// 成功
	SUCCESS = buildCode(200, "success")
	// 失败
	COMMON_FAIL = buildCode(-4396, "fail")
)

func buildCode(code int, msg string) MsgCode {
	return MsgCode{Code: code, Msg: msg}
}
