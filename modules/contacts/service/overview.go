package service

import "ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/dao"

// ContactSvc 联系人组的服务接口
type Service interface {
	Contacts() ContactSvc
}

var _ Service = &service{}

type service struct {
	store dao.Factory
}

func NewService() *service {
	return &service{
		store: dao.NewDataStore(nil),
	}
}

func (s *service) Contacts() ContactSvc {
	return newContactSvc(s)
}
