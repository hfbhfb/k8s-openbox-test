package ginutils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"

	"backend/codeasset/common/innhttpcode"
	. "backend/codeasset/common/share"
)

func Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func WriteString(c *gin.Context, code int, msg string, data interface{}) {
	res := map[string]interface{}{"code": code}
	res["msg"] = msg
	if msg == "" {
		res["msg"] = innhttpcode.ErrorText(code)
	}
	if data != nil && !IsNil(data) {
		res["data"] = data
	} else {
		res["data"] = []interface{}{}
	}
	if id, ok := c.Get("X-Request-ID"); ok {
		res["request_id"] = fmt.Sprintf("%v", id)
	}

	c.JSON(http.StatusOK, res)
}

func WriteStringEx(c *gin.Context, code int, data interface{}) {
	res := map[string]interface{}{"code": code}
	res["msg"] = ErrorText(code)
	if data != nil && !IsNil(data) {
		res["data"] = data
	} else {
		res["data"] = []interface{}{}
	}
	if id, ok := c.Get("X-Request-ID"); ok {
		res["request_id"] = fmt.Sprintf("%v", id)
	}

	c.JSON(http.StatusOK, res)
}

func WriteStringNull(c *gin.Context, code int, msg string, data interface{}) {
	res := map[string]interface{}{"code": code}
	res["msg"] = msg
	if data != nil && !IsNil(data) {
		res["data"] = data
	} else {
		res["data"] = nil
	}
	if id, ok := c.Get("request_id"); ok {
		res["request_id"] = fmt.Sprintf("%v", id)
	}

	c.JSON(http.StatusOK, res)
}

func IsNil(s interface{}) bool {
	defer func() {
		recover()
	}()
	return reflect.ValueOf(s).IsNil()
}

func ClientIP(ctx *gin.Context) string {
	if addr := ctx.Request.Header.Get("X-Forwarded-For"); addr != "" {
		return strings.TrimSpace(strings.Split(addr, ",")[0])
	}
	if addr := ctx.Request.Header.Get("X-Real-Ip"); addr != "" {
		return addr
	}
	if addr := ctx.Request.Header.Get("Remote"); addr != "" {
		return strings.TrimSpace(strings.Split(addr, ":")[0])
	}
	if addr := ctx.Request.Header.Get("X-Appengine-Remote-Addr"); addr != "" {
		return strings.TrimSpace(strings.Split(addr, ":")[0])
	}
	return ctx.ClientIP()
}
