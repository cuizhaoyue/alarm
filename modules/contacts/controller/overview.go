package controller

import (
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/service"
)

type ContactsController struct {
	svc service.Service
}

func NewContactsController() *ContactsController {
	return &ContactsController{
		svc: service.NewService(),
	}
}
