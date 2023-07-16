package storage

import (
	"sync"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
	err  error
)

// NewGormDB 创建gorm.DB实例，供临时测试使用
func NewGormDB() (*gorm.DB, error) {
	// 使用公共库获取数据库连接池
	once.Do(func() {
		var mysqlPool string
		if lib.GetEnvInfo("IS_DEV") {
			mysqlPool = consts.DefaultMysqlPool
		} else {
			mysqlPool = consts.ProMysqlPool
		}
		db, err = lib.GetGormPool(mysqlPool)
	})

	return db, err
}

// NewTestGormDB 测试时使用
func NewTestGormDB() (*gorm.DB, error) {

	// 创建数据库dsn
	dsn := "xxxxx:xxxxxxxx123@tcp(127.0.0.1:7001)/alert?charset=utf8mb4&parseTime=true&loc=Local"

	// 创建数据库连接池
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置MySQL最大连接数
	sqlDB.SetMaxOpenConns(100)

	// 设置MySQL空闲连接最大存活时间
	sqlDB.SetConnMaxLifetime(10)

	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(100)

	return db, nil
}
