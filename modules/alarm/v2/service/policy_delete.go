package service

import (
	"context"
	"strings"
	"sync"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dao"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/dto"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/modules/alarm/v2/model"
	"github.com/pkg/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeleteCollection 批量删除告警策略
func (s *alarmService) DeleteCollection(ctx context.Context, req *dto.DeletePolicyRequest) error {
	// 构造labelSelector, 格式为 `key in (v1, v2)`
	labelSelector := PromRuleNameKey + " in " + "(" + strings.Join(req.InstanceIds, ",") + ")"
	listOpts := metav1.ListOptions{
		LabelSelector: labelSelector,
	}

	// 删除PrometheusRule
	if err := s.store.PromRuleOperator().DeleteCollection(ctx, listOpts); err != nil {
		return errors.Wrap(err, "batch delete PrometheusRule failed")
	}

	// 删除AlertmanagerConfig
	if err := s.deleteAlertmanagerConfigs(ctx, req.InstanceIds); err != nil {
		return errors.Wrap(err, "batch delete AlertmanagerConfig failed")
	}

	// 删除数据库中的记录
	if err := s.store.Alarm().Tx(ctx, func(ctx context.Context, store dao.Factory) error {
		// 删除策略表中的记录
		if delErr := store.Alarm().DeleteCollection(ctx, req.InstanceIds); delErr != nil {
			return delErr
		}

		// 删除和策略关联的资源记录
		if delErr := store.ResourceOnPolicy().DeleteCollection(ctx, req.InstanceIds); delErr != nil {
			return delErr
		}

		return nil
	}); err != nil {
		// 如果数据库中的记录删除失败需要重建PrometheusRule和AlertmanagerConfig
		// 获取对应的告警策略
		policies, listErr := s.store.Alarm().ListByInstanceIds(ctx, req.InstanceIds)
		if listErr != nil {
			return errors.Wrap(err, "get policy list by instance id failed")
		}

		wg := sync.WaitGroup{}
		finished := make(chan struct{}, 1)

		for _, p := range policies {
			wg.Add(1)
			go func(policy *model.AlarmPolicy) {
				defer wg.Done()

				labels, annotations := buildCommonLablesAndAnnotations(policy)
				rules := buildRulesFromPolicy(policy, labels, annotations)

				// 构建PrometheusRule对象
				promRule := buildPrometheusRule(policy.InstanceId, policy.InstanceId, rules, policy.Limit)

				// 调用operator接口，创建PrometheusRule，获取返回的PrometheusRule实例
				// 如果创建过程中出现错误不会中断，需要把所有的告警策略都创建完成，dao层日志会打印未完成创建的PrometheusRule名称
				_, _ = s.store.PromRuleOperator().Create(ctx, promRule)

				var amcfgs []*model.AlertmanagerConfig
				if checkBuildAm(labels, annotations) {
					amcfgs = buildAlertmanagerConfig(rules)
				}
				for _, amcfg := range amcfgs {
					_ = s.store.AlertmanagerConfig().Create(ctx, amcfg)
				}
			}(p)
		}
		go func() {
			wg.Wait()
			close(finished)
		}()

		<-finished

		return errors.Wrap(err, "delete policy failed")
	}

	return nil
}

// 删除所有具体相同告警策略id的AlertmanagerConfig实例
func (s *alarmService) deleteAlertmanagerConfigs(ctx context.Context, policyIds []string) error {
	// 构造labelSelector, 格式为 `policy_id in (v1, v2)`
	labelSelector := consts.PolicyId + " in " + "(" + strings.Join(policyIds, ",") + ")"
	listOpts := metav1.ListOptions{
		LabelSelector: labelSelector,
	}

	err := s.store.AlertmanagerConfig().DeleteCollection(ctx, listOpts)
	if err != nil {
		return err
	}

	return nil
}
