package logger

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"backend/codeasset/utils/tlog"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

type bodyLogReader struct {
	io.ReadCloser
	bodyBuf *bytes.Buffer
}

func (r bodyLogReader) Read(b []byte) (int, error) {
	//memory copy here!
	n, err := r.ReadCloser.Read(b)
	if n != 0 && err == nil {
		fmt.Println(n)
		r.bodyBuf.Write(b[0:n])
	}
	fmt.Println("err:", err)
	// r.bodyBuf.Read(b)
	return n, err
}
func (r bodyLogReader) Close() error {
	//memory copy here!
	return r.ReadCloser.Close()
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		requestID := t.UnixNano() % 1000000000
		c.Set("X-Request-ID", requestID)

		blw := bodyLogWriter{bodyBuf: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		blr := bodyLogReader{bodyBuf: bytes.NewBufferString(""), ReadCloser: c.Request.Body}
		c.Request.Body = blr
		c.Next()

		b, err := io.ReadAll(blr.ReadCloser)
		if err == nil && len(b) > 0 {
			blr.bodyBuf.Write(b[0:])
		}
		latency := time.Now().Sub(t)
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		cliInfo := fmt.Sprintf(" request-id: %d | http-code: %3d | runtime: %v | clientip: %s | http-method: %s  path: %s  user-agent: %s ",
			requestID,
			statusCode,
			latency,
			clientIP,
			method, path,
			c.Request.Header.Get("User-Agent"),
		)
		cliInfo2 := fmt.Sprintf(" request-id: %d  body: %s resp: %s url: %s",
			requestID,
			blr.bodyBuf,
			blw.bodyBuf,
			c.Request.URL,
		)
		//fmt.Printf("cliInfo :%v\n", cliInfo)
		tlog.Info(cliInfo)
		// tlog.InfoCategory(cliInfo)
		tlog.InfoCategory(cliInfo2)
	}
}
