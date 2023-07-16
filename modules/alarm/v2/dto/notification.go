package dto

import "time"

// ListAlertsRequest 获取告警列表的请求数据结构
type ListAlertsRequest struct {
	PageOption
	// 所属区域 all-所有区域
	Region string `json:"Region"`
	// 所属可用区, all-所有可用区
	Az []string `json:"Az"`
	// 告警级别
	Level []string `json:"Level"`
	// 资源类型
	ResourceType []string `json:"ResourceType"`
	// 资源子类型
	ResourceSubType []string `json:"ResourceSubType"`
	// 持续时间
	// lessThanTenMin - 小于10分钟
	// tenMinToOneHour - 10分钟到1小时
	// oneHourToOneDay - 1小时到24小时
	// moreThanOneDay - 大于24小时
	Duration []string `json:"Duration"`
	// 模糊搜索, name-告警名称，policy_name-策略名称，resource-告警实例
	SearchKey   string `json:"SearchKey"`
	SearchValue string `json:"SearchValue"`
	// 告警状态，firing 或 resolved
	Status string `json:"Status"`
	// 按某个字段排序, start-按开始时间排序, end-按结束时间排序
	OrderBy string
	// 排序码, asce-升序, desc-降序
	OrderCode string
}

// ListAlertsResponse 获取告警列表的响应数据结构
type ListAlertsResponse struct {
	SuccessResponse
	Data ListAlertsResponseData `json:"Data"`
}

type ListAlertsResponseData struct {
	Total int64
	Items []*AlertInfo
}

type AlertInfo struct {
	Id int `json:"Id"`
	// 告警id, 也是Fingerprint
	AlertId string `json:"AlertId"`
	// 告警名称
	Name string `json:"Name"`
	// 策略名称
	PolicyName string `json:"PolicyName"`
	// 所属区域
	Region string `json:"Region"`
	// 所属可用区
	Az string `json:"Az"`
	// 告警级别
	Level string `json:"Level"`
	// 处理人
	Handler string `json:"Handler"`
	// 资源类型
	ResourceType string `json:"ResourceType"`
	// 资源子类型
	ResourceSubType string `json:"ResourceSubType"`
	// 告警实例
	Resource string `json:"Resource"`
	// 告警规则
	Expression string `json:"Expression"`
	// 阈值
	Threshold string `json:"Threshold"`
	// 当前值
	CurrentValue string `json:"CurrentValue"`
	// 开始时间
	StartsAt int64 `json:"StartsAt"`
	// 结束时间
	EndsAt int64 `json:"EndsAt"`
	// 告警持续时间
	Duration string `json:"Duration"`

	// 告警状态
	Status string `json:"Status"`
}

// OverviewRequest 告警总览请求数据结构
type OverviewRequest struct {
	PageOption
	// 所属区域
	Region string `json:"Region"`
}

type OverviewResponse struct {
	SuccessResponse
	Data OverviewResponseData `json:"Data"`
}

// OverviewResponseData 告警总览数据
type OverviewResponseData struct {
	// 按持续时间分类
	ByDuration []Value `json:"ByDuration"`

	// 按资源实例分类获取top10
	ByResource BaseInfo `json:"ByResource"`

	// 按告警等级分类
	ByLevel BaseInfo `json:"ByLevel"`

	// 按服务对告警分类
	ByService []ResourceTypeAlerts `json:"ByService"`

	// 按物理资源对告警分类
	ByPhysical []ResourceTypeAlerts `json:"ByPhysical"`

	// 按资源池对告警分类
	ByResourcePool []ResourceTypeAlerts `json:"ByResourcePool"`

	// 按云产品对告警分类
	ByCloudProduct []ResourceTypeAlerts `json:"ByCloudProduct"`
}

type BaseInfo struct {
	// 前端页面的提示信息
	Info Info `json:"Info"`
	// 前端页面需要的值
	Values []Value `json:"Values"`
}

type Info struct {
	Name string `json:"Name"`
}

type Value struct {
	// 英文名称
	Key string `json:"Key"`
	// 中文名称
	Name string `json:"Name"`
	// 告警数量
	Value int `json:"Value"`
}

// ResourceAlert 按资源实例对告警分类
type ResourceAlert struct {
	// 资源实例名称
	Name string `json:"Name"`
	// 告警数量
	Value int `json:"Value"`
}

type ResourceAlerts []ResourceAlert

func (a ResourceAlerts) Len() int           { return len(a) }
func (a ResourceAlerts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ResourceAlerts) Less(i, j int) bool { return a[i].Value > a[j].Value }

// ResourceTypeAlerts 按服务对告警分类
type ResourceTypeAlerts struct {
	Prefix string     `json:"Prefix"`
	Number int        `json:"Number"`
	Unit   string     `json:"Unit"`
	Label  string     `json:"Label"`
	List   []TileItem `json:"List"`
}

type TileItem struct {
	Number int    `json:"Number"`
	Unit   string `json:"Unit"`
	Label  string `json:"Label"`
	Kind   string `json:"Kind"`
}

// HandleRequest 为告警分配处理人的请求数据结构
type HandleRequest struct {
	Id int
	// 告警id, 也是Fingerprint
	AlertId string
	// 告警名称
	Name string
	// 策略名称
	PolicyName string
	// 策略id
	PolicyId string
	// 所属区域
	Region string
	// 所属可用区
	Az string
	// 告警级别
	Level string
	// 处理人
	Handler string
	// 资源类型
	ResourceType string
	// 资源子类型
	ResourceSubType string
	// 告警实例
	Resource string
	// 告警规则
	Expression string
	// 阈值
	Threshold string
	// 当前值
	CurrentValue string
	// 开始时间
	StartsAt time.Time
	// 结束时间
	EndsAt time.Time
	// 告警持续时间
	Duration     string
	DurationFlag string

	// 告警状态
	Status string
}
