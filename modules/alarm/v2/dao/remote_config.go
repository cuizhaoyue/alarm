package dao

import (
	"context"
	"errors"
	"fmt"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"gorm.io/gorm"
)

type RemoteConfig interface {
	// 短信配置操作
	GetSMSConfig(ctx context.Context) (*model.SMSConfig, error)
	UpdateSMSConfig(ctx context.Context, cfg *model.SMSConfig) error
	CreateSMSConfig(ctx context.Context, cfg *model.SMSConfig) error

	// 邮箱配置操作
	GetMailboxConfig(ctx context.Context) (*model.MailboxConfig, error)
	UpdateMailboxConfig(ctx context.Context, cfg *model.MailboxConfig) error
	CreateMailboxConfig(ctx context.Context, cfg *model.MailboxConfig) error

	// 企业微信配置操作
	GetWechatConfig(ctx context.Context) (*model.WechatConfig, error)
	UpdateWechatConfig(ctx context.Context, cfg *model.WechatConfig) error
	CreateWechatConfig(ctx context.Context, cfg *model.WechatConfig) error

	// 飞书配置操作
	GetFeishuConfig(ctx context.Context) (*model.FeishuConfig, error)
	UpdateFeishuConfig(ctx context.Context, cfg *model.FeishuConfig) error
	CreateFeishuConfig(ctx context.Context, cfg *model.FeishuConfig) error

	// 钉钉配置操作
	ListDingtalkConfig(ctx context.Context, req *dto.ListDingtalkConfigRequest) (*model.DingtalkConfigList, error)
	UpdateDingtalkConfig(ctx context.Context, cfg *model.DingtalkConfig) error
	CreateDingtalkConfig(ctx context.Context, cfg *model.DingtalkConfig) error
	DeleteDingtalkConfig(ctx context.Context, id int) error
}

var _ RemoteConfig = &remoteConfig{}

type remoteConfig struct {
	db *gorm.DB
}

func newRemoteConfig(ds *datastore) *remoteConfig {
	return &remoteConfig{ds.db}
}

// GetSMSConfig 获取短信配置
func (r *remoteConfig) GetSMSConfig(ctx context.Context) (*model.SMSConfig, error) {
	ret := &model.SMSConfig{}
	if err := r.db.First(ret).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
				"message": fmt.Errorf("get sms config failed, %s", err).Error(),
			})

			return nil, err
		} else {
			return nil, nil
		}
	}

	return ret, nil
}

// UpdateSMSConfig 更新短信配置
func (r *remoteConfig) UpdateSMSConfig(ctx context.Context, cfg *model.SMSConfig) error {
	if err := r.db.Save(cfg).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("update sms config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// CreateSMSConfig 创建短信配置
func (r *remoteConfig) CreateSMSConfig(ctx context.Context, cfg *model.SMSConfig) error {
	if err := r.db.Create(cfg).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("create sms config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// GetMailboxConfig 获取邮箱配置
func (r *remoteConfig) GetMailboxConfig(ctx context.Context) (*model.MailboxConfig, error) {
	ret := &model.MailboxConfig{}
	if err := r.db.First(ret).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
				"message": fmt.Errorf("get mailbox config failed, %s", err).Error(),
			})

			return nil, err
		} else {
			return nil, nil
		}
	}

	return ret, nil
}

// UpdateMailboxConfig 更新邮箱配置
func (r *remoteConfig) UpdateMailboxConfig(ctx context.Context, cfg *model.MailboxConfig) error {
	if err := r.db.Save(cfg).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("update mailbox config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// CreateMailboxConfig 创建邮箱配置
func (r *remoteConfig) CreateMailboxConfig(ctx context.Context, cfg *model.MailboxConfig) error {
	if err := r.db.Create(cfg).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("create mailbox config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// GetWechatConfig 获取企业微信配置
func (r *remoteConfig) GetWechatConfig(ctx context.Context) (*model.WechatConfig, error) {
	ret := &model.WechatConfig{}
	if err := r.db.First(ret).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
				"message": fmt.Errorf("get wechat config failed, %s", err).Error(),
			})

			return nil, err
		} else {
			return nil, nil
		}
	}

	return ret, nil
}

// UpdateWechatConfig 更新企业微信配置
func (r *remoteConfig) UpdateWechatConfig(ctx context.Context, cfg *model.WechatConfig) error {
	if err := r.db.Save(cfg).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("update Wechat config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// CreateWechatConfig 创建企业微信配置
func (r *remoteConfig) CreateWechatConfig(ctx context.Context, cfg *model.WechatConfig) error {
	if err := r.db.Create(cfg).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("create Wechat config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// GetFeishuConfig 获取飞书配置
func (r *remoteConfig) GetFeishuConfig(ctx context.Context) (*model.FeishuConfig, error) {
	ret := &model.FeishuConfig{}
	if err := r.db.First(ret).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
				"message": fmt.Errorf("get feishu config failed, %s", err).Error(),
			})

			return nil, err
		} else {
			return nil, nil
		}
	}

	return ret, nil
}

// UpdateFeishuConfig 更新飞书配置
func (r *remoteConfig) UpdateFeishuConfig(ctx context.Context, cfg *model.FeishuConfig) error {
	if err := r.db.Save(cfg).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("update feishu config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// CreateFeishuConfig 创建飞书配置
func (r *remoteConfig) CreateFeishuConfig(ctx context.Context, cfg *model.FeishuConfig) error {
	if err := r.db.Create(cfg).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("create feishu config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// ListDingtalkConfig 获取钉钉配置
func (r *remoteConfig) ListDingtalkConfig(ctx context.Context, req *dto.ListDingtalkConfigRequest) (*model.DingtalkConfigList, error) {
	ret := &model.DingtalkConfigList{}

	err := r.db.
		Offset(req.Offset).
		Limit(req.PageSize).
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.Total).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
				"message": fmt.Errorf("get Dingtalk config failed, %s", err).Error(),
			})

			return nil, err
		} else {
			return nil, nil
		}
	}

	return ret, nil
}

// UpdateDingtalkConfig 更新钉钉配置
func (r *remoteConfig) UpdateDingtalkConfig(ctx context.Context, cfg *model.DingtalkConfig) error {
	if err := r.db.Save(cfg).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("update Dingtalk config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// CreateDingtalkConfig 创建钉钉配置
func (r *remoteConfig) CreateDingtalkConfig(ctx context.Context, cfg *model.DingtalkConfig) error {
	if err := r.db.Create(cfg).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("create Dingtalk config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// DeleteDingtalkConfig 删除钉钉配置
func (r *remoteConfig) DeleteDingtalkConfig(ctx context.Context, id int) error {
	if err := r.db.Where("id = ?", id).Delete(&model.DingtalkConfig{}).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("delete Dingtalk config failed, %s", err).Error(),
		})

		return err
	}

	return nil
}
