package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
)

// AlertService 告警消息服务接口
type AlertService interface {
	CreateOrUpdate(ctx context.Context, req *model.NotificationData) error
	List(ctx context.Context, req *dto.ListAlertsRequest) (*dto.ListAlertsResponseData, error)
	Overview(ctx context.Context, req *dto.OverviewRequest) (*dto.OverviewResponseData, error)
}

var _ AlertService = &alertService{}

type alertService struct {
	store dao.Factory
}

func newAlertService(svc *service) *alertService {
	return &alertService{
		store: svc.store,
	}
}
