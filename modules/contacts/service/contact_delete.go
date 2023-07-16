package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/dto"
)

// Delete 删除联系人组
func (s *contactSvc) Delete(ctx context.Context, req *dto.DeleteContactRequest) error {
	if err := s.store.Contacts().Delete(ctx, req.Id); err != nil {
		return err
	}

	return nil
}

// RemoveUserFromContact 从用户组中删除某个用户
func (s *contactSvc) RemoveUserFromContact(ctx context.Context, req *dto.RemoveUserFromContactRequest) error {
	return s.store.Contacts().RemoveUserFromContacts(ctx, req)
}
