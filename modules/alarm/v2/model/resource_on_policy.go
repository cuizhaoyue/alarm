package model

// ResourceOnPolicy 创建策略时选择的资源实例
type ResourceOnPolicy struct {
	Id int `gorm:"primaryKey"`
	// 策略id
	PolicyId string
	// 资源实例的名称
	Name string
	// 资源所属region
	Region string
	// 资源所属az
	Az string
}

// TableName 设置表名
func (*ResourceOnPolicy) TableName() string {
	return "resources_on_policy"
}

type ResourceOnPolicyList struct {
	Total int64
	Items []*ResourceOnPolicy
}
