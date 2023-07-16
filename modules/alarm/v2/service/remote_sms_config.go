package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"
)

// GetSMSConfig 获取短信配置
func (r *remoteConfigService) GetSMSConfig(ctx context.Context) (*dto.GetSMSConfigResponseData, error) {
	cfg, err := r.store.RemoteConfig().GetSMSConfig(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 不返回AK SK
	data := &dto.GetSMSConfigResponseData{
		// Protocol: cfg.Protocol,
		// IP:       cfg.IP,
		URL: cfg.URL,
		// User:     cfg.User,
		Sign: cfg.Sign,
		// AK:       cfg.AK,
		// SK:       cfg.SK,
		Template: cfg.Template,
	}

	return data, nil
}

func (r *remoteConfigService) CreateOrUpdateSMSConfig(ctx context.Context, req *dto.UpdateSMSConfigRequest) error {
	cfg, err := r.store.RemoteConfig().GetSMSConfig(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// 如果配置信息不存在则先创建该配置数据
	if cfg == nil {
		cfg = &model.SMSConfig{
			// Protocol: req.Protocol,
			// IP:       req.IP,
			URL: req.URL,
			// User:     req.User,
			Sign:     req.Sign,
			AK:       req.AK,
			SK:       req.SK,
			Template: req.Template,
		}
		err := r.store.RemoteConfig().CreateSMSConfig(ctx, cfg)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	}

	// 更新配置信息
	// cfg.Protocol = req.Protocol
	// cfg.IP = req.IP
	cfg.URL = req.URL
	// cfg.User = req.User
	cfg.Sign = req.Sign
	cfg.AK = req.AK
	cfg.SK = req.SK
	cfg.Template = req.Template

	if err := r.store.RemoteConfig().UpdateSMSConfig(ctx, cfg); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
