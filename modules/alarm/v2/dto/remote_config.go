package dto

import "ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"

// GetSMSConfigResponse 获取短信配置的响应数据结构
type GetSMSConfigResponse struct {
	SuccessResponse
	Data *GetSMSConfigResponseData `json:"Data"`
}
type GetSMSConfigResponseData struct {
	Protocol string `json:"Protocol"`
	IP       string `json:"Ip"`
	URL      string `json:"Url"`
	User     string `json:"User"`
	Sign     string `json:"Sign"`
	AK       string `json:"AK"`
	SK       string `json:"SK"`
	Template string `json:"Template"`
}

// UpdateSMSConfigRequest 更新短信远程配置的请求数据结构
type UpdateSMSConfigRequest struct {
	Protocol string `json:"Protocol"`
	IP       string `json:"Ip"`
	URL      string `json:"Url"`
	User     string `json:"User"`
	Sign     string `json:"Sign"`
	AK       string `json:"AK"`
	SK       string `json:"SK"`
	Template string `json:"Template"`
}

// UpdateSMSConfigResponse 更新短信远程配置的响应数据结构
type UpdateSMSConfigResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// GetMailboxConfigResponse 获取邮箱配置的响应数据结构
type GetMailboxConfigResponse struct {
	SuccessResponse
	Data *GetMailboxConfigResponseData `json:"Data"`
}
type GetMailboxConfigResponseData struct {
	SMTPServer  string `json:"SmtpServer"`
	Port        int    `json:"Port"`
	DisplayName string `json:"DisplayName"`
	Username    string `json:"Username"`
	// Password    string `json:"Password"` // 不返回密码
	Template string `json:"Template"`
	SSL      bool   `json:"SSL"`
}

// UpdateMailboxConfigRequest 更新邮箱远程配置的请求数据结构
type UpdateMailboxConfigRequest struct {
	SMTPServer  string `json:"SmtpServer"`
	Port        int    `json:"Port"`
	DisplayName string `json:"DisplayName"`
	Username    string `json:"Username"`
	Password    string `json:"Password"`
	Template    string `json:"Template"`
	SSL         bool   `json:"SSL"`
}

// UpdateMailboxConfigResponse 更新邮箱远程配置的响应数据结构
type UpdateMailboxConfigResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// GetWechatConfigResponse 获取企业微信配置的响应数据结构
type GetWechatConfigResponse struct {
	SuccessResponse
	Data *GetWechatConfigResponseData `json:"Data"`
}
type GetWechatConfigResponseData struct {
	CompanyId string   `json:"CompanyId"` // 公司id
	AppId     string   `json:"AppId"`     // 应用id
	AppSecret string   `json:"AppSecret"` // 应用秘钥
	ToParty   []string `json:"ToParty"`   // 群组id，用于接收消息
	Template  string   `json:"Template"`  // 模板
}

// UpdateWechatConfigRequest 更新企业微信远程配置的请求数据结构
type UpdateWechatConfigRequest struct {
	CompanyId string   `json:"CompanyId"` // 公司id
	AppId     string   `json:"AppId"`     // 应用id
	AppSecret string   `json:"AppSecret"` // 应用秘钥
	ToParty   []string `json:"ToParty"`   // 群组id，用于接收消息
	Template  string   `json:"Template"`  // 模板
}

// UpdateWechatConfigResponse 更新企业微信远程配置的响应数据结构
type UpdateWechatConfigResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// GetFeishuConfigResponse 获取飞书配置的响应数据结构
type GetFeishuConfigResponse struct {
	SuccessResponse
	Data *GetFeishuConfigResponseData `json:"Data"`
}
type GetFeishuConfigResponseData struct {
	Url string `json:"Url"` // webhook地址
}

// UpdateFeishuConfigRequest 更新飞书远程配置的请求数据结构
type UpdateFeishuConfigRequest struct {
	Url string `json:"Url"` // webhook地址
}

// UpdateFeishuConfigResponse 更新飞书远程配置的响应数据结构
type UpdateFeishuConfigResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// ListDingtalkConfigRequest 获取钉钉配置列表的请求数据结构
type ListDingtalkConfigRequest struct {
	PageOption
}

// ListDingTalkConfigResponse 获取钉钉配置列表的响应数据结构
type ListDingTalkConfigResponse struct {
	SuccessResponse
	Data *model.DingtalkConfigList `json:"Data"`
}

// CreateDingtalkConfigRequest 创建钉钉配置列表的请求数据结构
type CreateDingtalkConfigRequest struct {
	// 机器人名称
	Name string `json:"Name"`
	// 机器人webhoook地址
	Url string `json:"Url"`
}

// CreateDingTalkConfigResponse 创建钉钉配置列表的响应数据结构
type CreateDingTalkConfigResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// DeleteDingtalkConfigRequest 删除钉钉配置列表的请求数据结构
type DeleteDingtalkConfigRequest struct {
	// 配置id
	Id int `json:"Id"`
}

// DeleteDingTalkConfigResponse 删除钉钉配置列表的响应数据结构
type DeleteDingTalkConfigResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// UpdateDingtalkConfigRequest 创建钉钉配置列表的请求数据结构
type UpdateDingtalkConfigRequest struct {
	Id int `json:"Id"`
	// 机器人名称
	Name string `json:"Name"`
	// 机器人webhoook地址
	Url string `json:"Url"`
}

// UpdateDingTalkConfigResponse 创建钉钉配置列表的响应数据结构
type UpdateDingTalkConfigResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}
