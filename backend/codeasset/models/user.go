package models

// type User struct {
// 	automodel.Gouser
// }

// type _UserOrmManager struct {
// 	*_BaseMgr
// }

// func NewUser(db *gorm.DB) *_UserOrmManager {
// 	if db == nil {
// 		panic(fmt.Errorf("GoordersMgr need init by db"))
// 	}
// 	ctx, cancel := context.WithCancel(context.Background())
// 	return &_UserOrmManager{_BaseMgr: &_BaseMgr{DB: db.Table("goorders"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
// }

// func (obj *_UserOrmManager) Com() (result User, err error) {
// 	err = obj.DB.WithContext(obj.ctx).Model(automodel.Gouser{}).Find(&result).Error
// 	return
// }
