package doc

import "ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"

// 发生错误时的返回信息
// swagger:response errResponse
type errResponseWrapper struct {
	// in:body
	Body dto.ErrResponse
}

// Create 创建告警策略
// swagger:route POST /alarm/v2/policy/create policy createPolicyRequest
//
// 创建告警策略
//
// 创建告警策略
//
//      Responses:
//        default: errResponse
//        200: createPolicyResponse

// 创建创建告警策略的请求数据
// swagger:parameters createPolicyRequest
type createPolicyRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.CreatePolicyRequest
}

// swagger:response createPolicyResponse
type createPolicyResponseWrapper struct {
	// in: body
	Body dto.CreatePolicyResponse
}

// List 获取告警策略列表
// swagger:route POST /alarm/v2/policy/list policy listPolicyRequest
//
// 获取告警策略列表
//
// 获取告警策略列表
//
//      Responses:
//        default: errResponse
//        200: listPolicyResponse

// swagger:parameters listPolicyRequest
type listPolicyRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.ListPolicyRequest
}

// swagger:response listPolicyResponse
type listPolicyResponseWrapper struct {
	// in: body
	Body dto.ListPolicyResponse
}

// Get 获取告警策略详情
// swagger:route GET /alarm/v2/policy/describe policy getPolicyRequest
//
// 获取告警策略详情
//
// 获取告警策略详情
//
//      Responses:
//        default: errResponse
//        200: getPolicyResponse

// swagger:parameters getPolicyRequest
type getPolicyRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: query
	InstanceId int `json:"InstanceId" form:"InstanceId"`
}

// swagger:response getPolicyResponse
type getPolicyResponseWrapper struct {
	// in: body
	Body dto.GetPolicyResponse
}

// DeleteCollection 删除告警策略
// swagger:route POST /alarm/v2/policy/delete policy deletePolicyRequest
//
// 删除告警策略
//
// 删除告警策略
//
//      Responses:
//        default: errResponse
//        200: deletePolicyResponse

// swagger:parameters deletePolicyRequest
type deletePolicyRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.DeletePolicyRequest
}

// swagger:response deletePolicyResponse
type deletePolicyResponseWrapper struct {
	// in: body
	Body dto.DeletePolicyResponse
}

// Update 编辑更新告警策略
// swagger:route POST /alarm/v2/policy/update policy updatePolicyRequest
//
// 编辑更新告警策略
//
// 编辑更新告警策略
//
//      Responses:
//        default: errResponse
//        200: updatePolicyResponse

// swagger:parameters updatePolicyRequest
type updatePolicyRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.UpdatePolicyRequest
}

// swagger:response updatePolicyResponse
type updatePolicyResponseWrapper struct {
	// in: body
	Body dto.UpdatePolicyResponse
}

// Switch 启动或停止告警策略
// swagger:route POST /alarm/v2/policy/switch policy switchPolicyRequest
//
// 启动或停止告警策略
//
// 启动或停止告警策略
//
//      Responses:
//        default: errResponse
//        200: switchPolicyResponse

// swagger:parameters switchPolicyRequest
type switchPolicyRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.SwitchPolicyRequest
}

// swagger:response switchPolicyResponse
type switchPolicyResponseWrapper struct {
	// in: body
	Body dto.SwitchPolicyResponse
}
