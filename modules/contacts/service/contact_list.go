package service

import (
	"context"
	"sync"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/model"
)

// List 获取联系人组列表
func (s *contactSvc) List(ctx context.Context, req *dto.ListContactRequest) (*dto.ListContactResponseData, error) {
	// 设置页码偏移量
	req.SetupOffset()

	// 先从联系人组表contacts中获取到联系人的数据
	contactList, err := s.store.Contacts().ListContacts(ctx, req)
	if err != nil {
		return nil, err
	}

	// 从关联表中获取联系人下的用户
	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)     // 记录错误
	finished := make(chan struct{}, 1) // 完成标志

	var m sync.Map

	for _, item := range contactList.Items {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			rlts, err := s.store.Contacts().ListRelatedDataByContactsId(ctx, id)
			if err != nil {
				errChan <- err
			}

			m.Store(id, rlts)
		}(item.Id)
	}

	// 等待任务执行完成
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case err := <-errChan:
		return nil, err
	case <-finished:
	}

	// 根据关联数据构建返回数据
	ret := &dto.ListContactResponseData{}
	ret.Total = contactList.Total
	var contactsInfos []*dto.ContactsInfo
	for _, item := range contactList.Items {
		// 构建用户列表
		var users []*dto.User
		rlts, _ := m.Load(item.Id)
		for _, rlt := range rlts.([]*model.ContactUserRelation) {
			users = append(users, &dto.User{
				BaseUserInfo: dto.BaseUserInfo{
					UserId:    rlt.UserId,
					Username:  rlt.Username,
					Email:     rlt.Email,
					Telephone: rlt.Telephone,
				},
				CreatedAt: rlt.CreatedAt.UnixMilli(),
			})
		}
		// 构建联系人组列表
		contactsInfos = append(contactsInfos, &dto.ContactsInfo{
			Id:        item.Id,
			Name:      item.Name,
			CreatedAt: item.CreatedAt.UnixMilli(),
			UpdatedAt: item.UpdatedAt.UnixMilli(),
			Users:     users,
		})
	}

	ret.Items = contactsInfos

	return ret, nil
}
