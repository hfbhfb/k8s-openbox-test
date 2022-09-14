package my

import (
	"time"

	"backend/codeasset/automodel"
	"backend/codeasset/common/ginutils"
	"backend/codeasset/common/innhttpcode"
	"backend/codeasset/common/middleware/jwtwrap"
	"backend/codeasset/models"
	"backend/codeasset/utils/tlog"

	"github.com/gin-gonic/gin"
)

type ReqMyConsumer struct {
	Page      int64     `json:"page" form:"page"`           //开始下标
	PageSize  int64     `json:"page_size" form:"page_size"` //返回条数
	StartTime time.Time `json:"start_time" form:"start_time"`
	EndTime   time.Time `json:"end_time" form:"end_time"`
}

type RespMyConsumer struct {
	ConsumerList []ConsumerItem `json:"consumerlist" ` //渲染币
}
type ConsumerItem struct {
	CreateTime  time.Time `json:"create_time"`
	FinishTime  time.Time `json:"finish_time"`
	Url         string    `json:"url"`
	Callbackurl string    `json:"callback_url"`
	Cost        int64     `json:"cost"`
	Status      int64     `json:"status"`
}

func UserConsumerHistory(c *gin.Context) {
	var err error

	var req ReqMyConsumer
	if err = c.Bind(&req); err != nil {
		tlog.Error(c, err)
		ginutils.WriteString(c, innhttpcode.ERR_INVAID_PARAM, innhttpcode.ErrorText(innhttpcode.ERR_INVAID_PARAM), "ksljfkdsjfl")
		return
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	var resp RespMyConsumer
	CommonUserConsumerHistory(c, &req, &resp, true)

	return
}

func CommonUserConsumerHistory(c *gin.Context, req *ReqMyConsumer, resp *RespMyConsumer, filterFinish bool) {

	var err error
	claim := jwtwrap.GetClient(c)
	mgrorder := automodel.GoordersMgr(models.BaseData)

	// tlog.Info(req.Page)
	// tlog.Info(req.PageSize)
	// orders, err = mgrorder.GetFromIDgouser(claim.UserId)
	dbselect := mgrorder.Model(automodel.Goorders{})
	dbselect = dbselect.Limit(int(req.PageSize)).Offset(int((req.Page - 1) * req.PageSize))
	if !req.StartTime.IsZero() {
		dbselect = dbselect.Where(" createtime > ? and createtime < ?  ", req.StartTime, req.EndTime)
	}
	results := make([]*automodel.Goorders, 0)

	dbselect = dbselect.Model(automodel.Goorders{}).Where("`idgouser` = ?", claim.UserId)
	if filterFinish {
		dbselect = dbselect.Where("`status` = 3")
	}
	err = dbselect.Order("createtime desc").Find(&results).Error

	if err != nil {
		ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), nil)
	} else {
		for _, item := range results {

			resp.ConsumerList = append(resp.ConsumerList, ConsumerItem{
				CreateTime:  item.Createtime,
				FinishTime:  item.Finishtime,
				Url:         item.URL,
				Callbackurl: item.Callbackurl,
				Cost:        int64(item.Cost),
				Status:      int64(item.Status),
			})
		}
		ginutils.WriteString(c, innhttpcode.STATUS_OK, innhttpcode.ErrorText(innhttpcode.STATUS_OK), resp)
		return
	}

	ginutils.WriteString(c, innhttpcode.STATUS_OK, innhttpcode.ErrorText(innhttpcode.STATUS_OK), nil)
	return
}
