package automodel

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _GouserMgr struct {
	*_BaseMgr
}

// GouserMgr open func
func GouserMgr(db *gorm.DB) *_GouserMgr {
	if db == nil {
		panic(fmt.Errorf("GouserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GouserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gouser"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GouserMgr) GetTableName() string {
	return "gouser"
}

// Get 获取
func (obj *_GouserMgr) Get() (result Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GouserMgr) Gets() (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GouserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gouser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIDgouser idgouser获取 id 帐号
func (obj *_GouserMgr) WithIDgouser(idgouser int64) Option {
	return optionFunc(func(o *options) { o.query["idgouser"] = idgouser })
}

// WithLoginName login_name获取 登录姓名
func (obj *_GouserMgr) WithLoginName(loginName string) Option {
	return optionFunc(func(o *options) { o.query["login_name"] = loginName })
}

// WithLoginPasswd login_passwd获取 密码: md5
func (obj *_GouserMgr) WithLoginPasswd(loginPasswd string) Option {
	return optionFunc(func(o *options) { o.query["login_passwd"] = loginPasswd })
}

// WithCreatedAt createdAt获取 创建时间
func (obj *_GouserMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["createdAt"] = createdAt })
}

// WithUpdateAt updateAt获取 更新时间
func (obj *_GouserMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updateAt"] = updateAt })
}

// WithName name获取 姓名
func (obj *_GouserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithGender gender获取 性别
func (obj *_GouserMgr) WithGender(gender int8) Option {
	return optionFunc(func(o *options) { o.query["gender"] = gender })
}

// WithEmail email获取
func (obj *_GouserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPhone phone获取
func (obj *_GouserMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithIDRole idRole获取 外键到goRole表
func (obj *_GouserMgr) WithIDRole(idRole int) Option {
	return optionFunc(func(o *options) { o.query["idRole"] = idRole })
}

// WithLastLoginIP lastLoginIp获取
func (obj *_GouserMgr) WithLastLoginIP(lastLoginIP string) Option {
	return optionFunc(func(o *options) { o.query["lastLoginIp"] = lastLoginIP })
}

// WithGousercol gousercol获取
func (obj *_GouserMgr) WithGousercol(gousercol string) Option {
	return optionFunc(func(o *options) { o.query["gousercol"] = gousercol })
}

// WithAccount account获取 渲染币
func (obj *_GouserMgr) WithAccount(account int64) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// GetByOption 功能选项模式获取
func (obj *_GouserMgr) GetByOption(opts ...Option) (result Gouser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GouserMgr) GetByOptions(opts ...Option) (results []*Gouser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDgouser 通过idgouser获取内容 id 帐号
func (obj *_GouserMgr) GetFromIDgouser(idgouser int64) (result Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`idgouser` = ?", idgouser).Find(&result).Error

	return
}

// GetBatchFromIDgouser 批量查找 id 帐号
func (obj *_GouserMgr) GetBatchFromIDgouser(idgousers []int64) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`idgouser` IN (?)", idgousers).Find(&results).Error

	return
}

// GetFromLoginName 通过login_name获取内容 登录姓名
func (obj *_GouserMgr) GetFromLoginName(loginName string) (result Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`login_name` = ?", loginName).Find(&result).Error

	return
}

// GetBatchFromLoginName 批量查找 登录姓名
func (obj *_GouserMgr) GetBatchFromLoginName(loginNames []string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`login_name` IN (?)", loginNames).Find(&results).Error

	return
}

// GetFromLoginPasswd 通过login_passwd获取内容 密码: md5
func (obj *_GouserMgr) GetFromLoginPasswd(loginPasswd string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`login_passwd` = ?", loginPasswd).Find(&results).Error

	return
}

// GetBatchFromLoginPasswd 批量查找 密码: md5
func (obj *_GouserMgr) GetBatchFromLoginPasswd(loginPasswds []string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`login_passwd` IN (?)", loginPasswds).Find(&results).Error

	return
}

// GetFromCreatedAt 通过createdAt获取内容 创建时间
func (obj *_GouserMgr) GetFromCreatedAt(createdAt time.Time) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`createdAt` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_GouserMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`createdAt` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过updateAt获取内容 更新时间
func (obj *_GouserMgr) GetFromUpdateAt(updateAt time.Time) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`updateAt` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 更新时间
func (obj *_GouserMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`updateAt` IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_GouserMgr) GetFromName(name string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 姓名
func (obj *_GouserMgr) GetBatchFromName(names []string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromGender 通过gender获取内容 性别
func (obj *_GouserMgr) GetFromGender(gender int8) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`gender` = ?", gender).Find(&results).Error

	return
}

// GetBatchFromGender 批量查找 性别
func (obj *_GouserMgr) GetBatchFromGender(genders []int8) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`gender` IN (?)", genders).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_GouserMgr) GetFromEmail(email string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找
func (obj *_GouserMgr) GetBatchFromEmail(emails []string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容
func (obj *_GouserMgr) GetFromPhone(phone string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找
func (obj *_GouserMgr) GetBatchFromPhone(phones []string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromIDRole 通过idRole获取内容 外键到goRole表
func (obj *_GouserMgr) GetFromIDRole(idRole int) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`idRole` = ?", idRole).Find(&results).Error

	return
}

// GetBatchFromIDRole 批量查找 外键到goRole表
func (obj *_GouserMgr) GetBatchFromIDRole(idRoles []int) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`idRole` IN (?)", idRoles).Find(&results).Error

	return
}

// GetFromLastLoginIP 通过lastLoginIp获取内容
func (obj *_GouserMgr) GetFromLastLoginIP(lastLoginIP string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`lastLoginIp` = ?", lastLoginIP).Find(&results).Error

	return
}

// GetBatchFromLastLoginIP 批量查找
func (obj *_GouserMgr) GetBatchFromLastLoginIP(lastLoginIPs []string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`lastLoginIp` IN (?)", lastLoginIPs).Find(&results).Error

	return
}

// GetFromGousercol 通过gousercol获取内容
func (obj *_GouserMgr) GetFromGousercol(gousercol string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`gousercol` = ?", gousercol).Find(&results).Error

	return
}

// GetBatchFromGousercol 批量查找
func (obj *_GouserMgr) GetBatchFromGousercol(gousercols []string) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`gousercol` IN (?)", gousercols).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容 渲染币
func (obj *_GouserMgr) GetFromAccount(account int64) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`account` = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量查找 渲染币
func (obj *_GouserMgr) GetBatchFromAccount(accounts []int64) (results []*Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`account` IN (?)", accounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GouserMgr) FetchByPrimaryKey(idgouser int64) (result Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`idgouser` = ?", idgouser).Find(&result).Error

	return
}

// FetchUniqueByNameUNIQUE primary or index 获取唯一内容
func (obj *_GouserMgr) FetchUniqueByNameUNIQUE(loginName string) (result Gouser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gouser{}).Where("`login_name` = ?", loginName).Find(&result).Error

	return
}
