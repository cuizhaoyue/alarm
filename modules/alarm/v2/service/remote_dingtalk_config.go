package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"
)

// ListDingtalkConfig 获取钉钉配置列表
func (r *remoteConfigService) ListDingtalkConfig(ctx context.Context, req *dto.ListDingtalkConfigRequest) (*model.DingtalkConfigList, error) {
	req.SetupOffset()

	dataList, err := r.store.RemoteConfig().ListDingtalkConfig(ctx, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dataList, nil
}

// CreateDingtalkConfig 创建钉钉配置列表
func (r *remoteConfigService) CreateDingtalkConfig(ctx context.Context, req *dto.CreateDingtalkConfigRequest) error {
	cfg := &model.DingtalkConfig{
		Name: req.Name,
		Url:  req.Url,
	}

	if err := r.store.RemoteConfig().CreateDingtalkConfig(ctx, cfg); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// DeleteDingtalkConfig 删除钉钉配置列表
func (r *remoteConfigService) DeleteDingtalkConfig(ctx context.Context, req *dto.DeleteDingtalkConfigRequest) error {
	if err := r.store.RemoteConfig().DeleteDingtalkConfig(ctx, req.Id); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// UpdateDingtalkConfig 更新钉钉配置列表
func (r *remoteConfigService) UpdateDingtalkConfig(ctx context.Context, req *dto.UpdateDingtalkConfigRequest) error {
	cfg := &model.DingtalkConfig{
		Id:   req.Id,
		Name: req.Name,
		Url:  req.Url,
	}

	if err := r.store.RemoteConfig().UpdateDingtalkConfig(ctx, cfg); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
