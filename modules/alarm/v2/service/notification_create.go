package service

import (
	"context"
	"time"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"
)

// CreateOrUpdate 创建或更新告警消息记录
func (s *alertService) CreateOrUpdate(ctx context.Context, req *model.NotificationData) error {
	alerts := req.Alerts
	var records []*model.Alert

	for _, a := range alerts {
		labels := a.Labels
		annotatinos := a.Annotations
		alert := &model.Alert{
			AlertId:         a.Fingerprint,
			Name:            labels[consts.AlertName],
			PolicyName:      labels[consts.PolicyName],
			PolicyId:        labels[consts.PolicyId],
			Region:          labels[consts.Region],
			Az:              labels[consts.Az],
			Level:           annotatinos[consts.Level],
			ResourceType:    labels[consts.ResourceType],
			ResourceSubType: labels[consts.ResourceSubType],
			Expression:      annotatinos[consts.Expression],
			Threshold:       annotatinos[consts.Threshold] + annotatinos[consts.Unit],
			CurrentValue:    annotatinos[consts.CurrentValue],
			StartsAt:        a.StartsAt,
			EndsAt:          a.EndsAt,
			Duration:        a.EndsAt.Sub(a.StartsAt).String(),
			Status:          a.Status,
		}

		// 确定持续时长标志
		if a.EndsAt.Sub(a.StartsAt) <= time.Minute*10 {
			alert.DurationFlag = model.LessThanTenMin
		} else if a.EndsAt.Sub(a.StartsAt) > time.Minute*10 && a.EndsAt.Sub(a.StartsAt) <= time.Hour {
			alert.DurationFlag = model.TenMinToOneHour
		} else if a.EndsAt.Sub(a.StartsAt) > time.Hour && a.EndsAt.Sub(a.StartsAt) <= time.Hour*24 {
			alert.DurationFlag = model.OneHourToOneDay
		} else {
			alert.DurationFlag = model.MoreThanOneDay
		}

		// 根据不同资源类型确定资源实例名称
		switch labels[consts.ResourceType] {
		case consts.PhysicalResource:
			// 物理资源
			alert.Resource = labels[consts.Hostname]
		case consts.ResourcePool:
			// 资源池
			alert.Resource = labels[consts.ResourcePoolKey]
		case consts.xxxxxService:
			// 服务
			alert.Resource = labels[consts.xxxxxServiceKey]
		}

		records = append(records, alert)
	}

	err := s.store.Alert().CreateOrUpdate(ctx, records)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
