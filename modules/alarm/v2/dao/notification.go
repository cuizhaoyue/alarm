package dao

import (
	"context"
	"fmt"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// AlertStore 告警通知接口
type AlertStore interface {
	CreateOrUpdate(ctx context.Context, alerts []*model.Alert) error
	List(ctx context.Context, req *dto.ListAlertsRequest) (*model.AlertList, error)
}

var _ AlertStore = &alertStore{}

type alertStore struct {
	db *gorm.DB
}

func newAlertStore(ds *datastore) *alertStore {
	return &alertStore{ds.db}
}

// CreateOrUpdate 创建或更新告警消息记录
func (a *alertStore) CreateOrUpdate(ctx context.Context, alerts []*model.Alert) error {
	err := a.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "alert_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"handler", "current_value", "ends_at", "duration", "duration_flag", "status"}),
	}).Create(alerts).Error

	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("create or update alert notification failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// List 获取告警列表
func (a *alertStore) List(ctx context.Context, req *dto.ListAlertsRequest) (*model.AlertList, error) {
	ret := &model.AlertList{}

	// 模糊搜索
	if req.SearchValue != "" {
		switch req.SearchKey {
		case "name":
			// 告警名称
			a.db = a.db.Where("name like ?", "%"+req.SearchValue+"%")
		case "policy_name":
			// 策略名称
			a.db = a.db.Where("policy_name like ?", "%"+req.SearchValue+"%")
		case "resource":
			// 告警实例
			a.db = a.db.Where("resource like ?", "%"+req.SearchValue+"%")
		}
	}

	if req.Region != "all" && req.Region != "" {
		a.db = a.db.Where("region = ?", req.Region)
	}

	if len(req.Az) > 0 {
		a.db = a.db.Where("az in ?", req.Az)
	}

	if len(req.Level) > 0 {
		a.db = a.db.Where("level in ?", req.Level)
	}

	if len(req.ResourceType) > 0 {
		a.db = a.db.Where("resource_type in ?", req.ResourceType)
	}

	if len(req.ResourceSubType) > 0 {
		a.db = a.db.Where("resource_sub_type in ?", req.ResourceSubType)
	}

	// 根据告警持续时长筛选
	if len(req.Duration) > 0 {
		a.db = a.db.Where("duration_flag in ?", req.Duration)
	}

	err := a.db.Where("status = ?", req.Status).
		Offset(req.Offset).
		Limit(req.PageSize).
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.Total).Error

	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("get alerts list failed, %s", err).Error(),
		})

		return nil, err
	}

	return ret, nil
}
