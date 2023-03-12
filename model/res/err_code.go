package res

type ErrorCode int

const (
	ErrorSettings ErrorCode = 1001 // 系统错误
	ErrorArgument ErrorCode = 1002 // 参数错误
)

var (
	// ErrorMap 错误码对应的错误信息msg
	ErrorMap = map[ErrorCode]string{
		ErrorSettings: "系统错误",
		ErrorArgument: "参数错误",
	}
)
