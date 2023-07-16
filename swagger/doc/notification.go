package doc

import "ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"

// ListAlerts 获取告警列表
// swagger:route POST /alarm/v2/alert/list alert listAlertRequest
//
// 获取告警列表
//
// 获取告警列表
//
//      Responses:
//        default: errResponse
//        200: listAlertResponse

// swagger:parameters listAlertRequest
type listAlertRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.ListAlertsRequest
}

// swagger:response listAlertResponse
type listAlertResponseWrapper struct {
	// in: body
	Body dto.ListAlertsResponse
}

// AlertOverview 告警总览
// swagger:route POST /alarm/v2/alert/overview alert alertOverviewRequest
//
// 告警总览
//
// 告警总览
//
//      Responses:
//        default: errResponse
//        200: alertOverviewResponse

// swagger:parameters alertOverviewRequest
type alertOverviewRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: body
	Body dto.OverviewRequest
}

// swagger:response alertOverviewResponse
type alertOverviewResponseWrapper struct {
	// in: body
	Body dto.OverviewResponse
}
