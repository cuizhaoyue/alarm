package doc

import "ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/contacts/dto"

// Create 创建联系人组
// swagger:route POST /contact/create contact createContactRequest
//
// 创建联系人组
//
// 创建联系人组
//
//      Responses:
//        default: errResponse
//        200: createContactResponse

// 创建创建告警策略的请求数据
// swagger:parameters createContactRequest
type createContactRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.CreateContactRequest
}

// swagger:response createContactResponse
type createContactResponseWrapper struct {
	// in: body
	Body dto.CreateContactResponse
}

// Update 编辑联系人组
// swagger:route POST /contact/update contact updateContactRequest
//
// 编辑联系人组
//
// 编辑联系人组，同时更新联系人组中的联系人信息
//
//      Responses:
//        default: errResponse
//        200: updateContactResponse

// 创建创建告警策略的请求数据
// swagger:parameters updateContactRequest
type updateContactRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.UpdateContactRequest
}

// swagger:response updateContactResponse
type updateContactResponseWrapper struct {
	// in: body
	Body dto.UpdateContactResponse
}

// List 获取联系人组列表
// swagger:route POST /contact/list contact listContactRequest
//
// 获取联系人组列表
//
// 获取联系人组列表
//
//      Responses:
//        default: errResponse
//        200: listContactResponse

// 创建创建告警策略的请求数据
// swagger:parameters listContactRequest
type listContactRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.ListContactRequest
}

// swagger:response listContactResponse
type listContactResponseWrapper struct {
	// in: body
	Body dto.ListContactResponse
}

// Delete 删除联系人组
// swagger:route POST /contact/delete contact deleteContactRequest
//
// 删除联系人组
//
// 删除联系人组
//
//      Responses:
//        default: errResponse
//        200: deleteContactResponse

// 创建创建告警策略的请求数据
// swagger:parameters deleteContactRequest
type deleteContactRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.DeleteContactRequest
}

// swagger:response deleteContactResponse
type deleteContactResponseWrapper struct {
	// in: body
	Body dto.DeleteContactResponse
}

// RemoveUserFromContact 从联系人组中移除用户
// swagger:route POST /contact/remove contact removeUserFromContactRequest
//
// 从联系人组中移除用户
//
// 从联系人组中移除用户
//
//      Responses:
//        default: errResponse
//        200: removeUserFromContactResponse

// 创建创建告警策略的请求数据
// swagger:parameters removeUserFromContactRequest
type removeUserFromContactRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.RemoveUserFromContactRequest
}

// swagger:response removeUserFromContactResponse
type removeUserFromContactResponseWrapper struct {
	// in: body
	Body dto.RemoveUserFromContactResponse
}
