package service

import (
	"context"
	"sort"
	"strconv"
	"strings"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"
)

// List 获取告警列表
func (s *alertService) List(ctx context.Context, req *dto.ListAlertsRequest) (*dto.ListAlertsResponseData, error) {
	req.SetupOffset()

	respData := &dto.ListAlertsResponseData{}
	var alertInfos []*dto.AlertInfo

	alerts, err := s.store.Alert().List(ctx, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	respData.Total = alerts.Total

	for _, a := range alerts.Items {
		info := &dto.AlertInfo{
			Id:              a.Id,
			AlertId:         a.AlertId,
			Name:            a.Name,
			PolicyName:      a.PolicyName,
			Region:          a.Region,
			Az:              a.Az,
			Level:           a.Level,
			Handler:         a.Handler,
			ResourceType:    a.ResourceType,
			ResourceSubType: a.ResourceSubType,
			Resource:        a.Resource,
			Expression:      a.Expression,
			Threshold:       a.Threshold,
			CurrentValue:    a.CurrentValue,
			StartsAt:        a.StartsAt.UnixMilli(),
			EndsAt:          a.EndsAt.UnixMilli(),
			Duration:        convertDuration(a.Duration),
			Status:          a.Status,
		}

		alertInfos = append(alertInfos, info)
	}

	// 对告警排序
	switch req.OrderBy {
	case consts.OrderByStart:
		// 按开始时间排序
		switch req.OrderCode {
		case consts.AsceOrder:
			// 升序排序
			asc := func(a1, a2 *dto.AlertInfo) bool {
				return a1.StartsAt < a2.StartsAt
			}
			By(asc).Sort(alertInfos)
		case consts.DescOrder:
			// 降序排序
			desc := func(a1, a2 *dto.AlertInfo) bool {
				return a1.StartsAt > a2.StartsAt
			}
			By(desc).Sort(alertInfos)
		}
	case consts.OrderByEnd:
		// 按结束时间排序
		switch req.OrderCode {
		case consts.AsceOrder:
			// 升序排序
			asc := func(a1, a2 *dto.AlertInfo) bool {
				return a1.EndsAt < a2.EndsAt
			}
			By(asc).Sort(alertInfos)
		case consts.DescOrder:
			// 降序排序
			desc := func(a1, a2 *dto.AlertInfo) bool {
				return a1.EndsAt > a2.EndsAt
			}
			By(desc).Sort(alertInfos)
		}
	}

	respData.Items = alertInfos

	return respData, nil
}

// 把持续时间转换成中文描述, 精确到分钟
func convertDuration(duration string) string {
	before, _, _ := strings.Cut(duration, "m")

	hoursOrAll, mins, found := strings.Cut(before, "h")
	if !found {
		min, _, _ := strings.Cut(hoursOrAll, "m")
		return min + model.TimeMap["m"]
	}

	intHour, _ := strconv.Atoi(hoursOrAll)

	if intHour < 24 {
		return hoursOrAll + model.TimeMap["h"] + mins + model.TimeMap["m"]
	}

	day, hour := intHour/24, intHour%24
	strDay := strconv.Itoa(day)
	strHour := strconv.Itoa(hour)
	return strDay + model.TimeMap["d"] + strHour + model.TimeMap["h"] + mins + model.TimeMap["m"]
}

// 对告警列表排序
type By func(a1, a2 *dto.AlertInfo) bool

func (by By) Sort(alerts []*dto.AlertInfo) {
	as := &alertsSorter{
		alerts: alerts,
		by:     by,
	}

	sort.Sort(as)
}

type alertsSorter struct {
	alerts []*dto.AlertInfo
	by     By
}

func (a *alertsSorter) Len() int           { return len(a.alerts) }
func (a *alertsSorter) Swap(i, j int)      { a.alerts[i], a.alerts[j] = a.alerts[j], a.alerts[i] }
func (a *alertsSorter) Less(i, j int) bool { return a.by(a.alerts[i], a.alerts[j]) }
