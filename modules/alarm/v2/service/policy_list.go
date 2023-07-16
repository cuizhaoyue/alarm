package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"
)

// List 获取告警策略列表数据
func (s *alarmService) List(ctx context.Context, req *dto.ListPolicyRequest) (*dto.ListPolicyResponseData, error) {
	// 设置页码offset值
	req.SetupOffset()

	data := &dto.ListPolicyResponseData{}

	// 获取策略列表
	var policyList []*model.AlarmPolicy
	if req.SearchValue != "" && req.SearchKey == "resource" {
		// 资源实例的模糊查询单独执行
		// 先按照资源实例名称查询到所有的告警策略
		listRsReq := &dto.ListResourcesOnPolicyRequest{
			PageOption: req.PageOption,
			Name:       req.SearchValue,
		}
		rsList, err := s.store.ResourceOnPolicy().ListResourceByName(ctx, listRsReq)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		// 获取所有的policy id
		var polyIds []string
		for _, item := range rsList.Items {
			polyIds = append(polyIds, item.PolicyId)
		}
		plist, err := s.store.Alarm().ListByInstanceIds(ctx, polyIds)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		policyList = plist
		data.Total = rsList.Total
	} else {
		policies, err := s.store.Alarm().List(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "error get policy list")
		}

		data.Total = policies.Total
		policyList = policies.Items
	}

	for _, p := range policyList {
		// 转换数据格式
		p.DB2FE()

		// 构建PolicyInfo对象
		info := &dto.PolicyInfo{
			Id:           p.Id,
			InstanceId:   p.InstanceId,
			Name:         p.Name,
			Comment:      p.Comment,
			Creator:      p.Creator,
			Enabled:      p.Enabled,
			Production:   p.Production,
			Labels:       p.Labels,
			Limit:        p.Limit,
			Type:         p.Type,
			FormPolicy:   p.FormPolicy,
			PromqlPolicy: p.PromqlPolicy,
			Receivers:    p.Receivers,
			NotifySetup:  p.NotifySetup,
			CreatedAt:    p.CreatedAt.UnixMilli(),
			UpdatedAt:    p.UpdatedAt.UnixMilli(),
		}

		data.Items = append(data.Items, info)
	}

	return data, nil
}
