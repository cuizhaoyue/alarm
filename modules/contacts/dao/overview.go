package dao

import (
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"gorm.io/gorm"
)

type Factory interface {
	Contacts() ContactStore
}

var _ Factory = &datastore{}

type datastore struct {
	db *gorm.DB
}

func NewDataStore(db *gorm.DB) *datastore {
	if db == nil {
		db = lib.GORMDefaultPool
	}
	return &datastore{db: db}
}

func (ds *datastore) Contacts() ContactStore {
	return newContactStore(ds)
}

// var (
// 	mysqlFactory Factory
// 	once         sync.Once
// )

// func GetMysqlFactory(db *gorm.DB) Factory {

// 	return newDataStore(db)
// }

// // GetTestMysqlFactory 测试时使用
// func GetTestMysqlFactory() (Factory, error) {
// 	var err error
// 	var dbIns *gorm.DB

// 	once.Do(func() {
// 		dbIns, err = storage.NewTestGormDB()
// 		mysqlFactory = newDataStore(dbIns)
// 	})

// 	if mysqlFactory == nil || err != nil {
// 		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
// 	}

// 	return mysqlFactory, nil
// }
