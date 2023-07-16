package consts

const (
	// 资源类型
	PhysicalResource = "physicalResource" // 物理资源
	ResourcePool     = "resourcePool"     // 资源池
	xxxxxService     = "service"          // 服务

	// 资源子类型
	PhysicalServer = "physicalServer" // 服务器
	PhysicalSwitch = "physicalSwitch" // 交换机

	// 资源池
	Kec     = "kec"     // 计算资源池
	Ssd     = "ssd3.0"  // 云硬盘3.0(SSD)资源池
	Ehdd    = "ehdd"    // 高效云盘资源池
	Ks3     = "ks3"     // 对象存储资源池
	Rds     = "rds"     // MySQL资源池
	Kcs     = "kcs"     // Redis资源池
	Xgw     = "xgw"     // XGW资源池
	Tengine = "tengine" // Tengine资源池
	Nat     = "nat"     // NAT资源池
	Kgw     = "kgw"     // KGW资源池
	Sgw     = "sgw"     // SGW资源池
	Pgw     = "pgw"     // PGW资源池
	Tgw     = "tgw"     // TGW资源池

	// 服务
	MilkyPlatForm = "milkyPlatForm" // 银河平台
	xxxxxPlatForm = "xxxxxPlatForm" // 鲁班平台

	// key名称
	Hostname        = "hostname"
	ResourcePoolKey = "resource_pool"
	xxxxxServiceKey = "service"
	CurrentValue    = "current_value"
	Region          = "region"
	Az              = "az"
)

// 创建PrometheusRule时必须添加以下label，否则生成Prometheus告警
const (
	DefaultLabelname  = "release"
	DefaultLabelvalue = "pm"
)

const (
	// label或annotation的key名称
	PolicyName      = "policy_name"       // 策略名称
	PolicyId        = "policy_id"         // 策略id
	PolicyType      = "policy_type"       // 策略类型
	ResourceType    = "resource_type"     // 资源类型
	ResourceSubType = "resource_sub_type" // 资源子类型
	RuleId          = "rule_id"           // 告警规则id

	AlertName          = "alert_name"           // 告警显示名称
	MonitorName        = "monitor_name"         // 监控项名称, e.g. compute_pool_cpu_usage
	MonitorDisplayName = "monitor_display_name" // 监控项显示名称, e.g. CPU使用率
	Level              = "level"                // 告警等级
	For                = "for"                  // 持续时间
	BinaryOperator     = "binary_operator"      // 运算操作符
	Threshold          = "threshold"            // 阈值
	Unit               = "unit"                 // 单位
	SendResolved       = "send_resolved"        // 是否发送告警恢复通知
	RepeatInterval     = "repeat_interval"      // 告警间隔
	GroupWait          = "group_wait"           // 等待时间
	Expression         = "exporession"          // 告警规则描述

	ToEmails       = "to_emails"       // 告警接收邮箱
	ToSms          = "to_sms"          // 告警短信接收用户
	EnableWechat   = "enable_wechat"   // 是否通知微信
	EnableFeishu   = "enable_feishu"   // 是否通知飞书
	EnableDingtalk = "enable_dingtalk" // 是否通知钉钉
	ToDingtalk     = "to_dingtalk"     // 接收告警的钉钉机器人
)

const (
	// 排序码
	// 升序
	AsceOrder = "asce"
	// 降序
	DescOrder = "desc"

	// 按开始时间排序
	OrderByStart = "start"
	// 按结束时间排序
	OrderByEnd = "end"
)
