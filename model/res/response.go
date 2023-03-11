package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success = 0
	Error   = 7
)

// Response 返回响应的封装数据结构
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// Result 返回响应的封装方法
func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Ok 返回成功的响应
func Ok(msg string, data any, c *gin.Context) {
	Result(Success, msg, data, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, "成功", data, c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, msg, map[string]any{}, c)
}

// Fail 返回成功的响应
func Fail(msg string, data any, c *gin.Context) {
	Result(Error, msg, data, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, msg, map[string]any{}, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code] // 从错误码对应的错误信息map中获取错误信息
	if !ok {
		Result(int(code), msg, map[string]any{}, c)
	}
	Result(Error, "未知错误", map[string]any{}, c)
}
