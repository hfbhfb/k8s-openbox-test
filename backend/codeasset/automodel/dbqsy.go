package automodel

import (
	"time"
)

// Goorders [...]
type Goorders struct {
	IDorder     int64     `gorm:"primaryKey;column:idorder" json:"-"`
	IDgouser    int64     `gorm:"column:idgouser" json:"idgouser"`       // idgouser和url构成唯一主键
	Createtime  time.Time `gorm:"column:createtime" json:"createtime"`   // 订单创建时间
	Finishtime  time.Time `gorm:"column:finishtime" json:"finishtime"`   // 订单完成时间
	Cost        int       `gorm:"column:cost" json:"cost"`               // 单一订单的消耗
	Status      int       `gorm:"column:status" json:"status"`           // 状态 1: 未处理  2:正在处理 3:已完成 4:失败
	Callbackurl string    `gorm:"column:callbackurl" json:"callbackurl"` // 处理完成后,用户可以通过这个标识获取下载路径
	URL         string    `gorm:"column:url" json:"url"`                 // idgouser和url构成唯一主键
	Parsetitle  string    `gorm:"column:parsetitle" json:"parsetitle"`
	Workerid    string    `gorm:"column:workerid" json:"workerid"` // 表示哪个渲染机器在使用
}

// TableName get sql table name.获取数据库表名
func (m *Goorders) TableName() string {
	return "goorders"
}

// GoordersColumns get sql column name.获取数据库列名
var GoordersColumns = struct {
	IDorder     string
	IDgouser    string
	Createtime  string
	Finishtime  string
	Cost        string
	Status      string
	Callbackurl string
	URL         string
	Parsetitle  string
	Workerid    string
}{
	IDorder:     "idorder",
	IDgouser:    "idgouser",
	Createtime:  "createtime",
	Finishtime:  "finishtime",
	Cost:        "cost",
	Status:      "status",
	Callbackurl: "callbackurl",
	URL:         "url",
	Parsetitle:  "parsetitle",
	Workerid:    "workerid",
}

// Gouser [...]
type Gouser struct {
	IDgouser    int64     `gorm:"primaryKey;column:idgouser" json:"-"`    // id 帐号
	LoginName   string    `gorm:"column:login_name" json:"loginName"`     // 登录姓名
	LoginPasswd string    `gorm:"column:login_passwd" json:"loginPasswd"` // 密码: md5
	CreatedAt   time.Time `gorm:"column:createdAt" json:"createdAt"`      // 创建时间
	UpdateAt    time.Time `gorm:"column:updateAt" json:"updateAt"`        // 更新时间
	Name        string    `gorm:"column:name" json:"name"`                // 姓名
	Gender      int8      `gorm:"column:gender" json:"gender"`            // 性别
	Email       string    `gorm:"column:email" json:"email"`
	Phone       string    `gorm:"column:phone" json:"phone"`
	IDRole      int       `gorm:"column:idRole" json:"idRole"` // 外键到goRole表
	LastLoginIP string    `gorm:"column:lastLoginIp" json:"lastLoginIp"`
	Gousercol   string    `gorm:"column:gousercol" json:"gousercol"`
	Account     int64     `gorm:"column:account" json:"account"` // 渲染币
}

// TableName get sql table name.获取数据库表名
func (m *Gouser) TableName() string {
	return "gouser"
}

// GouserColumns get sql column name.获取数据库列名
var GouserColumns = struct {
	IDgouser    string
	LoginName   string
	LoginPasswd string
	CreatedAt   string
	UpdateAt    string
	Name        string
	Gender      string
	Email       string
	Phone       string
	IDRole      string
	LastLoginIP string
	Gousercol   string
	Account     string
}{
	IDgouser:    "idgouser",
	LoginName:   "login_name",
	LoginPasswd: "login_passwd",
	CreatedAt:   "createdAt",
	UpdateAt:    "updateAt",
	Name:        "name",
	Gender:      "gender",
	Email:       "email",
	Phone:       "phone",
	IDRole:      "idRole",
	LastLoginIP: "lastLoginIp",
	Gousercol:   "gousercol",
	Account:     "account",
}
