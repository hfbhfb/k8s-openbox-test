package automodel

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _GoordersMgr struct {
	*_BaseMgr
}

// GoordersMgr open func
func GoordersMgr(db *gorm.DB) *_GoordersMgr {
	if db == nil {
		panic(fmt.Errorf("GoordersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoordersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("goorders"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoordersMgr) GetTableName() string {
	return "goorders"
}

// Get 获取
func (obj *_GoordersMgr) Get() (result Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoordersMgr) Gets() (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoordersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Goorders{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIDorder idorder获取
func (obj *_GoordersMgr) WithIDorder(idorder int64) Option {
	return optionFunc(func(o *options) { o.query["idorder"] = idorder })
}

// WithIDgouser idgouser获取 idgouser和url构成唯一主键
func (obj *_GoordersMgr) WithIDgouser(idgouser int64) Option {
	return optionFunc(func(o *options) { o.query["idgouser"] = idgouser })
}

// WithCreatetime createtime获取 订单创建时间
func (obj *_GoordersMgr) WithCreatetime(createtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtime"] = createtime })
}

// WithFinishtime finishtime获取 订单完成时间
func (obj *_GoordersMgr) WithFinishtime(finishtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["finishtime"] = finishtime })
}

// WithCost cost获取 单一订单的消耗
func (obj *_GoordersMgr) WithCost(cost int) Option {
	return optionFunc(func(o *options) { o.query["cost"] = cost })
}

// WithStatus status获取 状态 1: 未处理  2:正在处理 3:已完成 4:失败
func (obj *_GoordersMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCallbackurl callbackurl获取 处理完成后,用户可以通过这个标识获取下载路径
func (obj *_GoordersMgr) WithCallbackurl(callbackurl string) Option {
	return optionFunc(func(o *options) { o.query["callbackurl"] = callbackurl })
}

// WithURL url获取 idgouser和url构成唯一主键
func (obj *_GoordersMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithParsetitle parsetitle获取
func (obj *_GoordersMgr) WithParsetitle(parsetitle string) Option {
	return optionFunc(func(o *options) { o.query["parsetitle"] = parsetitle })
}

// WithWorkerid workerid获取 表示哪个渲染机器在使用
func (obj *_GoordersMgr) WithWorkerid(workerid string) Option {
	return optionFunc(func(o *options) { o.query["workerid"] = workerid })
}

// GetByOption 功能选项模式获取
func (obj *_GoordersMgr) GetByOption(opts ...Option) (result Goorders, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoordersMgr) GetByOptions(opts ...Option) (results []*Goorders, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDorder 通过idorder获取内容
func (obj *_GoordersMgr) GetFromIDorder(idorder int64) (result Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`idorder` = ?", idorder).Find(&result).Error

	return
}

// GetBatchFromIDorder 批量查找
func (obj *_GoordersMgr) GetBatchFromIDorder(idorders []int64) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`idorder` IN (?)", idorders).Find(&results).Error

	return
}

// GetFromIDgouser 通过idgouser获取内容 idgouser和url构成唯一主键
func (obj *_GoordersMgr) GetFromIDgouser(idgouser int64) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`idgouser` = ?", idgouser).Find(&results).Error

	return
}

// GetBatchFromIDgouser 批量查找 idgouser和url构成唯一主键
func (obj *_GoordersMgr) GetBatchFromIDgouser(idgousers []int64) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`idgouser` IN (?)", idgousers).Find(&results).Error

	return
}

// GetFromCreatetime 通过createtime获取内容 订单创建时间
func (obj *_GoordersMgr) GetFromCreatetime(createtime time.Time) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`createtime` = ?", createtime).Find(&results).Error

	return
}

// GetBatchFromCreatetime 批量查找 订单创建时间
func (obj *_GoordersMgr) GetBatchFromCreatetime(createtimes []time.Time) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`createtime` IN (?)", createtimes).Find(&results).Error

	return
}

// GetFromFinishtime 通过finishtime获取内容 订单完成时间
func (obj *_GoordersMgr) GetFromFinishtime(finishtime time.Time) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`finishtime` = ?", finishtime).Find(&results).Error

	return
}

// GetBatchFromFinishtime 批量查找 订单完成时间
func (obj *_GoordersMgr) GetBatchFromFinishtime(finishtimes []time.Time) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`finishtime` IN (?)", finishtimes).Find(&results).Error

	return
}

// GetFromCost 通过cost获取内容 单一订单的消耗
func (obj *_GoordersMgr) GetFromCost(cost int) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`cost` = ?", cost).Find(&results).Error

	return
}

// GetBatchFromCost 批量查找 单一订单的消耗
func (obj *_GoordersMgr) GetBatchFromCost(costs []int) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`cost` IN (?)", costs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态 1: 未处理  2:正在处理 3:已完成 4:失败
func (obj *_GoordersMgr) GetFromStatus(status int) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态 1: 未处理  2:正在处理 3:已完成 4:失败
func (obj *_GoordersMgr) GetBatchFromStatus(statuss []int) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCallbackurl 通过callbackurl获取内容 处理完成后,用户可以通过这个标识获取下载路径
func (obj *_GoordersMgr) GetFromCallbackurl(callbackurl string) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`callbackurl` = ?", callbackurl).Find(&results).Error

	return
}

// GetBatchFromCallbackurl 批量查找 处理完成后,用户可以通过这个标识获取下载路径
func (obj *_GoordersMgr) GetBatchFromCallbackurl(callbackurls []string) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`callbackurl` IN (?)", callbackurls).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 idgouser和url构成唯一主键
func (obj *_GoordersMgr) GetFromURL(url string) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找 idgouser和url构成唯一主键
func (obj *_GoordersMgr) GetBatchFromURL(urls []string) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromParsetitle 通过parsetitle获取内容
func (obj *_GoordersMgr) GetFromParsetitle(parsetitle string) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`parsetitle` = ?", parsetitle).Find(&results).Error

	return
}

// GetBatchFromParsetitle 批量查找
func (obj *_GoordersMgr) GetBatchFromParsetitle(parsetitles []string) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`parsetitle` IN (?)", parsetitles).Find(&results).Error

	return
}

// GetFromWorkerid 通过workerid获取内容 表示哪个渲染机器在使用
func (obj *_GoordersMgr) GetFromWorkerid(workerid string) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`workerid` = ?", workerid).Find(&results).Error

	return
}

// GetBatchFromWorkerid 批量查找 表示哪个渲染机器在使用
func (obj *_GoordersMgr) GetBatchFromWorkerid(workerids []string) (results []*Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`workerid` IN (?)", workerids).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoordersMgr) FetchByPrimaryKey(idorder int64) (result Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`idorder` = ?", idorder).Find(&result).Error

	return
}

// FetchUniqueIndexByUserURL primary or index 获取唯一内容
func (obj *_GoordersMgr) FetchUniqueIndexByUserURL(idgouser int64, url string) (result Goorders, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Goorders{}).Where("`idgouser` = ? AND `url` = ?", idgouser, url).Find(&result).Error

	return
}
