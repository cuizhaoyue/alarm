package model

import (
	"time"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	amtpl "github.com/prometheus/alertmanager/template"
)

// type AlertStatus string

const (
	// 告警状态
	AlertFiring   = "firing"
	AlertResolved = "resolved"

	// 告警持续时长
	// 小于10分钟
	LessThanTenMin = "lessThanTenMin"
	// 10分钟到1小时
	TenMinToOneHour = "tenMinToOneHour"
	// 1小时到24小时
	OneHourToOneDay = "oneHourToOneDay"
	// 大于24小时
	MoreThanOneDay = "moreThanOneDay"

	// 告警等级
	P0 = "p0"
	P1 = "p1"
	P2 = "p2"
	P3 = "p3"
)

// 前端需要的提示信息
const (
	FELevelMsg    = "总告警数"
	FEResourceMsg = "告警数"
	FEPrefix      = "总告警数"
)

var (
	TimeMap = map[string]string{
		"d": "天",
		"h": "小时",
		"m": "分钟",
	}

	DurationMap = map[string]string{
		"lessThanTenMin":  "小于10分钟",
		"tenMinToOneHour": "10分钟到1小时",
		"oneHourToOneDay": "1小时到一天",
		"moreThanOneDay":  "大于2一天",
	}

	LevelMap = map[string]string{
		"p0": "紧急",
		"p1": "重要",
		"p2": "次要",
		"p3": "提醒",
	}
	KindMap = map[string]string{
		"p0": "error",
		"p1": "warn",
		"p2": "minor",
		"p3": "info",
	}

	ServiceList      = []string{consts.MilkyPlatForm, consts.xxxxxPlatForm}
	PhysicalList     = []string{consts.PhysicalServer, consts.PhysicalSwitch}
	ResourcePoolList = []string{consts.Kec, consts.Ssd, consts.Ehdd, consts.Ks3, consts.Rds, consts.Kcs,
		consts.Xgw, consts.Tengine, consts.Nat, consts.Kgw, consts.Sgw, consts.Pgw, consts.Tgw}

	ResourceSubTypeMap = map[string]string{
		// 服务
		consts.MilkyPlatForm: "银河平台",
		consts.xxxxxPlatForm: "GMS运维平台",
		// 资源子类型
		consts.PhysicalServer: "服务器",
		consts.PhysicalSwitch: "交换机",
		// 资源池
		consts.Kec:     "计算资源池",
		consts.Ssd:     "云硬盘3.0(SSD)资源池",
		consts.Ehdd:    "高效云盘资源池",
		consts.Ks3:     "对象存储资源池",
		consts.Rds:     "MySQL资源池",
		consts.Kcs:     "Redis资源池",
		consts.Xgw:     "XGW资源池",
		consts.Tengine: "Tengine资源池",
		consts.Nat:     "NAT资源池",
		consts.Kgw:     "KGW资源池",
		consts.Sgw:     "SGW资源池",
		consts.Pgw:     "PGW资源池",
		consts.Tgw:     "TGW资源池",
	}
)

type (
	NotificationData = amtpl.Data
)

// Alert 告警记录
type Alert struct {
	Id int `gorm:"primaryKey"`
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

func (Alert) TableName() string {
	return "alert"
}

type AlertList struct {
	Total int64
	Items []*Alert
}
