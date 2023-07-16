package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
)

type RemoteConfigService interface {
	// 短信配置操作
	GetSMSConfig(ctx context.Context) (*dto.GetSMSConfigResponseData, error)
	CreateOrUpdateSMSConfig(ctx context.Context, req *dto.UpdateSMSConfigRequest) error
	// 邮箱配置操作
	GetMailboxConfig(ctx context.Context) (*dto.GetMailboxConfigResponseData, error)
	CreateOrUpdateMailboxConfig(ctx context.Context, req *dto.UpdateMailboxConfigRequest) error
	// 企业微信配置操作
	GetWechatConfig(ctx context.Context) (*dto.GetWechatConfigResponseData, error)
	CreateOrUpdateWechatConfig(ctx context.Context, req *dto.UpdateWechatConfigRequest) error
	// 飞书配置操作
	GetFeishuConfig(ctx context.Context) (*dto.GetFeishuConfigResponseData, error)
	CreateOrUpdateFeishuConfig(ctx context.Context, req *dto.UpdateFeishuConfigRequest) error
	// 钉钉配置操作
	ListDingtalkConfig(ctx context.Context, req *dto.ListDingtalkConfigRequest) (*model.DingtalkConfigList, error)
	CreateDingtalkConfig(ctx context.Context, req *dto.CreateDingtalkConfigRequest) error
	DeleteDingtalkConfig(ctx context.Context, req *dto.DeleteDingtalkConfigRequest) error
	UpdateDingtalkConfig(ctx context.Context, req *dto.UpdateDingtalkConfigRequest) error
}

var _ RemoteConfigService = &remoteConfigService{}

type remoteConfigService struct {
	store dao.Factory
}

func newRemoteConfigService(svc *service) *remoteConfigService {
	return &remoteConfigService{
		store: svc.store,
	}
}
