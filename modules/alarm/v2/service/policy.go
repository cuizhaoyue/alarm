package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
)

type AlarmService interface {
	Create(ctx context.Context, req *dto.CreatePolicyRequest) error
	Get(ctx context.Context, insId string) (*dto.GetPolicyResponseData, error)
	List(ctx context.Context, req *dto.ListPolicyRequest) (*dto.ListPolicyResponseData, error)
	DeleteCollection(ctx context.Context, req *dto.DeletePolicyRequest) error
	Update(ctx context.Context, req *dto.UpdatePolicyRequest) error
	Switch(ctx context.Context, req *dto.SwitchPolicyRequest) error
}

var _ AlarmService = &alarmService{}

type alarmService struct {
	store dao.Factory
}

func newAlarmService(svc *service) *alarmService {
	return &alarmService{
		store: svc.store,
	}
}
