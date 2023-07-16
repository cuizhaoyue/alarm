package dao

import (
	"context"
	"fmt"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"gorm.io/gorm"
)

// ResourceOnPolicyStore 告警策略选择的资源相关的接口
type ResourceOnPolicyStore interface {
	Create(ctx context.Context, rs []*model.ResourceOnPolicy) error
	Delete(ctx context.Context, polId string) error
	DeleteCollection(ctx context.Context, polIds []string) error
	ListResourceByName(ctx context.Context, req *dto.ListResourcesOnPolicyRequest) (*model.ResourceOnPolicyList, error)
}

var _ ResourceOnPolicyStore = &resourceOnPolicyStore{}

type resourceOnPolicyStore struct {
	db *gorm.DB
}

func newResourceOnPolicyStore(ds *datastore) *resourceOnPolicyStore {
	return &resourceOnPolicyStore{
		db: ds.db,
	}
}

// Create 保存告警策略选择的资源记录到数据库
func (r *resourceOnPolicyStore) Create(ctx context.Context, rs []*model.ResourceOnPolicy) error {
	if err := r.db.Create(rs).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("create resources on policy failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// Delete 通过告警策略删除记录
func (r *resourceOnPolicyStore) Delete(ctx context.Context, polId string) error {
	if err := r.db.Where("policy_id = ?", polId).Delete(&model.ResourceOnPolicy{}).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("delete resources by policy id [%s] failed, %s", polId, err).Error(),
		})

		return err
	}
	return nil
}

// DeleteCollection 批量删除
func (r *resourceOnPolicyStore) DeleteCollection(ctx context.Context, polIds []string) error {
	if err := r.db.Where("policy_id in ?", polIds).Delete(&model.ResourceOnPolicy{}).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("delete a batch of resources failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// List 根据资源名称做模糊搜索
func (r *resourceOnPolicyStore) ListResourceByName(ctx context.Context, req *dto.ListResourcesOnPolicyRequest) (*model.ResourceOnPolicyList, error) {
	ret := &model.ResourceOnPolicyList{}

	err := r.db.Where("name like ?", "%"+req.Name+"%").
		Offset(req.Offset).
		Limit(req.PageSize).
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.Total).Error

	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("list resources on policy failed, %s", err).Error(),
		})

		return nil, err
	}

	return ret, nil
}
