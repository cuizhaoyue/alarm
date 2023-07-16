package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/model"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"gorm.io/gorm"
)

// ContactStore 联系人组的存储接口
type ContactStore interface {
	Create(ctx context.Context, contact *model.Contacts, users []*dto.BaseUserInfo) error
	Update(ctx context.Context, contacts *model.Contacts, addUsers, subUsers []*dto.BaseUserInfo) error
	Delete(ctx context.Context, id int) error
	ListContacts(ctx context.Context, req *dto.ListContactRequest) (*model.ContactList, error)
	RemoveUserFromContacts(ctx context.Context, req *dto.RemoveUserFromContactRequest) error
	ListRelatedDataByContactsId(ctx context.Context, id int) ([]*model.ContactUserRelation, error)
}

var _ ContactStore = &contacts{}

type contacts struct {
	db *gorm.DB
}

func newContactStore(ds *datastore) *contacts {
	return &contacts{ds.db}
}

// Create 创建联系人组
func (c *contacts) Create(ctx context.Context, contact *model.Contacts, users []*dto.BaseUserInfo) error {
	// 在事务中执行创建操作
	err := c.db.Transaction(func(tx *gorm.DB) error {
		// 添加联系人组
		if err := tx.Create(contact).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
					"message": fmt.Errorf("the name of contact is existed, %s", err).Error(),
				})
			}

			return err
		}

		// 构建关联数据
		var rlts []*model.ContactUserRelation
		for _, u := range users {
			rlts = append(rlts, &model.ContactUserRelation{
				ContactId: contact.Id,
				UserId:    u.UserId,
				Username:  u.Username,
				Telephone: u.Telephone,
				Email:     u.Email,
				CreatedAt: time.Now(),
			})
		}

		// 关联表中添加对应数据
		if err := tx.Create(rlts).Error; err != nil {
			lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
				"message": fmt.Errorf("create relation record for contact and user failed, %s", err).Error(),
			})

			return err
		}

		return nil
	})

	return err
}

// Update 更新联系人组及其对应的用户
func (c *contacts) Update(ctx context.Context, contacts *model.Contacts, addUsers, subUsers []*dto.BaseUserInfo) error {
	// 在事务中对联系人组表和关联表同时更新
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(contacts).Error; err != nil {
			lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
				"message": fmt.Errorf("update contacts failed, %s", err).Error(),
			})

			return err
		}

		// 添加增加后的关联数据
		if len(addUsers) > 0 {
			// 构建关联数据
			var rlts []*model.ContactUserRelation
			for _, u := range addUsers {
				rlts = append(rlts, &model.ContactUserRelation{
					ContactId: contacts.Id,
					UserId:    u.UserId,
					Username:  u.Username,
					Telephone: u.Telephone,
					Email:     u.Email,
					CreatedAt: time.Now(),
				})
			}

			// 关联表中添加对应数据
			if err := tx.Create(rlts).Error; err != nil {
				lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
					"message": fmt.Errorf("add relation record for contact and user failed, %s", err).Error(),
				})

				return err
			}
		}

		// 删除需要去除的关联数据
		if len(subUsers) > 0 {
			var uids []string
			for _, u := range subUsers {
				uids = append(uids, u.UserId)
			}
			if err := tx.Unscoped().
				Where("user_id in ?", uids).
				Delete(&model.ContactUserRelation{}).Error; err != nil {

				lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
					"message": fmt.Errorf("remove relation record for contact and user failed, %s", err).Error(),
				})

				return err
			}
		}

		return nil
	})
}

// Delete 删除联系人组
func (c *contacts) Delete(ctx context.Context, id int) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		// 删除联系人组Contacts表中的数据
		if err := tx.Unscoped().Where("id = ?", id).Delete(&model.Contacts{}).Error; err != nil {
			lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
				"message": fmt.Errorf("delete contacts failed, %s", err).Error(),
			})

			return err
		}

		// 删除关联表中的数据
		if err := tx.Unscoped().Where("contact_id = ?", id).Delete(&model.ContactUserRelation{}).Error; err != nil {
			lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
				"message": fmt.Errorf("delete associated data for contacts and users failed, %s", err).Error(),
			})

			return err
		}

		return nil
	})
}

// RemoveUserFromContacts 删除用户组中的某个用户
func (c *contacts) RemoveUserFromContacts(ctx context.Context, req *dto.RemoveUserFromContactRequest) error {
	if err := c.db.Unscoped().
		Where("contact_id = ? and user_id = ?", req.Id, req.UserId).
		Delete(&model.ContactUserRelation{}).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("remove user from contacts failed, %s", err).Error(),
		})

		return err
	}

	return nil
}

// ListContacts 获取联系人组列表
func (c *contacts) ListContacts(ctx context.Context, req *dto.ListContactRequest) (*model.ContactList, error) {
	// 模糊搜索
	if req.SearchValue != "" {
		switch req.SearchKey {
		case "name":
			c.db = c.db.Where("name like ?", "%"+req.SearchValue+"%")
		}
	}

	// 获取联系人组数据
	ret := &model.ContactList{}
	if err := c.db.Offset(req.Offset).
		Limit(req.PageSize).
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.Total).Error; err != nil {

		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("list contacts failed, %s", err).Error(),
		})

		return nil, err
	}

	return ret, nil
}

// ListRelatedDataByContactsId 通过联系人组的id查找联系人组和用户的关联数据
func (c *contacts) ListRelatedDataByContactsId(ctx context.Context, id int) ([]*model.ContactUserRelation, error) {
	var ret []*model.ContactUserRelation
	if err := c.db.Where("contact_id = ?", id).
		Find(&ret).Error; err != nil {
		lib.Log.TagError(lib.GetTraceContext(ctx), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("get relation data by contacts id failed, %s", err).Error(),
		})
		return nil, err
	}

	return ret, nil
}
