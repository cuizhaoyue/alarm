package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
)

// Get 获取告警策略详情
func (s *alarmService) Get(ctx context.Context, insId string) (*dto.GetPolicyResponseData, error) {
	ret := &dto.GetPolicyResponseData{}

	policy, err := s.store.Alarm().Get(ctx, insId)
	if err != nil {
		return nil, err
	}

	// 转换数据结构
	policy.DB2FE()

	policyInfo := &dto.PolicyInfo{
		Id:           policy.Id,
		InstanceId:   policy.InstanceId,
		Name:         policy.Name,
		Creator:      policy.Creator,
		Enabled:      policy.Enabled,
		Type:         policy.Type,
		FormPolicy:   policy.FormPolicy,
		PromqlPolicy: policy.PromqlPolicy,
		Receivers:    policy.Receivers,
		NotifySetup:  policy.NotifySetup,
	}

	ret.PolicyInfo = policyInfo

	return ret, nil
}
