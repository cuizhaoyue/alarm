package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"
)

// GetMailboxConfig 获取邮箱配置
func (r *remoteConfigService) GetMailboxConfig(ctx context.Context) (*dto.GetMailboxConfigResponseData, error) {
	cfg, err := r.store.RemoteConfig().GetMailboxConfig(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	data := &dto.GetMailboxConfigResponseData{
		SMTPServer:  cfg.SMTPServer,
		Port:        cfg.Port,
		DisplayName: cfg.DisplayName,
		Username:    cfg.Username,
		Template:    cfg.Template,
		SSL:         cfg.SSL,
	}

	return data, nil
}

// CreateOrUpdateMailboxConfig 创建或更新邮箱配置
func (r *remoteConfigService) CreateOrUpdateMailboxConfig(ctx context.Context, req *dto.UpdateMailboxConfigRequest) error {
	// 先从数据库中获取配置
	cfg, err := r.store.RemoteConfig().GetMailboxConfig(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// 判断数据库中是否存在配置，如果不存在则直接创建
	if cfg == nil {
		cfg = &model.MailboxConfig{
			SMTPServer:  req.SMTPServer,
			Port:        req.Port,
			DisplayName: req.DisplayName,
			Username:    req.Username,
			Password:    req.Password,
			Template:    req.Template,
			SSL:         req.SSL,
		}

		if err := r.store.RemoteConfig().CreateMailboxConfig(ctx, cfg); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}

	// 更新邮箱配置到数据库
	cfg.SMTPServer = req.SMTPServer
	cfg.Port = req.Port
	cfg.DisplayName = req.DisplayName
	cfg.Username = req.Username
	cfg.Password = req.Password
	cfg.Template = req.Template
	cfg.SSL = req.SSL

	if err := r.store.RemoteConfig().UpdateMailboxConfig(ctx, cfg); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
