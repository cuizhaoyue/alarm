package model

// SMSConfig 短信配置
type SMSConfig struct {
	Id int `gorm:"primaryKey"`
	// Protocol string
	// IP       string `gorm:"column:ip"`
	URL string `gorm:"column:url"`
	// User     string
	Sign     string
	AK       string `gorm:"column:ak"`
	SK       string `gorm:"column:sk"`
	Template string
}

func (SMSConfig) TableName() string {
	return "sms_config"
}

// MailboxConfig 邮箱配置
type MailboxConfig struct {
	Id int `gorm:"primaryKey"`
	// 邮箱服务器地址
	SMTPServer string `gorm:"column:smtp_server"`
	// 邮箱服务器端口
	Port int
	// 显示名称
	DisplayName string
	// 邮箱用户名
	Username string
	// 邮箱密码
	Password string
	// SMTP是否加密
	SSL bool `gorm:"column:ssl"`
	// 邮件通知模板
	Template string
}

func (MailboxConfig) TableName() string {
	return "mailbox_config"
}

// WechatConfig 微信配置
type WechatConfig struct {
	Id int `gorm:"primaryKey"`
	// 公司id
	CompanyId string
	// 应用id
	AppId string
	// 应用秘钥
	AppSecret string
	// 群组id，用于接收消息
	ToParty       []string `gorm:"-"`
	ToPartyShadow string   `gorm:"column:to_party"`
	// 模板
	Template string
}

func (WechatConfig) TableName() string {
	return "wechat_config"
}

// FeiShuConfig 飞书配置
type FeishuConfig struct {
	Id  int    `gorm:"primaryKey"`
	Url string // webhook地址
}

func (FeishuConfig) TableName() string {
	return "feishu_config"
}

// DingtalkConfig 钉钉配置
type DingtalkConfig struct {
	Id int `gorm:"primaryKey"`
	// 机器人名称
	Name string
	// 机器人webhoook地址
	Url string
}

func (DingtalkConfig) TableName() string {
	return "dingtalk_config"
}

type DingtalkConfigList struct {
	Total int64             `json:"Total"`
	Items []*DingtalkConfig `json:"Items"`
}
