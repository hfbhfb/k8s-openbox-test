package detail

import (
	"time"

	"backend/codeasset/controls/my"

	"github.com/gin-gonic/gin"
)

func HistoryItems(c *gin.Context) {
	var req my.ReqMyConsumer
	var resp my.RespMyConsumer

	year, month, day := time.Now().Date()
	currentTime := time.Now()
	req.StartTime = currentTime.AddDate(0, -12, 0)
	req.EndTime = time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	req.Page = 1
	req.PageSize = 20
	my.CommonUserConsumerHistory(c, &req, &resp, false)
	return
}
