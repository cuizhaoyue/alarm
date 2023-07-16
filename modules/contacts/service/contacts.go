package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/dao"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/dto"
)

type ContactSvc interface {
	Create(ctx context.Context, req *dto.CreateContactRequest) error
	Update(ctx context.Context, req *dto.UpdateContactRequest) error
	List(ctx context.Context, req *dto.ListContactRequest) (*dto.ListContactResponseData, error)
	Delete(ctx context.Context, req *dto.DeleteContactRequest) error
	RemoveUserFromContact(ctx context.Context, req *dto.RemoveUserFromContactRequest) error
}

var _ ContactSvc = &contactSvc{}

type contactSvc struct {
	store dao.Factory
}

func newContactSvc(svc *service) *contactSvc {
	return &contactSvc{store: svc.store}
}
