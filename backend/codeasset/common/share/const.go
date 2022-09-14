package share

import "time"

const (
	NAME    = "k8s-openbox-test"
	VERSION = "v0.0.1"
)

const (
	EXPIREDTIME    = (24 * time.Hour) //
	TIMELAYOUT     = "2006-01-02 15:04:05"
	BLOCKDAYLAYOUT = "20060102"
)

const (
	FABRIC_VERSION    = "Version"
	FABRIC_GETDATA    = "GetData"
	FABRIC_SAVEDATA   = "SaveData"
	FABRIC_UPDATEDATE = "UpdateData"
	FABRIC_DELETEDATA = "DeleteData"
	FABRIC_GETHISTORY = "GetHistory"
)

const (
	BcApiAddUser        = "/api/user"
	BcApiAddResource    = "/api/unsigndata"
	BcApiDelResource    = "/api/data/"   //api/data/{{record-id}}
	BcApiModifyResource = "/api/undata/" //api/undata/{{record-id}}

	CompanyNameMaxLong   = 20 //公司名称长度限制
	NodeNameMaxLong      = 20 //节点名称长度限制
	ResourceTitleMaxLong = 20 //资源标题长度限制
	ResourceDescMaxLong  = 50 //资源描述长度限制
)
