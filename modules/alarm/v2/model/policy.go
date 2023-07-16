package model

import (
	"encoding/json"
	"time"
)

type AlarmPolicyType string

// 定义告警策略的类型
const (
	FormAlarmPolicyType AlarmPolicyType = "form"
	PromAlarmPolicyType AlarmPolicyType = "promql"
)

// AlarmPolicy 告警策略
type AlarmPolicy struct {
	Id int `gorm:"primaryKey;column:id"`
	// 实例id
	InstanceId string
	// 告警策略名称
	Name string
	// 告警策略的创建人
	Creator string
	// 告警策略的最后修改人
	Updater string
	// 是否启用告警策略，true-启用，false-不启用
	Enabled bool
	// 资源类型, 当Type为form时此字段才有值
	ResourceType string
	// 资源子类型, 当Type为form时此字段才有值
	ResourceSubType string
	// 业务线名称
	Production string
	// 策略描述
	Comment string
	// 策略标签
	Labels       map[string]string `gorm:"-"`
	LabelsShadow string            `gorm:"column:labels"`
	// 每条告警规则的最大告警次数
	Limit int
	// 告警策略类型, form-表单告警策略, promql-PromQL告警策略
	Type AlarmPolicyType
	// 表单格式的告警策略, Type为form时传此字段
	FormPolicy *FormPolicy `gorm:"-"`
	// 序列化表单格式告警策略，保存于数据库中
	FormPolicyShadow string `gorm:"column:form_policy"`
	// PromQL告警策略, Type为promql时传此字段
	PromqlPolicy *PromqlPolicy `gorm:"-"`
	// 序列化PromQL告警策略，保存于数据库中
	PromqlPolicyShadow string `gorm:"column:promql_policy"`
	// 告警接收人
	Receivers       *Receivers `gorm:"-"`
	ReceiversShadow string     `gorm:"column:receivers"`
	// 通知设置
	NotifySetup       *NotifySetup `gorm:"-"`
	NotifySetupShadow string       `gorm:"column:notify_setup"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (*AlarmPolicy) TableName() string {
	return "alarm_policy"
}

// FE2DB 把前端传入的数据转换成数据库可存储的字符串形式
func (ap *AlarmPolicy) FE2DB() error {
	if len(ap.Labels) > 0 {
		labelsByts, err := json.Marshal(ap.Labels)
		if err != nil {
			return err
		}
		ap.LabelsShadow = string(labelsByts)
	}
	if ap.FormPolicy != nil {
		formPolicy, err := json.Marshal(ap.FormPolicy)
		if err != nil {
			return err
		}
		ap.FormPolicyShadow = string(formPolicy)
	}
	if ap.PromqlPolicy != nil {
		promqlPolicy, err := json.Marshal(ap.PromqlPolicy)
		if err != nil {
			return err
		}
		ap.PromqlPolicyShadow = string(promqlPolicy)
	}
	if ap.Receivers != nil {
		receivers, err := json.Marshal(ap.Receivers)
		if err != nil {
			return err
		}
		ap.ReceiversShadow = string(receivers)
	}
	if ap.NotifySetup != nil {
		notifySetup, err := json.Marshal(ap.NotifySetup)
		if err != nil {
			return err
		}
		ap.NotifySetupShadow = string(notifySetup)
	}

	return nil
}

// DB2FE 把数据库中的字符串形式反序列化为go结构类型
func (ap *AlarmPolicy) DB2FE() {
	if ap.LabelsShadow != "" {
		ap.Labels = map[string]string{}
		_ = json.Unmarshal([]byte(ap.LabelsShadow), &ap.Labels)
	}
	if ap.FormPolicyShadow != "" {
		ap.FormPolicy = &FormPolicy{}
		_ = json.Unmarshal([]byte(ap.FormPolicyShadow), ap.FormPolicy)
	}
	if ap.PromqlPolicyShadow != "" {
		ap.PromqlPolicy = &PromqlPolicy{}
		_ = json.Unmarshal([]byte(ap.PromqlPolicyShadow), ap.PromqlPolicy)
	}
	if ap.ReceiversShadow != "" {
		ap.Receivers = &Receivers{}
		_ = json.Unmarshal([]byte(ap.ReceiversShadow), ap.Receivers)
	}
	if ap.NotifySetupShadow != "" {
		ap.NotifySetup = &NotifySetup{}
		_ = json.Unmarshal([]byte(ap.NotifySetupShadow), ap.NotifySetup)
	}
}

// AlarmPolicyList 告警策略列表
type AlarmPolicyList struct {
	Total int64
	Items []*AlarmPolicy
}

// FormPolicy 表单告警策略
type FormPolicy struct {
	// 资源类型
	ResourceType string `json:"ResourceType"`
	// 资源类型的中文名称
	ResourceTypeName string `json:"ResourceTypeName"`
	// 资源子类型
	ResourceSubType string `json:"ResourceSubType"`
	// 资源子类型的中文名称
	ResourceSubTypeName string `json:"ResourceSubTypeName"`
	// 告警规则详情
	Rules []*AlertRule `json:"Rules"`
	// 选择的资源
	Resources Resources `json:"Resources"`
}

type AlertRule struct {
	// 告警规则唯一id
	RuleId string `json:"RuleId"`
	// 告警显示名称
	Name string `json:"Name"`
	// 监控项名称，英文名称，e.g. cpu_use_percent
	MonitorName string `json:"MonitorName"`
	// 监控项显示名，中文名称, e.g. cpu使用率
	DisplayName string `json:"DisplayName"`
	// 自定义标签
	Labels map[string]string `json:"Labels"`
	// 告警等级
	Level string `json:"Level"`
	// 持续时间
	For string `json:"For"`
	// 统计方式
	Algo string `json:"Algo"`
	// 统计方式显示名
	AlgoDisplayName string `json:"AlgoDisplayName"`
	// 运算操作符
	Operator string `json:"Operator"`
	// 阈值
	Threshold float64 `json:"Threshold"`
	// 告警间隔
	Interval string `json:"Interval"`
	// 等待时间
	GroupWaitTime string `json:"GroupWaitTime"`
	// 中文表达式
	ExpressionWithChinese string `json:"ExpressionWithChinese"`
	// 告警策略的单位
	Unit string `json:"Unit"`
	// 规则中的PromQL告警模板
	PromqlTpl *PromqlTemplate `json:"PromqlTpl"`
}

type Resources struct {
	// 资源选择方式, 1-按服务树选择, 2-按范围选择, 3-按对象选择
	Index int `json:"Index"`
	// 资源实例列表
	Instances []*Resource `json:"Instances"`
	// 选择的资源实例数量
	Count int `json:"Count"`
}

// Resource 表示被选择的资源(资源结构后续可能会发生变化，暂时按下面结构处理)
type Resource struct {
	Key                 string `json:"Key"`
	Value               string `json:"Value"`
	Name                string `json:"Name"`
	Region              string `json:"Region"`
	RegionName          string `json:"RegionName"`
	Az                  string `json:"Az"`
	AzName              string `json:"AzName"`
	Lab                 string `json:"Lab"`
	LabName             string `json:"LabName"`
	ResourceTypeName    string `json:"ResourceTypeName"`
	ResourceSubTypeName string `json:"ResourceSubTypeName"`
}

// PromqlPolicy 告警策略
type PromqlPolicy struct {
	// 告警规则id
	RuleId string `json:"RuleId"`
	// prometheus 查询语句
	Promql string `json:"Promql"`
	// 告警显示名称
	Name string `json:"Name"`
	// 告警等级
	Level string `json:"Level"`
	// 持续时间
	For string `json:"For"`
	// 告警间隔
	Interval string `json:"Interval"`
	// 告警次数
	AlarmCount int `json:"AlarmCount"`
	// 等待时间
	GroupWaitTime string `json:"GroupWaitTime"`
	// 中文表达式
	ExpressionWithChinese string `json:"ExpressionWithChinese"`
}

type Receivers struct {
	// 是否发送告警恢复通知
	SendResolved bool `json:"SendResolved"`
	// 联系人
	NoticeUsers []*NoticeUser `json:"NoticeUsers"`
	// 联系人组
	ContactsGroup []*Contacts `json:"ContactsGroup"`
}

// NoticeUser 联系人
type NoticeUser struct {
	// 用于标识创建顺序
	Index int `json:"Index"`
	// 用户id
	UserId string `json:"UserId"`
	// 用户名称
	Username string `json:"Username"`
	// 邮箱
	Email string `json:"Email"`
	// 电话
	Telephone string `json:"Telephone"`
	// 是否发送邮件
	EnableEmail bool `json:"SendEmail"`
	// 是否发送短信
	EnableSms bool `json:"SendSms"`
}

// Contacts 联系人组
type Contacts struct {
	// 联系人组id
	Id int `json:"Id"`
	// 用于标识创建顺序
	Index int `json:"Index"`
	// 联系人组名称
	Name string `json:"Name"`
	// 是否发送邮件
	EnableEmail bool `json:"EnableEmail"`
	// 是否发送短信
	EnableSms bool `json:"EnableSms"`
	// 联系人组中的用户
	Users []*User `json:"Users"`
}

type User struct {
	// 用户id
	UserId string `json:"UserId"`
	// 用户名称
	Username string `json:"Username"`
	// 邮箱
	Email string `json:"Email"`
	// 电话
	Telephone string `json:"Telephone"`
}

// NotifySetup 通知设置
type NotifySetup struct {
	// 是否发送到企业微信
	EnableWecom bool `json:"EnableWecom"`
	// 是否发送到飞书
	EnableFeishu bool `json:"EnableFeishu"`
	// 飞书业务群
	FeishuGroup string `json:"FeishuGroup"`
	// 是否发送到钉钉机器人
	EnableDingtalk bool              `json:"EnableDingtalk"`
	DingtalkRobots []*DingtalkConfig `json:"DingtalkRobots"`
}
