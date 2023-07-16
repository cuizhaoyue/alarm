package service

import (
	"context"
	"time"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

// Update 编辑更新告警策略
func (s *alarmService) Update(ctx context.Context, req *dto.UpdatePolicyRequest) error {
	// 构造告警实例
	policy, err := buildPolicyInstanceFromUpdateReq(req)
	if err != nil {
		return errors.Wrapf(err, "build policy instance [%s] from update request failed", req.Name)
	}

	labels, annotations := buildCommonLablesAndAnnotations(policy)

	// 构建所选资源列表
	resources := buildResourceOnPolicy(policy)

	// 获取原来的告警策略实例并判断是否需要重新创建AlertmanagerConfig实例
	var alarmChange bool
	oriPolicy, err := s.store.Alarm().Get(ctx, req.InstanceId)
	if err != nil {
		return errors.Wrap(err, "get origin policy failed")
	}
	if oriPolicy.FormPolicyShadow != policy.FormPolicyShadow || oriPolicy.PromqlPolicyShadow != policy.PromqlPolicyShadow {
		alarmChange = true
	}

	// 更新数据库记录
	err = s.store.Alarm().Tx(ctx, func(ctx context.Context, store dao.Factory) error {
		if err := store.Alarm().Update(ctx, policy); err != nil {
			return err
		}
		// 表单格式告警策略时，要更新资源表
		if req.Type == model.FormAlarmPolicyType {
			// 先删除原来所有的资源记录，再将新的资源记录更新进去
			if err := store.ResourceOnPolicy().Delete(ctx, policy.InstanceId); err != nil {
				return err
			}
			if err := store.ResourceOnPolicy().Create(ctx, resources); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "update policy instance [%s] failed", req.Name)
	}

	// 获取PrometheusRule实例
	promRule, err := s.store.PromRuleOperator().Get(ctx, req.InstanceId)
	if err != nil {
		return errors.Wrapf(err, "get PrometheusRule ")
	}

	// 构建告警规则列表
	rules := buildRulesFromPolicy(policy, labels, annotations)

	if checkBuildAm(labels, annotations) {
		// 告警规则有更改则需要更新AlertmanagerConfig
		if alarmChange {
			// 先删除之前具有相同告警策略id的AlertmanagerConfig实例
			if err := s.deleteAlertmanagerConfigs(ctx, []string{req.InstanceId}); err != nil {
				return errors.Wrapf(err, "delete AlertmanagerConfig for policy [%s] failed", req.InstanceId)
			}

			// 重新为告警规则创建AlertmanagerConfig
			amcfgs := buildAlertmanagerConfig(rules)
			for _, amcfg := range amcfgs {
				if err := s.store.AlertmanagerConfig().Create(ctx, amcfg); err != nil {
					return errors.Wrap(err, "create AlertmanagerConfig failed")
				}
			}
		}
	} else {
		if err := s.deleteAlertmanagerConfigs(ctx, []string{req.InstanceId}); err != nil {
			return errors.Wrapf(err, "delete AlertmanagerConfig for policy [%s] failed", req.InstanceId)
		}
	}

	// 更新PrometheusRule实例
	promRule.Spec = model.PrometheusRuleSpec{
		Groups: []model.RuleGroup{
			{
				Name:     policy.InstanceId,
				Interval: EvaluateInterval,
				Rules:    rules,
			},
		},
	}

	newPR, err := s.store.PromRuleOperator().Update(ctx, promRule)
	if err != nil {
		return errors.Wrapf(err, "update PrometheusRule [%s] failed", newPR.Name)
	}

	return nil
}

// 根据更新请求构造AlarmPolicy实例
func buildPolicyInstanceFromUpdateReq(req *dto.UpdatePolicyRequest) (*model.AlarmPolicy, error) {
	policy := &model.AlarmPolicy{
		Id:          req.Id,
		InstanceId:  req.InstanceId,
		Name:        req.Name,
		Comment:     req.Comment,
		Creator:     req.Creator,
		Updater:     req.Updater,
		Limit:       req.Limit,
		Production:  req.Production,
		Type:        req.Type,
		NotifySetup: req.NotifySetup,
		UpdatedAt:   time.Now(),
	}

	// 判断告警策略类型
	switch req.Type {
	case model.FormAlarmPolicyType:
		// 表单类型
		policy.FormPolicy = req.FormPolicy
		policy.ResourceType = req.FormPolicy.ResourceType
		policy.ResourceSubType = req.FormPolicy.ResourceSubType

		// 添加告警规则id
		for i := 0; i < len(policy.FormPolicy.Rules); i++ {
			if policy.FormPolicy.Rules[i].RuleId == "" {
				policy.FormPolicy.Rules[i].RuleId = xid.New().String()
			}
		}
	case model.PromAlarmPolicyType:
		// promql类型
		policy.PromqlPolicy = req.PromqlPolicy

		// 添加告警规则id
		if policy.PromqlPolicy.RuleId == "" {
			policy.PromqlPolicy.RuleId = xid.New().String()
		}
	}

	// 判断是否会传入了联系人
	if req.Receivers == nil {
		policy.Receivers = req.Receivers
	}

	// 转换数据格式
	err := policy.FE2DB()
	if err != nil {
		return nil, err
	}

	return policy, nil
}
