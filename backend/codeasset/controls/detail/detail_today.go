package detail

import (
	"time"

	"backend/codeasset/controls/my"

	"github.com/gin-gonic/gin"
)

func TodayItems(c *gin.Context) {
	var req my.ReqMyConsumer
	var resp my.RespMyConsumer

	year, month, day := time.Now().Date()
	req.StartTime = time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	req.EndTime = time.Date(year, month, day, 23, 59, 59, 0, time.Now().Location())
	req.Page = 1
	req.PageSize = 1000000
	// tlog.Info(req)
	// tlog.Info(req)
	// tlog.Info(req)
	// tlog.Info(req)
	my.CommonUserConsumerHistory(c, &req, &resp, false)
	return
}
