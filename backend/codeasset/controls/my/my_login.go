package my

import (
	"backend/codeasset/automodel"
	"backend/codeasset/common/ginutils"
	"backend/codeasset/common/innhttpcode"
	"backend/codeasset/common/middleware/jwtwrap"
	"backend/codeasset/models"
	"backend/codeasset/utils/lib"
	"backend/codeasset/utils/tlog"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	// return
	var err error
	var resp struct {
		AccessToken string      `json:"access_token" ` //用户昵称
		MenuInfo    interface{} `json:"menu_info" `
	}
	var req struct {
		UserName string `json:"username" ` //用户昵称
		UserPwd  string `json:"password"`  //用户密码
	}
	if err = c.BindJSON(&req); err != nil {
		tlog.Error(c, err)
		ginutils.WriteString(c, innhttpcode.ERR_INVAID_PARAM, innhttpcode.ErrorText(innhttpcode.ERR_INVAID_PARAM), "ksljfkdsjfl")
		return
	}

	mgr := automodel.GouserMgr(models.BaseData)
	var user automodel.Gouser
	user, err = mgr.GetFromLoginName(req.UserName)
	if err != nil {
		ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), nil)
		return
	} else {

		tlog.Debug(req.UserName)
		tlog.Debug(req.UserPwd)
		tlog.Debug(lib.Md5(req.UserPwd, false))

		if lib.Md5(req.UserPwd, false) != user.LoginPasswd {
			ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), nil)
			return
		} else {

			clientJwt := jwtwrap.Claims{}
			clientJwt.UserId = user.IDgouser
			clientJwt.Username = user.LoginName
			resp.AccessToken, err = jwtwrap.NewJWT().CreateToken(clientJwt)
			if err == nil {
				ginutils.WriteString(c, innhttpcode.STATUS_OK, innhttpcode.ErrorText(innhttpcode.STATUS_OK), resp)
				return
			} else {
				ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), "token jwt生成失败")
				return
			}
		}
	}

	/*
		userModel := models.NewUser(models.BaseData)
		userModel.LoginName = req.UserName
		err = userModel.GetUser()
		if err != nil {
			ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), nil)
			return
		} else {
			if userModel.LoginPasswd == "" {
				ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), nil)
				return
			} else {

				fmt.Println(lib.Md5(req.UserPwd, false))
				if lib.Md5(req.UserPwd, false) != userModel.LoginPasswd {
					fmt.Println(lib.Md5(req.UserPwd, false))
					ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), nil)
					// ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), lib.Md5(req.UserPwd, false))
					// ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), userModel.LoginPassword)
					return
				} else {

					clientJwt := jwtwrap.Claims{}
					clientJwt.UserId = userModel.IDgouser
					clientJwt.Username = userModel.LoginName
					resp.AccessToken, err = jwtwrap.NewJWT().CreateToken(clientJwt)
					if err == nil {
						ginutils.WriteString(c, innhttpcode.STATUS_OK, innhttpcode.ErrorText(innhttpcode.STATUS_OK), resp)
						return
					} else {
						ginutils.WriteString(c, innhttpcode.ERR_USERNAME_OR_PASSWD, innhttpcode.ErrorText(innhttpcode.ERR_USERNAME_OR_PASSWD), "token jwt生成失败")
					}
					return
				}
			}
		}
	*/
	ginutils.WriteString(c, innhttpcode.STATUS_OK, "", nil)
	return
}
