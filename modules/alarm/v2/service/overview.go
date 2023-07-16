package service

import "ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao"

type Service interface {
	AlarmSvc() AlarmService
	PromqlTplSvc() PromqlTplSvc
	RemoteConfigSvc() RemoteConfigService
	AlertSvc() AlertService
}

var _ Service = &service{}

type service struct {
	store dao.Factory
}

// func NewService(store dao.Factory) *service {
func NewService() *service {
	// return &service{store: store}
	return &service{
		store: dao.NewDataStore(nil),
	}
}

func (svc *service) AlarmSvc() AlarmService {
	return newAlarmService(svc)
}

func (svc *service) PromqlTplSvc() PromqlTplSvc {
	return newPromqlTplSvc(svc)
}

func (svc *service) RemoteConfigSvc() RemoteConfigService {
	return newRemoteConfigService(svc)
}

func (svc *service) AlertSvc() AlertService {
	return newAlertService(svc)
}
