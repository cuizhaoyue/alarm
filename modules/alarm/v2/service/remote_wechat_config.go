package service

import (
	"context"
	"strings"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"
)

// GetWechatConfig 获取企业微信配置
func (r *remoteConfigService) GetWechatConfig(ctx context.Context) (*dto.GetWechatConfigResponseData, error) {
	cfg, err := r.store.RemoteConfig().GetWechatConfig(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	data := &dto.GetWechatConfigResponseData{
		CompanyId: cfg.CompanyId,
		AppId:     cfg.AppId,
		// AppSecret: cfg.AppSecret, // 不返回密钥
		ToParty:  strings.Split(cfg.ToPartyShadow, ","),
		Template: cfg.Template,
	}

	return data, nil
}

// CreateOrUpdateWechatConfig 创建或更新企业微信配置
func (r *remoteConfigService) CreateOrUpdateWechatConfig(ctx context.Context, req *dto.UpdateWechatConfigRequest) error {
	// 先从数据库中获取配置
	cfg, err := r.store.RemoteConfig().GetWechatConfig(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// 判断数据库中是否存在配置，如果不存在则直接创建
	if cfg == nil {
		cfg = &model.WechatConfig{
			CompanyId:     req.CompanyId,
			AppId:         req.AppId,
			AppSecret:     req.AppSecret,
			ToPartyShadow: strings.Join(req.ToParty, ","),
			Template:      req.Template,
		}

		if err := r.store.RemoteConfig().CreateWechatConfig(ctx, cfg); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}

	// 更新配置到数据库
	cfg.CompanyId = req.CompanyId
	cfg.AppId = req.AppId
	cfg.AppSecret = req.AppSecret //
	cfg.ToPartyShadow = strings.Join(req.ToParty, ",")
	cfg.Template = req.Template

	if err := r.store.RemoteConfig().UpdateWechatConfig(ctx, cfg); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
