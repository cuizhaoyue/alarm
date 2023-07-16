package controller

import (
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/service"
)

type AlarmController struct {
	svc service.Service
}

// func NewAlarmController(store dao.Factory) *AlarmController {
func NewAlarmController() *AlarmController {
	return &AlarmController{
		svc: service.NewService(),
	}
}
