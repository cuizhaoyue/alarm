package dao

import (
	"fmt"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao/alertmanagerconfig"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao/clientset"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao/prometheusoperator"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"gorm.io/gorm"

	clientversioned "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned"
)

type Factory interface {
	Alarm() AlarmStore
	PromRuleOperator() prometheusoperator.PromRuleOperator
	AlertmanagerConfig() alertmanagerconfig.AlertmanagerConfiger
	PromQLTpl() PromQLTplStore
	ResourceOnPolicy() ResourceOnPolicyStore
	RemoteConfig() RemoteConfig
	Alert() AlertStore
}

var _ Factory = &datastore{}

type datastore struct {
	db        *gorm.DB
	clientSet *clientversioned.Clientset
}

func NewDataStore(db *gorm.DB) Factory {
	clientSet, err := clientset.PromOperatorClientSet()
	if err != nil {
		panic(fmt.Errorf("get clientset error: %s", err))
	}

	if db == nil {
		db = lib.GORMDefaultPool
	}

	return &datastore{
		db:        db,
		clientSet: clientSet,
	}
}

// Alarm 创建告警策略接口
func (ds *datastore) Alarm() AlarmStore {
	return newAlarmStore(ds)
}

// PromRuleOperator 创建PrometheusRule的操作接口
func (ds *datastore) PromRuleOperator() prometheusoperator.PromRuleOperator {
	return prometheusoperator.NewPromRuleOperator(ds.clientSet)
}

// AlertmanagerConfig 创建AlertmanagerConfig的操作接口
func (ds *datastore) AlertmanagerConfig() alertmanagerconfig.AlertmanagerConfiger {
	return alertmanagerconfig.NewAlertmanagerConfiger(ds.clientSet)
}

// PromQLTpl 创建模板接口
func (ds *datastore) PromQLTpl() PromQLTplStore {
	return newPromTplStore(ds)
}

// ResourceOnPolicy 创建资源接口
func (ds *datastore) ResourceOnPolicy() ResourceOnPolicyStore {
	return newResourceOnPolicyStore(ds)
}

// RemoteConfig 远程配置接口
func (ds *datastore) RemoteConfig() RemoteConfig {
	return newRemoteConfig(ds)
}

// Alert 告警消息接口
func (ds *datastore) Alert() AlertStore {
	return newAlertStore(ds)
}

// var (
// 	factoryIns Factory
// 	once       sync.Once
// )

// // GetFactoryInstance 获取全局Factory实例
// func GetFactoryInstance(db *gorm.DB) Factory {
// 	if factoryIns == nil {
// 		factoryIns = newDataStore(db)
// 	}

// 	return factoryIns
// }

// // GetTestMysqlFactory 测试时使用
// func GetTestMysqlFactory() (Factory, error) {
// 	var err error
// 	var dbIns *gorm.DB

// 	once.Do(func() {
// 		dbIns, err = storage.NewTestGormDB()
// 		factoryIns = newDataStore(dbIns)
// 	})

// 	if factoryIns == nil || err != nil {
// 		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", factoryIns, err)
// 	}

// 	return factoryIns, nil
// }
