package models

import (
	"fmt"

	"strconv"

	"backend/codeasset/automodel"
	. "backend/codeasset/common/share"
	"backend/codeasset/utils/config"
	"backend/codeasset/utils/db"
	"backend/codeasset/utils/snowflake"
	"backend/codeasset/utils/tlog"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

const (
	defaultInstanceName = "base" // 默认的实例名称
)

var (
	Config    *config.Config
	IDNode    *snowflake.Node
	BaseData  *gorm.DB
	BaseRedis *redis.Client
)

func initMysql() {

	Config = config.Cfg
	// logger init
	tlog.Logger(Config.Env, NAME, strconv.Itoa(Config.NodeID))

	if Config.DB.Alias == "sqlite" {
		tlog.Debug("sqlite prepare")
	}

	if Config.DB.Host == "" {
		return
	}
	fmt.Printf("\n models init 1\n")
	// mysql init
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", Config.DB.User, Config.DB.Password, Config.DB.Host, Config.DB.Port, Config.DB.Database)
	dbConfig := &db.Config{
		Alias:      Config.DB.Alias,
		DataSource: source,
	}
	err := db.Init([]*db.Config{dbConfig})
	if err != nil {
		// tlog.Panic("init db error", err)
		tlog.Error("init mysql db error")
	}
	db.RegisterLogger(dbConfig.Alias, tlog.GetLogger())
	setBaseData(dbConfig.Alias) // set base

}

func initRedis() {

	if Config.RDB.Host == "" {
		return
	}
	var err error
	rConfig := &db.RConfig{
		Host:     Config.RDB.Host,
		Port:     Config.RDB.Port,
		Password: Config.RDB.Password,
		DB:       Config.RDB.DB,
	}
	fmt.Printf("\n models init rConfig = %+v\n", rConfig)
	BaseRedis, err = db.InitRedis(rConfig)
	if err != nil {
		// tlog.Panic("InitRedis init panic")
		tlog.Error("InitRedis  redis init panic")
	}

	// snowflake init
	IDNode, err = snowflake.NewNode(int64(Config.NodeID))
	if err != nil {
		tlog.Panic("snowflake init panic")
	}

}

func init() {

	initMysql()
	initRedis()
}

// save basedata
func setBaseData(instanceName string) {
	if db, err := db.GetInstanceSession(instanceName); err == nil {
		BaseData = db
	} else {
		tlog.Panic("set basedata error", err)
	}
	if instanceName == "sqlite" {
		// sqlite自己把库准备好
		BaseData.AutoMigrate(
			&automodel.Goorders{},
		)
	}
	//BaseData.AutoMigrate(
	//	&classicmodelsv1.Customers{},
	//	&classicmodelsv1.Employees{},
	//	&classicmodelsv1.Offices{},
	//	&classicmodelsv1.Orderdetails{},
	//	&classicmodelsv1.Orders{},
	//	&classicmodelsv1.Payments{},
	//	&classicmodelsv1.Productlines{},
	//	&classicmodelsv1.Products{},
	//)
}
