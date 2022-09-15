package softinfo

import (
	"backend/codeasset/common/ginutils"
	"backend/codeasset/common/innhttpcode"

	. "backend/codeasset/common/share"

	"github.com/gin-gonic/gin"
)

func VersionInfo(c *gin.Context) {

	ginutils.WriteString(c, innhttpcode.STATUS_OK, NAME+":"+VERSION, nil)
	return
}
