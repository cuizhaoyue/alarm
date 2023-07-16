package dao

import (
	"context"
	"fmt"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"gorm.io/gorm"
)

type AlarmStore interface {
	Tx(ctx context.Context, f AlarmTxFunc) error
	Create(ctx context.Context, policy *model.AlarmPolicy) error
	Update(ctx context.Context, policy *model.AlarmPolicy) error
	Get(ctx context.Context, InsId string) (*model.AlarmPolicy, error)
	List(ctx context.Context, req *dto.ListPolicyRequest) (*model.AlarmPolicyList, error)
	ListByInstanceIds(ctx context.Context, ids []string) ([]*model.AlarmPolicy, error)
	DeleteCollection(ctx context.Context, instanceIds []string) error
}

type AlarmTxFunc func(ctx context.Context, store Factory) error

var _ AlarmStore = &alarmStore{}

type alarmStore struct {
	db *gorm.DB
}

func newAlarmStore(ds *datastore) *alarmStore {
	return &alarmStore{ds.db}
}

// Tx 用于执行事务的接口
func (as *alarmStore) Tx(ctx context.Context, f AlarmTxFunc) error {
	return as.db.Transaction(func(tx *gorm.DB) error {
		// 使用开启事务的tx重新初始化一个Factory
		txStore := NewDataStore(tx)

		// 执行注册的事务函数
		return f(ctx, txStore)
	})
}

// Create 新建告警策略数据
func (as *alarmStore) Create(ctx context.Context, policy *model.AlarmPolicy) error {
	err := as.db.Create(policy).Error

	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("create alarm policy failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// Update 更新数据库
func (as *alarmStore) Update(ctx context.Context, policy *model.AlarmPolicy) error {
	if err := as.db.Save(policy).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("update policy failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// Get 获取告警策略详情
func (as *alarmStore) Get(ctx context.Context, InsId string) (*model.AlarmPolicy, error) {
	ret := &model.AlarmPolicy{}
	if err := as.db.Where("instance_id = ?", InsId).First(ret).Error; err != nil {

		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("get alarm policy instance failed, %s", err).Error(),
		})
		return nil, err
	}

	return ret, nil
}

// List 获取告警策略列表
func (as *alarmStore) List(ctx context.Context, req *dto.ListPolicyRequest) (*model.AlarmPolicyList, error) {
	ret := &model.AlarmPolicyList{}

	// 精确搜索
	if req.Type == model.FormAlarmPolicyType {
		if len(req.ResourceType) > 0 {
			as.db = as.db.Where("resource_type in ?", req.ResourceType)
		}
		if len(req.ResourceSubType) > 0 {
			as.db = as.db.Where("resource_sub_type in ?", req.ResourceSubType)
		}
	}

	// 模糊匹配搜索
	if req.SearchValue != "" {
		switch req.SearchKey {
		case "name":
			// 告警策略名称
			as.db = as.db.Where("name like ?", "%"+req.SearchValue+"%")
		case "creator":
			// 创建者
			as.db = as.db.Where("creator like ?", "%"+req.SearchValue+"%")
		case "production":
			// 业务线
			as.db = as.db.Where("production like ?", "%"+req.SearchValue+"%")
		}
	}

	// 查询数据库
	err := as.db.Where("type = ?", req.Type).
		Offset(req.Offset).
		Limit(req.PageSize).
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.Total).Error
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("get policy list failed, %s", err).Error(),
		})

		return nil, err
	}

	return ret, nil
}

// ListByInstanceIds 根据instance id获取告警策略
func (as *alarmStore) ListByInstanceIds(ctx context.Context, ids []string) ([]*model.AlarmPolicy, error) {
	var ret []*model.AlarmPolicy
	if err := as.db.Where("instance_id in ?", ids).Find(&ret).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("get policy list by instance id failed, %s", err).Error(),
		})

		return nil, err
	}

	return ret, nil
}

// DeleteCollection 批量删除告警策略
func (as *alarmStore) DeleteCollection(ctx context.Context, instanceIds []string) error {
	if err := as.db.Unscoped().
		Where("instance_id in ?", instanceIds).
		Delete(&model.AlarmPolicy{}).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("delete policies failed, %s", err).Error(),
		})

		return err
	}

	return nil
}
