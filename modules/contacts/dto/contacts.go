package dto

type BaseUserInfo struct {
	// 用户id
	UserId string `json:"UserId"`
	// 用户名称
	Username string `json:"Username"`
	// 用户邮箱
	Email string `json:"Email"`
	// 用户电话
	Telephone string `json:"Telephone"`
}

// CreateContactRequest 创建联系人组的请求数据
type CreateContactRequest struct {
	// 联系人组的名称
	Name string `json:"Name"`
	// 添加到联系人组的用户列表
	Users []*BaseUserInfo `json:"Users"`
}

// CreateContactResponse 创建联系人组的响应数据结构
type CreateContactResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// type ContactResponseData struct {
// 	// 联系人组的id
// 	Id int `json:"Id"`
// 	// 联系人组的自定义索引
// 	Index int `json:"Index"`
// 	// 联系人组的名称
// 	Name string `json:"Name"`
// }

// UpdateContactRequest 编辑联系人组的请求数据结构
type UpdateContactRequest struct {
	// 正在编辑的联系人组id
	Id int `json:"Id"`
	// 联系人组名称
	Name string `json:"Name"`
	// 编辑后要加入联系人组的用户列表
	Users []*BaseUserInfo `json:"Users"`
}

// UpdateContactResponse 编辑联系人组的响应数据结构
type UpdateContactResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// ListContactRequest 获取联系人组列表的请求数据结构
type ListContactRequest struct {
	PageOption
	// 模糊搜索的key，允许值：name
	SearchKey string `json:"SearchKey" example:"name"`
	// 模糊搜索的value
	SearchValue string `json:"SearchValue"`
}

// ListContactResponse 获取联系人组列表的响应数据结构
type ListContactResponse struct {
	SuccessResponse
	Data *ListContactResponseData `json:"Data"`
}

type ListContactResponseData struct {
	Total int64           `json:"Total"`
	Items []*ContactsInfo `json:"Items"`
}

type ContactsInfo struct {
	Id        int     `json:"Id"`   // 联系人组id
	Name      string  `json:"Name"` // 联系人组名称
	CreatedAt int64   `json:"CreatedAt"`
	UpdatedAt int64   `json:"UpdatedAt"`
	Users     []*User `json:"Users"`
}

type User struct {
	BaseUserInfo
	// 加入用户组的时间
	CreatedAt int64 `json:"CreatedAt"`
}

// DeleteContactRequest 删除联系人组的请求数据结构
type DeleteContactRequest struct {
	// 要删除的联系人组的id
	Id int `json:"Id"`
}

// DeleteContactResponse 删除联系人组的响应数据结构
type DeleteContactResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}

// RemoveUserFromContactRequest 从联系人组中删除用户的请求数据结构
type RemoveUserFromContactRequest struct {
	Id     int    `json:"Id"`
	UserId string `json:"UserId"`
}

// RemoveUserFromContactResponse 从联系人组中删除用户的响应数据结构
type RemoveUserFromContactResponse struct {
	SuccessResponse
	Data struct{} `json:"Data"`
}
