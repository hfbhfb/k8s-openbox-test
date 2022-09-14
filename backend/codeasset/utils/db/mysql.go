package db

import (
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	defaultMaxOpenConns = 50     // 最大连接数
	defaultMaxIdleConns = 10     // 最大空闲连接
	defaultAlias        = "base" // 默认的实例名称

	defaultSession *gorm.DB
	pool           = make(map[string]*gorm.DB)
	l              sync.RWMutex
)

type Config struct {
	Alias      string
	DataSource string
}

func Init(conf []*Config) error {
	l.Lock()
	defer l.Unlock()

	var err error
	for _, v := range conf {
		if _, ok := pool[v.Alias]; !ok {
			pool[v.Alias], err = NewSession(v.DataSource, v.Alias)
			if err != nil {
				return err
			}
		}
	}

	if session, ok := pool[defaultAlias]; ok {
		defaultSession = session
	}
	return nil
}

func NewSession(dataSource string, alias string) (*gorm.DB, error) {
	if alias == "sqlite" {
		conn, err := gorm.Open(sqlite.Open("gormsqlite.db"), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return conn, nil
	} else {
		//conn, err := gorm.Open("mysql", dataSource)
		for i := 0; i < 80; i++ {
			conn, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
			if err != nil {
				if i < 80 {
					time.Sleep(time.Second)
					continue
				}
				return nil, err
			}
			db, _ := conn.DB()
			db.SetMaxOpenConns(defaultMaxOpenConns)
			db.SetMaxIdleConns(defaultMaxIdleConns)
			db.SetConnMaxLifetime(time.Second * 60)
			return conn, nil
		}
	}

	return nil, errors.Errorf("mysql init")
}

// 获取默认DB
func DefaultSession() *gorm.DB {
	return defaultSession
}

// 设置数据库最大连接数
func SetMaxOpenConns(alias string, maxOpenConns int) error {
	session, err := GetInstanceSession(alias)

	if err != nil {
		return err
	}
	db, _ := session.DB()
	db.SetMaxOpenConns(maxOpenConns)
	return nil
}

// 设置数据库最大空闲连接数
func SetMaxIdleConns(alias string, maxIdleConns int) error {
	session, err := GetInstanceSession(alias)
	if err != nil {
		return err
	}
	db, _ := session.DB()
	db.SetMaxIdleConns(maxIdleConns)
	return nil
}

// 获取指定实例名称的session
func GetInstanceSession(name string) (*gorm.DB, error) {
	l.RLock()
	defer l.RUnlock()
	if session, ok := pool[name]; ok {
		return session, nil
	}
	return nil, errors.New("unknown DataBase alias name :" + name)
}

func RegisterLogger(name string, loger *logrus.Logger) error {
	l.RLock()
	defer l.RUnlock()
	//if session, ok := pool[name]; ok {
	//gorm.LogMode(true)
	//session.SetLogger(loger)
	//}
	return errors.New("unknown DataBase alias name :" + name)
}
