package dto

type ListResourcesOnPolicyRequest struct {
	PageOption
	// 资源名称
	Name string `json:"Name"`
}
