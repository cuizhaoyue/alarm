package model

import "time"

// PromQL模板
type PromqlTemplate struct {
	// 模板id
	Id int `gorm:"primaryKey" json:"-"`
	// 模板英文标识
	Name string
	// 模板显示名称, e.g. CPU使用率
	DisplayName string
	// PromQL 模板
	Promql string
	// 单位
	Unit string
	// 聚合方式、统计方式
	Algo string
	// 统计方式的名称
	AlgoDisplayName string
	// 资源子类型的英文标识, 根据子类型来查找对应的模板
	ResourceSubType string
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

func (*PromqlTemplate) TableName() string {
	return "promql_template"
}

type PromqlTemplateList struct {
	Total int64
	Items []*PromqlTemplate
}
