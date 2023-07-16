package doc

import "ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"

// ListPromqlTpl 获取PromQL模板
// swagger:route GET /alarm/v2/promql/list promql listPromqlTplRequest
//
// 获取PromQL模板
//
// 获取PromQL模板
//
//      Responses:
//        default: errResponse
//        200: listPromqlTplResponse

// swagger:parameters listPromqlTplRequest
type listPromqlTplRequestWrapper struct {
	// in:query
	Action string `json:"Action" form:"Action"`
	// in: query
	ResourceSubType string `json:"resource_sub_type" form:"resource_sub_type"`
}

// swagger:response listPromqlTplResponse
type listPromqlTplResponseWrapper struct {
	// in: body
	Body dto.ListPromqlTplResponse
}
