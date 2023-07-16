package dto

import (
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
)

// CreatePolicyRequest 创建告警策略的请求数据结构
type CreatePolicyRequest struct {
	// 告警策略名称
	Name string `json:"Name"`
	// 告警策略的创建人
	Creator string `json:"Creator"`
	// 告警策略的最后修改人
	Updater string `json:"Updater"`
	// 业务线名称
	Production string `json:"Production"`
	// 策略描述
	Comment string `json:"Comment"`
	// 策略标签
	Labels map[string]string `json:"Labels"`
	// 每条告警规则的最大告警次数
	Limit int `json:"Limit"`
	// Example: "form"
	Type model.AlarmPolicyType `json:"Type"`
	// 表单格式的告警策略, Type为form时传此字段
	FormPolicy *model.FormPolicy `json:"FormPolicy,omitempty"`
	// PromQL告警策略, Type为promql时传此字段
	PromqlPolicy *model.PromqlPolicy `json:"PromqlPolicy,omitempty"`
	// 告警接收人
	Receivers *model.Receivers `json:"Receivers,omitempty"`
	// 通知设置
	NotifySetup *model.NotifySetup `json:"NotifySetup"`
}

// CreatePolicyResponse 创建告警策略的响应数据结构
type CreatePolicyResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// ListPolicyRequest 获取告警策略列表的请求数据结构
type ListPolicyRequest struct {
	PageOption
	// 告警策略类型, form-表单类型, promql-PromQL类型
	Type model.AlarmPolicyType `json:"Type"`
	// 资源类型
	ResourceType []string `json:"ResourceType"`
	// 资源子类型
	ResourceSubType []string `json:"ResourceSubType"`
	// 模糊搜索的key，允许值: name-策略名称 creator-创建者 production-业务线 resource-关联的实例
	SearchKey   string `json:"SearchKey"`
	SearchValue string `json:"SearchValue"`
}

// ListPolicyResponse 获取告警策略列表的响应数据
type ListPolicyResponse struct {
	SuccessResponse
	Data *ListPolicyResponseData `json:"Data"`
}

type ListPolicyResponseData struct {
	Total int64         `json:"Total"`
	Items []*PolicyInfo `json:"Items"`
}

type PolicyInfo struct {
	Id int `json:"Id"`
	// 实例id
	InstanceId string `json:"InstanceId"`
	// 告警策略名称
	Name string `json:"Name"`
	// 策略描述
	Comment string `json:"Comment"`
	// 告警策略的创建人
	Creator string `json:"Creator"`
	// 告警策略的最后修改人
	Updater string `json:"Updater"`
	// 是否启用告警策略，true-启用，false-不启用
	Enabled bool `json:"Enabled"`
	// 业务线名称
	Production string `json:"Production"`
	// 策略标签
	Labels map[string]string `json:"Labels"`
	// 每条告警规则的最大告警次数
	Limit int `json:"Limit"`
	// 告警策略类型, form-表单告警策略, promql-PromQL告警策略
	Type model.AlarmPolicyType `json:"Type"`
	// 表单格式的告警策略, Type为form时传此字段
	FormPolicy *model.FormPolicy `json:"FormPolicy,omitempty"`
	// PromQL告警策略, Type为promql时传此字段
	PromqlPolicy *model.PromqlPolicy `json:"PromqlPolicy,omitempty"`
	// 告警接收人
	Receivers *model.Receivers `json:"Receivers"`
	// 通知设置
	NotifySetup *model.NotifySetup `json:"NotifySetup"`
	CreatedAt   int64              `json:"CreatedAt"`
	UpdatedAt   int64              `json:"UpdatedAt"`
}

// GetPolicyResponse 获取告警策略详情的响应数据结构
type GetPolicyResponse struct {
	SuccessResponse
	Data GetPolicyResponseData `json:"Data"`
}

type GetPolicyResponseData struct {
	*PolicyInfo
}

// DeletePolicyRequest 删除告警策略的请求数据结构
type DeletePolicyRequest struct {
	InstanceIds []string `json:"InstanceIds"`
}

// DeletePolicyResponse 删除告警策略的响应数据结构
type DeletePolicyResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// UpdatePolicyRequest 更新告警策略的请求数据结构
type UpdatePolicyRequest struct {
	Id int `json:"Id"`
	// 实例id
	InstanceId string `json:"InstanceId"`
	// 告警策略名称
	Name string `json:"Name"`
	// 告警策略的创建人
	Creator string `json:"Creator"`
	// 告警策略的最后修改人
	Updater string `json:"Updater"`
	// 业务线名称
	Production string `json:"Production"`
	// 策略描述
	Comment string `json:"Comment"`
	// 策略标签
	Labels map[string]string `json:"Labels"`
	// 每条告警规则的最大告警次数
	Limit int `json:"Limit"`
	// 告警策略类型, form-表单告警策略, promql-PromQL告警策略
	Type model.AlarmPolicyType `json:"Type"`
	// 表单格式的告警策略, Type为form时传此字段
	FormPolicy *model.FormPolicy `json:"FormPolicy"`
	// PromQL告警策略, Type为promql时传此字段
	PromqlPolicy *model.PromqlPolicy `json:"PromqlPolicy"`
	// 告警接收人
	Receivers *model.Receivers `json:"Receivers"`
	// 通知设置
	NotifySetup *model.NotifySetup `json:"NotifySetup"`
	CreatedAt   int64              `json:"CreatedAt"`
	UpdatedAt   int64              `json:"UpdatedAt"`
}

// UpdatePolicyResponse 更新告警策略的响应数据结构
type UpdatePolicyResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// SwitchPolicyRequest 启动或停止告警策略的请求数据结构
type SwitchPolicyRequest struct {
	// 告警策略的实例id
	InstanceId string `json:"InstanceId"`
	// 启动或停止告警策略, true-启动 false-停止
	Enable bool `json:"Enable"`
}

// SwitchPolicyResponse 启动或停止告警策略的响应数据结构
type SwitchPolicyResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}
