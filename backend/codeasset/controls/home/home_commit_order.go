package home

import (
	"errors"
	"time"

	"backend/codeasset/automodel"
	"backend/codeasset/common/ginutils"
	"backend/codeasset/common/innhttpcode"
	"backend/codeasset/common/middleware/jwtwrap"
	"backend/codeasset/models"
	"backend/codeasset/utils/tlog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	CONST_COST = 0
)

func CommitOrder(c *gin.Context) {
	var err error
	var resp struct {
		CallBackUrl string `json:"call_back_url" ` //用户昵称
	}
	var req struct {
		Url string `json:"url"` //
	}
	if err = c.BindJSON(&req); err != nil {
		tlog.Error(c, err)
		ginutils.WriteString(c, innhttpcode.ERR_INVAID_PARAM, innhttpcode.ErrorText(innhttpcode.ERR_INVAID_PARAM), nil)
		return
	}
	if req.Url == "" {
		ginutils.WriteString(c, innhttpcode.ERR_INVAID_PARAM, innhttpcode.ErrorText(innhttpcode.ERR_INVAID_PARAM), nil)
		return
	}

	claim := jwtwrap.GetClient(c)
	mgrorder := automodel.GoordersMgr(models.BaseData)
	var order automodel.Goorders
	mgruser := automodel.GouserMgr(models.BaseData)
	opt1 := mgruser.WithIDgouser(claim.UserId)
	var user automodel.Gouser
	user, err = mgruser.GetByOption(opt1)
	if err != nil {
		ginutils.WriteString(c, innhttpcode.ERR_COMMON, err.Error(), nil)
		return
	}
	tlog.Info(user)
	tlog.Info(claim.UserId, req.Url)

	order, err = mgrorder.FetchUniqueIndexByUserURL(claim.UserId, req.Url)
	if err != nil {
		ginutils.WriteString(c, innhttpcode.ERR_COMMON, err.Error(), nil)
		return
	} else {
		tlog.Info(order)
		if order.IDorder == 0 {

			err = mgrorder.Transaction(func(tx *gorm.DB) error {
				// return nil will commit the whole transaction
				if (user.Account - int64(CONST_COST)) < 0 {
					return errors.New("余额不足")
				}
				// mgruser.WithAccount()
				err = mgruser.DB.Model(automodel.Gouser{}).Where(map[string]interface{}{
					"idgouser": claim.UserId,
				}).Update("account", user.Account-int64(CONST_COST)).Error

				err2 := tx.Create(&automodel.Goorders{
					IDgouser:   claim.UserId,
					URL:        req.Url,
					Cost:       CONST_COST,
					Createtime: time.Now(),
					Finishtime: time.Now(),
					Status:     1,
				}).Error
				if err2 != nil {
					return err2
					// errors.New("创建出错")
				}
				return nil
			})

			if err != nil {
				ginutils.WriteString(c, innhttpcode.ERR_COMMON, err.Error(), nil)
			} else {
				ginutils.WriteString(c, innhttpcode.STATUS_OK, "", nil)
			}
			return
		} else {

			ginutils.WriteString(c, innhttpcode.ERR_COMMON, "重复提交", nil)
			return
		}
	}

	ginutils.WriteString(c, innhttpcode.STATUS_OK, innhttpcode.ErrorText(innhttpcode.STATUS_OK), resp)
	return
}
