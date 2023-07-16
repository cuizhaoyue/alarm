package service

import (
	"context"
	"time"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/model"
)

func (s *contactSvc) Update(ctx context.Context, req *dto.UpdateContactRequest) error {
	// 构建联系人组对象
	contacts := &model.Contacts{
		Id:        req.Id,
		Name:      req.Name,
		UpdatedAt: time.Now(),
	}

	// 根据联系人组的id获取原来的用户关联数据
	oriRlts, err := s.store.Contacts().ListRelatedDataByContactsId(ctx, req.Id)
	if err != nil {
		return err
	}
	var oriUsers []*dto.BaseUserInfo
	for _, rlt := range oriRlts {
		oriUsers = append(oriUsers, &dto.BaseUserInfo{
			UserId:    rlt.UserId,
			Username:  rlt.Username,
			Email:     rlt.Email,
			Telephone: rlt.Telephone,
		})
	}

	addUsers := difference(req.Users, oriUsers)
	subUsers := difference(oriUsers, req.Users)

	err = s.store.Contacts().Update(ctx, contacts, addUsers, subUsers)

	return err
}

// 获取公共元素外的其它元素
func difference(users1, users2 []*dto.BaseUserInfo) []*dto.BaseUserInfo {
	m := make(map[string]*dto.BaseUserInfo)

	// 将users1的元素添加到map中
	for _, u := range users1 {
		m[u.UserId] = u
	}

	// 删除两个列表的相同元素
	for _, u := range users2 {
		delete(m, u.UserId)
	}

	var result []*dto.BaseUserInfo
	for _, user := range m {
		result = append(result, user)
	}

	return result
}
