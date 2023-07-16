package doc

import "ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"

// GetSMSConfig 获取短信配置
// swagger:route GET /alarm/v2/remoteconfig/sms remote_config getSMSConfigRequest
//
// 获取短信配置
//
// 获取短信配置
//
//      Responses:
//        default: errResponse
//        200: getSMSConfigResponse

// swagger:parameters getSMSConfigRequest
type getSMSConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
}

// swagger:response getSMSConfigResponse
type getSMSConfigResponseWrapper struct {
	// in: body
	Body dto.GetSMSConfigResponse
}

// UpdateSMSConfig 更新短信配置
// swagger:route POST /alarm/v2/remoteconfig/sms remote_config updateSMSConfigRequest
//
// 更新短信配置
//
// 更新短信配置
//
//      Responses:
//        default: errResponse
//        200: updateSMSConfigResponse

// swagger:parameters updateSMSConfigRequest
type updateSMSConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.UpdateSMSConfigRequest
}

// swagger:response updateSMSConfigResponse
type updateSMSConfigResponseWrapper struct {
	// in: body
	Body dto.UpdateSMSConfigResponse
}

// GetMailboxConfig 获取邮箱配置
// swagger:route GET /alarm/v2/remoteconfig/mailbox remote_config getMailboxConfigRequest
//
// 获取邮箱配置
//
// 获取邮箱配置
//
//      Responses:
//        default: errResponse
//        200: getMailboxConfigResponse

// swagger:parameters getMailboxConfigRequest
type getMailboxConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
}

// swagger:response getMailboxConfigResponse
type getMailboxConfigResponseWrapper struct {
	// in: body
	Body dto.GetMailboxConfigResponse
}

// UpdateMailboxConfig 更新邮箱配置
// swagger:route POST /alarm/v2/remoteconfig/mailbox remote_config updateMailboxConfigRequest
//
// 更新邮箱配置
//
// 更新邮箱配置
//
//      Responses:
//        default: errResponse
//        200: updateMailboxConfigResponse

// swagger:parameters updateMailboxConfigRequest
type updateMailboxConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.UpdateMailboxConfigRequest
}

// swagger:response updateMailboxConfigResponse
type updateMailboxConfigResponseWrapper struct {
	// in: body
	Body dto.UpdateMailboxConfigResponse
}

// GetWechatConfig 获取企业微信配置
// swagger:route GET /alarm/v2/remoteconfig/wechat remote_config getWechatConfigRequest
//
// 获取企业微信配置
//
// 获取企业微信配置
//
//      Responses:
//        default: errResponse
//        200: getWechatConfigResponse

// swagger:parameters getWechatConfigRequest
type getWechatConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
}

// swagger:response getWechatConfigResponse
type getWechatConfigResponseWrapper struct {
	// in: body
	Body dto.GetWechatConfigResponse
}

// UpdateWechatConfig 更新企业微信配置
// swagger:route POST /alarm/v2/remoteconfig/wechat remote_config updateWechatConfigRequest
//
// 更新企业微信配置
//
// 更新企业微信配置
//
//      Responses:
//        default: errResponse
//        200: updateWechatConfigResponse

// swagger:parameters updateWechatConfigRequest
type updateWechatConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.UpdateWechatConfigRequest
}

// swagger:response updateWechatConfigResponse
type updateWechatConfigResponseWrapper struct {
	// in: body
	Body dto.UpdateWechatConfigResponse
}

// GetFeishuConfig 获取飞书配置
// swagger:route GET /alarm/v2/remoteconfig/feishu remote_config getFeishuConfigRequest
//
// 获取飞书配置
//
// 获取飞书配置
//
//      Responses:
//        default: errResponse
//        200: getFeishuConfigResponse

// swagger:parameters getFeishuConfigRequest
type getFeishuConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
}

// swagger:response getFeishuConfigResponse
type getFeishuConfigResponseWrapper struct {
	// in: body
	Body dto.GetFeishuConfigResponse
}

// UpdateWechatConfig 更新飞书配置
// swagger:route POST /alarm/v2/remoteconfig/feishu remote_config updateFeishuConfigRequest
//
// 更新飞书配置
//
// 更新飞书配置
//
//      Responses:
//        default: errResponse
//        200: updateFeishuConfigResponse

// swagger:parameters updateFeishuConfigRequest
type updateFeishuConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.UpdateFeishuConfigRequest
}

// swagger:response updateFeishuConfigResponse
type updateFeishuConfigResponseWrapper struct {
	// in: body
	Body dto.UpdateFeishuConfigResponse
}

// ListDingtalkConfig 获取钉钉配置列表
// swagger:route POST /alarm/v2/remoteconfig/dingtalk/list remote_config listDingtalkConfigRequest
//
// 获取钉钉配置列表
//
// 获取钉钉配置列表
//
//      Responses:
//        default: errResponse
//        200: listDingtalkConfigResponse

// swagger:parameters listDingtalkConfigRequest
type listDingtalkConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in:body
	Body dto.ListDingtalkConfigRequest
}

// swagger:response listDingtalkConfigResponse
type listDingtalkConfigResponseWrapper struct {
	// in: body
	Body dto.ListDingTalkConfigResponse
}

// CreateDingtalkConfig 创建钉钉配置
// swagger:route POST /alarm/v2/remoteconfig/dingtalk/create remote_config createDingtalkConfigRequest
//
// 创建钉钉配置
//
// 创建钉钉配置
//
//      Responses:
//        default: errResponse
//        200: createDingtalkConfigResponse

// swagger:parameters createDingtalkConfigRequest
type createDingtalkConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in:body
	Body dto.CreateDingtalkConfigRequest
}

// swagger:response createDingtalkConfigResponse
type createDingtalkConfigResponseWrapper struct {
	// in: body
	Body dto.CreateDingTalkConfigResponse
}

// DeleteDingtalkConfig 删除钉钉配置
// swagger:route POST /alarm/v2/remoteconfig/dingtalk/delete remote_config deleteDingtalkConfigRequest
//
// 删除钉钉配置
//
// 删除钉钉配置
//
//      Responses:
//        default: errResponse
//        200: deleteDingtalkConfigResponse

// swagger:parameters deleteDingtalkConfigRequest
type deleteDingtalkConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in:body
	Body dto.DeleteDingtalkConfigRequest
}

// swagger:response deleteDingtalkConfigResponse
type deleteDingtalkConfigResponseWrapper struct {
	// in: body
	Body dto.DeleteDingTalkConfigResponse
}

// UpdateDingtalkConfig 更新钉钉配置
// swagger:route POST /alarm/v2/remoteconfig/dingtalk/update remote_config updateDingtalkConfigRequest
//
// 更新钉钉配置
//
// 更新钉钉配置
//
//      Responses:
//        default: errResponse
//        200: updateDingtalkConfigResponse

// swagger:parameters updateDingtalkConfigRequest
type updateDingtalkConfigRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in:body
	Body dto.UpdateDingtalkConfigRequest
}

// swagger:response updateDingtalkConfigResponse
type updateDingtalkConfigResponseWrapper struct {
	// in: body
	Body dto.UpdateDingTalkConfigResponse
}
