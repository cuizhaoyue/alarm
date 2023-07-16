package dto

type ListPromqlTplResponse struct {
	SuccessResponse
	Data ListPromqlTplResponseData `json:"Data"`
}

type ListPromqlTplResponseData struct {
	Items []*TplInfo
}

type TplInfo struct {
	// 模板英文标识
	Name string `json:"Name"`
	// 模板显示名称, e.g. CPU使用率
	DisplayName string `json:"DisplayName"`
	// PromQL 模板
	Promql string `json:"Promql"`
	// 单位
	Unit string `json:"Unit"`
	// 聚合方式、统计方式
	Algo string `json:"Algo"`
	// 统计方式的名称
	AlgoDisplayName string `json:"AlgoDisplayName"`
	// 资源子类型的英文标识, 根据子类型来查找对应的模板
	ResourceSubType string `json:"ResourceSubType"`
}
