package service

import (
	"context"
	"time"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/model"
	"github.com/pkg/errors"
)

// Create 创建联系人组
func (s *contactSvc) Create(ctx context.Context, req *dto.CreateContactRequest) error {
	// 构建联系人组信息
	contacts := &model.Contacts{
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 创建联系人组
	err := s.store.Contacts().Create(ctx, contacts, req.Users)
	if err != nil {
		return errors.Wrap(err, "create contacts error")
	}

	return nil
}
