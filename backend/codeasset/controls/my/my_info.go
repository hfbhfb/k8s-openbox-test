package my

import (
	"backend/codeasset/automodel"
	"backend/codeasset/common/ginutils"
	"backend/codeasset/common/innhttpcode"
	"backend/codeasset/common/middleware/jwtwrap"
	"backend/codeasset/models"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	var err error
	var resp struct {
		Account int64  `json:"account" ` //渲染币
		Name    string `json:"name" `    //渲染币
	}

	claim := jwtwrap.GetClient(c)
	mgr := automodel.GouserMgr(models.BaseData)
	var user automodel.Gouser
	user, err = mgr.GetFromIDgouser(claim.UserId)
	if err != nil {
		ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), nil)
	} else {
		resp.Account = user.Account
		resp.Name = user.LoginName
		ginutils.WriteString(c, innhttpcode.STATUS_OK, innhttpcode.ErrorText(innhttpcode.STATUS_OK), resp)
	}

	return
}
