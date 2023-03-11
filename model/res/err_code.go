package res

type ErrorCode int

const (
	ErrorSettings ErrorCode = 1001 // 系统错误
)

var (
	// ErrorMap 错误码对应的错误信息msg
	ErrorMap = map[ErrorCode]string{
		ErrorSettings: "系统错误",
	}
)
