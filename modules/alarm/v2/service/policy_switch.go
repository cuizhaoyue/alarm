package service

import (
	"context"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"github.com/pkg/errors"
)

// Switch 启动或停止告警策略
func (s *alarmService) Switch(ctx context.Context, req *dto.SwitchPolicyRequest) error {
	// 获取policy实例
	policy, err := s.store.Alarm().Get(ctx, req.InstanceId)
	if err != nil {
		return errors.WithStack(err)
	}

	if policy.Enabled == req.Enable {
		return nil
	}

	policy.DB2FE()

	policy.Enabled = req.Enable

	if err := s.store.Alarm().Update(ctx, policy); err == nil {
		labels, annotations := buildCommonLablesAndAnnotations(policy)
		switch req.Enable {
		case false:
			// 停止告警策略，删除PrometheusRule和AlertmanagerConfig
			if err := s.store.PromRuleOperator().Delete(ctx, policy.InstanceId); err != nil {
				return errors.WithStack(err)
			}
			if checkBuildAm(labels, annotations) {
				if err := s.deleteAlertmanagerConfigs(ctx, []string{req.InstanceId}); err != nil {
					return errors.WithStack(err)
				}
			}
		case true:
			// 启动告警策略，创建PrometheusRule
			// 根据告警策略构造告警规则列表
			rules := buildRulesFromPolicy(policy, labels, annotations)

			// 构建PrometheusRule对象
			promRule := buildPrometheusRule(policy.InstanceId, policy.InstanceId, rules, policy.Limit)

			// 调用operator接口，创建PrometheusRule，获取返回的PrometheusRule实例
			_, err := s.store.PromRuleOperator().Create(ctx, promRule)
			if err != nil {
				return errors.WithStack(err)
			}

			// 创建AlertmanagerConfig
			if checkBuildAm(labels, annotations) {
				amcfgs := buildAlertmanagerConfig(rules)
				for _, amcfg := range amcfgs {
					if err := s.store.AlertmanagerConfig().Create(ctx, amcfg); err != nil {
						return errors.Wrap(err, "create AlertmanagerConfig failed")
					}
				}
			}
		}
	} else {
		return errors.WithStack(err)
	}

	return nil
}
