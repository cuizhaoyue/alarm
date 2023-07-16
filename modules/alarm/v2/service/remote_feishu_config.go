package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"
)

// GetFeishuConfig 获取企业微信配置
func (r *remoteConfigService) GetFeishuConfig(ctx context.Context) (*dto.GetFeishuConfigResponseData, error) {
	cfg, err := r.store.RemoteConfig().GetFeishuConfig(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	data := &dto.GetFeishuConfigResponseData{
		Url: cfg.Url,
	}

	return data, nil
}

// CreateOrUpdateFeishuConfig 创建或更新企业微信配置
func (r *remoteConfigService) CreateOrUpdateFeishuConfig(ctx context.Context, req *dto.UpdateFeishuConfigRequest) error {
	// 先从数据库中获取配置
	cfg, err := r.store.RemoteConfig().GetFeishuConfig(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// 判断数据库中是否存在配置，如果不存在则直接创建
	if cfg == nil {
		cfg = &model.FeishuConfig{
			Url: req.Url,
		}

		if err := r.store.RemoteConfig().CreateFeishuConfig(ctx, cfg); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}

	// 更新配置到数据库
	cfg.Url = req.Url
	if err := r.store.RemoteConfig().UpdateFeishuConfig(ctx, cfg); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
