package service

import (
	"context"
	"sort"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"
)

// Overview 获取告警总览
func (s *alertService) Overview(ctx context.Context, req *dto.OverviewRequest) (*dto.OverviewResponseData, error) {
	req.Offset = -1
	req.PageSize = -1

	lstReq := &dto.ListAlertsRequest{
		PageOption: req.PageOption,
		Status:     model.AlertFiring,
	}
	alertList, err := s.store.Alert().List(ctx, lstReq)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 按持续时长对告警分类
	durationAlerts := buildAlertsByDuration(alertList.Items)

	// 按资源实例对告警分类
	resourceAlerts := buildAlertsByResource(alertList.Items)

	// 按告警级别对告警分类
	levelAlerts := buildAlertsByLevel(alertList.Items)

	// 按服务对告警分类
	serviceAlerts := buildAlertsByType(alertList.Items, consts.xxxxxService)

	// 按物理资源对告警分类
	physicalAlerts := buildAlertsByType(alertList.Items, consts.PhysicalResource)

	// 按资源池对告警分类
	resourcePoolAlerts := buildAlertsByType(alertList.Items, consts.ResourcePool)

	data := &dto.OverviewResponseData{
		ByDuration:     durationAlerts,
		ByResource:     resourceAlerts,
		ByLevel:        levelAlerts,
		ByService:      serviceAlerts,
		ByPhysical:     physicalAlerts,
		ByResourcePool: resourcePoolAlerts,
	}

	return data, nil
}

// 按资源类型对告警分类
func buildAlertsByType(alerts []*model.Alert, rtype string) []dto.ResourceTypeAlerts {
	var rtypeAlerts []dto.ResourceTypeAlerts

	svcMap := make(map[string]map[string]int)

	switch rtype {
	case consts.PhysicalResource:
		for _, v := range model.PhysicalList {
			svcMap[v] = map[string]int{
				model.P0: 0, model.P1: 0, model.P2: 0, model.P3: 0,
			}
		}
	case consts.ResourcePool:
		for _, v := range model.ResourcePoolList {
			svcMap[v] = map[string]int{
				model.P0: 0, model.P1: 0, model.P2: 0, model.P3: 0,
			}
		}
	case consts.xxxxxService:
		for _, v := range model.ServiceList {
			svcMap[v] = map[string]int{
				model.P0: 0, model.P1: 0, model.P2: 0, model.P3: 0,
			}
		}
	}

	// 先按资源子类型分类再按告警等级分类
	for _, alert := range alerts {
		if rtype == alert.ResourceType {
			svcMap[alert.ResourceSubType][alert.Level]++
		}
	}

	for svc, levelMap := range svcMap {
		switch svc {
		case consts.xxxxxService:
		case consts.PhysicalResource:
		case consts.ResourcePool:

		}
		var items []dto.TileItem
		var num int
		for level, n := range levelMap {
			num += n
			items = append(items, dto.TileItem{
				Number: n,
				Unit:   "",
				Label:  model.LevelMap[level],
				Kind:   model.KindMap[level],
			})
		}
		rtypeAlerts = append(rtypeAlerts, dto.ResourceTypeAlerts{
			Prefix: model.FEPrefix,
			Number: num,
			Unit:   "",
			Label:  model.ResourceSubTypeMap[svc],
			List:   items,
		})
	}

	return rtypeAlerts
}

// 按告警等级对告警进行分类
func buildAlertsByLevel(alerts []*model.Alert) dto.BaseInfo {
	alertMap := make(map[string]int)

	for _, a := range alerts {
		alertMap[a.Level]++
	}

	values := []dto.Value{
		{
			Key:   model.P0,
			Name:  model.LevelMap[model.P0],
			Value: alertMap[model.P0],
		},
		{
			Key:   model.P1,
			Name:  model.LevelMap[model.P1],
			Value: alertMap[model.P1],
		},
		{
			Key:   model.P2,
			Name:  model.LevelMap[model.P2],
			Value: alertMap[model.P2],
		},
		{
			Key:   model.P3,
			Name:  model.LevelMap[model.P3],
			Value: alertMap[model.P3],
		},
	}

	levelAlerts := dto.BaseInfo{
		Info: dto.Info{
			Name: model.FELevelMsg,
		},
		Values: values,
	}

	return levelAlerts
}

// 按资源实例对告警分类
func buildAlertsByResource(alerts []*model.Alert) dto.BaseInfo {
	alertMap := make(map[string]int)

	for _, a := range alerts {
		alertMap[a.Resource]++
	}

	var ras []dto.ResourceAlert
	for k, v := range alertMap {
		ras = append(ras, dto.ResourceAlert{
			Name:  k,
			Value: v,
		})
	}

	sort.Sort(dto.ResourceAlerts(ras))
	if len(ras) > 10 {
		ras = ras[:9]
	}

	values := make([]dto.Value, len(ras))
	for i := 0; i < len(ras); i++ {
		values[i] = dto.Value{
			Name:  ras[i].Name,
			Value: ras[i].Value,
		}
	}

	resourceAlerts := dto.BaseInfo{
		Info: dto.Info{
			Name: model.FEResourceMsg,
		},
		Values: values,
	}

	return resourceAlerts
}

// 按持续时长对告警分类
func buildAlertsByDuration(alerts []*model.Alert) []dto.Value {
	alertMap := make(map[string]int)

	for _, a := range alerts {
		alertMap[a.DurationFlag]++
	}

	values := []dto.Value{
		{
			Key:   model.LessThanTenMin,
			Name:  model.DurationMap[model.LessThanTenMin],
			Value: alertMap[model.LessThanTenMin],
		},
		{
			Key:   model.TenMinToOneHour,
			Name:  model.DurationMap[model.TenMinToOneHour],
			Value: alertMap[model.TenMinToOneHour],
		},
		{
			Key:   model.OneHourToOneDay,
			Name:  model.DurationMap[model.OneHourToOneDay],
			Value: alertMap[model.OneHourToOneDay],
		},
		{
			Key:   model.MoreThanOneDay,
			Name:  model.DurationMap[model.MoreThanOneDay],
			Value: alertMap[model.MoreThanOneDay],
		},
	}

	return values
}
