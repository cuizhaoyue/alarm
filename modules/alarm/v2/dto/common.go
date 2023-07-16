package dto

// ErrResponse 错误响应
type ErrResponse struct {
	// Example: 1001
	Code int `json:"Code"`
	// Example: 失败
	Message string `json:"Message" example:"失败"`
	// Example: {}
	Data struct{} `json:"Data"`
}

// SuccessResponse 成功响应
type SuccessResponse struct {
	// 响应码，0表示请求成功
	// Example: 0
	Code int `json:"Code"`
	// Example: 成功
	Message string `json:"Message" example:"成功"`
}

// PageOption 分页
type PageOption struct {
	// 每页数据个数，默认10个, 当值为-1时不分页
	PageSize int  `json:"PageSize"`
	// 当前页码，从1开始
	PageNo int  `json:"PageNo"`
	// 数据偏移量
	Offset int `json:"-"`
}

// SetupOffset 根据页码设置offset的值
func (p *PageOption) SetupOffset() {
	if p.PageSize == -1 {
		p.Offset = -1
		p.PageSize = -1

		return
	}
	if p.PageNo == 0 && p.PageSize == 0 {
		p.PageNo = 1
		p.PageSize = 10
	}

	p.Offset = (p.PageNo - 1) * p.PageSize
}
